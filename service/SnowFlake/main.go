package main

import (
	"sync"
	"time"
)

/*
 * 1 符号位  |  41 时间戳                                    | 2 节点   | 11 （秒内）自增ID
 * 0        |  000000 00000000 00000000 00000000 00000000 000 |   00   | 000000 00000
 * 按照此方法，每秒每个节点可以产生的ID数量为: 2048个，即2^11个ID，三个节点每秒可以产生 2048 * 3 = 6144个ID
 * 此方法可以确保生成的ID位数为12位的数字，若想在每秒生成更多的ID，则产生的ID的位数会大于12位
 */

const (
	epoch          = int64(1577808000000) // 设置起始时间(时间戳/毫秒)：2020-01-01 00:00:00，有效期69年
	timestampBits  = uint(41)             //时间戳位数
	datacenterBits = uint(5)              //数据中心id位数
	workerIdBits   = uint(5)              //机器id所占位数
	sequenceBits   = uint(12)             //自增序列所占位数

	timeStampMax      = int64(-1 ^ (-1 << timestampBits))          //时间戳最大值
	datacenterMax     = int64(-1 ^ (-1 << datacenterBits))         // 支持的最大数据中心id数量
	workeridMax       = int64(-1 ^ (-1 << workerIdBits))           // 支持的最大机器id数量
	sequenceMax       = int64(-1 ^ (-1 << sequenceBits))           // 支持的最大序列id数量
	workeridShift     = sequenceBits                               // 机器id左移位数
	datacenteridShift = sequenceBits + workerIdBits                // 数据中心id左移位数
	timestampShift    = sequenceBits + workerIdBits + workerIdBits // 时间戳左移位数
)

type IdWorker struct {
	sync.Mutex
	timestamp    int64
	workid       int64
	datacenterid int64
	sequence     int64
}
type ArgInalid func(data int64) error

func IsValidDataCenter(data int64) bool {
	if data < 0 || data > datacenterMax {
		return false
	}
	return true
}

func IsValidWorkerId(data int64) bool {
	if data < 0 || data > workeridMax {
		return false
	}
	return true
}

func NewIdWoker(datacenterid, workerID int64) (*IdWorker, bool) {
	if !IsValidDataCenter(datacenterid) || !IsValidWorkerId(workerID) {
		return nil, false
	}
	instance := &IdWorker{
		timestamp:    0,
		datacenterid: datacenterid,
		workid:       workerID,
		sequence:     0,
	}
	return instance, true
}

func (worker *IdWorker) NextId() int64 {
	worker.Lock()
	defer worker.Unlock()
	now := time.Now().UnixNano() / 1e6
	if worker.timestamp == now {
		worker.sequence = (worker.sequence + 1) & sequenceMax
		if worker.sequence == 0 {
			for now <= worker.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		worker.sequence = 0
	}
	t := now - epoch
	if t > timeStampMax {
		return 0
	}
	worker.timestamp = now
	r := int64(t)<<int64(timestampShift) | (worker.datacenterid << int64(datacenteridShift)) | (worker.workid << int64(workeridShift)) | (worker.sequence)
	return r
}

func GetTimeStamp(id int64) int64 {
	return id >> int64(timestampShift)
}

func GetGenTimestamp(id int64) int64 {
	return GetTimeStamp(id) + epoch
}

func GetGetnTime(id int64) string {
	return time.Unix(GetGenTimestamp(id)/1000, 0).Format("2023-08-017 21:53:22")
}

func s() {
	var x sync.WaitGroup
	x.Add(1)
}
