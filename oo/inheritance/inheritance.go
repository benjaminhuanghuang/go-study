package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (person Person) getNameAge() (string, int) {
	return person.name, person.age
}

type Student struct {
	Person
	class_id   int
	speciality string
}

func (student Student) getSpeciality() string {
	return student.speciality
}

func main() {
	s := new(Student)
	s.name = "Ben"
	s.age = 100

	name, age := s.getNameAge()

	fmt.Println(name, age)

	fmt.Println(" ", s.getSpeciality())
}
