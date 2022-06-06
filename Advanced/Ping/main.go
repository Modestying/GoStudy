package main

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
)

func ServerPing(target string) (bool, error) {
	var ICMPCOUNT = 2
	var PINGTIME = time.Duration(1)
	pinger, err := ping.NewPinger(target)
	if err != nil {
		return false, err
	}
	pinger.Count = ICMPCOUNT
	pinger.Timeout = time.Duration(PINGTIME * time.Millisecond)
	pinger.SetPrivileged(true)
	pinger.Run() // blocks until finished
	stats := pinger.Statistics()
	// 有回包，就是说明IP是可用的
	if stats.PacketsRecv >= 1 {
		return true, nil
	}
	return false, nil
}
func main() {
	pinger, err := ping.NewPinger("192.168.10.1")
	pinger.SetPrivileged(true)
	if err != nil {
		panic(err)
	}
	pinger.Count = 1
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	fmt.Println(stats)
	if stats.PacketsRecv > 0 {
		fmt.Println("设备在线")
	}
}
