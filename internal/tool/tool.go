package tool

import (
	"fmt"
	"os"
)

func FlagMode() {
	option := os.Args[1]

	switch option {
	case "--help", "-h":
		printHelpManual()
	case "--version", "-v":
		fmt.Println("0.0.1")
	}
}

func printHelpManual() {
	fmt.Println("Usage: dk [options]")
	fmt.Printf("  %-20s %s\n", "dk", "Run the program")
	fmt.Printf("  %-20s %s\n", "dk [--help | -h]", "Show this help message")
}
