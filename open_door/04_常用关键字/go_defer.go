package main

import "fmt"

func main() {
	//defer就是try catch finally中的finally，会在return之后执行
	for i := 0; i < 10; i++ {
		fmt.Println(abcdef(i))
	}
	fmt.Println("##############")
	//如果defer调用的函数是闭包呢？
	for i := 0; i < 10; i++ {
		fmt.Println(abcdefClosePackage(i))
	}
}

func abcdef(a int) bool {
	var returnBool = false
	//defer定义在这里时，当代码运行到这里时，returnBool的值就确定为false了，只是这句话没被执行而已。必须等到这个函数return命令执行完了的前一刻才被执行
	defer fmt.Println("the finally return=", returnBool)
	if a > 1 && a < 5 {
		returnBool = true
		//defer fmt.Println("the finally return=", returnBool)
		//如果defer写在上面那一行，returnBool的值就会是true，并且，只有在这个条件执行完毕且return之前才会被执行，else的代码的return不会执行上面那句话
		return returnBool
	} else {
		returnBool = false
		return returnBool
	}

}

func abcdefClosePackage(a int) bool {
	var returnBool = false
	//defer定义在这里时，当代码运行到这里时，函数没入参，也不执行，等执行时，拿取returnBool的值，那个时候returnBool是什么值，就取什么值
	/*defer func(){
		fmt.Println("the finally return=", returnBool)
	}()*/
	if a > 1 && a < 5 {
		returnBool = true
		defer func() {
			fmt.Println("the finally return=", returnBool)
		}()
		//如果defer写在上面那一行，returnBool的值就会是true，并且，只有在这个条件执行完毕且return之前才会被执行，else的代码的return不会执行上面那句话
		return returnBool
	} else {
		returnBool = false
		return returnBool
	}

}
