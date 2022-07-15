package mapper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MapperTest struct {
	suite.Suite
}

func (s *MapperTest) TestSuccess() {
	s.Run("When is a valid string",
		func() {
			skipString, _ := NewSkipString(3, "google.com")
			expected := "goOglE.cOm"

			MapString(skipString)
			result := fmt.Sprint(skipString)
			s.Equal(expected, result)
		},
	)
	s.Run("When is a diff step",
		func() {
			skipString, _ := NewSkipString(2, "google.com")
			expected := "gOoGlE.CoM"

			MapString(skipString)
			result := fmt.Sprint(skipString)
			s.Equal(expected, result)
		},
	)
	s.Run("When is a step equal to 1",
		func() {
			skipString, _ := NewSkipString(1, "google.com")
			expected := "GOOGLE.COM"

			MapString(skipString)
			result := fmt.Sprint(skipString)
			s.Equal(expected, result)
		},
	)
	s.Run("When is a string is UpperCase",
		func() {
			skipString, _ := NewSkipString(3, "GOOGLE.COM")
			expected := "goOglE.cOm"

			MapString(skipString)
			result := fmt.Sprint(skipString)
			s.Equal(expected, result)
		},
	)
	s.Run("When is a string with numbers",
		func() {
			skipString, _ := NewSkipString(3, "google.666.com")
			expected := "goOglE.666.Com"

			MapString(skipString)
			result := fmt.Sprint(skipString)
			s.Equal(expected, result)
		},
	)
	s.Run("When is a string without alphanumerics",
		func() {
			skipString, _ := NewSkipString(3, "google.?-/.c>om")
			expected := "goOglE.?-/.C>oM"

			MapString(skipString)
			result := fmt.Sprint(skipString)
			s.Equal(expected, result)
		},
	)
	s.Run("When is a valid empty string",
		func() {
			skipString, _ := NewSkipString(3, "")
			expected := ""
			MapString(skipString)
			result := fmt.Sprint(skipString)
			s.Equal(expected, result)
		},
	)
}

func (s *MapperTest) TestErrors() {
	s.Run("When is a invalid step 0",
		func() {
			_, err := NewSkipString(0, "google.com")
			s.Error(err)
		},
	)
}

func TestSuiteMapper(t *testing.T) {
	suite.Run(t, new(MapperTest))
}
