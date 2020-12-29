package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

// func TestTicker(t *testing.T) {
// 	ticker := time.NewTicker(time.Second)
// 	Ticker(ticker)
// 	// go Ticker(ticker)
// 	// time.Sleep(time.Second * 4)
// 	// ticker.Stop()
// }

func TestTickerController(t *testing.T) {
	tc := TickerController{ticker: time.NewTicker(time.Second)}
	go tc.StartTicker()
	time.Sleep(time.Second * 3)
	fmt.Println("此时的ticker是: ", tc.ticker)
	tc.StopTicker()
}

func TestExec(t *testing.T) {
	// cmd := exec.Command("git", "add", ".")
	// output, _ := cmd.CombinedOutput()
	// res := string(output)
	res := ExecCommand("git add .")
	if strings.Contains(res, "Unable to create") &&
		strings.Contains(res, "index.lock") &&
		strings.Contains(res, "File exists") {
		fmt.Println("yeah")
	} else {
		fmt.Println("non")
	}
	// fmt.Println(res)
}

func TestGetwd(t *testing.T) {
	p, _ := os.Getwd()
	fmt.Println(p)
}

func TestDeleteFile(t *testing.T) {
	resolveRes := ExecCommand("del/f/s/q F:\test\new folder\\index.txt")
	fmt.Println(resolveRes)
}
