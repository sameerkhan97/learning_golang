/* * As we saw in the discussion of interfaces above, a type need not declare explicitly that it implements an interface.
	 Instead, a type implements the interface just by implementing the interface's methods
   * It's necessary only to ask whether a type implements an interface, without actually using the interface itself, 
     perhaps as part of an error check, use the blank identifier to ignore the type-asserted value:
   * This situation arises is when it is necessary to guarantee within the package implementing the type that 
      it actually satisfies the interface
*/

package main

import (
	"fmt"
	"json"
)

func main() {

	if _, ok := val.(json.Marshaler); ok {
		fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
	}
	//One place this situation arises is when it is necessary to guarantee within the package implementing
	//the type that it actually satisfies the interface.
}
