package main

import (
	"os"
	"path"
	"strings"
)

// TryToResolve : resolve error output string , and try to resolve it
func TryToResolve(out string) string {
	//Log
	Log("", "try to resolve problem : "+out)

	//another process are operating git repository, try to remove .git/index/lock and git add . again
	if strings.Contains(out, "Unable to create") &&
		strings.Contains(out, "index.lock") &&
		strings.Contains(out, "File exists") {
		basePath, _ := os.Getwd()
		basePath = strings.ReplaceAll(path.Join(basePath, ".git", "index.lock"), "/", "\\")
		resolveRes := ExecCommand("del/f/s/q " + basePath)
		resolveRes = ExecCommand("git add .")
		return resolveRes
	}

	return ""
}
