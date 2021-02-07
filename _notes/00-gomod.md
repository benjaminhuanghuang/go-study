go path的问题：
- 代码开发必须在go path src目录下，不然，就有问题。
- 依赖手动管理
- 依赖包没有版本可言

从 v1.5 开始开始引入 vendor 包模式，如果项目目录下有 vendor 目录，那么 go 工具链会优先使用 vendor 内的包进行编译、测试等

govendor 的问题：
- 解决了包依赖，一个配置文件就管理
- 依赖包全都下载到项目vendor下，每个项目都有一份


go modules 是 golang 1.11 新加的特性

```
  go env -w GO111MODULE=on
```
GO111MODULE 有三个值：off, on和auto（默认值）。
- GO111MODULE=off，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。
- GO111MODULE=on，go命令行会使用modules，而一点也不会去GOPATH目录下查找。
- GO111MODULE=auto，默认值，go命令行将会根据当前目录来决定是否启用module功能。

这种情况下可以分为两种情形：
- 当前目录在GOPATH/src之外且该目录包含go.mod文件
- 当前文件在包含go.mod文件的子目录下面。

当modules功能启用时，依赖包的存放位置变更为$GOPATH/pkg，允许同一个package多个版本并存，且多个项目可以共享缓存的 module

gomod 会在 $GOPATH/pkg/mod 目录下保存包源码和链接库。

## go mod command
```
  go mod download 下载模块到本地缓存，缓存路径是 $GOPATH/pkg/mod/cache, 多个项目可以共享缓存的module。

  go mod edit 是提供了命令版编辑 go.mod 的功能，例如 go mod edit -fmt go.mod 会格式化 go.mod
  go mod graph 把模块之间的依赖图显示出来
  go mod init 初始化模块（例如把原本dep管理的依赖关系转换过来）
  go mod tidy 增加缺失的包，移除没用的包
  go mod vendor 把依赖拷贝到 vendor/ 目录下
  go mod verify 确认依赖关系
  go mod why 解释为什么需要包和模块
```

## go module 管理
```
  go mod init   <project id>

  go build             # 编译安装

  go list -m list      # 查看所有以升级依赖版本

  go get

  go mod tiny          # 删除不必要的库
```
