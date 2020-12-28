package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

// InitGitPro ...
func InitGitPro(gConfig GitConfig, remoteBranchName *string) {

	//init user info
	fmt.Print("Initing git user info.")
	Log("", "testLog...")
	tc1 := TickerController{ticker: time.NewTicker(time.Second)}
	go tc1.StartTicker()
	tc1.StopTicker()
	if gConfig.gitUsername != "" {
		ExecCommand("git config --global user.name \"" + gConfig.gitUsername + "\"")
		fmt.Println(">>>[git config username]")
	}
	if gConfig.gitPassword != "" {
		ExecCommand("git config --global user.password \"" + gConfig.gitPassword + "\"")
		fmt.Println(">>>[git config password]")
	}
	if gConfig.gitUserEmail != "" {
		ExecCommand("git config --global user.email \"" + gConfig.gitUserEmail + "\"")
		fmt.Println(">>>[git config email]")
	}

	//init git project
	fmt.Print("Initing git project.")
	tc := TickerController{ticker: time.NewTicker(time.Second)}
	go tc.StartTicker()
	res := ExecCommand("git config --global http.postBuffer 524288000")
	res = ExecCommand("git status")
	tc.StopTicker()

	//init git ignore file
	if !PathIsExists(".gitignore") {
		res = ExecCommand("type nul>.gitignore")
		ignoreStr := "APP_saveMan.exe\ngit_config.properties\nlogs"
		err := ioutil.WriteFile(".gitignore", []byte(ignoreStr), 0x666)
		if err != nil {
			fmt.Printf("generate ignore file failed : %v\n", err)
			return
		}
	}

	//did not init git project
	if strings.TrimSpace(res) == "exit status 128" {
		//init git project
		fmt.Println(">>>[git init]")
		res = ExecCommand("git init")
		// add remote address
		fmt.Println(">>>[git remote add origin " + gConfig.remoteStoreAddress + "]")
		res = ExecCommand("git remote add origin " + gConfig.remoteStoreAddress)
		// add work file to index
		fmt.Println(">>>[git add .]")
		res = ExecCommand("git add .")
		// add index to local repository
		fmt.Println(">>>[git commit -m " + gConfig.commitPrefix + time.Now().Format("2006_01_02#15:04:05") + "]")
		res = ExecCommand("git commit -m " + gConfig.commitPrefix + time.Now().Format("2006_01_02#15:04:05"))
		// add local branch
		fmt.Println(">>>[git branch " + gConfig.localBranchName + "]")
		res = ExecCommand("git branch " + gConfig.localBranchName)
		// switch to local branch
		fmt.Println(">>>[git checkout " + gConfig.localBranchName + "]")
		res = ExecCommand("git checkout " + gConfig.localBranchName)
		// get remote branchs
		fmt.Println(">>>[git fetch]")
		res = ExecCommand("git fetch")
		// update binding remote branch
		if gConfig.forceRecover == "0" {
			var isc bool
			for {
				res := ExecCommand("git branch -a")
				if isc = strings.Contains(res, "remotes/origin/"+*remoteBranchName); isc {
					*remoteBranchName = *remoteBranchName + "1"
				} else {
					//update config file
					UpdateConfigFile("remote_branch_name", *remoteBranchName)
					break
				}
			}
		}
		// push local repository to remote
		fmt.Println(">>>[git push -u origin " + gConfig.localBranchName + ":" + *remoteBranchName + " --force]")
		res = ExecCommand("git push -u origin " + gConfig.localBranchName + ":" + *remoteBranchName + " --force")
		// bind local branch with remote branch
		fmt.Println(">>>[git branch --set-upstream-to=origin/" + *remoteBranchName + " " + gConfig.localBranchName + "]")
		res = ExecCommand("git branch --set-upstream-to=origin/" + *remoteBranchName + " " + gConfig.localBranchName)

		fmt.Println("Initing git project is done, now your local files is connecting with your git store")
	} else {
		//For the case: user does not start saveMan, and he change his file, when he starts saveMan, we should git push to sync files
		fmt.Print("sync with cloud data")
		tc := TickerController{ticker: time.NewTicker(time.Second)}
		go tc.StartTicker()
		res = ExecCommand("git branch --set-upstream-to=origin/" + *remoteBranchName + " " + gConfig.localBranchName)
		if res == "exit status 128" {
			ExecCommand("git branch " + gConfig.localBranchName)
			ExecCommand("git checkout " + gConfig.localBranchName)
		}
		res = ExecCommand("git add .")
		res = ExecCommand("git commit -m " + gConfig.commitPrefix + time.Now().Format("2006_01_02#15:04:05"))
		res = ExecCommand("git push -u origin " + gConfig.localBranchName + ":" + gConfig.remoteBranchName + " --force")
		ExecCommand("git branch --set-upstream-to=origin/" + *remoteBranchName + " " + gConfig.localBranchName)
		tc.StopTicker()
	}
	fmt.Println("Local files have connected with your git store, just edit it!")
}
