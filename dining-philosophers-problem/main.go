// Dining philosophers problem
// https://en.wikipedia.org/wiki/Dining_philosophers_problem

package main

import "fmt"
import "time"
import "sync"

func philo(index int, sticks []bool, wg *sync.WaitGroup) {
	left := index
	right := (index + 1) % len(sticks)

	for {
		if sticks[left] && sticks[right] {
			sticks[right], sticks[left] = false, false
			fmt.Printf("Philo %d is eating\n", index)
			time.Sleep(2000 * time.Millisecond)
			sticks[right], sticks[left] = true, true

			fmt.Printf("Philo %d is resting\n", index)
			time.Sleep(3000 * time.Millisecond)
		} else {
			fmt.Printf("Philo %d is waiting\n", index)
			time.Sleep(1000 * time.Millisecond)
		}
	}
	wg.Done()
}

func main() {
	sticks := []bool{true, true, true, true, true}
	var wg sync.WaitGroup
	wg.Add(len(sticks))
	
	go philo(0, sticks, &wg)
	go philo(1, sticks, &wg)
	go philo(2, sticks, &wg)
	go philo(3, sticks, &wg)
	go philo(4, sticks, &wg)

	wg.Wait()
}