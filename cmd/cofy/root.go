package main

import (
	"fmt"
	"github.com/overbool/cofy/common"
	"github.com/overbool/cofy/core/storage/mysql"
	"github.com/overbool/cofy/model"
	"github.com/overbool/cofy/router"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var root string

var logger = log.New(os.Stdout, "", log.LstdFlags)

func init() {
	rootCMD.Flags().StringVar(&root, "root", "./", "Show all version info")
}

var rootCMD = &cobra.Command{
	Use:   "serve [option]",
	Short: "serve - Static file serving and directory listing",
	Run: func(cmd *cobra.Command, args []string) {
		repoRoot, err := detectRepoRoot()
		if err != nil {
			panic(err)
		}

		err = common.InitConfig(filepath.Join(repoRoot, "config.toml"))
		if err != nil {
			panic(err)
		}

		if !viper.GetBool("server.debug") {
			gin.SetMode(gin.ReleaseMode)
		}

		r := gin.Default()
		router.Load(r)

		// db
		db, err := mysql.New()
		if err != nil {
			panic(err)
		}
		defer db.Close()

		db.DB().AutoMigrate(&model.User{})

		port := fmt.Sprintf(":%s", viper.GetString("server.port"))
		log.Printf("Start to listening the incoming requests on http address: %s", port)
		log.Printf(http.ListenAndServe(port, r).Error())
	},
}
