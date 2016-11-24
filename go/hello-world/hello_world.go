package hello

import "fmt"

const testVersion = 2

func HelloWorld(name string) string {
	if name != "" {
		return fmt.Sprintf("Hello, %s!", name)
	} else {
		return "Hello, World!"
	}
}
