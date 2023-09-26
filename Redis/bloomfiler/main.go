package main

import (
	"fmt"
	"math"

	"github.com/spaolacci/murmur3"
)

func main() {
	fmt.Println(34 >> 5)
	fmt.Println(34 & 31)

	A := 1
	A = 1 << 1
	fmt.Println(A)
	A = 1
	A = A | (1 << 3)
	// old: 0001
	// new: 1000
	fmt.Println(A)
	return
	fmt.Println("bloom filter...")
	encry := NewEncryptor()
	fmt.Println(encry.Encrypt("xx"))
	fmt.Println(encry.Encrypt("xx"))

}

type Encryptor struct {
}

func NewEncryptor() *Encryptor {
	return &Encryptor{}
}

func (e *Encryptor) Encrypt(origin string) int32 {
	hasher := murmur3.New32()
	hasher.Write([]byte(origin))
	return int32(hasher.Sum32() % math.MaxInt32)

}

type LocalBloomService struct {
	m, k, n   int32
	bitmap    []int
	encryptor *Encryptor
}

func NewLocalBloomService(m, k int32, encryptor *Encryptor) *LocalBloomService {
	return &LocalBloomService{
		m:         m,
		k:         k,
		bitmap:    make([]int, m/32),
		encryptor: encryptor,
	}
}

func (l *LocalBloomService) getEncrypted(val string) []int32 {
	encrypteds := make([]int32, 0, l.k)
	origin := val
	for i := 0; int32(i) < l.k; i++ {
		encrypted := l.encryptor.Encrypt(origin)
		encrypteds = append(encrypteds, encrypted%l.m)
		if int32(i) == l.k-1 {
			break
		}
		origin = string(encrypted)
	}
	return encrypteds
}

func (l *LocalBloomService) Exist(val string) bool {
	for _, offset := range l.getEncrypted(val) {
		index := offset >> 5     // == /32
		bitOffset := offset & 31 // == %32
		if l.bitmap[index]&(1<<bitOffset) == 0 {
			return false
		}
	}
	return true
}

func (l *LocalBloomService) Set(val string) {
	l.n++
	for _, offset := range l.getEncrypted(val) {
		index := offset >> 5     // 索引
		bitOffset := offset & 31 // 获取余数，计算偏差量
		l.bitmap[index] = l.bitmap[index] | (1 << bitOffset)
	}
}
