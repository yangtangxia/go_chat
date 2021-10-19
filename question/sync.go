package main

import "sync"

type Info struct {
	mu sync.Mutex // 互斥锁，它的作用是守护在临界区入口来确保同一时间只能有一个线程进入临界区。
	Str string
}

func Update(info *Info)  {
	info.mu.Lock()
	info.Str = "1"
	info.mu.Unlock()
}