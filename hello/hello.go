package hello

import "fmt"

func Hello(name string) string {
	greeting := fmt.Sprintf("Hello, %s!", name)
	return greeting
}
