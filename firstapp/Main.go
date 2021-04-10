package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "io/ioutil"
	// "log"
	// "net/http"
)

// 3 ways to declare variables
// function level 	var i int = 42 or i := 42
// package level var i int = 42 (not i := 42)

// package variable must be declare like this:  var i int = 42
// lower case 1st letter variable can only be accessed within the package,
// upper case 1st letter can be accessed from outside
// 3 level of variable scope: block, package (lower case), global (upper case)

// func main() {
// 	// var i int
// 	// i = 24
// 	// var i int = 24
// 	// i := 24
// 	var i int = 24
// 	var j string = strconv.Itoa(i)

// 	fmt.Printf("%v , %T \n", i, i)
// 	fmt.Printf("%v , %T \n", j, j)
// }

func main() {
	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

// const (
// 	isAdmin uint8 = 1 << iota
// 	isHeadquarters
// 	canSeeFinancials

// 	canSeeAfrica
// 	canSeeAsia
// 	canSeeEurope
// 	canSeeNorthAmerica
// 	canSeeSouthAmerica
// )

// func main() {
// 	var roles byte = isAdmin | canSeeFinancials | canSeeEurope //1000101

// 	fmt.Printf("%v %T\n", isAdmin, isAdmin)
// 	fmt.Printf("%v %T\n", canSeeFinancials, canSeeFinancials)
// 	fmt.Printf("%v %T\n", canSeeEurope, canSeeEurope)
// 	fmt.Printf("%v %T\n", roles, roles)
// 	fmt.Println("isAdmin? ", isAdmin & roles == isAdmin)
// 	fmt.Println("isHeadquarters? ", roles & isHeadquarters == isHeadquarters)
// }

// func main() {
// 	var p = make(map[string]uint)
// 	p["Califonia"] = 10
// 	p["Florida"] =  20
// 	p["New York"]=  30
// 	fmt.Println(p)
// }

// func main() {
// 	var msg string = "Hello_1"
// 	fmt.Println(msg)
// 	updateMessage(&msg)
// 	fmt.Println(msg)
// }
// func updateMessage(msg *string) {
// 	*msg = "Hello_2"
// }

// ~/.bash_profile
// # Include the location of the go app bin (/usr/local/go/)
// export PATH="/usr/local/go/bin:/Users/harjinder.chaudry/golib/bin:/opt/vertica/bin:$PATH"
// # (workspaces - must contain a src directory) 3rd party libraries + personal workspace
// # first segment is used by the 'go get' command to get files from github
// export GOPATH="/Users/harjinder.chaudry/golib:/Users/harjinder.chaudry/code"

/*
go run ~/code/src/github.com/hchaudry/firstapp/Main.go
go build github.com/hchaudry/firstapp
go install github.com/hchaudry/firstapp
*/
