package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
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

// PathIsExists : return the path whether exists.
func PathIsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil || os.IsExist(err) {
		return true
	}
	return false
}

// Log : print log, return error if there is error
func Log(basePath string, logContent string) error {

	logPath := path.Join(basePath, "logs",
		strconv.Itoa(time.Now().Year()),
		time.Now().Month().String(),
		strconv.Itoa(time.Now().Day()))
	if !PathIsExists(logPath) {
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	file, err := os.OpenFile(path.Join(logPath, "data.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	defer file.Close()
	if err != nil {
		return err
	}

	logger := log.New(file, "", log.LstdFlags)
	logger.Println(logContent)

	return nil
}
