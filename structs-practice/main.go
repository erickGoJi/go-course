package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	t, b, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote, err := note.New(t, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Error saving the note: ", err)
		return
	}

	fmt.Println("Note saved successfully!")

}

func getNoteData() (string, string, error) {
	title, err := getUserData("Enter the title of the note: ", "Title")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	body, err := getUserData("Enter the body of the note: ", "Body")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return title, body, nil
}

func getUserData(prompText, fieldname string) (string, error) {
	fmt.Print(prompText)

	reader := bufio.NewReader(os.Stdin)
	// recieve a parameter to stop reading
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	// remove the newline character from the input
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	if input == "" {
		errFormat := fmt.Sprintf("%v field is required", fieldname)
		return "", errors.New(errFormat)
	}

	return input, nil
}
