package cmd

import (
	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2022/internal/aoeinput"
	"github.com/unixlab/AoC2022/internal/dayXX"
)

// dayXXCmd represents the dayXX command
var dayXXCmd = &cobra.Command{
	Use: "dayXX",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		dayXX.RunPart1(cmd.Use, aoeinput.Read(cmd.Use, example))
		dayXX.RunPart2(cmd.Use, aoeinput.Read(cmd.Use, example))
	},
}

func init() {
	rootCmd.AddCommand(dayXXCmd)
}
