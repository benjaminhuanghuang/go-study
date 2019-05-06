- [Everything you need to know about Packages in Go](https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc)

- https://thenewstack.io/understanding-golang-packages/


## Rules
- go run: cannot run non-main package

- don't use an underscore in package name (go-lint)
- don't use MixedCaps in package name; stringStudy should be stringstudy (go-lint)

- package 的约定是使用小写字符。
- package 以由多个文件组成，但是使用相同的 package 这一行。
- 名称以大写字母起始的是可导出的，可以在包的外部调用
- 一个包不可以存在于两个文件夹下，两个文件夹下的包是两个不同包
一个package的文件不能在多个文件夹下。

如果多个文件夹下有重名的package，它们其实是彼此无关的package。
如果一个go文件需要同时使用不同目录下的同名package，需要在import这些目录时为每个目录指定一个package的别名。