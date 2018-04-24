package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

const dir = "./prank/"

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func changeFileNames() {
	re := regexp.MustCompile("[^A-za-z]")

	file, err := os.Open(dir)
	handleErr(err)
	defer file.Close()

	list, err := file.Readdirnames(0) // 0 to read all files and folders
	handleErr(err)

	for _, name := range list {
		fmt.Println("\nOld Name: ", name)
		newName := re.ReplaceAllString(name, " ")
		fmt.Println("New Name: ", newName)
		os.Rename(dir+name, dir+newName)
	}

	fmt.Println("File names have been changed")
}

func main() {
	changeFileNames()
}

//http://www.golangprograms.com/how-to-read-names-of-all-files-and-folders-in-current-directory.html
