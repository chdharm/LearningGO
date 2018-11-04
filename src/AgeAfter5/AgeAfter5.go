package main
import "fmt"

// one verison all in main...
// another version uses a different func

func main() {
    var name string
    var age int
    var age_back int
    var age_plus int

    fmt.Println("Hello. What is your name?")
    fmt.Scanf("%s", &name)
    fmt.Printf("Hi %s! How old are you? ", name)
    fmt.Scanf("%d", &age)

    age_back = age - 5
    fmt.Printf("Did you know that five years ago, you were be %d years old?\n", age_back)
    age_plus = age + 5
    fmt.Printf("...and that five years from now, you will be %d years old. Imagine that!\n", age_plus)
}