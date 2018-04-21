
package main
 
import (
    "log"
    "os"
    "fmt"
    "strings"
)
func readCurrentDir() {
    file, err := os.Open("C:\\Users\\g\\Desktop\\prank")
    if err != nil {
        log.Fatalf("failed opening directory: %s", err)
    }
    defer file.Close()
 
    list,_ := file.Readdirnames(0) // 0 to read all files and folders
    for _, name := range list {
        fmt.Println("Old Name - ", name)
        fmt.Println("New Name - ", strings.Replace(name, "0123456789", " ", -1))
    }
}
 
func main() {
    readCurrentDir()
}

//http://www.golangprograms.com/how-to-read-names-of-all-files-and-folders-in-current-directory.html