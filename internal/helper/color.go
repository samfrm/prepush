package helper

import (
	"fmt"
)

// This helper will help to print in console with colors

func PrintSuccess(text string) {
	// print in console with bright green color
	fmt.Println("\u001B[1;92m", text, "\u001B[0m")
}

func PrintError(text string) {
	// print in console with red color
	fmt.Println("\u001B[1;91m", text, "\u001B[0m")
}

func PrintInfo(text string) {
	// print in console with bright blue color
	fmt.Println("\u001B[1;94m", text, "\u001B[0m")
}

func PrintWarning(text string) {
	// print in console with bright yellow color
	fmt.Println("\u001B[1;93m", text, "\u001B[0m")
}
