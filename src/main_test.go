package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CapitalizeEveryThirdAlphanumericCharTest struct {
	suite.Suite
}

func (s *CapitalizeEveryThirdAlphanumericCharTest) SetupTest() {
}

func (s *CapitalizeEveryThirdAlphanumericCharTest) TestSuccess() {
	s.Run("When is a valid string",
		func() {
			url := "google.com"
			expected := "goOglE.cOm"

			result := CapitalizeEveryThirdAlphanumericChar(url)
			s.Equal(expected, result)
		},
	)
	s.Run("When is a string with UpperCase",
		func() {
			url := "GOOGLE.COM"
			expected := "goOglE.cOm"

			result := CapitalizeEveryThirdAlphanumericChar(url)
			s.Equal(expected, result)
		},
	)
	s.Run("When is a string with numbers",
		func() {
			url := "google.666.com"
			expected := "goOglE.666.Com"

			result := CapitalizeEveryThirdAlphanumericChar(url)
			s.Equal(expected, result)
		},
	)
	s.Run("When is a string with not alphanumerics",
		func() {
			url := "google.?-/.c>om"
			expected := "goOglE.?-/.C>oM"

			result := CapitalizeEveryThirdAlphanumericChar(url)
			s.Equal(expected, result)
		},
	)

}

func TestSuiteMapper(t *testing.T) {
	suite.Run(t, new(CapitalizeEveryThirdAlphanumericCharTest))
}
