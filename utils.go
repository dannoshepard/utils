package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadStdin() string {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var stdin string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin += scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		return stdin
	}

	return ""
}
