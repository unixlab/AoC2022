package cmd

import (
	"fmt"

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
		input := aoeinput.Read("", cmd.Use, example)
		fmt.Printf("day05 part 1 => %s\n", day05.RunPart1(input))
		fmt.Printf("day05 part 2 => %s\n", day05.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day05Cmd)
}
