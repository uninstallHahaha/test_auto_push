package main

import (
	"fmt"
	"os"
	"time"
)

//IsFile : file return true
func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

//Ticker : start dot ticker
func Ticker(ticker *time.Ticker) {
	n := 0
	for {
		if n == 30 {
			ticker.Stop()
			ticker.Stop()
			break
		}
		<-ticker.C
		fmt.Print(".")
		n++
	}
}

//TickerController : generate new ticker,  controll ticker start and stop it.
//	before dot:
//tc := TickerController{ticker: time.NewTicker(time.Second)}
//go tc.StartTicker()
//	after dot:
//tc.StopTicker()
type TickerController struct {
	ticker *time.Ticker
}

//StartTicker : start ticker
func (t TickerController) StartTicker() {
	Ticker(t.ticker)
}

//StopTicker : stop the ticker of TickerController
func (t TickerController) StopTicker() {
	fmt.Println("Done")
	t.ticker.Stop()
}
