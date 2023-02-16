package main

import "fmt"

func main() {
	//go中的if条件判断和java中一样
	var aaa = 11
	if aaa < 5 {
		fmt.Println("go here aaa<5")
	} else if aaa == 5 {
		fmt.Println("go here aaa=5")
	} else {
		fmt.Println("go here aaa>5")
	}
}
