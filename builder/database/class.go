package database

type Class struct {
	Name string `json:"name,omitempty"`
	Room string `json:"room,omitempty"`
}

func NewClass(name string, room string) *Class {
	return &Class{Name: name, Room: room}
}

func GetClassByName(name string) (*Class, error) {

	/* database query operation ...
	   SELECT * FROM class WHERE class.name = 'NAME'
	*/

	return NewClass(name, "1.2.3"), nil
}
