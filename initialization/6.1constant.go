/* * Constants in Go are created at compile time, even when defined as locals in functions.
   * Constants can only be numbers, characters (runes), strings or booleans.
   * Because of the compile-time restriction, the expressions that define them must be constant expressions, evaluatable by the compiler for example : 1<<3 is a constant
     expression, while math.Sin(math.Pi/4) is not because the function call to math.Sin needs to happen at run time.
   * In Go, enumerated constants are created using the iota enumerator. Since iota can be part of an expression and expressions can be implicitly repeated, it is easy to
     build intricate sets of values.
   * use 'const' keyword to declare constant   */

package main

import "fmt"

func main() {
	const a int64 = 42 //TYPED CONSTANT
	fmt.Printf("%v,%T\n", a, a)

	const b = 4.12 //UNTYPED CONSTANT
	fmt.Printf("%v,%T\n", b, b)
	//statement a=a+1/b=b+1 cause error bcoz we cant modify constants

	const c = false //UNTYPED CONSTANT
	fmt.Printf("%v,%T\n", c, c)
}
