package main

import (
	"fmt"
	"github.com/lizongying/go-utils"
)

func main() {
	a := []string{"1", "3", "2"}
	utils.SliceAsc(a)

	//[1 2 3]
	fmt.Println(a)

	b := []uint64{1, 3, 2}
	utils.SliceDesc(b)

	//[3 2 1]
	fmt.Println(b)
}
