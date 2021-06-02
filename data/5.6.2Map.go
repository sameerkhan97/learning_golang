package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var m1 = map[string]Vertex{
	"Bell Labs": Vertex{40.12, -4854.854},
	"Google ":   Vertex{455, 7854.454},
}

func main() {
	//1 :
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.612, -415.413,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m1)

	//2 :
	m2 := make(map[string]int)

	m2["Answer"] = 42
	fmt.Println("The Value : ", m2["Answer"])
	m2["Answer"] = 48
	fmt.Println("The Value : ", m2["Answer"])
	delete(m2, "Answer")
	fmt.Println("The Value : ", m2["Answer"])
	v, ok := m2["Answer"]
	fmt.Println("the Value : ", v, "Present ?", ok)
}
