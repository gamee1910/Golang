package di

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}

func LogError(writer io.Writer, error string){
	fmt.Fprintf(writer, "ERROR: %s", error)
}