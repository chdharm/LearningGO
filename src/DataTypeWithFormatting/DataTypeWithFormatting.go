package main
import "fmt"

func main() {
    var first_name string
    var last_name string
    var grade int
    var student_id int
    var login string
    var gpa float64

    fmt.Println("Please enter the following information:")
    fmt.Print("First name: ")
    fmt.Scanf("%s", &first_name)
    fmt.Print("Last name: ")
    fmt.Scanf("%s", &last_name)
    fmt.Print("Grade (9 - 12): ")
    fmt.Scanf("%d", &grade)
    fmt.Print("Student ID: ")
    fmt.Scanf("%d", &student_id)
    fmt.Print("Log-in: ")
    fmt.Scanf("%s", &login)
    fmt.Print("GPA (0.0 - 4.0): ")
    fmt.Scanf("%f", &gpa)

    fmt.Println("Your information:")
    fmt.Printf("\tLogin: %s\n", login)
    fmt.Printf("\tID: %s\n", student_id)
    fmt.Printf("\tName: %s %s\n", first_name, last_name)
    fmt.Printf("\tGPA: %f\n", gpa)
    fmt.Printf("\tGrade: %d\n", grade)
}