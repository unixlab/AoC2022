package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2022/internal/aoeinput"
	"github.com/unixlab/AoC2022/internal/day10"
)

// day10Cmd represents the day10 command
var day10Cmd = &cobra.Command{
	Use: "day10",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aoeinput.Read("", cmd.Use, example)
		part1, part2 := day10.Run(input, true)
		fmt.Printf("day10 part 1 => %d\n", part1)
		fmt.Printf("day10 part 2 => \n%s", part2)
	},
}

func init() {
	rootCmd.AddCommand(day10Cmd)
}
