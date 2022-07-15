package mapper

import (
	"errors"
	"strings"
	"unicode"
)

type skipString struct {
	step           uint
	value          string
	valueColection []rune
}

func NewSkipString(st uint, s string) (*skipString, error) {
	if st == 0 {
		return nil, errors.New("Step Invalid")
	}
	return &skipString{
		step:  st,
		value: strings.ToLower(s),
	}, nil
}

func (ss *skipString) TransformRune(pos int) {
	if uint(pos)%ss.step == (ss.step - 1) {
		ss.valueColection[pos] = unicode.To(unicode.UpperCase, ss.valueColection[pos])
	}
}

func (ss *skipString) GetValueAsRuneSlice() []rune {
	ss.valueColection = []rune(ss.value)
	return ss.valueColection
}

func (ss skipString) String() string {
	return string(ss.valueColection)
}
