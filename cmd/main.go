package main

import (
	"fmt"
	"utils/utils"
)

func main() {
	or := make(map[int]utils.Order)
	fmt.Println(or)
	or[1] = utils.Order{}
	fmt.Println(or[1])
}
