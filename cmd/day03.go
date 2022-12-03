package cmd

import (
	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2022/internal/aoeinput"
	"github.com/unixlab/AoC2022/internal/day03"
)

// day03Cmd represents the day03 command
var day03Cmd = &cobra.Command{
	Use: "day03",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		day03.RunPart1(cmd.Use, aoeinput.Read(cmd.Use, example))
		day03.RunPart2(cmd.Use, aoeinput.Read(cmd.Use, example))
	},
}

func init() {
	rootCmd.AddCommand(day03Cmd)
}
