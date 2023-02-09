// 每个go project都是从main package的main方法开始的
// 类的文件名不一定和package的名字相同，最好相同
// go项目，每一个文件夹都是一个独立的个体，独立的模块。负责一个独立的职责
package main

//导入需要使用的包
import "fmt"

// main 方法
// 如果这个project不是建在GOPATH路径下，则可以使用go mod来进行处理
// go mod init {具体工程名} //用来创建一个新的go module
// go mod tidy //可以把需要依赖的lib都从远端下载下来，下载到GOPATH的pkg文件夹下
// go mod vendor //就是把所需的依赖拉到该module下
// build deploy方法
// go build xxx.go //生成当下操作系统下的可执行文件
// set GOARCH=amd64; set GOOS=linux //设置完了这个就可以生成linux下的可执行文件了
// go生成的工程可执行文件都是很大的，因为每个都包含runtime
func main() {
	fmt.Println("hello world111")
}
