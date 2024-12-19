package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	path := "path to utilities" + os.Args[1]

	if len(os.Args) < 2 {
		fmt.Println("Specify the name util as an argument")
		return
	}

	if _, err := os.Stat(path); err == nil {
		fmt.Println("The folder already exists:", path)
	} else if os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			log.Fatal("Error creating directory:", err)
		}
		fmt.Println("Directory created:", path)
	} else {
		log.Fatal("Error checking directory:", err)
	}

	cmd := exec.Command("code", path)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error when running VS Code:", err)
	} else {
		fmt.Println("The folder is open in VS Code")
	}
}
