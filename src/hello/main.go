package main

import (
	_ "fmt"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	log.Println("Hello Go world.")
}
