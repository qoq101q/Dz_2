package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("poem.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Unable to close file", err)
			os.Exit(1)
		}
	}(file)

	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}

		var letters []string
		for i := 0; i < n; i++ {
			letters = append(letters, string(data[i]))
			fmt.Println(letters[i])
		}
	}

	// Create file and write data
	/*text := "Hello Go!"
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Unable to write", err)
		os.Exit(1)
	}*/

	//Delete file
	/*err = os.Remove("hello.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File delete successfully.")
	*/
}
