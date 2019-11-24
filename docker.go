package main

import (
	"time"

	"github.com/spf13/cobra"
)

func docker() *cobra.Command {
	return &cobra.Command{
		Use: "docker",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("Hey there!", prettyTime)
			return nil
		},
	}
}
