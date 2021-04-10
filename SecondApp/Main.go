// primative types

// int8		-128	127
// uint8	0		255	(alias = byte)
// string	array of uint8 (utf-8)	e.g. "a" double quotes
// rune		array of int32 (utf-32)	e.g. 'a' single quotes

package main

import "fmt"

func main() {
	var s byte = 255 // all variables are initialized to zero, false
	var t byte
	t = s & 3 // bitwise AND
	fmt.Printf("%v, %T\n", t, t)

	var j string
	j = "this is a string"
	k := []byte(j)
	l := string(k)
	var m rune = 'a'

	fmt.Printf("%v, %T \n", j, j)
	fmt.Printf("%v, %T \n", k, k)
	fmt.Printf("%v, %T \n", l, l)
	fmt.Printf("%v, %T \n", m, m)

}

/*
go run ~/code/src/github.com/hchaudry/SecondApp/Main.go
go build github.com/hchaudry/SecondApp
go install github.com/hchaudry/SecondApp
*/
