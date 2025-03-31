package worker

import ("fmt"
		"time"
	)

func (w) Worker {
	id int
	jobs <-chan int
	results chan<- int
}