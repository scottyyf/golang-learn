# 第二章 程序的结构

1. 命名。
    - 保留命名
      ```
      break[python] map[python] else[python]
      if[python] type[python] continue[python]  
      for[python] import[python] return[python]
      range[python]
      
      default func interface select case var
      defer go  struct chan  goto packge switch
      const fallthrough
      ```
    
    - 内建
      * 常量： true false iota nil
      * 内建类型： int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr float32 float64 complex128
         complex64 bool byte rune string error
      * 内置函数： make len cap new append copy close delete complex real imag panic recover

2. 作用域： 函数内部定义，包中定义。变量名大写开头的包级名字，包外可访问
3. 驼峰命名。与python推荐不一样
4. 声明。四种声明语句
   - var 变量
   - const 常量
   - type 类型
   - func 函数对象
   
5. .go结尾的源文件，语句开始声明属于哪个包，然后import，最后func定义
6. 交换值 i, j = j, i。同python
7. 指针。
   - 如果p是指针，则*ｐ对应指针指向的变量的值，*p可以出现在表达式的左边，表示更新指针指向的变量的值
   ```go
   var x int
   &x 指向x的内存地址，指针的类型是*int 即指向int类型的指针
   ```

   - 指针的初始值为nil。指针相等：只有指向同一个变量或都为nil才相等
   
8. flag包可用来解析命令行。
   ```go
   var n = flag.Bool("n", false, "omit taiiling newline") //n是指针m,是bool的，默认false
   var sep = flag.String("s", ", ", "separator") // s是指针,string的，默认", "
   
   func main()  {
        flag.Parse() //初始化
        fmt.Print(strings.Join(flag.Args(), *sep)) //连接string
        if !*n{
            fmt.Println()
        }
   }
   ```

9. new函数。new函数将创建一个匿名变量，返回的是变量地址。多次执行相同的new函数，返回的东西并不是同一个东西，
这个和python多次初始化一个不可变类型不同
    ```go
    p := new(int) // p是*int类型的指针，默认值为0
    println(*p) // 返回指针对应的内容，0
    ```
   
