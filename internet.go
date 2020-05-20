package faker

func (f *Faker) Email() (mail string) {
	check := func(provider *Provider) bool {
		return provider.Internet == nil || len(provider.Internet.FreeEmailDomains) == 0
	}
	p := f.GetProviderWithCheck(check)

	var userNameFormatTemplate, emailFormatTemplate string
	if p.Internet != nil {
		if len(p.Internet.UserNameFormatTemplates) > 0 {
			userNameFormatTemplate = f.RandomStringElement(p.Internet.UserNameFormatTemplates)
		} else {
			userNameFormatTemplate = "{{.LastName}}{{.FirstName}}"
		}

		if p.Internet.EmailFormatTemplate != "" {
			emailFormatTemplate = p.Internet.EmailFormatTemplate
		} else {
			emailFormatTemplate = "{{.UserName}}@{{.EmailDomain}}"
		}
	}

	return f.Format(emailFormatTemplate, map[string]interface{}{
		"UserName": f.Format(userNameFormatTemplate, map[string]interface{}{
			"LastName":  f.PersonFirstNameWithI18nLanguage(I18nLanguageEnUs),
			"FirstName": f.PersonLastNameWithI18nLanguage(I18nLanguageEnUs),
		}),
		"EmailDomain": f.RandomStringElement(p.Internet.FreeEmailDomains),
	})
}
