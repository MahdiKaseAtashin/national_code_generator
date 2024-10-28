/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify the following national code",
	Long: `This command takes a national code as an argument and verifies its validity.
The national code should follow the standard format and will be checked
to ensure it meets all necessary criteria for validity.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a national code")
			return
		}
		if nationalCodeValidator(args[0]) {
			fmt.Println("Valid :)")
		} else {
			fmt.Println("Invalid national code!")
		}
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}

func nationalCodeValidator(val string) bool {
	// List of all invalid national codes where all digits are the same
	allDigitEqual := []string{
		"0000000000", "1111111111", "2222222222", "3333333333",
		"4444444444", "5555555555", "6666666666", "7777777777",
		"8888888888", "9999999999",
	}

	// Check if the value is in the list of invalid codes
	for _, code := range allDigitEqual {
		if val == code {
			return false
		}
	}

	// Regular expression to match exactly 10 digits
	codeMelliPattern := regexp.MustCompile(`^[0-9]{10}$`)
	if !codeMelliPattern.MatchString(val) {
		return false
	}

	// Convert the string to a slice of runes (to access digits individually)
	chArray := []rune(val)

	// Calculate the weighted sum of the first 9 digits
	num0 := int(chArray[0]-'0') * 10
	num1 := int(chArray[1]-'0') * 9
	num2 := int(chArray[2]-'0') * 8
	num3 := int(chArray[3]-'0') * 7
	num4 := int(chArray[4]-'0') * 6
	num5 := int(chArray[5]-'0') * 5
	num6 := int(chArray[6]-'0') * 4
	num7 := int(chArray[7]-'0') * 3
	num8 := int(chArray[8]-'0') * 2
	a := int(chArray[9] - '0')

	b := num0 + num1 + num2 + num3 + num4 + num5 + num6 + num7 + num8
	c := b % 11

	// Return true if the validation passes, otherwise false
	return (c < 2 && a == c) || (c >= 2 && (11-c) == a)
}
