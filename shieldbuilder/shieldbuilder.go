package shieldbuilder

import (
	"strings"
)

type shield struct {
	front bool
	back  bool
	left  bool
	right bool
}

type shieldBuilder struct {
	code string
}

// NewShieldBuilder constructs a shield builder
func NewShieldBuilder() *shieldBuilder {
	return new(shieldBuilder)
}

func (sh *shieldBuilder) RaiseFront() *shieldBuilder {
	sh.code += "F"
	return sh
}

func (sh *shieldBuilder) RaiseBack() *shieldBuilder {
	sh.code += "B"
	return sh
}

func (sh *shieldBuilder) RaiseLeft() *shieldBuilder {
	sh.code += "L"
	return sh
}

func (sh *shieldBuilder) RaiseRight() *shieldBuilder {
	sh.code += "R"
	return sh
}

func (sh *shieldBuilder) Build() *shield {
	code := sh.code

	return &shield{
		front: strings.Contains(code, "F"),
		back:  strings.Contains(code, "B"),
		left:  strings.Contains(code, "L"),
		right: strings.Contains(code, "R"),
	}
}
