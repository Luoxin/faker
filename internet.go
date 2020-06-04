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

func (f *Faker) Image(width, height uint32) string {
	const (
		maxWidth  = 1024
		maxHeight = 1024
	)

	check := func(provider *Provider) bool {
		return provider.Internet == nil || len(provider.Internet.ImagePlaceholderServiceTemplateList) == 0
	}
	p := f.GetProviderWithCheck(check)
	if check(p) {
		return ""
	}

	if width <= 0 || width > maxWidth {
		width = uint32(f.IntBetween(1, maxWidth))
	}

	if height <= 0 || height > maxHeight {
		height = uint32(f.IntBetween(1, maxHeight))
	}

	return f.Format(f.RandomStringElement(p.Internet.ImagePlaceholderServiceTemplateList), map[string]interface{}{
		"WIDTH":  width,
		"HEIGHT": height,
	})
}
