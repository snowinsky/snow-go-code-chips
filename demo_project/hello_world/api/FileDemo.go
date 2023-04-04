package api

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

/*
https://zhuanlan.zhihu.com/p/260661233
*/
func FileDemo() {
	f := createAndWriteTxtFile()
	openFileAndAppendTxtFile(f)
	openAndReadLineTxtFile(f)
}

func openAndReadLineTxtFile(f string) {
	ff, err := os.Open(f)
	if err != nil {
		panic("open file err" + err.Error())
	}
	defer ff.Close()

	r := bufio.NewReader(ff)
	for {
		line, isPrefix, lineErr := r.ReadLine()
		if lineErr != nil {
			if lineErr == io.EOF {
				break
			}
			panic("read line error" + lineErr.Error())
		}
		log.Println(isPrefix)
		log.Println(string(line))
	}

	ff.Seek(0, 0)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(" read line by readString error" + err.Error())
		}
		log.Println("ReadString:" + line)
	}
	ff.Seek(0, 0)
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(" read line by readString error" + err.Error())
		}
		log.Println("ReadBytes:" + string(line))
	}
}

/*
第一个参数表示打开文件的路径，这里必须是全路径

第二个参数表示模式，常见的模式有

O_RDONLY(只读模式)
O_WRONLY(只写模式),
O_RDWR( 可读可写模式)
O_APPEND(追加模式)。
第三个参数，表示权限，取值范围（0-7）表示如下：

0：没有任何权限
1：执行权限 x(如果是可执行文件，是可以运行的)
2：写权限 r
3：写权限与执行权限 xr
4：读权限 w
5：读权限与执行权限 xw
6：读权限与写权限 wr
7：读权限，写权限，执行权限 xwr
*/
func openFileAndAppendTxtFile(f string) {
	log.Println(f)
	ff, err := os.OpenFile(f, os.O_APPEND, 2)
	if err != nil {
		panic("open file error")
	}
	defer ff.Close()
	for i := 0; i < 10; i++ {
		ff.WriteString(strconv.Itoa(i) + "::append data::" + "\r\n")
	}
}

func createAndWriteTxtFile() string {
	newFileName := "abc.txt"
	/*
	  create表示有文件则清空数据，文件不存在则创建
	*/
	f, err := os.Create(newFileName)
	if err != nil {
		panic("create file error")
	}
	defer f.Close()

	for i := 0; i < 10; i++ {
		f.WriteString(strconv.Itoa(i) + "::insert new record::" + "\r\n")
	}

	return newFileName
}
