package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var root string
var port uint32

var logger = log.New(os.Stdout, "", log.LstdFlags)

func init() {
	rootCMD.Flags().StringVar(&root, "root", "./", "Show all version info")
	rootCMD.Flags().Uint32Var(&port, "port", 8000, "Listening port")
}

var rootCMD = &cobra.Command{
	Use:   "serve [option]",
	Short: "serve - Static file serving and directory listing",
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
		log.Printf(http.ListenAndServe(":8080", r).Error())
	},
}
