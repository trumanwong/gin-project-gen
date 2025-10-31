package upgrade

import (
	"fmt"

	"github.com/trumanwong/gin-project-gen/internal/base"

	"github.com/spf13/cobra"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the gin-project-gen tools",
	Long:  "Upgrade the gin-project-gen tools. Example: gin-project-gen upgrade",
	Run:   Run,
}

// Run upgrade the gin-project-gen tools.
func Run(_ *cobra.Command, _ []string) {
	err := base.GoInstall(
		"github.com/trumanwong/gin-project-gen@latest",
		"github.com/trumanwong/protoc-gen-go-gin@latest",
		"github.com/trumanwong/protoc-gen-go-gin-errors@latest",
		"github.com/trumanwong/protobuf-go/cmd/protoc-gen-go@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
		"github.com/google/gnostic/cmd/protoc-gen-openapi@latest",
	)
	if err != nil {
		fmt.Println(err)
	}
}
