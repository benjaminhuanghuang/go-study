##
brew uninstall go
brew install go

code ~/.zshrc 
code ~/.bashrc    # for original shell

export GOPATH=$HOME/Documents/Workspace/Go
export PATH="$PATH:user/local/opt/go/libexec/bin:GOPATH/bin"
