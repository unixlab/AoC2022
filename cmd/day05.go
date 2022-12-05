package cmd

import (
	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2022/internal/aoeinput"
	"github.com/unixlab/AoC2022/internal/day05"
)

// day05Cmd represents the day05 command
var day05Cmd = &cobra.Command{
	Use: "day05",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		day05.RunPart1(cmd.Use, aoeinput.Read(cmd.Use, example))
		day05.RunPart2(cmd.Use, aoeinput.Read(cmd.Use, example))
	},
}

func init() {
	rootCmd.AddCommand(day05Cmd)
}
