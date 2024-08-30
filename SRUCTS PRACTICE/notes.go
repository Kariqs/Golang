package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"notes.com/create-note/note"
)

func main() {
	var title string = fetchUserInput("Enter the title:")
	var content string = fetchUserInput("Enter the content of the note:")

	var userNote, err = note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = userNote.Save()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Note saved successfully.")
}

func fetchUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
