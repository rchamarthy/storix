package main

import (
	"os"

	"github.com/spf13/cobra"
)

var version string

func main() {
	if err := cli().Execute(); err != nil {
		os.Exit(1)
	}
}

func cli() *cobra.Command {
	root := &cobra.Command{
		Use:     "appdctl",
		Short:   "appdctl is a tool for managing AppDynamics VMs",
		Version: version + "\n",
	}

	root.SetVersionTemplate(version + "\n")
	root.PersistentFlags().StringP("server", "s",
		"/var/run/appd-os.sock",
		"appd-os server socket path",
	)
	root.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	root.PersistentFlags().BoolP("yaml", "y", false, "yaml output")
	root.PersistentFlags().BoolP("json", "j", false, "json output")

	root.AddCommand(show())
	return root
}

func show() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show information about AppDynamics VMs",
	}

	cmd.AddCommand(NewShowDisks())
	cmd.AddCommand(NewShowPartitions())

	return cmd
}
