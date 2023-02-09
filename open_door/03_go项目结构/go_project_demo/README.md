# GO 语言项目目录

- api 放对外的api文件，比如http的，protobuf的，thirft的等等
- assets
- build 执行go build命令
- cmd 放可执行文件，一般里面只放一个main.go来引用internal中或者pkg中的代码来实现项目的启动
- configs 放一些配置文件
- deployments 放一些部署用的脚本和工具的
- docs 项目文档
- examples 使用这个项目的一些例子
- init 放项目的一些初始化脚本的，例如初始化数据库，初始化文件之类的
- internal 自己的私有代码，不能被别的项目引用的，一般放internal/app下面
- pkg 可以给其他项目引用的代码
- test 放集成测试的代码的，一般单元测试都是在源代码同目录下，加_test
- third_party 放第三方的jar
- tools 放常用的工具
- vendor 引用的外部的资源，被下载到本地后放入的地方

> 此外，源代码还可以平铺在项目根目录下，不放文件夹里


- scripts 一些脚本文件的存放地
- Makefile 触发scripts中的脚本

## GO 语言模块划分

> GO语言模块划分并不是横向分层模式，不是从controller到service到dao这种模式。而是以职责为模块划分基础，user-manager，person-manager，customer-manager。
> GO语言的模块划分是从职责角度纵向分层模式
> GO语言项目中的每一个文件夹都是一个独立的包。

## GO 语言崇尚显式的初始化，错误处理，方法调用

> GO语言不鼓励使用init函数，就是不鼓励在类的构造函数中写入大量的逻辑。尽量把init函数中的逻辑抠出来放入type中，然后在适当的位置调用
> GO语言不鼓励使用try catch来操作error，鼓励把error当成一个普通的参数，可以传输或者抛出，每一层都要判断并处理
> GO语言鼓励使用type 。。。interface来暴露接口
>
> 
