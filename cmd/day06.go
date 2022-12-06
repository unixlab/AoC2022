package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2022/internal/aoeinput"
	"github.com/unixlab/AoC2022/internal/day06"
)

// day06Cmd represents the day06 command
var day06Cmd = &cobra.Command{
	Use: "day06",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		for _, line := range aoeinput.Read("", cmd.Use, example) {
			fmt.Printf("day06 part 1 => %d\n", day06.Run(line, 4))
			fmt.Printf("day06 part 2 => %d\n", day06.Run(line, 14))
		}
	},
}

func init() {
	rootCmd.AddCommand(day06Cmd)
}
