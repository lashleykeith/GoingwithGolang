package main
 
import (
    "log"
    "os"
    "fmt"
    "regexp"
)
func readCurrentDir() {
    dir := "C:\\Users\\g\\Desktop\\prank"
    file, err := os.Open(dir)
    if err != nil {
        log.Fatalf("failed opening directory: %s", err)
    }
    defer file.Close()
 
    list,_ := file.Readdirnames(0) // 0 to read all files and folders
    for _, name := range list {
        oldName := name
        fmt.Println("Old Name - ", oldName)
        re := regexp.MustCompile( "[^A-za-z]" )
        newName := re.ReplaceAllString( oldName, " ")
        fmt.Println("New Name - ", newName)        
        os.Rename(dir+oldName, dir+newName)
        fmt.Println("File names have been changed")

    }
}
 
func main() {
    readCurrentDir()

}


//http://www.golangprograms.com/how-to-read-names-of-all-files-and-folders-in-current-directory.html