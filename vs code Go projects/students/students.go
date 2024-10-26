package students

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) printData() {
	fmt.Printf("Name: %s, Age: %d\n", s.Name, s.Age)
}
