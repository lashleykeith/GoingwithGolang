/*
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
*/

/* 
package main 

import "fmt"


func add(x float64, y float64) float64 {
	return x + y
}

func substract(x float64, y float64) float64{
	return x - y
}

func divide(x float64, y float64) float64{
	return x / y
}

func multiply(x float64, y float64) float64{
	return x * y
}

func main(){
	fmt.Print("Enter inital price of the stock: ")
	var price float64
	fmt.Scanln(&price)
	fmt.Print("Enter quantity of stocks: ")
	var quantity float64
	fmt.Scanln(&quantity)
	fmt.Println(quantity)
	fmt.Println(price)
	stock_price := multiply(price, quantity)
	fmt.Println("The initial stock is ", stock_price) 
	increase := multiply(stock_price, 0.01227)
	fmt.Println(increase)
	new_price1 := add(stock_price, increase)
	fmt.Println("The new stock price for the first week will be ", new_price1)
	//new_price1 := add(stock_price, increase)

}


*/

/*
package main

import "fmt"


func add(x float64, y float64) float64 {
	return x + y
}

func substract(x float64, y float64) float64 {
	return x - y	
}

func divide(x float64, y float64) float64 {
	return x/y
}

func multiply(x float64, y float64) float64{
	return x * y
}



func main() {
	fmt.Print("Enter initial price of stock : ")
	var price float64
	fmt.Scanln(&price)
	fmt.Print("Enter quantity of stocks: ")
	var quantity float64
	fmt.Scanln(&quantity)
	fmt.Print(quantity,"\n")
	fmt.Print(price, "\n")
	//happy := add(5,10)
	//fmt.Print(happy)
	stock_price := multiply(price, quantity)
	fmt.Print("\n")
	fmt.Print("The initial stock price is ", stock_price)
	fmt.Print("\n")
	increase := multiply(stock_price, 0.01227)
	new_price1 := add(stock_price, increase)
	fmt.Print("The new stock price will be for the first week ", new_price1)
	fmt.Print("\n")
	increase = multiply(new_price1, 0.01227)
	new_price2 := add(new_price1, increase)
	fmt.Print("The new stock price will be for the second week ", new_price2)
	fmt.Print("\n")
	payday1 := add(new_price2, new_price2)
	fmt.Print("You now have brought two more shares now you have ", payday1)
	increase = multiply(payday1, 0.01227)
	new_price3 := add(payday1, increase)
	fmt.Print("\n")
	fmt.Print("The new stock price will be for the third week ", new_price3)
	fmt.Print("\n")
	increase = multiply(new_price3, 0.01227)
	new_price4 := add(new_price3, increase)
	fmt.Print("The new stock price will be for the fourth week ", new_price4)
	fmt.Print("\n")
	payday2 := add(new_price4, new_price4)
	fmt.Print("You now have brought two more shares now you have ", payday2)


}

*/ 