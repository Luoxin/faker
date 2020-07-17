package faker

func (f *Faker) UserAgent() (ua string) {
	return f.UserAgentWithLanguage(f.Language)
}

func (f *Faker) UserAgentWithLanguage(language I18nLanguage) (ua string) {
	check := func(provider *Provider) bool {
		return provider.Internet == nil || len(provider.Internet.UserAgent) == 0
	}

	p := f.GetProviderWithCheckI18nLanguage(language, check)
	if check(p) {
		return "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36"
	}

	return f.RandomStringElement(p.Internet.UserAgent)
}
