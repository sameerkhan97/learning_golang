//Variables can be initialized just like constants but the initializer can be a general expression computed at run time.

package main

import "fmt"

func main() {
	var num1 int = 1
	var num2 int = 2
	var num3 int
	var num4 int
	num3 = 3
	num4 = 4

	var num5, num6 int
	num5, num6 = 5, 6
	//if dont use var key word then
	num7 := 7 //var num7=9

	//constant variable
	const num8 = 8

	x, y, z, s := 8, true, 7.2431584, "string"
	fmt.Print(x, y, z, s)

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)
	fmt.Println(num6)
	fmt.Println(num7)
	fmt.Println(num8)
	var (
		home  = "HOME"
		user  = 1021
		model = 12.1
	)
	fmt.Println(home)
	fmt.Println(user)
	fmt.Println(model)
}
