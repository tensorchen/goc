package goc

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Version is dynamically set by the toolchain.
var Version = "DEV"

// BuildDate is dynamically set at build time by the toolchain.
var BuildDate = "" // YYYY-MM-DD


var versionCmd = &cobra.Command{
	Use:    "version",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(versionOutput)
	},
}

func VersionCmd() *cobra.Command {
	return versionCmd
}

func SetVersionCmd(vc *cobra.Command) {
	versionCmd = vc
}

