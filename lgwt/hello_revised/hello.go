package hello

const (
	// Languages Supported
	English  = "English"
	Spanish  = "Spanish"
	French   = "French"
	Japanese = "Japanese"
	// Hello greeting for each language
	englishHelloPrefix  = "Hello"
	spanishHelloPrefix  = "Hola"
	frenchHelloPrefix   = "Bonjour"
	japaneseHelloPrefix = "こんにちは"
)

// Hello provides a greeting to a specified name in a specified language.
// Default language is 'English' and default name is 'World'.
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := getHelloPrefix(language)

	return prefix + " " + name
}

func getHelloPrefix(language string) (prefix string) {
	return getHelloPrefixSwitch(language)
}

func getHelloPrefixSwitch(language string) (prefix string) {
	switch language {
	case English:
		prefix = englishHelloPrefix
	case Spanish:
		prefix = spanishHelloPrefix
	case French:
		prefix = frenchHelloPrefix
	case Japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

var helloPrefixes = map[string]string{
	English:  englishHelloPrefix,
	Spanish:  spanishHelloPrefix,
	French:   frenchHelloPrefix,
	Japanese: japaneseHelloPrefix,
}

func getHelloPrefixMap(language string) (prefix string) {
	prefix, ok := helloPrefixes[language]
	if !ok {
		return englishHelloPrefix
	}
	return prefix
}
