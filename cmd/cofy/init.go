package main

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
)

var initCMD = &cobra.Command{
	Use:   "init",
	Short: "Initial project",
	Run: func(cmd *cobra.Command, args []string) {
		p, err := detectRepoRoot()
		if err != nil {
			panic(err)
		}

		// check whether the repo root is existent
		if _, err := os.Stat(p); !os.IsNotExist(err) {
			cmd.Println("cofy configuration file already exists")
			return
		}

		cmd.Printf("initializing cofy at %s", p)
		if err := initConfig(p); err != nil {
			panic(err)
		}
	},
}

func initConfig(repoRoot string) error {
	box := packr.New("cofy", "../../config")
	return box.Walk(func(s string, file packd.File) error {
		p := filepath.Join(repoRoot, s)
		dir := path.Dir(p)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}
		}
		return ioutil.WriteFile(p, []byte(file.String()), 0644)
	})
}
