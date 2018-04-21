/*package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    files, err := ioutil.ReadDir(".")
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fmt.Println(file.Name())
    }
}
*/


/*
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "a space-separated string"
    str = strings.Replace(str, " ", ",", -1)
    fmt.Println(str)
}
*/


package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", " ", -1))
}




/*package main

import (
    "fmt"
    "os"
)

func main() {

    before := "/home/sam/test.txt"
    after := "/home/sam/optimized.txt"

    // Rename or move file from one location to another.
    os.Rename(before, after)
    fmt.Println("DONE")
}
*/