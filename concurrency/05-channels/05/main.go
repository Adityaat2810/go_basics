package main

import "fmt"

/**
Closing Channels
close(ch)
x, ok := <-ch // ok-> false (since channel is close )


---------------------------
ch := make(chan int)

go func() {
	ch <- 1
	ch <- 2
	close(ch)
}()

for {
	val, ok := <-ch
	if !ok {
		break
	}
	fmt.Println(val)
}
------------------------

close(ch)
ch <- 10
-- never send on closed channel (cause panic)
**/


func main(){
	ch := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
	//Automatically exits when channel closes
    //Cleaner than manual ok check


}

// channels directions
func sendOnly(ch chan<- int) {
	ch <- 10
}

func receiveOnly(ch <-chan int) {
	fmt.Println(<-ch)
}

