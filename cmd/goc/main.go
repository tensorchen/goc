package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tensorchen/goc"
)

func main() {
	goc.New("goc").Run(func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	})
}
