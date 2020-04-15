/*
+Documentation: https://golang.org

+build a go file and an executable:
go build Basics.go
./Basics

- if you want to give the exe a different name:
go build -o myapp.exe

-if you want an executable for a different OS and CPU, use options:
GOOS=windows GOARCH=amd64 go build -o winApp.exe

-for mac, GOOS=darwin

+just run the go file:
go run Basics.go

+find go environment variable values in terminal:
go env

+correct the indetation of the go file in terminal:
gofmt -w Basics.go //-w writes the updates

-if you want all files in the dir to be formatted:
go fmt

*/
package main

import (
	"fmt"
	"math"
)

const y = 4

func main() {
	fmt.Println("Hello There..")
	x := 5.2
	fmt.Printf("%2.2f to the power of 8 is %.4f", x, math.Pow(x, 8))
}
