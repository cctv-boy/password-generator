package main

import (
	"PasswordGenerator/pkg/password_generator"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var balnk string

func main() {

	color.Blue("Hi, this program will help you generate a random password")

	fmt.Println(generate())
	fmt.Scan(&balnk)

	os.Exit(1)
}

func generate() string {
	var (
		characters string
		length     int
		count      int
		result     string
	)
	color.Blue("Enter the characters that must be present in the password (you can just randomly press the keyboard)")
	fmt.Scan(&characters)
	if len(characters) <= 2 {
		color.Yellow("There must be at least 3 characters in your list")
		generate()
	} else {
		color.Blue("Great, now enter the number of characters in the password (it is recommended to use passwords that use more than 7 characters)")
		_, err := fmt.Scan(&length)
		if err != nil {
			color.Red("What you entered is not a number, the program ends :(")
			fmt.Scan(&balnk)
			os.Exit(1)
		}
		color.Blue("Great, now enter the number of passwords I will generate")
		_, err = fmt.Scan(&count)

		if err != nil {
			color.Red("What you entered is not a number, the program ends :(")
			fmt.Scan(&balnk)
			os.Exit(1)
		}
		for i := 0; i < count; i++ {
			result = result + password_generator.GeneratePassword(length, characters) + "\n"
		}

	}
	color.Blue("Great, here's a list of passwords:\n")
	return result
}
