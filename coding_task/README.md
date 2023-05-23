# Credit Card Validator
This repository contains a simple credit card validator program written in Go. The program checks whether a given credit card number is valid or not based on specific criteria.

## Functionality

The program includes the following functions:

## `isValidCreditCardNumber(cardNumber string) bool`

- The **`isValidCreditCardNumber`** function takes a credit card number as its inpurt and returns a bool indicating wthether the card number is valid or not. Then regular expression pattern **`"^(4|5|6)\d{3}(-?\d{4}){3}$"`** is used to validate the credit card number. The **`regexp.MatchString`** funcion is used to check if the card number matches the pattern. If this is no match then the function returns **`false`** indicating invalid card number.

- If the card number matches the pattern then next step is to remove any hyphens from the card number using the **`strings.ReplaceAll`** function then checks for consecutive repeated digits in the card number. It iterates over the characters of the card number up to the third-to-last Character.
- If four consecutive characters are the same then function returns **`false`** indicating an invalid card number.

- If the card number passes all the checks, then the function returns **`true`** that indicate a valid card number.

## `main()`

- In the **`main`** function, the program first reads the number of credit card numbers to validate from the user.

- Then, a loop is used to read each credit card number and validate it using the **`isValidCreditCardNumber`** function.

- If a card number is valid, then the program prints **"Valid"**, otherwise it prints **"Invalid"**.