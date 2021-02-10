package main

import (
	"fmt"
	"github.com/aibotsoft/clipboard"
)

func main() {
	fmt.Println("fuck off")
	all, err := clipboard.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(all)
}
