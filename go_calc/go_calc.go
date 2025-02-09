package main

import (
	"bufio"
	"os"
	"os/exec"
	s "strings"
	"fmt"
	"regexp"
	"strconv"
)

const quit = "q"
const clear = "c"

var p = fmt.Println

func renderMenu() {
	p("(q) - quit | (c) - clear saved result\n")
}

func quitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func quitOnQ(input string) {
	if s.ToLower(input) == quit {
		p("Bye!")
		os.Exit(0)
	}
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	savedResult := ""

	for {
		clearTerminal()
		renderMenu()

		if savedResult != "" {
			p(savedResult, "\n")
		}

		scanner.Scan()
		quitOnError(scanner.Err())

		input := scanner.Text();

		if input == "" {
			continue
		}

		quitOnQ(input)

		if s.ToLower(input) == clear {
			savedResult = ""
		} else {
			validTokens := make([]string, 0)
			inputTokens := s.Split(input, " ")

			for i := range len(inputTokens) {
				token := s.TrimSpace(inputTokens[i])
				r := regexp.MustCompile("[^\\d\\+\\-\\*\\/\\.]")
				token = r.ReplaceAllString(token, "")
				if token != "" {
					validTokens = append(validTokens, token)
				}
			}

			val1, _ := strconv.Atoi(validTokens[0])
			operator := validTokens[1]
			val2, _ := strconv.Atoi(validTokens[2])

			switch operator {
			case "+":
				savedResult = strconv.Itoa(val1 + val2)
			case "-":
				savedResult = strconv.Itoa(val1 - val2)
			case "*":
				savedResult = strconv.Itoa(val1 * val2)
			case "/":
				savedResult = strconv.Itoa(val1 / val2)
			}
		}

		p("")
	}
}
