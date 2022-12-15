package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2022/internal/aoeinput"
	"github.com/unixlab/AoC2022/internal/day15"
)

// day15Cmd represents the day15 command
var day15Cmd = &cobra.Command{
	Use: "day15",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aoeinput.Read("", cmd.Use, example)
		fmt.Printf("day15 part 1 => %d\n", day15.RunPart1(input, example))
		fmt.Printf("day15 part 2 => %d\n", day15.RunPart2(input, example))
	},
}

func init() {
	rootCmd.AddCommand(day15Cmd)
}
