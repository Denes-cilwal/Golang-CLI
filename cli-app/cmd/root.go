package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "add",
	Short: "add values passed to function",
	Long:  `Demo application to demonstrat`,
	Run: func(cmd *cobra.Command, args []string) {
		sum := 0
		for _, args := range args {
			num, err := strconv.Atoi(args)

			if err != nil {
				fmt.Println(err)
			}
			sum = sum + num
		}
		fmt.Println("result of addition is", sum)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
