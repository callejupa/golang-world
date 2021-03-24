package main

import "fmt"

func main() {
	for i := 1; i < 10; i += 2 {
		fmt.Println(i)
		if i == 5 {
			fmt.Printf(" pasa el continue")
			continue
		}
		fmt.Printf(" pasó por aqui")
	}
}
