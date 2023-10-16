package main

import "github.com/spf13/cobra"

func OverviewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "overview",
		Short: "Prints and overview of cluster pools and clusters",
		Long:  `Prints and overview of cluster pools and clusters.`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Usage()
		},
	}

	return cmd
}
