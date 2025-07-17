package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/spf13/cobra"
)

var (
	length    int
	uppercase bool
	numbers   bool
	symbols   bool
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a secure random password",
	Long:  `Generate a secure random password with customizable options like length, uppercase letters, numbers, and symbols.`,
	Run: func(cmd *cobra.Command, args []string) {
		password := generatePassword(length, uppercase, numbers, symbols)
		fmt.Println("Generated password:", password)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntVarP(&length, "length", "l", 12, "Length of the password")
	generateCmd.Flags().BoolVarP(&uppercase, "uppercase", "u", false, "Include uppercase letters")
	generateCmd.Flags().BoolVarP(&numbers, "numbers", "n", false, "Include numbers")
	generateCmd.Flags().BoolVarP(&symbols, "symbols", "s", false, "Include symbols")
}

func generatePassword(length int, useUpper, useNumbers, useSymbols bool) string {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	num := "0123456789"
	sym := "!@#$%^&*()-_=+[]{}<>?/|"

	charset := lower
	if useUpper {
		charset += upper
	}
	if useNumbers {
		charset += num
	}
	if useSymbols {
		charset += sym
	}

	if length < 1 || len(charset) == 0 {
		return ""
	}

	var sb strings.Builder
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		sb.WriteByte(charset[index.Int64()])
	}

	return sb.String()
}
