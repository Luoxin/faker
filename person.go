package faker

type Person struct {
	f *Faker
}

func NewPerson() *Person {
	return &Person{}
}

func NewPersonWithFaker(f *Faker) *Person {
	p := NewPerson()
	p.SetFaker(f)
	return p
}

func (p *Person) SetFaker(f *Faker) {
	p.f = f
}

func (p Person) FirstNameMale() string {
	index := p.f.IntBetween(0, len(firstNameMale)-1)
	return firstNameMale[index]
}

func (p Person) FirstNameFemale() string {
	index := p.f.IntBetween(0, len(firstNameFemale)-1)
	return firstNameFemale[index]
}

func (p Person) FirstName() string {
	names := append(firstNameMale, firstNameFemale...)
	return p.Faker.RandomStringElement(names)
}

func (p Person) LastName() string {
	index := p.f.IntBetween(0, len(lastName)-1)
	return lastName[index]
}
