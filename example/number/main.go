package main

import (
	"fmt"
	"github.com/lizongying/go-utils"
)

func main() {
	a := []int8{1, 3, 2}

	//3
	fmt.Println(utils.Max(a...))

	b := []uint64{1, 3, 2}

	//1
	fmt.Println(utils.Min(b...))
}
