package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isValidCreditCardNumber(cardNumber string) bool {
	// Regular expression pattern to validate credit card number
	pattern := `^(4|5|6)\d{3}(-?\d{4}){3}$`

	// Check if the card number matches the pattern
	match, _ := regexp.MatchString(pattern, cardNumber)
	if !match {
		return false
	}

	// Remove hyphens from the card number
	cardNumber = strings.ReplaceAll(cardNumber, "-", "")

	// Check for consecu<ve repeated digits
	for i := 0; i < len(cardNumber)-3; i++ {
		if cardNumber[i] == cardNumber[i+1] && cardNumber[i] == cardNumber[i+2] && cardNumber[i] == cardNumber[i+3] {
			return false
		}
	}
	return true
}

func main() {
	// Read the number of credit card numbers
	var n int
	fmt.Scanln(&n)

	// Read each credit card number and validate
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		cardNumber := scanner.Text()
		if isValidCreditCardNumber(cardNumber) {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
	}
}
