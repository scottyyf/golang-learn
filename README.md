# golang-learn

1. go语言中不能随意换行。比如 x + y
2. for 循环。for的条件没有小括号，且左大括号必须和for在同一行;for init;condition;post这里的init是可以省略的
    ```go
    for i:=1;i<3; i++{
        fmt.Prinln(i)
    }
    
    var i = 1
    for ;i<3; i++{
        // do something
    }
   
    for i < len(os.Args) { // 等价while loop
        s += sep + os.Args[i]
        sep = " "
        i ++
    }
   
    for {
    } //死循环
    ```

3. var声明变量。var 变量名 变量类型，声明时如果没有给值，那么默认都是对应类型的0值
    ```go
    var s, sep string
    ```
   
4. 赋值。支持赋值运算符 +=这种。不支持--i, 也不支持 x = i--;支持 i += 1, i++
    ```go
    s += sep
    s := "" //一般在函数中使用 推荐
    var s string //s = ""空字符 推荐
    var s = "" 同时声明多个变量时采用
    var s string = "" 同时声明多个变量时采用
    ```
   
5. 循环字符串
   ```go
   s, sep := "", ""
   for index, arg := range os.Args[1:] {//range的return值是index和value，但index是从0开始，不受后面的影响
        s += sep + arg
        sep = " "
        fmt.Println(index, arg)
   }
   ```
   
6. 字符串连接。普通字符串相加，太多则耗内存
   ```go
   strings.Join(os.Args[1:], " ")
   ```
   
7. if。if也不需要括号,且和左大括号必须同行
   ```go
   if n > 1{
        // do some thing
   }
   ```      

8. map数据结构。key/value的数据；key支持任意类型，只要能用==进行比较，string是最常用的key。
9. make函数用来创建空map，除创建空map外，make还有其他用处
   ```go
   counts := make(map[string]int) //map中key为str类型，value为int类型
   ```
   
10. 自动加1
   ```go
   counts[input.Text()]++ // counts中没有对应的key时，默认初始化为int的初始值0
   //等价
   line := input.Text()
   counts[line] = counts[line] + 1
   ```

11. bufio package。处理程序的输入和输出，它的Scanner类型，可实现接收输入或将输入打散成行或单词
12. os.Open返回两个值。一个是打开的文件类型，一个是内置的error类型，且error等于nil时，说明文件成功打开
   ```go
    f, err := os.Open(file_path)
    if err != nil {
        // failed open file
        continue
    }
    // do some thing
    f.Close() // 显示关闭文件
   ```

13. map对象被作为参数传递给一个函数时，使用的是引用，指向同一个内存地址。所以在外面改，也会影响原map对象