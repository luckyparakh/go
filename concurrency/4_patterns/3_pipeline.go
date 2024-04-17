package main

import "fmt"
/*
dataToChannel is a function that takes a slice of integers as input and returns a receive-only channel of integers.

Parameters:
- nums: A slice of integers that will be sent to the channel.

Returns:
- out: A receive-only channel of integers.

Description:
The dataToChannel function creates a new channel called 'out' using the make function. 
It then creates a goroutine to send each element of the 'nums' slice to the 'out' channel. Once all the numbers have been sent, the channel is closed using the close function. 
The function returns the 'out' channel even if the goroutine is still processing.
*/
func dataToChannel(nums []int) <-chan int {
	// Create new out channel 
	out := make(chan int)

	// Create a go routine, so that parent func is not blocked
	go func() {
		for i := range nums {
			out <- nums[i]
		}
		// Close this channel once all number are processed
		close(out)
	}()

	// Return output channel even previous GO routine is still processing
	return out
}
/*
sq is a function that takes a receive-only channel of integers as input and returns a receive-only channel of integers.

Parameters:
- in: A receive-only channel of integers.

Returns:
- out: A receive-only channel of integers.

Description:
The sq function creates a new channel called 'out' using the make function. 
It then creates a goroutine to receive each element from the 'in' channel, square the value, and send it to the 'out' channel. Once all the values have been processed, the 'out' channel is closed using the close function. 
The function returns the 'out' channel even if the goroutine is still processing.
*/
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for value := range in {
			out <- value * value
		}
		close(out)
	}()
	return out
}

func main() {
	// Stage one from where data is sent to channel
	s1 := dataToChannel([]int{1, 2, 3})
	// Stage two where data is squared
	s2 := sq(s1)
	for v := range s2 {
		fmt.Println(v)
	}
}
