# 章节1

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

   字符串分割
   ```go
   index, text := strings.Split(STR, "\n")
   ```   

   if !strings.HasPrefix(STR, PREFIX)。字符前缀判断

   组合字符串
   ```go
   fmt.Sprintf("%s%d", aStr, aInt)//返回字符串。将不同类型的数据组合成字符串
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
14. ioutil.ReadFile读取文件的内容，返回值data,err
15. ioutil.ReadAll(resp.Body)从http请求体中读取内容, 返回data,err

16. import 与使用。 import image/color,使用时直接color.White
17. http.Get(url)从web中读取内容，返回resp,err
18. 常量和变量声明一般在包级别。这样整个包都可以使用这个变量；常量值可以是数字，字符串，bool值
19. 复合声明。实例化go中复合类型
   ```go
   var palette = []color.Color{color.White, color.Black}//生成一个slice切片
   ```

```go
anim := gif.GIF{LoopCount: nframes} 
//生成一个struct结构体;gif.GIF是一个struct类型，时一组数据集合
// anim的LoopCount字段将设置为nframes，
//其他为默认值;struct内部变量可通过点来访问，如anim.Delay
```

20. go FUNC表示创建一个新的goroutine,并在这个新的goroutine中执行这个函数
21. channel类型，是go中的一个核心类型。类似python中的collections.deque
    * 操作符是<-。
       ```go
      ch <- v // 发送值v到channel ch中
      v := <- ch // 从channel ch中接收数据，并将数据赋值给v
       ```

    * 创建channel
      ```go
      ch := make(chan int, 100) // 可接收和发送类型为int的数据,容量是100
      ch := make(chan <- float64) // 只可以用来发送float64类型的数据到ch
      ch := make(<- chan int) //只可以用来给其他参数接收int类型的数据
      ```

22. time.Since(time.Now()).Seconds()获取时间差。t1-t0
23. 定义锁。var mu sync.Mutex;mu.Lock();mu.Unlock()
24. fmt.Fprintf(w io.writer, A_string) 将string写入w中去。这里w是一个io.Writer实例