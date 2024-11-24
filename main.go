package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	bot "bitwarn.com/bit_warn/bot"
)

func main() {
	// Open keys file
	f, err := os.Open("keys.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find and set bot token
	r := bufio.NewReader(f)
	token := false
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		if token {
			bot.BotToken = strings.TrimSpace(line)
			break
		}

		if strings.Contains(line, "BotToken") {
			token = true
		}
	}

	// Run bot
	bot.Run()
}