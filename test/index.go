package main

import "fmt"


// Multiple Constants Decalration block 
const (
	PRODUCT = "Apple"
	QUANTIY = 100000
	PRICE = 100.00
	STOCK = true
)
// Global variable declarations 
var (
	m int
	n int 
)

func main () {

    var x int = 1 //interger datatype decalration 
	var y int 
	var a,b,c = 2,4.5,8.5 //Multiple float 32 variable declarations 

    fmt.Println(x)
	fmt.Println(a,b,c)
	fmt.Println(y)


	city := "New York" //String Variable Declaration
	Country := "United States of America"

	fmt.Println(city)
	fmt.Println(Country) // Variable Nmes are case sensetive 



	food,drink,price:= "Cheeseburger", "Inu Wai", 20.25 //Multiple type of variable declaration in same line 
	fmt.Println(food,drink,price)
	m,n= 1,2
	fmt.Println(m,n)





	fmt.Println(PRICE)
	fmt.Println(PRODUCT)
	fmt.Println(QUANTIY)
	fmt.Println(STOCK)
}