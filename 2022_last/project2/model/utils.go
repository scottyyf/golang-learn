package model12

type person struct {
	Name string
	age uint8
	salary float32
}

func NewPerson(name string) *person{
	return &person{Name: name}
}

func (p *person) SetAge(age uint8)  {
	if age >0 && age < 150{
		p.age = age
	}else {
		p.age = 111
	}
}

func (p *person) GetAge() uint8  {
	return p.age
}

//func (p person) name()  {
//
//}
