package main

import (
    "fmt"
)

// common interface
type Person interface {
    GetFullName() string
}

// university teacher
type Teacher struct {
    FirstName string
    LastName  string
    Department string
	Salary float64
}

// getting the full name of teacher
func (t Teacher) GetFullName() string {
    return t.FirstName + " " + t.LastName
}
//bonus for teachers 
func (t Teacher) CalculateBonus() float64 {
    return t.Salary * 0.1 
}

// university students
type Student struct {
    FirstName string
    LastName  string
    StudentID string
	GPA      float64
}
//checking isHonorStudent or not 
func (s Student) IsHonorStudent() bool {
    return s.GPA > 3.5
}
type CourseGrade struct {
    CourseName string
    Grade      float64
}
// full name and last name student
func (s Student) GetFullName() string {
    return s.FirstName + " " + s.LastName +" " +s.StudentID
}

// all about government organization)
type GovernmentAttachment struct {
    Person     Person
    Position   string
    Department string

}

// government person full name 
func (ga GovernmentAttachment) GetFullName() string {
    return ga.Person.GetFullName()
}

func main() {
    teacher  := Teacher{FirstName: "Daulet", LastName: "Di", Department: "Computer Science"}
	teacher1 := Teacher{FirstName: "Aruzhan", LastName: "Keulenzhanova", Department: "Information systems"}
    student1 := Student{FirstName: "Jelly", LastName: "Fish", StudentID: "12345"}
	student2 := Student{FirstName: "Aru", LastName: "Jelly", StudentID: "21010"}

    attachment1 := GovernmentAttachment{Person: teacher, Position: "Lecturer", Department: "Education"}
    attachment2 := GovernmentAttachment{Person: student1, Position: "Intern", Department: "Government Agency"}
	attachment3 := GovernmentAttachment{Person: student2, Position: "Intern", Department: "Club Fair"}
	attachment4 := GovernmentAttachment{Person: teacher1, Position: "Practice", Department: "ACM"}

    people1 := []Person{teacher, student1, attachment1, attachment2}
	people2 := []Person{teacher1, student2, attachment3, attachment4}

    for _, person := range people1 {
        fmt.Printf("Name: %s\n", person.GetFullName())

        // Checking for gov attachment
        if attachment, ok := person.(GovernmentAttachment); ok {
            fmt.Printf("Position: %s, Department: %s\n", attachment.Position, attachment.Department)
        }
        fmt.Println()
    }
	for _, person := range people2 {
        fmt.Printf("Name: %s\n", person.GetFullName())

        if attachment, ok := person.(GovernmentAttachment); ok {
            fmt.Printf("Position: %s, Department: %s\n", attachment.Position, attachment.Department)
        }
        switch p := person.(type) {
        case Teacher:
            bonus := p.CalculateBonus()
            fmt.Printf("Bonus: $%.2f\n", bonus)
        case Student:
            if p.IsHonorStudent() {
                fmt.Println("This student is an honor student.")
            }
        }
        fmt.Println()
	}
}
