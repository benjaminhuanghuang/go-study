##
brew uninstall go
brew install go

code ~/.zshrc 
code ~/.bashrc    # for original shell
code ~/.bash_profile   # for mac

export GOPATH=$HOME/Documents/Workspace/Go
export PATH="$PATH:user/local/opt/go/libexec/bin:GOPATH/bin"

echo $GOPATH

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