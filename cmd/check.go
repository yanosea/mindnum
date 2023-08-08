/*
Copyright © 2023 yanosea <myanoshi0626@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var mindNum int

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "You can check the character from the mind number.",
	Long: `You can check the character from the mind number.

You must input the mind number in the range of 1 to 9.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if mindNum < 1 || 9 < mindNum {
			return fmt.Errorf(`
  Invalid arg: You must input the mind number in the range of 1 to 9.
`)
		}
		var descriptionPath = resourcePath + "descriptions/"
		switch mindNum {
		case 1:
			descriptionPath += "1.txt"
		case 2:
			descriptionPath += "2.txt"
		case 3:
			descriptionPath += "3.txt"
		case 4:
			descriptionPath += "4.txt"
		case 5:
			descriptionPath += "5.txt"
		case 6:
			descriptionPath += "6.txt"
		case 7:
			descriptionPath += "7.txt"
		case 8:
			descriptionPath += "8.txt"
		case 9:
			descriptionPath += "9.txt"
		default:
			panic(nil)
		}

		if description, err := resources.ReadFile(descriptionPath); err == nil {
			fmt.Print(string(description))
			return nil
		} else {
			panic(nil)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().IntVarP(&mindNum, "mindNum", "n", 0, "Mind number you got from 'mindnum get'")
	if err := checkCmd.MarkFlagRequired("mindNum"); err != nil {
		panic("panic!")
	}
}
