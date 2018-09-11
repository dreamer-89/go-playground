package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/radlinskii/go-playground/utils/file_utils"
	"github.com/radlinskii/go-playground/utils/number_utils"
	"github.com/radlinskii/go-playground/utils/string_utils"
)

func main() {
	fmt.Printf("reversed message: %s\n", string_utils.Reverse("!oG ,olleH"))

	fmt.Printf("uppercase message: %s\n", string_utils.ToUpperCase("Hello, Go!"))

	var num float64 = -3
	sqrt, err := number_utils.Sqrt(num)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Printf("my sqrt of %g: %g\n", num, sqrt)
	}

	fmt.Printf("is %d a odd number? %t\n", int(num), number_utils.IsOdd(int(num)))

	file_utils.Copy("./file_utils/new_file.txt", "./file_utils/file.txt")

	fmt.Println(string_utils.WordCount("Hello World!! World! I said Hello !!\n"))

	number_utils.PrintFibonacci()

	var name string
	flag.StringVar(&name, "name", "", "Who should i greet?")
	flag.Parse()
	if name == "" {
		fmt.Println("Hello World!")
	} else {
		fmt.Printf("Hello %s!\n", name)
	}

	if len(os.Args) == 2 {
		file_utils.Write(os.Args[1])
	}

	ipAddr := string_utils.IPAddr{192, 168, 182, 157}
	fmt.Printf("stringified ip address type: %v\n", ipAddr)

	s := strings.NewReader("Uryyb Jbeyq!\n")
	r := string_utils.Rot13Reader{s}
	fmt.Print("decoded message: ")
	io.Copy(os.Stdout, &r)

	file_utils.RemoveTabs("./file_utils/file.txt")
}
