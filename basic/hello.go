/**
To run the program, put the code in hello-world.go and use go run.
Sometimes we’ll want to build our programs into binaries. We can do this using go build.
We can then execute the built binary directly.

$ go run hello.go
Hello Go!

$ go build hello.go
$ ls
hello	hello.go

$ ./hello
Hello Go!
*/

package main

import "fmt"

func main() {
	fmt.Println("Hello Go!")
}