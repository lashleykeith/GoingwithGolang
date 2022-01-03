
// package main

// import "fmt"

// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

// func main() {
// 	fmt.Println(split(17))
// }


// package main
  
// import "fmt"
  
// // Main Method
// func main() {
  
//     // calling the function, here
//     // function returns two values
//     m, d := calculator(105, 7)
  
//     fmt.Println("105 x 7 = ", m)
//     fmt.Println("105 / 7 = ", d)
// }
  
// // function having named arguments
// func calculator(a, b int) (mul int, div int) {
  
//     // here, simple assignment will
//     // initialize the values to it
//     mul = a * b
//     div = a / b
  
//     // here you have return keyword
//     // without any resultant parameters
//     return
// }






package main 

import (
        "fmt"
        "strconv"
        )

func add(x, y float64) float64 {
    return x + y
}

func substract(x, y float64) float64 {
    return x - y
}

func divide(x, y float64) float64 {
    return x / y
}

func multiply(x, y float64) float64 {
    return x * y
}

func swap(x string, y string)(string, string){
    return y, x
}

func calculator(a, b float64) (mul float64, sub float64) {
  
    // here, simple assignment will
    // initialize the values to it
    mul = a * b
    sub = a - b
    
  
    // here you have return keyword
    // without any resultant parameters
    return
}
func main(){
    fmt.Println("Enter the initial price of the stock: ")
    var price float64
    fmt.Scanln(&price)
    fmt.Println("Your initial stock price was ", price, " dollars")
    fmt.Println("Enter quantity of stocks: ")
    var quantity float64
    fmt.Scanln(&quantity)
    fmt.Println("You brought ", quantity, " stocks at", price,"dollars each.")
    investment_init := multiply(price, quantity)
    fmt.Println("The initial investment was ", investment_init, "dollars")
    increase := multiply(investment_init, 0.01227)
    fmt.Println("profit ",increase)
    new_investment1 := add(investment_init, increase)
    fmt.Println("The investment for the first week will be", new_investment1)
    increase = multiply(new_investment1, 0.01227)
    new_investment2 := add(new_investment1, increase)
    fmt.Println("The new investment for the second week will be ", new_investment2)
    payday1 := add(new_investment2, new_investment2)
    fmt.Println(" You now have brought one more shares now you have ", payday1)
    increase = multiply(payday1, 0.01227)
    new_investment3 := add(payday1, increase)
    fmt.Println("The new investment for the third week will be ", new_investment3)
    increase = multiply(new_investment3, 0.01227)
    new_investment4 := add(new_investment3, increase)
    fmt.Println("The new investment fo the fourth week will be", new_investment4)
    payday2 := add(new_investment4, new_investment4)
    fmt.Println(" You now have brought one more shares now you have ", payday2)
    s1 := strconv.FormatInt(int64(payday2),10)

    a, b := "Dollars ", s1
    fmt.Println(a, b)

    a,b = swap(" Dollars ", s1)
    fmt.Println(a,b)

    fmt.Println("New Bank")

    payday3 := float64(payday2)
    g, m := calculator(payday3, .05)
    total := add(payday3, .05)

    fmt.Println("PayDay x interest = ", g)
    fmt.Println("PayDay - interest = ", m)
    fmt.Println("Money you have now = ", total )


}

