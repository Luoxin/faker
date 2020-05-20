package faker

func (f *Faker) PhoneNumber() (phoneNumber string) {
	check := func(provider *Provider) bool {
		return provider.PhoneNumber == nil || len(provider.PhoneNumber.PhonePrefixes) == 0
	}
	p := f.GetProviderWithCheck(check)
	if p.PhoneNumber == nil {
		return
	}

	formatTemplate := p.PhoneNumber.PhoneFormatTemplate
	if formatTemplate == "" {
		formatTemplate = "###-###-###"
	}

	args := map[string]interface{}{}

	if p.PhoneNumber != nil {
		args["PhonePrefix"] = f.RandomStringElement(p.PhoneNumber.PhonePrefixes)
	}

	return f.Format(formatTemplate, args)
}
