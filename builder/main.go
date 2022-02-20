package main

import (
	"DesignPatterns/builder/database"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

func main() {

	studentBuilder := database.NewStudentBuilder().
		SetFirstname("Hannes").
		SetSurname("Mayer")

	class, err := database.GetClassByName("CHIF19")
	if err != nil {
		panic(errors.Wrap(err, "could not find class"))
	}
	studentBuilder.SetClass(class)

	studentBuilder.SetGrades([]*database.Grade{
		database.NewGrade(database.German, 1),
		database.NewGrade(database.English, 1),
		database.NewGrade(database.Programming, 5),
	})

	student := studentBuilder.Build()
	PrintStudent(student)
	PrintStudentGPA(student)
}

func PrintStudent(student *database.Student) {
	j, err := json.MarshalIndent(student, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}

func PrintStudentGPA(student *database.Student) {
	fmt.Print("Student " + student.Firstname + " got the grades ")

	sum := 0

	for i, grade := range student.Grades {
		if i == len(student.Grades)-1 {
			fmt.Print(grade.Grade, " ")
		} else {
			fmt.Print(grade.Grade, ", ")
		}
		sum += grade.Grade
	}

	fmt.Printf("which has an average of %0.2f\n", float64(sum)/float64(len(student.Grades)))
}
