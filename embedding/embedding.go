/* * Go does not provide the typical, type-driven notion of subclassing, but it does have the ability to “borrow” pieces 
     of an implementation by embedding types within a struct or interface.
   * The embedded elements are pointers to structs and of course must be initialized to point to valid structs before they can be used   */

package main

import "fmt"

type person struct {
	name, add string
	phone     int
}
type employe struct {
	person     //Embedding person inside employe
	id, salary int
	dept       string
}

func (E *employe) setInfo() {
	fmt.Scan(&E.name, &E.add, &E.phone, &E.salary, &E.dept, &E.id)
}
func (E employe) getInfo() {
	fmt.Println("Name   	: ", E.name)
	fmt.Println("Id     	: ", E.id)
	fmt.Println("Dept   	: ", E.dept)
	fmt.Println("Salary 	: ", E.salary)
	fmt.Println("Phone No. : ", E.phone)
	fmt.Println("Adress    : ", E.add)
}
func main() {
	emp := employee
	emp.setInfo()
	emp.getInfo()
}
