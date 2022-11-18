package cmd

import (
	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2022/internal/aoeinput"
	"github.com/unixlab/AoC2022/internal/day01"
)

// day01Cmd represents the day01 command
var day01Cmd = &cobra.Command{
	Use: "day01",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		day01.RunPart1(aoeinput.Read(cmd.Use, example))
		day01.RunPart2(aoeinput.Read(cmd.Use, example))
	},
}

func init() {
	rootCmd.AddCommand(day01Cmd)
}
