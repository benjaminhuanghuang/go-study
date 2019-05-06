## 应用程序入口
- 必须是 main package
- 必须是 main() method
- 文件名不一定是 main.go

- main() method 没有返回值, using os.Exit(1) return status.

- main() 不支持传入参数，using os.Args 
```
  package main

  func main(){
    os.Args[1]

    os.Exit(-1)
  }
```
