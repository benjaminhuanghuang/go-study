## Install go on mac
```
  brew uninstall go
  brew install go

  code ~/.zshrc 
  code ~/.bashrc    # for original shell
  code ~/.bash_profile   # for mac

  export GOPATH=$HOME/Documents/Workspace/Go
  export PATH="$PATH:user/local/opt/go/libexec/bin:GOPATH/bin"

  echo $GOPATH

  go env

  go version
```

## Update go on Mac
```
brew uninstall -f go

brew update

brew install go
```    

## Set go env
```
  go env -w GOBIN=/Users/youdi/go/bin
  go env -w GO111MODULE=on
  go env -w GOPROXY=https://goproxy.cn,direct // 墙内专用，使用七牛云的
```
 go env -w 会将配置写到 GOENV="/Users/<username>/Library/Application Support/go/env"

## Setup VS code
  Preferences -> settings, add go settings
  ```
  "go.gopath":"/Users/hhuang/Documents/Workspace/Go"
  ```

  CMD + Shift + P
  > go
  Go: Install/Update Tools

## Debugger 
  brew install go-delve/delve/delve
  or
  $ go get github.com/derekparker/delve/cmd/dlv

## GO
  go语言的工作空间其实就是一个文件目录，目录中必须包含src、pkg、bin三个目录。


