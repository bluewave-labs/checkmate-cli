package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StdIn(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading input: %v", err)
	}

	input = strings.TrimSpace(input)
	return input, nil
}
