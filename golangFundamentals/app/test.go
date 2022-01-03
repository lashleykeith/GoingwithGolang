


package main

import "fmt"

var i, j int = 1, 2

func main() {
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java)
}

package main

import (
    "fmt"
    "reflect"
    )

var i, j int = 1, 2
var b float64 = 3.05

func main(){
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java, b)
    fmt.Println("i = ", reflect.TypeOf(i))
    fmt.Println("j = ", reflect.TypeOf(j))
    fmt.Println("c = ", reflect.TypeOf(c))
    fmt.Println("python = ", reflect.TypeOf(python))
    fmt.Println("java = ", reflect.TypeOf(java))
    fmt.Println("b = ", reflect.TypeOf(b))
  
}


///////////////////////////////////////////////////////



// package main

// import "fmt"

// func main() {
//     var i, j int = 1, 2
//     k := "Peter"
//     i = 2
//     c, python, java := true, false, "no!"

//     fmt.Println(i, j, k, c, python, java)
// }


// // package main

// // import "fmt"

// // func main() {
// //     var i, j int = 1, 2
// //     k := "Peter"
// //     i = 2
// //     c, python, java := true, false, "no!"

// //     fmt.Println(i, j, k, c, python, java)
// // }
