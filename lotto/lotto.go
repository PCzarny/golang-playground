package main

import "fmt"
import "math/rand"
import "time"

func main() {
	var guess int
    r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	score := r1.Intn(1000)
	for {
		fmt.Println("Type your proposition:")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println(err)
			break
		}
		if score == guess {
			fmt.Println("You win!!")
			fmt.Print("\n\n")
			break
		} else if score > guess {
			fmt.Println("Too small")
		} else {
			fmt.Println("Too big")
		}
		fmt.Print("\n")
	}
}