package mock

import "fmt"

// Talker ...
type Talker interface {
	SayHello(word string) (response string)
}

// Person ...
type Person struct {
	name string
}

// NewPerson ...
func NewPerson(name string) *Person {
	return &Person{
		name: name,
	}
}

// SayHello ...
func (p *Person) SayHello(name string) (word string) {
	return fmt.Sprintf("Hello %s, welcome come to our company, my name is %s", name, p.name)
}

// Company ...
type Company struct {
	Usher Talker
}

// NewCompany ...
func NewCompany(t Talker) *Company {
	return &Company{
		Usher: t,
	}
}

// Meeting ...
func (c *Company) Meeting(gusetName string) string {
	return c.Usher.SayHello(gusetName)
}
