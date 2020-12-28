package main

// GitConfig ...
type GitConfig struct {
	localBranchName  string
	remoteBranchName string

	remoteStoreAddress string

	commitPrefix string
	saveDuration int

	gitUsername  string
	gitPassword  string
	gitUserEmail string

	forceRecover string

	logdir string
}
