/*
Copyright © 2024 NAME HERE MAHDIKASEATASHIN@GMAIL.COM
*/
package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a national code",
	Long: `Generates a national code in the format of 10 digits.

The flag -r or --round will generate a rounded national code which is
the remainder of the division of the sum of the digits by 11. The
result is a single digit from 0 to 9. The command does not check if
the generated code is valid or not.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get the flag value, its default value is false
		rStatus, _ := cmd.Flags().GetBool("round")
		if rStatus { // if status is true, call addFloat
			fmt.Println(generateRoundedNationalCode())
		} else {
			fmt.Println(generateNationalCode())
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().BoolP("round", "r", false, "Generate rounded national code")
}

func generateNationalCode() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	list := make([]int, 0)
	sum := 0

	for i := 10; i > 1; i-- {
		j := r.Intn(10)
		list = append(list, j)
		sum += j * i
	}

	s := sum % 11
	if s < 2 {
		list = append(list, s)
	} else {
		list = append(list, 11-s)
	}

	result := ""
	for _, num := range list {
		result += fmt.Sprintf("%d", num)
	}

	return result
}

func generateRoundedNationalCode() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	list := make([]int, 0)
	sum := 0
	j := 10

	for i := 10; i > 1; i-- {
		// Generate a random number between 0 and j (j cannot be less than 2)
		if j < 2 {
			j = 2
		}
		j = r.Intn(j)
		list = append(list, j)
		sum += j * i
	}

	s := sum % 11
	if s < 2 {
		list = append(list, s)
	} else {
		list = append(list, 11-s)
	}

	// Check if all elements in the list are the same as the first element
	allSame := true
	for _, value := range list {
		if value != list[0] {
			allSame = false
			break
		}
	}

	// If all elements are the same, call the function recursively
	if allSame {
		return generateRoundedNationalCode()
	}

	result := ""
	for _, num := range list {
		result += fmt.Sprintf("%d", num)
	}

	return result
}
