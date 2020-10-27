package main
import (
    "fmt"
    "io/ioutil"
)

// Helper function for logging errors
func check(err error) {
    if err != nil {
		panic(err)
	}
}

func main() {
    
    // Reading content from file.
    content, err := ioutil.ReadFile("example.txt")
	check(err)
	fmt.Printf("File contents: %s\n", content)

    // Writing to a new file newFile.txt.
	data := []byte("Hello, Educative!")
	err = ioutil.WriteFile("newFile.txt", data, 0644)
	check(err)

    // Printing all files in the current directory.
    // Notice a new newFile.txt file (that we created above).
    files, err := ioutil.ReadDir(".")
	check(err)
    fmt.Println("Files in the current directory:")
    for _, file := range files {
		fmt.Println(file.Name())
	}
}