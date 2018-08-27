package main

// 1. 工作空间：src / bin / pkg 三个文件夹

// 2. 导入包的多种方法： 别名法； 初始化方法：目标包


// 3. 组织结构：


// 4. 权限： 所有成员包内可见， 但是包外只能使用大写开头的成员。 如果想使用，必须使用 指针结构体转换。


// 5. 初始化： init / main


// 6. 内部包： internal 存放在 lib 目录下（src 总目录下），使用时必须用完整路径：lib/internal/internalA， internal 目录下的各个包能相互访问，也可以访问外部包


// 7. 外部包：vender 优先级比标准库优先级更高

func main(){

}
