## Reference
- https://www.digitalocean.com/community/tutorials/understanding-the-gopath


先通过go env查看go的环境变量
```
  GOBIN=""
  GOPATH="/Users/bhuang/go"
  GOROOT="/usr/local/go"
```

GOROOT 是golang 的安装路径


$GOPATH 可以理解为工作目录
这个目录指定了需要从哪个地方寻找GO的包、可执行程序等
GOPATH之下主要包含三个目录: bin、pkg、src：
```
  ├── bin
  ├── pkg
  └── src
    └── github.com/foo/bar
      └── bar.go
```
- bin目录包含了可执行程序，注意是可执行的，不需要解释执行。
- pkg目录包含了使用的包或者说库。
- src里面包含了go的代码源文件，其中仍按包的不同进行组织。

GOPATH可以是一个目录列表, go get下载的第三方库, 一般都会下载到列表的第一个目录里面


go get会做两件事：
1. 从远程下载需要用到的包
2. 执行go install

go install 编译的是可执行文件,会生成可执行文件直接放到bin目录下，如果是一个普通的包，会被编译生成到pkg目录下该文件是.a结尾

