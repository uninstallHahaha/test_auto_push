package main

import (
	"fmt"
	"testing"
)

func TestGetBranchs(t *testing.T) {
	res := ExecCommand("git branch")
	fmt.Println("res", res)
}
