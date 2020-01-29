package testCases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//test case with subtest cases
func TestAddNumbers(t *testing.T) {

	assert.Equal(t, 10, AddNumbers(5, 5))

	t.Run("firstTest", func(t *testing.T) {
		assert.Equal(t, 10, 10)
	})

	t.Run("secondTest", func(t *testing.T) {
		assert.Equal(t, 10, 1)
	})

}

func TestWithRequire(t *testing.T) {

	assert.Equal(t, 10, AddNumbers(5, 5))

	t.Run("firstTest", func(t *testing.T) {
		require.Equal(t, 10, 20) //it will return from this line
		assert.Equal(t, 10, 10)  // this line is not executed beacause require condition is failed
	})

	t.Run("secondTest", func(t *testing.T) {
		//executed both statement
		assert.Equal(t, 10, 1)
		assert.Equal(t, 10, 2)
	})

}

//run single test case : go  test -v -run TestCaseFileMocking

type MockFile struct {
	Data     []byte
	Err      error
	FileName string
}

// pointer type
func (m *MockFile) ReadFile(filePath string) ([]byte, error) {

	m.FileName = filePath
	if m.Err != nil {
		return nil, m.Err
	}

	return m.Data, nil
}

//test case: for reading file error
func TestCaseFileMocking(t *testing.T) {

	t.Run("Error Case-ReadFile", func(t *testing.T) {
		m := MockFile{}

		m.Err = errors.New("READ_ERROR")

		//check for error
		byteArray, err := ReadFile("./config.json", &m)

		assert.Error(t, err)
		assert.Nil(t, byteArray)
		assert.EqualError(t, err, "READ_ERROR")

		//to check filenpath
		assert.Equal(t, "./config.json", m.FileName)
	})

	t.Run("Success Case-ReadFile", func(t *testing.T) {
		m := MockFile{}

		m.Data = []byte("kunal")

		byteArray, err := ReadFile("./config.json", &m)

		if assert.NoError(t, err) {
			assert.Equal(t, []byte("kunal"), byteArray)
		}

	})

}

func TestReadFile(t *testing.T) {
	type ip struct {
		filePath string
		ds       readFileDS
	}

	type op struct {
		bs  []byte
		err error
	}
	tests := []struct {
		name   string
		input  ip
		output op
	}{
		{"Error Case-ReadFile",
			ip{"./config.json", &MockFile{Err: assert.AnError}},
			op{nil, assert.AnError},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile(tt.input.filePath, tt.input.ds)
			assert.Equal(t, tt.output.bs, got)
			assert.Equal(t, tt.output.err, err)
		})
	}
}
