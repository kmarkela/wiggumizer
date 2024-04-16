package fuzz

import (
	"fmt"
)

func updateProgress(hei, wdi, rl int, ii <-chan struct{}) {

	fmt.Printf("* Umount of Requests to Fuzz: %d\n", hei)
	fmt.Printf("* Umount of words in wordlist: %d\n", wdi)
	if rl != 0 {
		fmt.Printf("* Rate Limit: %d req/s\n", rl)
	}
	fmt.Println("===============================================")
	fmt.Println("# Requests Sent: 0")
	i := 1
	for range ii {
		moveCursorUp(1) // Move cursor up by one line
		clearLine()     // Clear the entire line
		i++
		fmt.Printf("# Requests Sent: %d\n", i)

	}

}

// moveCursorUp moves the cursor up by n lines.
func moveCursorUp(n int) {
	fmt.Printf("\033[%dA", n)
}

// clearLine clears the entire line.
func clearLine() {
	fmt.Print("\033[2K")
}
