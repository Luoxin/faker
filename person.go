package faker

func (f *Faker) PersonFirstName() (firstName string) {
	return f.PersonFirstNameWithI18nLanguage(f.Language)
}

func (f *Faker) PersonFirstNameWithI18nLanguage(language I18nLanguage) (firstName string) {
	check := func(provider *Provider) bool {
		return provider.Person == nil || len(provider.Person.FirstNames) == 0
	}

	p := f.GetProviderWithCheckI18nLanguage(language, check)
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
	return f.PersonLastNameWithI18nLanguage(f.Language)
}

func (f *Faker) PersonLastNameWithI18nLanguage(language I18nLanguage) (lastName string) {
	check := func(provider *Provider) bool {
		return provider.Person == nil || len(provider.Person.LastNames) == 0
	}
	p := f.GetProviderWithCheckI18nLanguage(language, check)
	if check(p) {
		return
	}
	return f.RandomStringElement(p.Person.LastNames)
}

func (f *Faker) PersonName() (name string) {
	return f.PersonNameWithI18nLanguage(f.Language)
}

func (f *Faker) PersonNameWithI18nLanguage(language I18nLanguage) (name string) {
	check := func(provider *Provider) bool {
		return provider.Person == nil || len(provider.Person.LastNames) == 0
	}

	p := f.GetProviderWithCheckI18nLanguage(language, check)

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
		nameFormatTemplate = "{{.FirstName}}{{.LastName}}"
	}

	return f.Format(nameFormatTemplate, map[string]interface{}{
		"FirstName": f.RandomStringElement(p.Person.FirstNamesFemale),
		"LastName":  f.RandomStringElement(p.Person.LastNames),
	})
}
