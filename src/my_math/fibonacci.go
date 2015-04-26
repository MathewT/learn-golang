package my_math

import (
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func Fib(n int) (val int64, err int) {
	err = 0
	if n > 99 {
		err = -1
		log.Printf("Invalid input: %d.  Input must be <= 99", n)
		return
	}
	val = 0
	//Compute the n'th Fibonacci number
	var current, previous, past int64
	current = 1
	previous = 1

	i := 0
	for i < n {
		past = previous
		previous = current
		current = past + previous
		i += 1
	}
	log.Printf("Fib number is: %d ", current)
	val = current
	return
}
