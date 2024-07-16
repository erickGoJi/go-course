package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// saver interface is used to save the data to the file
type saver interface {
	Save() error
}

//	type displayer interface {
//		Display()
//	}
//

// outputtable interface is used to display the data, and embeds the saver interface
type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {

	printSomething(1)
	printSomething("Hello")
	printSomething(1.5)

	t, b, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}

	todoContent, err := getTodoData()
	if err != nil {
		fmt.Println(err)
		return
	}

	todoItem, err := todo.New(todoContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todoItem)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(t, b)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		fmt.Println(err)
		return
	}

}

// Using an empty interface to accept any type of data
func printSomething(value interface{}) {
	intValue, ok := value.(int)
	if ok {
		fmt.Println("This is an integer: ", intValue)
		return
	}

	floatValue, ok := value.(float64)
	if ok {
		fmt.Println("This is an float: ", floatValue)
		return
	}

	stringValue, ok := value.(string)
	if ok {
		fmt.Println("This is an string: ", stringValue)
		return
	}
	// switch value.(type) {
	// case int:
	// 	// ... do something
	// 	fmt.Println("This is an integer")
	// case float64:
	// 	// ... do something
	// 	fmt.Println("This is a float")
	// case string:
	// 	// ... do something
	// 	fmt.Println("This is a string")
	// default:
	// 	// ...
	// 	fmt.Println("Unknown type")
	// }
	// fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// saveData saves the data to the file
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Error saving the note: ", err)
		return err
	}

	fmt.Println("Note saved successfully!")
	return nil
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

func getTodoData() (string, error) {
	content, err := getUserData("Enter the todo content: ", "Content")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return content, nil
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
