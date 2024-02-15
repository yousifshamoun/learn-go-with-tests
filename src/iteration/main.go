package iteration

func Repeat(character string, iterations int) string {
	var repeated string
	if iterations == 0 {
		iterations = 5
	}
	for i := 0; i < iterations; i++ {
		repeated += character
	}
	return repeated
}

