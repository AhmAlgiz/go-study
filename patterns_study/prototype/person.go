package person

type Person struct {
	name    string
	age     int
	address string
}

func (p *Person) Copy() *Person {
	return &Person{
		name:    p.name,
		age:     p.age,
		address: p.address,
	}
}

var defaultPerson = &Person{
	name:    "undefined",
	age:     0,
	address: "undefined",
}

func Create(name string, age int, address string) *Person {
	newPerson := defaultPerson.Copy()
	newPerson.name = name
	newPerson.age = age
	newPerson.address = address
	return newPerson
}
