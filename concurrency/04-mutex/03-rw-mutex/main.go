package main

import "sync"

/***

====== Why RW Mutex ? ===========
sometimes
  - many goroutines reads
  - few goroutines write

Mutex block everything rw mutex is samrter

syntax 'var mu sync.RWMutex'

mu.RLock()   // multiple readers allowed
mu.RUnlock()

mu.Lock()    // exclusive writer
mu.Unlock()



***/

// example
type Store struct{
  data map[string] int
  mu sync.RWMutex
}

func (s *Store) Get(key string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data[key]
}

func (s *Store) Set(key string, val int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = val
}

// If writes are frequent → RWMutex may hurt performance

func main(){

}

// mutex is not re-entrant
/*
-> means
A goroutine cannot lock the same mutex twice

mu.Lock()
mu.Lock() // ❌ DEADLOCK

Why?
Go mutex does not track goroutine ownership

*/


/*
🔟 Deadlock — classic interview trap 💀

eg 1
mu.Lock()
defer mu.Unlock()

funcA() // funcA also tries mu.Lock()

eg 2
go func() {
	mu.Lock()
}()

mu.Lock() // main waits forever

Deadlock conditions (INTERVIEW THEORY)
- Mutual exclusion ✔
- Hold and wait ✔
- No preemption ✔
- Circular wait ✔
*/