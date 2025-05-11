package main

import (
	"fmt"
	"github.com/go-ping/ping"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 255; i++ {
		ip := fmt.Sprintf("150.1.1.%d", i)
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			pinger, err := ping.NewPinger(ip)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				return
			}
			pinger.Count = 3
			pinger.Run()
			stats := pinger.Statistics()
			fmt.Printf("%s\t%d\t%d\t%d\t%s\n", ip, stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss, stats.AvgRtt)
		}(ip)
	}
	wg.Wait()
}
