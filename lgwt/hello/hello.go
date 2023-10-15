package hello

const (
	english  = "English"
	spanish  = "Spanish"
	french   = "French"
	japanese = "Japanese"

	englishHelloPrefix  = "Hello"
	spanishHelloPrefix  = "Hola"
	frenchHelloPrefix   = "Bonjour"
	japaneseHelloPrefix = "こんにちは"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefixFromMap(language) + " " + name
}

var greetingPrefixes = map[string]string{
	english:  englishHelloPrefix,
	spanish:  spanishHelloPrefix,
	french:   frenchHelloPrefix,
	japanese: japaneseHelloPrefix,
}

func greetingPrefixFromMap(language string) (prefix string) {
	prefix, ok := greetingPrefixes[language]
	if !ok {
		prefix = englishHelloPrefix
	}
	return prefix
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}
