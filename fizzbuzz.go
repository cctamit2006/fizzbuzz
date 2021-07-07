package fizzbuzz

// Fizzbuzz takes number and returns fizz and buzz
// Caller must print number if the return is false
// This is a conventional method of commenting on a
// function starting with "Function name: Fizzbuzz"
// i.e. the exported name, even for types
func Fizzbuzz(num int) (string, bool) {
	if num%15 == 0 {
		return "fizzbuzz", true
	}
	if num%3 == 0 {
		return "fizz", true
	}
	if num%5 == 0 {
		return "buzz", true
	}
	return "", false
}
