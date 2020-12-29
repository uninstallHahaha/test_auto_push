package main

import (
	"bytes"
	"os/exec"
	"strings"
)

//ExecCommand ...
func ExecCommand(command string) (res string) {
	//to log
	Log(gConfig.logdir, ">>>["+command+"]")
	// fmt.Println(">>>["+command+"]")

	//spilt command
	cList := strings.Split(command, " ")
	cmd := exec.Command(cList[0], cList[1:]...)
	//set out buffer and err buffer
	var out bytes.Buffer
	// var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	//run command
	err := cmd.Run()
	if err != nil {
		Log(gConfig.logdir, "cmd run error: "+res)
		TryToResolve(out.String())
	}
	res = out.String()
	return
}
