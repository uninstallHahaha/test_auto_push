package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify"
	_ "github.com/uninstallHahaha/deamon"
)

// go build -ldflags “-H=windowsgui”

//git config
var gConfig = GitConfig{}

func main() {

	//auto init git project
	if !ReadGitConfig(&gConfig) {
		return
	}

	//init git project
	InitGitPro(gConfig, &gConfig.remoteBranchName)

	//start file change monitor
	startMonitor()

}

//watcher
var watcher *fsnotify.Watcher

//add file path to watcher
func watchDir(path string, fi os.FileInfo, err error) error {
	if strings.HasPrefix(path, ".git") ||
		strings.HasPrefix(path, "APP_saveMan") ||
		strings.HasPrefix(path, "git_config") ||
		strings.HasPrefix(path, "logs") ||
		strings.HasPrefix(path, "node_modules") {

		return nil
	}
	// fmt.Println(path)
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}
	return nil
}

//monitor file change event
func startMonitor() {
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	//walk all paths
	if err := filepath.Walk("./", watchDir); err != nil {
		fmt.Println("walk folder path error", err)
	}

	done := make(chan bool)

	go func() {
		//task key
		var taskKey int
		for {
			select {
			case event := <-watcher.Events:
				//exclude .git
				if event.Name == ".git" {
					continue
				}
				//file change, add and commit
				fmt.Printf("Probe file change-> Path: %#v, code: %#v\n", event.Name, event.Op)
				//when add a new folder, add to watcher
				if event.Op == 0x1 {
					if !IsFile(event.Name) {
						watcher.Add(event.Name)
					}
				}
				var timer = time.NewTimer(time.Second * time.Duration(gConfig.saveDuration))
				rand.Seed(time.Now().UnixNano())
				taskKey = rand.Int() //本次任务的key
				go func(key int) {
					<-timer.C
					if key == taskKey {
						ExecCommand("git checkout " + gConfig.localBranchName)
						fmt.Println(">>>[git add .]")
						ExecCommand("git add .")
						fmt.Println(">>>[git commit -m " + gConfig.commitPrefix + time.Now().Format("2006_01_02#15:04:05") + "]")
						ExecCommand("git commit -m " + gConfig.commitPrefix + time.Now().Format("2006_01_02#15:04:05"))
						fmt.Println(">>>[git push]")
						ExecCommand("git push --force")
						taskKey = 0
						fmt.Println(">>>[push success]")
						fmt.Println("Waiting for your next edit...")
					}
				}(taskKey)
			case err := <-watcher.Errors:
				fmt.Println("monitor event error", err)
			}
		}
	}()

	<-done
}
