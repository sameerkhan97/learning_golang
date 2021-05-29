/*1.The make function returns a map of the given type, initialized and ready for use
  2.Mutating Maps
	Insert or update an element in map m:
		m[key] = elem
	Retrieve an element:
		elem = m[key]
	Delete an element:
		delete(m, key)
	Test that a key is present with a two-value assignment:
		elem, ok = m[key]
	If key is in m, ok is true. If not, ok is false.
	If key is not in the map, then elem is the zero value for the map's element type.
	Note: If elem or ok have not yet been declared you could use a short declaration form:
		elem, ok := m[key]
*/

package main
import "fmt"
func main(){
	statePop:=map[string]int{	//key of type string and value of type int
		"California"   : 31654 ,
		"Texas"       : 23145 ,
		"Florida"     : 54987 ,
		"New York"    : 71475 ,
		"Pennsylvania" : 36841 ,
		"Illinois"    : 16414 ,
		"Ohio"        : 46345 ,
	}
	fmt.Println(statePop)
	//another declaration using make method
	statePop1:=make(map[string]int,10)
	statePop1=map[string]int{
		"California"   : 31654 ,
		"Texas "       : 23145 ,
		"Florida "     : 54987 ,
		"New York "    : 71475 ,
		"Pennsylvania" : 36841 ,
		"Illinois "    : 16414 ,
		"Ohio "        : 46345 ,
	}
	fmt.Println(statePop1)
	//Note : arrays can be used as a key-type but not map 
	m:=map[[3]int]string{}
	fmt.Println(statePop,m)	

	//manipulation : 
	fmt.Println("Populataion of Ohio " ,statePop["Ohio"])
	statePop["Ohio"]=78964
	fmt.Println("New Populataion of Ohio " ,statePop["Ohio"])
	
	//checking if key is present or not
	fmt.Println("Searching in Map")
	pop,ok:=statePop["Delhi"];
	fmt.Println(pop,ok)
	pop,ok=statePop["New York"];
	fmt.Println(pop,ok)

	//size of map
	fmt.Println("No of Cities : ",len(statePop))

	//deleting from map
	fmt.Println("deleting from map")
	sp:=statePop;
	fmt.Println(sp)
	delete(sp,"Ohio")
	fmt.Println(sp)
	fmt.Println(statePop)

}