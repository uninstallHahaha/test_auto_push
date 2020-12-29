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
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	//run command
	err := cmd.Run()
	if err != nil {
		// fmt.Printf("run failed : %v, stderr: %v\n", err, stderr.String())
		res = err.Error()
		Log(gConfig.logdir, "error: "+res)
		return
	}
	res = out.String()
	return
}
