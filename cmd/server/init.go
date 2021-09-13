package server

import (
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs the Server",
	Long:  `Runs the Server`,
}
