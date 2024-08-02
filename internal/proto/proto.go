package proto

import (
	"github.com/spf13/cobra"
	"github.com/trumanwong/gin-project-gen/internal/proto/add"
	"github.com/trumanwong/gin-project-gen/internal/proto/client"
	"github.com/trumanwong/gin-project-gen/internal/proto/server"
	"github.com/trumanwong/gin-project-gen/internal/upgrade"
)

// CmdProto represents the proto command.
var CmdProto = &cobra.Command{
	Use:   "proto",
	Short: "Generate the proto files",
	Long:  "Generate the proto files.",
}

func init() {
	CmdProto.AddCommand(add.CmdAdd)
	CmdProto.AddCommand(client.CmdClient)
	CmdProto.AddCommand(server.CmdServer)
	CmdProto.AddCommand(upgrade.CmdUpgrade)
}