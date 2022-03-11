package main

import (
	"fmt"
	"log"
)
import "github.com/emersion/go-imap/client"

func main() {
	log.Println("connecting to server...")
	c, err := client.DialTLS("smtp.exmail.qq.com:465", "xx")
	fmt.Println(c, err)
	fmt.Println("xxxx")
}
