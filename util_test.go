package main

import (
	"fmt"
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