package hello

const (
	English  = "English"
	Spanish  = "Spanish"
	French   = "French"
	Japanese = "Japanese"
)

const (
	englishHelloPrefix  = "Hello"
	spanishHelloPrefix  = "Hola"
	frenchHelloPrefix   = "Bonjour"
	japaneseHelloPrefix = "こんにちは"
)

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
