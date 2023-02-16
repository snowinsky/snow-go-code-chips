package main

import "fmt"

func main() {
	var aaa = 11
	//switch 语句中的case里面自带break,匹配到case之后就不会执行下面的了，如果想让它继续把下面的都跑完，就加fallthrough
	switch aaa {
	case 1:
		fmt.Println("go here aaa=1")
	case 2:
		fmt.Println("go here aaa=2")
	case 3:
		fmt.Println("go here aaa=3")
	case 4:
		fmt.Println("go here aaa=4")
	case 5:
		fmt.Println("go here aaa=5")
	case 6:
		fmt.Println("go here aaa=6")
	case 7:
		fmt.Println("go here aaa=7")
	case 8:
		fmt.Println("go here aaa=8")
	case 9:
		fmt.Println("go here aaa=9")
	case 10:
		fmt.Println("go here aaa=10")
	case 11: //比如aaa等于11时，会先执行aa下面的语句，然后遇上fallthrough之后不会break，而是直接执行12下面的语句，直到跳出switch范围
		fmt.Println("go here aaa=11")
		fallthrough
	case 12:
		fmt.Println("go here aaa=12")
		fallthrough
	default:
		fmt.Println("go here aaa=default")
	}
}
