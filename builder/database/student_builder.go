package database

type StudentBuilder struct {
	Firstname string   // Hannes
	Surname   string   // Mayer
	Class     *Class   // NewClass("CHIF19", "1.2.3")
	Grades    []*Grade // [NewGrade(German, 1), /*...*/]
}

func NewStudentBuilder() *StudentBuilder {
	return &StudentBuilder{}
}

func (b *StudentBuilder) SetFirstname(firstname string) *StudentBuilder {
	b.Firstname = firstname
	return b
}
func (b *StudentBuilder) SetSurname(surname string) *StudentBuilder {
	b.Surname = surname
	return b
}
func (b *StudentBuilder) SetClass(class *Class) *StudentBuilder {
	b.Class = class
	return b
}
func (b *StudentBuilder) SetGrades(grades []*Grade) *StudentBuilder {
	b.Grades = grades
	return b
}

func (b *StudentBuilder) Build() *Student {
	return &Student{
		Firstname: b.Firstname,
		Surname:   b.Surname,
		Class:     b.Class,
		Grades:    b.Grades,
	}
}
