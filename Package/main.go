// 1) Package creation
// 2) Creation of the module file "go.mod" with the command go mod init <module name>
// 3) Creating a "modulo" binary file with the go build command. Compile the entire project and run it with the command ./modulo
// 4) If a function starts with a lowercase letter, it will only be visible inside the package it is in (similar to private). If it starts with a uppercase letter, it can be exported by other packages (similar to public).

package main

import (
	"fmt"
	"modulo/support"
)

func main() {
	fmt.Println("Writing from main file")
	support.Write()
}
