package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https") {
			url = "https://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		data, err := ioutil.ReadAll(resp.Body) //从resp中读取内容
		//out := os.Stdout
		//data, err := io.Copy(out, resp.Body)
		//println(resp.Status)
		resp.Body.Close() //关闭请求体body流

		if err != nil {
			fmt.Println(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//
		fmt.Println(resp.Status)
		fmt.Printf("%s", data)
	}
}
