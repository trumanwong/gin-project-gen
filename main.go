package main

import (
	"github.com/trumanwong/gin-project-gen/internal/project"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gin-project-gen",
	Short:   "gin-project-gen: An elegant toolkit for Gin project.",
	Long:    `gin-project-gen: An elegant toolkit for Gin project.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
