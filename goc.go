package goc

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime/debug"
	"strings"
)
var versionOutput = ""

func New(name string) *Goc {
	if Version == "DEV" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}
	Version = strings.TrimPrefix(Version, "v")

	rootCmd := &cobra.Command{}

	if BuildDate == "" {
		rootCmd.Version = Version
	} else {
		rootCmd.Version = fmt.Sprintf("%s (%s)", Version, BuildDate)
	}
	versionOutput = fmt.Sprintf("%s version %s\n", name, rootCmd.Version)
	rootCmd.AddCommand(versionCmd)
	rootCmd.SetVersionTemplate(versionOutput)

	g := &Goc{
		rootCmd: rootCmd,
	}
	return xg
}

type Goc struct {
	rootCmd *cobra.Command
}

func (g *Goc) Run(handle func(cmd *cobra.Command, args []string)) {
	if g.rootCmd == nil {
		return
	}
	g.rootCmd.Run = handle
	err := g.rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func (g *Goc) SetRootCmd(rc *cobra.Command) {
	g.rootCmd = rc
}

func (g *Goc) RootCmd() *cobra.Command {
	return g.rootCmd
}


