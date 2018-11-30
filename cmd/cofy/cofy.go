package main

import (
	"os"

	"github.com/mitchellh/go-homedir"
)

const (
	// DefaultPathName is the default config dir name
	DefaultPathName = ".cofy"
	// DefaultPathRoot is the path to the default config dir location.
	DefaultPathRoot = "~/" + DefaultPathName
	// EnvDir is the environment variable used to change the path root.
	EnvDir = "COPY_PATH"
)

func detectRepoRoot() (string, error) {
	p := DefaultPathRoot
	if os.Getenv(EnvDir) != "" {
		p = os.Getenv(EnvDir)
	}
	p, err := homedir.Expand(p)
	if err != nil {
		return "", err
	}
	return p, nil
}

func main() {
	rootCMD.AddCommand(initCMD)
	rootCMD.AddCommand(versionCMD)
	rootCMD.Execute()
}
