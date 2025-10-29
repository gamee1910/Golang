package helloworld

import "fmt"

const (
	spanish = "Spanish"
	vietnam = "Viet Nam"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	vietnamHelloPrefix = "Xin chào, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language, name)
}

func greetingPrefix(language string, name string) string {
	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case vietnam:
		prefix = vietnamHelloPrefix
	}

	return prefix + name
}
func main() {
	fmt.Println(Hello("world", "english"))
}
