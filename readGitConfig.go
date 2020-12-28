package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// ReadGitConfig read git config
func ReadGitConfig(gConfig *GitConfig) bool {
	//tmp map for read git config
	var configMapTmp = make(map[string]string)
	bs, err := ioutil.ReadFile("./git_config.properties")
	if err != nil {
		fmt.Printf("未找到git_config.properties配置文件, 请将该文件放置到同级目录下, 然后尝试重新启动\n")
		return false
	}
	lines := strings.Split(string(bs[:]), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		kv := strings.Split(line, "=")
		configMapTmp[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
	}
	gConfig.commitPrefix = configMapTmp["commit_prefix"]
	gConfig.remoteStoreAddress = configMapTmp["remote_store_address"]
	duration, err := strconv.Atoi(configMapTmp["save_duration"])
	if err != nil {
		fmt.Printf("配置文件中save_duration应当为数字类型: %v\n", err)
		return false
	}
	gConfig.saveDuration = duration
	gConfig.gitUsername = configMapTmp["git_username"]
	gConfig.gitPassword = configMapTmp["git_password"]
	gConfig.gitUserEmail = configMapTmp["git_useremail"]

	gConfig.logdir = configMapTmp["logdir"] 

	gConfig.forceRecover = configMapTmp["force_recover"]

	if configMapTmp["local_branch_name"] == "" {
		gConfig.localBranchName = "master"
		if configMapTmp["remote_branch_name"] != "" {
			gConfig.remoteBranchName = configMapTmp["remote_branch_name"]
		} else {
			gConfig.remoteBranchName = "master"
		}
	} else {
		gConfig.localBranchName = configMapTmp["local_branch_name"]
		if configMapTmp["remote_branch_name"] != "" {
			gConfig.remoteBranchName = configMapTmp["remote_branch_name"]
		} else {
			gConfig.remoteBranchName = configMapTmp["local_branch_name"]
		}
	}

	if gConfig.remoteStoreAddress == "" {
		fmt.Println("请在配置文件中设置 remote_store_address 仓库地址")
		return false
	}

	return true
}
