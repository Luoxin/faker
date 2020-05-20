package faker

func (f *Faker) PersonFirstName() (firstName string) {
	check := func(provider *Provider) bool {
		return provider.Person == nil || len(provider.Person.FirstNames) == 0
	}
	p := f.GetProviderWithCheck(check)
	if check(p) {
		return
	}

	return f.RandomStringElement(p.Person.FirstNames)
}

func (f *Faker) PersonFirstNameMale() (firstName string) {
	check := func(provider *Provider) bool {
		return provider.Person == nil || len(provider.Person.FirstNamesMale) == 0
	}
	p := f.GetProviderWithCheck(check)
	if check(p) {
		return
	}
	return f.RandomStringElement(p.Person.FirstNames)
}

func (f *Faker) PersonFirstNameFemale() (lastName string) {
	check := func(provider *Provider) bool {
		return provider.Person == nil || len(provider.Person.FirstNamesFemale) == 0
	}

	p := f.GetProviderWithCheck(check)

	if check(p) {
		return
	}

	return f.RandomStringElement(p.Person.FirstNames)
}

func (f *Faker) PersonLastName() (lastName string) {
	check := func(provider *Provider) bool {
		return provider.Person == nil || len(provider.Person.LastNames) == 0
	}
	p := f.GetProviderWithCheck(check)
	if check(p) {
		return
	}
	return f.RandomStringElement(p.Person.LastNames)
}

func (f *Faker) PersonName() (name string) {
	check := func(provider *Provider) bool {
		return provider.Person == nil || len(provider.Person.FirstNames) == 0 || len(provider.Person.LastNames) == 0
	}
	p := f.GetProviderWithCheck(check)
	if check(p) {
		return
	}

	nameFormatTemplate := p.Person.NameFormatTemplate
	if nameFormatTemplate == "" {
		nameFormatTemplate = "{{FirstName}}{{LastName}}"
	}

	return f.Format(nameFormatTemplate, map[string]interface{}{
		"FirstName": f.RandomStringElement(p.Person.FirstNames),
		"LastName":  f.RandomStringElement(p.Person.LastNames),
	})
}

func (f *Faker) PersonNameMale() (name string) {
	check := func(provider *Provider) bool {
		return provider.Person == nil || len(provider.Person.FirstNamesMale) == 0 || len(provider.Person.LastNames) == 0
	}
	p := f.GetProviderWithCheck(check)
	if check(p) {
		return
	}

	nameFormatTemplate := p.Person.NameFormatTemplate
	if nameFormatTemplate == "" {
		nameFormatTemplate = "{{FirstName}}{{LastName}}"
	}

	return f.Format(nameFormatTemplate, map[string]interface{}{
		"FirstName": f.RandomStringElement(p.Person.FirstNamesMale),
		"LastName":  f.RandomStringElement(p.Person.LastNames),
	})
}

func (f *Faker) PersonNameFemale() (name string) {
	check := func(provider *Provider) bool {
		return provider.Person == nil || len(provider.Person.FirstNamesFemale) == 0 || len(provider.Person.LastNames) == 0
	}
	p := f.GetProviderWithCheck(check)
	if check(p) {
		return
	}

	nameFormatTemplate := p.Person.NameFormatTemplate
	if nameFormatTemplate == "" {
		nameFormatTemplate = "{{FirstName}}{{LastName}}"
	}

	return f.Format(nameFormatTemplate, map[string]interface{}{
		"FirstName": f.RandomStringElement(p.Person.FirstNamesFemale),
		"LastName":  f.RandomStringElement(p.Person.LastNames),
	})
}
