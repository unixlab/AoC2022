package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2022/internal/aoeinput"
	"github.com/unixlab/AoC2022/internal/day12"
)

// day12Cmd represents the day12 command
var day12Cmd = &cobra.Command{
	Use: "day12",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aoeinput.Read("", cmd.Use, example)
		fmt.Printf("day12 part 1 => %d\n", day12.Run(input, 1))
		fmt.Printf("day12 part 2 => %d\n", day12.Run(input, 2))
	},
}

func init() {
	rootCmd.AddCommand(day12Cmd)
}
