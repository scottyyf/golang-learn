package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)


func fetch(url string, ch chan<-string){ // ch只用来存string，
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil{
		ch <- fmt.Sprint()
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // ioutil.Discard是垃圾桶，可在里面写一些不需要的数据
	resp.Body.Close()
	if err != nil{
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds() // 获取时间差
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url) // 组合字符串并发送给channel
}


func main()  {
	start := time.Now()
	ch := make(chan string) // 创建双向str通道
	for _, url := range os.Args[1:]{
		go fetch(url, ch) //执行一个goroutine（异步），并在里面执行fetch操作
	}
	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Println("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Println(time.Now())
}
