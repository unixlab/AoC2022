package cmd

import (
	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2022/internal/aoeinput"
	"github.com/unixlab/AoC2022/internal/day04"
)

// day04Cmd represents the day04 command
var day04Cmd = &cobra.Command{
	Use: "day04",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		day04.RunPart1(cmd.Use, aoeinput.Read(cmd.Use, example))
		day04.RunPart2(cmd.Use, aoeinput.Read(cmd.Use, example))
	},
}

func init() {
	rootCmd.AddCommand(day04Cmd)
}
