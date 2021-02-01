

https://www.jianshu.com/p/697581546eac



.vscode是vscode的配置文件所在目录
```
"go.inferGopath": true,
"go.goroot": "C:/Go/",
"go.gopath": "C:/Go/gopath/go/",
"go.toolsGopath":"C:/Go/gopath/go/",
```
vscode的 可以用上面四行设置中的路径覆盖环境变量。
其中 go.toolsGopath 也设置上比较好。不然会在本地项目中创建bin目录，麻烦。指定一个固定位置，多个项目可以共享这些工具。


"go.inferGopath": true 的作用是把本地项目的src也包含到gopath中，但是不污染全局gopath

## launch.json

让vscode的调试和运行菜单关联到运行的代码
```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceRoot}/src/main.go",
            "env": {
                //"GOPATH":"C:/Go/gopath/go/",
                //"GOROOT":"C:/Go/",

            },
            "args": []
        }
    ]
}
```


检查当前vscode工程的GOPATH路径的命令 
```
vscode 菜单 view -> command palette -> GO:current GOPATH
```
