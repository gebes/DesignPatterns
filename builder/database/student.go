package database

type Student struct {
	Firstname string   `json:"firstname,omitempty"` // Hannes
	Surname   string   `json:"surname,omitempty"`   // Mayer
	Class     *Class   `json:"class,omitempty"`     // NewClass("CHIF19", "1.2.3")
	Grades    []*Grade `json:"grades,omitempty"`    // [NewGrade(German, 1), /*...*/]
}
