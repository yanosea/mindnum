/*
Copyright © 2023 yanosea <myanoshi0626@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

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
		var filePath = "resources/descriptions/"
		switch mindNum {
		case 1:
			filePath += "1.txt"
		case 2:
			filePath += "2.txt"
		case 3:
			filePath += "3.txt"
		case 4:
			filePath += "4.txt"
		case 5:
			filePath += "5.txt"
		case 6:
			filePath += "6.txt"
		case 7:
			filePath += "7.txt"
		case 8:
			filePath += "8.txt"
		case 9:
			filePath += "9.txt"
		default:
			panic(nil)
		}
		f, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().IntVarP(&mindNum, "mindNum", "n", 0, "Mind number you got from 'mindnum get'")
	if err := checkCmd.MarkFlagRequired("mindNum"); err != nil {
		panic("panic!")
	}
}
