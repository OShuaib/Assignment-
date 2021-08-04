package main

import (
	"fmt"
	"github.com/elliotchance/orderedmap"
)

func main(){
	m := orderedmap.NewOrderedMap()

	m.Set("foo", "bar")
	m.Set("qux", 1.23)
	m.Set(123, true)
	m.Set("Ade", 23)
	m.Set("Shuaib", true )

	m.Delete("123")

	for _, key := range m.Keys() {
		value, _:= m.Get(key)
		fmt.Println(key, value)
	}
	// Iterate through all elements from oldest to newest:
	/*for el := m.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key, el.Value)
	}*/

	// You can also use Back and Prev to iterate in reverse:
	for el := m.Back(); el != nil; el = el.Prev() {
		fmt.Println(el.Key, el.Value)
	}
	m.GetElement("qux")
}