package main

/*

1. Basics
  A goroutine is a lightweight concurrent function execution managed
  by the Go runtime, not by the OS.

  go doSomething()
    That go keyword tells the Go runtime:
    “Run this function concurrently, schedule it efficiently, don’t block my thread.”

2. Why cheap
  1. Dynamic stack growth
    - goRoutines start with 2kb size
    - stack grow automatically when needed
    - stack shrinks when unused
  2. M:N Scheduling Model
    - Go uses m:n scheduling
	- m (goroutines)
	- n(os threds)
	- The runtime multiplexes many goroutines onto fewer threads.


****************** G-P-M model ( for interviews ) *******************
G - goroutine
P - processor(logical schedular context)
M- OS Thread(machine)

G needs a P to run
P needs an M

M executes goroutines assigned by P
G → P → M → CPU

Why P exists?
  - Limits parallelism
  - Each P has its own run queue
  - Reduces lock contention

GOMAXPROCS
  runtime.GOMAXPROCS(n)
  - Controls number of P
  - Defaults to number of CPU cores
  - More P = more parallel execution

Goroutines are concurrent by default, parallel only if GOMAXPROCS > 1.
[IMP] Concurrency: Dealing with many tasks at once (switching).
Parallelism: Doing many tasks at once (simultaneously).


You create a G. It joins a queue inside a P.
An M grabs a P to start working.
The M pulls a G from that P's queue and runs it.


*/


