package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2022/internal/aoeinput"
	"github.com/unixlab/AoC2022/internal/day09"
)

// day09Cmd represents the day09 command
var day09Cmd = &cobra.Command{
	Use: "day09",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aoeinput.Read("", cmd.Use, example)
		fmt.Printf("day09 part 1 => %d\n", day09.Run(input, 2))
		fmt.Printf("day09 part 2 => %d\n", day09.Run(input, 10))
	},
}

func init() {
	rootCmd.AddCommand(day09Cmd)
}
