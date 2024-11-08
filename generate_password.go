package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

func main() {
	// Define flags
	length := flag.Int("length", 16, "Length of the password")
	includeUppercase := flag.Bool("uppercase", true, "Include uppercase letters")
	includeLowercase := flag.Bool("lowercase", true, "Include lowercase letters")
	includeDigits := flag.Bool("digits", true, "Include digits")
	includeSpecialChars := flag.Bool("specialchars", false, "Include special characters")

	// Parse the flags
	flag.Parse()

	// Define character sets
	var charset string
	if *includeUppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if *includeLowercase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if *includeDigits {
		charset += "0123456789"
	}
	if *includeSpecialChars {
		charset += "!@#$%^&*()_+-=[]{}|;:,.<>?`~"
	}

	// Check if any character set is included
	if len(charset) == 0 {
		fmt.Println("Error: At least one character set must be specified.")
		return
	}

	// Generate the password
	password := generatePassword(*length, charset)
	fmt.Println("Generated Password:", password)
}

func generatePassword(length int, charset string) string {
	n := big.NewInt(int64(len(charset)))
	var bytes = make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, n)
		if err != nil {
			panic(err)
		}
		bytes[i] = charset[index.Int64()]
	}
	return string(bytes)
}
