package test

import (
	"golang-learn/findPingAbleIp/goImpl/utils"
	"testing"
)

func TestShowPingAble(t *testing.T) {
	var ipChan chan string
	var pingAbleChan chan string
	var exitChan chan bool

	ipChan = make(chan string, 256)
	pingAbleChan = make(chan string, 256)
	exitChan = make(chan bool, 256)

	go utils.Write2Chan(ipChan, "192.168.4", 1, 255)
	for i := 0; i < cap(exitChan); i++ {
		go utils.GetPingAble(ipChan, pingAbleChan, exitChan)
	}
	go utils.Wait(exitChan, pingAbleChan)
	utils.ShowPingAble(pingAbleChan)
}

