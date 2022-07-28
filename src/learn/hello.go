package learn

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World" // this is ridiculous
	}

	return englishHelloPrefix + name
}
