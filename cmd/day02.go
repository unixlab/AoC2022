package cmd

import (
	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2022/internal/aoeinput"
	"github.com/unixlab/AoC2022/internal/day02"
)

// day02Cmd represents the day02 command
var day02Cmd = &cobra.Command{t
	Use: "day02",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		day02.RunPart1(cmd.Use, aoeinput.Read(cmd.Use, example))
		day02.RunPart2(cmd.Use, aoeinput.Read(cmd.Use, example))
	},
}

func init() {
	rootCmd.AddCommand(day02Cmd)
}
