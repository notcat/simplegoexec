package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/notcat/simplegoexec/commands"
)

// https://tutorialedge.net/golang/reading-console-input-golang/

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Execute a command/path to executable")
	fmt.Println("----------------------------------------")

	for {
		fmt.Print("-> ")
		cmd, _ := reader.ReadString('\n')
		// convert CRLF to LF
		cmd = strings.Replace(cmd, "\n", "", -1)

		// Put it in its own gothread to prevent the blocking of opening another file
		go run(cmd)
	}
}

func run(cmd string) {
	out, err := commands.Exec(cmd)
	if err != nil {
		fmt.Println(string(out))
		log.Fatal(err)
	}
	fmt.Printf(string(out))
}
