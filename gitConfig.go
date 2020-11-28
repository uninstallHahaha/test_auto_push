package main

// GitConfig ...
type GitConfig struct {
	localBranchName    string
	remoteBranchName   string
	
	commitPrefix       string
	remoteStoreAddress string
	saveDuration       int
	gitUsername        string
	gitPassword        string
	gitUserEmail       string

	forceRecover string
}
