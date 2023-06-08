package main

import (
	"fmt"
	"github.com/lizongying/go-utils"
)

func main() {
	a := []string{"1", "3", "2"}

	//true
	fmt.Println(utils.InSlice("1", a))
}
