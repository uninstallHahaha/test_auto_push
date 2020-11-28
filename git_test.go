package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetBranchs(t *testing.T) {
	res := ExecCommand("git branch")
	fmt.Println("res", res)
}

func TestGitCannotPush(t *testing.T) {
	res := ExecCommand("git push")
	fmt.Printf("%#v\n", res)
}

func TestGitCannotPull(t *testing.T) {
	res := ExecCommand("git pull")
	fmt.Printf("%#v\n", res)
}

func TestAllBranchs(t *testing.T) {
	res := ExecCommand("git branch -a")
	isc := strings.Contains(res, "main")
	// fmt.Printf("%v\n", res)
	fmt.Printf("%v\n", isc)
}

func TestUpdateConfigFile(t *testing.T) {
	UpdateConfigFile("local_branch_name", "vlice")
}

func TestNoBranch(t *testing.T) {
	res := ExecCommand("git branch --set-upstream-to=origin/master master")
	fmt.Println(res)
}
