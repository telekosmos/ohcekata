package main

import (
	"fmt"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

func executor(in string) {
	in = strings.TrimSpace(in)

	// var method, body string
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "exit":
		fmt.Println("Bye!")
		os.Exit(0)
	}

	fmt.Println("We go on...")
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

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(promptName+"> "),
		// prompt.OptionLivePrefix(livePrefix),
		prompt.OptionTitle("ehco"),
	)
	p.Run()
}
