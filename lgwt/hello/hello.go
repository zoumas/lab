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

// Hello returns a greeting to a specified name in a specified language.
// An empty name defaults to 'World'.
// An empty language defaults to 'English'.
func Hello(name, language string) (greeting string) {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + " " + name
}

func greetingPrefix(language string) (prefix string) {
	return greetingPrefixSwitch(language)
}

// 100 ns/op
func greetingPrefixSwitch(language string) (prefix string) {
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
	return prefix
}

var greetingPrefixes = map[string]string{
	English:  englishHelloPrefix,
	Spanish:  spanishHelloPrefix,
	French:   frenchHelloPrefix,
	Japanese: japaneseHelloPrefix,
}

// 130 ns/op
func greetingPrefixMap(language string) (prefix string) {
	prefix, ok := greetingPrefixes[language]
	if !ok {
		return englishHelloPrefix
	}
	return prefix
}
