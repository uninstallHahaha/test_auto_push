package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
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

// UpdateConfigFile : provide new value for config field, update the git_config file
// field : the field would like to be updated
// value : the new value of field
func UpdateConfigFile(field string, value string) {
	var tmpStr string
	fs, err := os.Open("./git_config.properties")
	if err != nil {
		fmt.Printf("open config file failed : %v\n", err)
		return
	}
	defer fs.Close()

	reader := bufio.NewReader(fs)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				if strings.HasPrefix(strings.TrimSpace(line), field) {
					tmpStr += field + " = " + value + "\n"
				} else {
					tmpStr += line
				}
			}
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		if strings.HasPrefix(strings.TrimSpace(line), field) {
			tmpStr += field + " = " + value + "\n"
			continue
		}
		tmpStr += line
	}
	err = ioutil.WriteFile("./git_config.properties", []byte(tmpStr), 666)
	if err != nil {
		fmt.Printf("rewrite config file failed : %v\n", err)
		return
	}
}
