/*
Copyright © 2023 yanosea <myanoshi0626@gmail.com>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var birthday string

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "You can get the mind number from the birthday (yyyyMMdd).",
	Long: `You can get the mind number from the birthday (yyyyMMdd).

You must input the birthday in the format "yyyyMMdd".`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := time.Parse("20060102", birthday); err != nil {
			return fmt.Errorf(`
  Invalid date format: You must input the birthday in the format "yyyyMMdd".
`)
		}
		var mindNum int
		tmpNum := birthday
		for {
			tmpSlice := strings.Split(tmpNum, "")
			for _, v := range tmpSlice {
				n, _ := strconv.Atoi(v)
				mindNum += int(n)
			}
			if mindNum < 10 {
				break
			} else {
				tmpNum = strconv.Itoa(mindNum)
				mindNum = 0
			}
		}
		fmt.Printf("Your mind number is %s !\n", strconv.Itoa(mindNum))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&birthday, "birthday", "b", "", "Birthday in the format 'yyyyMMdd'")
	if err := getCmd.MarkFlagRequired("birthday"); err != nil {
		panic("panic!")
	}
}
