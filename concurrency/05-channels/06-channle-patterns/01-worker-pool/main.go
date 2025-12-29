package main

import "fmt"

func main(){

	jobs := make(chan int)


	go func() {
		// recieving
		for job := range jobs {
			fmt.Println("Processing", job)
		} // stops when channel close 
	}()

	// sending
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs) // close after sending

}