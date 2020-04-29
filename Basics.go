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

+GO Lang is statically typed, meaning that the variable types will get checked
at compile time. similar to Java and C++.
dynamically typed checked languauges such as Python are checked at run time.

+arrays cannot be constants

+slices : it saves the values in a backing array. slice contains a header which has the address, length and capacity of the slice
- a slice and array with the same items will have different sizes and always array will take more space because slice only has 3 elements
- when you append new values to a slice, capacity will grow by 2^len of slice. this is if the capacity doesn't have enoigh space for the new item
-you cannot assing a value to a slice with length zero, i.e. len(cities) == 0 => cities[0] = "Toronto" will fail
*/
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var a = 5   //if not used compilation will fail
const y = 4 //if not used NO error

const (
	min1 = 100
	min2 // since we didnt provide value for min2 and 3, they will be defaulted to 100 like min1
	min3
)

func main() {
	fmt.Println("Hello There..")
	x := 5.2 //:= is used within blocks. you cannot use the same name used before
	fmt.Printf("%2.2f to the power of 8 is %.4f", x, math.Pow(x, 8))
	fmt.Println("\n", x, " to the power of 8 is", math.Pow(x, 8)) //use comma for concatenation

	y, z := "xyz", 2020
	_, _ = y, z //this is a way to mute the compiler for not using a variable

	var (
		name    string
		visible bool
		salary  float64
	)
	fmt.Println("\n")
	fmt.Println(name, visible, salary)

	var a, b, c int
	fmt.Println(a, b, c) // 0 0 0

	const p = 2 //untyped const
	const q = 3.5
	fmt.Println(p * q)

	/*
	  const r int = 2//typed const
	  fmt.Println(q*r//you cannot multiply typed an untyped constants
	*/

	var p2 float64 = p   //convert int to float64
	fmt.Printf("%T", p2) //find the type by %T

	const (
		c1 = iota //iota starts from 0
		c2
		c3
	)
	fmt.Println("\n")
	fmt.Println(c1, c2, c3) // 0 1 2

	const (
		c4 = (iota * 2) + 1 //iota starts from 0
		_                   //blank identifier will be 3 but skipped
		c5
		c6
	)
	fmt.Println("\n")
	fmt.Println(c4, c5, c6) // 1   5 7

	/*
	   data types
	   int, int8, int16,..
	   unit, unit8, unit16,...
	   float,..
	   rune alias for int32
	   byte alias for unit8
	   complex64, complex128
	   bool
	   string
	   array
	   slice : if you don't declare the size of an array it's called slice
	   map is gthe dict in other languauges
	   strunct is like a class
	   functions are types too
	*/

	var nums = [4]int{1, 2, 3, 4}
	var cities = []string{"A", "B"}
	var salary2 = map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}

	type person struct {
		name string
		age  int
	}
	var me person
	me.name = "Mehrdad"

	var g int = 2
	ptr := &g // pointing to the address of g

	fmt.Println(nums, cities, salary2, me.name, ptr)

	//converting types
	fmt.Println(string(2))
	fmt.Printf("%f", 2.2) //cannot convert float like above

	/*
	  convert string to int float etc:
	  var i = strconv.ParseInt
	  strconv.ParseBool
	  strconv.ParseFloat

	*/

	type xyz = int8
	var ab xyz = 1
	_ = ab

	//conditional stetements
	var condition = 100
	if condition == 90 {
		fmt.Println("90")
	} else if condition > 90 {
		fmt.Println("greater")
	} else {
		fmt.Println("less")
	}

	fmt.Println("OS library showing arguments passed from CLI, the pas and arguments are: ", os.Args) // this is an array, you can reach the first element with os.Args[0]

	if i, err := strconv.Atoi("20"); err == nil { //this type of if initialization is useful for error checking
		fmt.Println("no error, i is ", i)
	} else {
		fmt.Println("error")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//create seperation between output for readability
	fmt.Println(strings.Repeat("#", 20))

	//for is is flexible. continue and break works the same
	//go has no while loop
	i := 10
	for i<15 {
		fmt.Println(i)
		i++
	}
	fmt.Println(strings.Repeat("#", 20))

	//labels are for jumping to other sections of the program used in GoTo, break and continue
	//you can use the same label name in a var, there will be no conflict
	//labels must be used
	friends := [3]string{"A", "B", "C"}
	myfriends := [3]string{"A", "B"}

	outer:
		for index, name := range friends {
			for _, myFriedName := range myfriends{
				if name == myFriedName {
					fmt.Println("friends' index ", index)
					break outer//terminates outer loop
				}
			}
		}//out put is only 0
		fmt.Println(strings.Repeat("#", 20))

		//goto can be used as a loop. notice, you can only use the varibales that were
		//already in the scope when goto appeared, anything after goto cannot be handled
		//todo label should be before the place it's being called
		//goto is discouraged in GoLang too

		i = 0
		loop:
			if i < 5 {
				fmt.Println(i)
				i++
				goto loop
			}
			fmt.Println(strings.Repeat("#", 20))


			//arrays
			var numbers [3]int //initializing to zero
			fmt.Printf("%v", numbers) // output: [0 0 0]
			fmt.Printf("%#v", numbers) // output: [3]int{0, 0, 0}
			fmt.Println(strings.Repeat("#", 20))

			var strings = [4]string{"A", "B"}
			fmt.Printf("%#v", strings) // output: [4]string{"A", "B", "", ""}
			fmt.Println("/n")

			var numbers2 = [...]int{2,4,1,6,7,4} //you can initiate with as many values using ...
			fmt.Println("length is ", len(numbers2))
			fmt.Println("/n")

			var numbers3 = [2][3]int{
				{1,2,3},
				{4,5,6},
			}
			fmt.Println("multi line array is ", numbers3)

			//define array with values at specific index. if you leave index empty it will initialize on its ow
			var numbers4 = [...]int{
				0: 1, //value at index zero is 1
				8: 2,
				3, //index is not provided so it will use previous index + 1
				5: 5, //missed indexes 1,2,4,6 and 7 so they will be initialized to zero

			}
			fmt.Printf("array %d has length %d \n", numbers4, len(numbers4))

			fmt.Println("---------------Slices--------------\n")
			//slices
			var citiez []string
			fmt.Printf("citiez type is : %#v, and the length is %d \n", citiez, len(citiez))
			//citiez[0] = "toronto" //this will fail because slice is empty

			num5 := []int{1,2,3,4}
			

			num6 := make([]int, 2) // make function take a type and length
			fmt.Printf("num6 type is %#v \n", num6)

			num7 := append(num5, num6...) //... is required for concat
			fmt.Printf("concat num5 and num6 = %d \n", num7)

			/*
			num8 := num5
			fmt.Println("is num5 and num8 comparable? ", num8==num5)//this will through an error because slices cannot be compared like this
			*/

			fmt.Println("print part of num7", num7[1:3])//index 1 is included, and 3 is excluded
			fmt.Println("print part of num7", num7[1:]) //[2,3,4,0,0]
			fmt.Println("print part of num7", num7[:3]) //[1,2,3]
			fmt.Println("print part of num7", num7[:]) //entire num7
			fmt.Println("print part of num7", append(num7[:3], 100)) // [1,2,3,100]

			fmt.Printf("num7 before change is : %d\n", num7)
			num9, num10 := num7[1:3], num7[0:2]
			num10[1] = 500//this will modify num9 and num7 
			num11 := []int{}
			num11 = append(num11, num7[0:3]...)
			num7[0] = 999 //num11 won't get modified because it was declared with its own type
			fmt.Printf("num7 is : %d\nnum9 is : %d\nnum10 is : %d\nnum11 is : %d\n", num7, num9, num10, num11)
			fmt.Println(len(num9), cap(num9)) // 2 7 -> length is 2, but backing array has a cap of 5, because num7 has up to index 5
			fmt.Printf("%p, % \n", &num7, &num9) // both will have the same address
			/*
			output:
			num7 before change is : [1 2 3 100 0 0]
			num7 is : [999 500 3 100 0 0]
			num9 is : [500 3]
			num10 is : [999 500]
			num11 is : [1 500 3]
			*/


			n1 := []int{10, 20, 30, 40, 50}
		    n2 := append(n1, 100)
		    n2[0] = 66
		    fmt.Println(n1) // [10 20 30 40 50]


}
