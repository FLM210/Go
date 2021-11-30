package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Start")
	s1 := "1222"
	n1, _ := strconv.Atoi(s1)           //字符串转int
	n2, _ := strconv.ParseFloat(s1, 32) //字符串转float 32为float32
	fmt.Print(n1, n2)
}
