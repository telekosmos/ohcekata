package main

import (
	"fmt"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

func executor(in string) {
	in = strings.TrimSpace(in)

	blocks := strings.Split(in, " ")
	input := blocks[0]
	switch input {
	case "exit!":
		fmt.Println("Bye!")
		os.Exit(0)

	default:
		rev := reverse(input)
		palindrome := isPalindrome(input)
		fmt.Printf("%s\n", rev)
		if palindrome {
			fmt.Println("Nice word!")
		}
	}
}

func completer(t prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{
		{Text: ""},
	}
}

func main() {
	var promptName string
	if len(os.Args) == 2 {
		promptName = os.Args[1]
	} else {
		promptName = "noname"
	}

	fmt.Printf("Hola %s!\nType 'exit!' to quit\n", promptName)
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(promptName+"> "),
		// prompt.OptionLivePrefix(livePrefix),
		prompt.OptionTitle("ehco"),
	)
	p.Run()
}
