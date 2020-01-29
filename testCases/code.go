package testCases

func AddNumbers(a ...int) int {
	result := 0
	for _, number := range a {
		result = result + number
	}
	return result
}

type readFileDS interface {
	ReadFile(string) ([]byte, error)
}

func ReadFile(filePath string, ds readFileDS) ([]byte, error) {

	return ds.ReadFile(filePath)
}
