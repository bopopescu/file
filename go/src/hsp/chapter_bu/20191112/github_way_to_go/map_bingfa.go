package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
	lock     sync.Mutex
}

func NewBan(ctx context.Context) *Ban {
	o := &Ban{visitIPs: make(map[string]time.Time)}
	go func() {
		ad := 1
		timer := time.NewTimer(time.Minute * 1)
		for {
			select {
			case <-timer.C:
				fmt.Println(123)
				o.lock.Lock()
				for k, v := range o.visitIPs {
					if time.Now().Sub(v) >= time.Minute*1 {
						delete(o.visitIPs, k)
					}
				}
				o.lock.Unlock()
				timer.Reset(time.Minute * 1)
			case <-ctx.Done():
				fmt.Println(321)
				return
			default:
				//fmt.Println(ad)
				ad++

			}
		}
	}()
	return o
}

func (o *Ban) visit(ip string) bool {
	o.lock.Lock()
	defer o.lock.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}

func main() {
	fuck := time.Now()
	success := int64(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ban := NewBan(ctx)

	wait := &sync.WaitGroup{}

	fk := 1
	wait.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wait.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				fk++
				fmt.Println(ip, fk)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}

	}
	wait.Wait()

	time.Sleep(1e9)
	fmt.Println(time.Minute, time.Second)
	fmt.Println(fuck)
	fmt.Println(time.Now())
	fmt.Println(time.Now().Sub(fuck))
	fmt.Println("success:", success)
}
