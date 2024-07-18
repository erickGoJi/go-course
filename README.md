# go-course

Package main is the entry point for the program.
```go
package main

import "fmt"

// greet prints a greeting message to the console.
func greet(name string) {
    fmt.Println("Hello, " + name + "!")
}

// main is the main function that gets executed when the program starts.
func main() {
    greet("World")
}
```

