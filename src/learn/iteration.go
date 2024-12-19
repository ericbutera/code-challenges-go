package learn

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ { //nolint:intrange
		repeated += character
	}
	return repeated
}
