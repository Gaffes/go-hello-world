// https://www.golangprograms.com/string-to-boolean-data-type-conversion-in-go.html

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	s1 := "true"
	b1, _ := strconv.ParseBool(s1)
	fmt.Printf("%T, %v\n", b1, b1)

	s2 := ""
	b2, err := strconv.ParseBool(s2)
	fmt.Printf("%T, %v\n", b2, b2)
	fmt.Println(err);

	s3 := "0"
	b3, err := strconv.ParseBool(s3)
	fmt.Printf("%T, %v\n", b3, b3)
	fmt.Println(err);

	s4 := "F"
	b4, _ := strconv.ParseBool(s4)
	fmt.Printf("%T, %v\n", b4, b4)

var b bool = true
fmt.Println(reflect.TypeOf(b))
// Boolean value can be converted to a string.
var s string = strconv.FormatBool(true)
fmt.Println(reflect.TypeOf(s))

}
