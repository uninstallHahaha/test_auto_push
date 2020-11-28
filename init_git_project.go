package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

// InitGitPro ...
func InitGitPro(gConfig GitConfig) {

	//init user info
	fmt.Print("Initing git user info.")
	tc1 := TickerController{ticker: time.NewTicker(time.Second)}
	go tc1.StartTicker()
	tc1.StopTicker()
	if gConfig.gitUsername != "" {
		ExecCommand("git config --global user.name \"" + gConfig.gitUsername + "\"")
		fmt.Printf(">>>[git config username]\n")
	}
	if gConfig.gitPassword != "" {
		ExecCommand("git config --global user.password \"" + gConfig.gitPassword + "\"")
		fmt.Printf(">>>[git config password]\n")
	}
	if gConfig.gitUserEmail != "" {
		ExecCommand("git config --global user.email \"" + gConfig.gitUserEmail + "\"")
		fmt.Printf(">>>[git config email]\n")
	}

	//init git project
	fmt.Print("Initing git project.")
	tc := TickerController{ticker: time.NewTicker(time.Second)}
	go tc.StartTicker()
	res := ExecCommand("git config --global http.postBuffer 524288000")
	res = ExecCommand("git status")
	tc.StopTicker()

	//did not init git project
	if strings.TrimSpace(res) == "exit status 128" {
		//init git ignore file
		res = ExecCommand("type nul>.gitignore")
		ignoreStr := "APP_saveMan.exe\ngit_config.properties"
		err := ioutil.WriteFile(".gitignore", []byte(ignoreStr), 0x666)
		if err != nil {
			fmt.Printf("generate ignore file failed : %v\n", err)
			return
		}
		//init git project
		fmt.Println(">>>[git init]")
		res = ExecCommand("git init")
		// fmt.Println(res)
		fmt.Println(">>>[git remote add origin " + gConfig.remoteStoreAddress + "]")
		res = ExecCommand("git remote add origin " + gConfig.remoteStoreAddress)
		// fmt.Println(res)
		fmt.Println(">>>[git branch " + gConfig.localBranchName + "]")
		res = ExecCommand("git branch " + gConfig.localBranchName)
		// fmt.Println(res)
		fmt.Println(">>>[git checkout " + gConfig.localBranchName + "]")
		res = ExecCommand("git checkout " + gConfig.localBranchName)
		// fmt.Println(res)
		fmt.Println(">>>[git add .]")
		res = ExecCommand("git add .")
		// fmt.Println(res)
		fmt.Println(">>>[git commit -m " + gConfig.commitPrefix + time.Now().Format("2006_01_02#15:04:05") + "]")
		res = ExecCommand("git commit -m " + gConfig.commitPrefix + time.Now().Format("2006_01_02#15:04:05"))
		// fmt.Println(res)
		fmt.Println(">>>[git push -u origin " + gConfig.localBranchName + "]")
		res = ExecCommand("git push -u origin " + gConfig.localBranchName)
		// fmt.Println(res)
		fmt.Println("Initing git project is done, now your local files is connecting with your git store")
	} else {
		//For the case: user does not start saveMan, and he change his file, when he starts saveMan, we should git push to sync files
		fmt.Print("sync with cloud data")
		tc := TickerController{ticker: time.NewTicker(time.Second)}
		go tc.StartTicker()
		res = ExecCommand("git add .")
		res = ExecCommand("git commit -m " + gConfig.commitPrefix + time.Now().Format("2006_01_02#15:04:05"))
		// res = ExecCommand("git pull")
		res = ExecCommand("git push -u origin " + gConfig.localBranchName + " --force")
		tc.StopTicker()
	}
	fmt.Println("Local files have connected with your git store, just edit it!")
}
