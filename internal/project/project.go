package project

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// CmdNew represents the new command.
var CmdNew = &cobra.Command{
	Use:   "new",
	Short: "Create a gin project template",
	Long:  "Create a gin project using the repository template. Example: gin-project-gen new helloworld",
	Run:   run,
}

var (
	repoURL string
	branch  string
	timeout string
)

func init() {
	if repoURL = os.Getenv("GIN_LAYOUT_REPO"); repoURL == "" {
		repoURL = "https://github.com/trumanwong/gin-layout.git"
	}
	timeout = "60s"
	CmdNew.Flags().StringVarP(&repoURL, "repo-url", "r", repoURL, "layout repo")
	CmdNew.Flags().StringVarP(&branch, "branch", "b", branch, "repo branch")
	CmdNew.Flags().StringVarP(&timeout, "timeout", "t", timeout, "time out")
}

func run(_ *cobra.Command, args []string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	t, err := time.ParseDuration(timeout)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	name := ""
	if len(args) == 0 {
		prompt := &survey.Input{
			Message: "What is project name ?",
			Help:    "Created project name.",
		}
		err = survey.AskOne(prompt, &name)
		if err != nil || name == "" {
			return
		}
	} else {
		name = args[0]
	}
	projectName, workingDir := processProjectParams(name, wd)
	p := &Project{Name: projectName}
	done := make(chan error, 1)
	go func() {
		done <- p.New(ctx, workingDir, repoURL, branch)
	}()
	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			fmt.Fprint(os.Stderr, "\033[31mERROR: project creation timed out\033[m\n")
			return
		}
		fmt.Fprintf(os.Stderr, "\033[31mERROR: failed to create project(%s)\033[m\n", ctx.Err().Error())
	case err = <-done:
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[31mERROR: Failed to create project(%s)\033[m\n", err.Error())
		}
	}
}

func processProjectParams(projectName string, workingDir string) (projectNameResult, workingDirResult string) {
	_projectDir := projectName
	_workingDir := workingDir
	// Process ProjectName with system variable
	if strings.HasPrefix(projectName, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			// cannot get user home return fallback place dir
			return _projectDir, _workingDir
		}
		_projectDir = filepath.Join(homeDir, projectName[2:])
	}

	// check path is relative
	if !filepath.IsAbs(projectName) {
		absPath, err := filepath.Abs(projectName)
		if err != nil {
			return _projectDir, _workingDir
		}
		_projectDir = absPath
	}

	return filepath.Base(_projectDir), filepath.Dir(_projectDir)
}
