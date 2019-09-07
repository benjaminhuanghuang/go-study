# Struct

- 仅支持封装，不支持继承和多态
- Struct 没有构造函数，需要指定 factory method,factory method返回局部变量的地址，在外部依然有效。Go语言来决定在heap还是在
stack上分配内存
- Sturct 作为函数参数是值传递
