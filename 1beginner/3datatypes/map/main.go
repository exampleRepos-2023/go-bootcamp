package main

import "fmt"

func main() {
	// A map is unordered, growing collection of key to value pairs
	// ["Tommy" -> "Ninja", "Bobby" -> "Sensei", "Johnny" -> "Ninja"]
	// Keys are unique, you can search by keys in constant time
	// Values are not unique

	var m map[string]string // map[(key type)](value type)
	if m == nil {
		fmt.Println("nil is map's default zero value")
	}
	m = map[string]string{"Tommy": "Ninja", "Bobby": "Sensei"}
	fmt.Println(m)

	// Individual elements
	fmt.Println(m["Tommy"])
	m["Tommy"] = "Not a Ninja"
	fmt.Println(m["Tommy"])
	m["Andy"] = "Ninja to be"
	fmt.Println(m)
	delete(m, "Andy")
	fmt.Println(m)
	value := m["Johnny"]
	fmt.Printf("value=[%s]\n", value)
	value, ok := m["Johnny"]
	fmt.Printf("value=[%s], ok=[%t]\n", value, ok)
	m["Johnny"] = ""
	value, ok = m["Johnny"]
	fmt.Printf("value=[%s], ok=[%t]\n", value, ok)

	// Iterations
	for key, value := range m {
		fmt.Printf("%s: %s | ", key, value)
	}
	fmt.Println()
	fmt.Println(len(m))

	// Pre-allocation
	m = make(map[string]string, 5)
	fmt.Println(m)
	fmt.Println(len(m))

	arrayToMap := map[[2]int]map[string]string{
		[...]int{1, 2}: {"one": "two"},
	}
	fmt.Println(arrayToMap)
}
