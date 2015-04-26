package main

import (
	_ "fmt"
	"log"
	"my_math"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	my_math.Fib(8)
}
