package utils

import (
	"fmt"
	"os/exec"
	"strconv"
	"syscall"
)

func Write2Chan(ipChan chan string, prefix string, start, end int) {
	for i := start; i <= end; i++ {
		ip := prefix + "." + strconv.Itoa(i)
		ipChan <- ip
	}
	close(ipChan)
}

func GetPingAble(ipChan chan string,
	pingAbleChan chan string, exitChan chan bool) {
	for {
		res, ok := <-ipChan
		if !ok {
			break
		}
		if PingAble(res) {
			pingAbleChan <- res
		}
	}
	exitChan <- true
}

func PingAble(ip string) bool {
	cmd := exec.Command("ping", "-c", "1", ip, "-w", "1")
	if err := cmd.Start(); err != nil {
		return false
	}

	if err := cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				//fmt.Println(ip, " exit with code : ", status.ExitStatus())
				if status.ExitStatus() == 0 {
					return true
				}else {
					return false
				}
			}
		}
	}
	return true
}

func Wait(exitChan chan bool, pingAbleChan chan string) {
	for i := 0; i < cap(exitChan); i++ {
		<-exitChan
	}
	close(pingAbleChan)
}

func ShowPingAble(pingAbleChan chan string) {
	for {
		if res, ok := <-pingAbleChan; !ok {
			break
		} else {
			fmt.Printf("%v\t ", res)
		}
	}
}
