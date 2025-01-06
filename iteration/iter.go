package iteration

const repeatCount int = 5

/*
Provides interface that implements the iteration protocol (in theory),
technically allowing you to loop through an iterable or a sequence
*/
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
