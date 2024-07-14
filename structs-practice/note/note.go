package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Note is a struct, it will hold the note data, addding a json tag to the fields
type Note struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

// Save is a method of the Note type
func (n Note) Save() error {
	filename := strings.ReplaceAll(n.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)

}

// New is a constructor/creation function
func New(title, body string) (Note, error) {
	if title == "" || body == "" {
		return Note{}, errors.New("all fields are required")
	}

	return Note{
		title,
		body,
		time.Now(),
	}, nil
}

// Display is a method of the Note type
func (n Note) Display() {
	fmt.Printf("Your note title: %v has the following content:\n\n%v\n", n.Title, n.Body)
}
