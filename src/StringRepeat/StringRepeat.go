package main
import "fmt"
import "strings"
 
func main() {
    snow := "blue"
    fmt.Printf("Mary had a little lamb.\nIts fleece as white as %s\n", snow)
    fmt.Println(strings.Repeat(".", 10))
}