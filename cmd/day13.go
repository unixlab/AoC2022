package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2022/internal/aoeinput"
	"github.com/unixlab/AoC2022/internal/day13"
)

// day13Cmd represents the day13 command
var day13Cmd = &cobra.Command{
	Use: "day13",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aoeinput.Read("", cmd.Use, example)
		fmt.Printf("day13 part 1 => %d\n", day13.RunPart1(input))
		fmt.Printf("day13 part 2 => %d\n", day13.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day13Cmd)
}
