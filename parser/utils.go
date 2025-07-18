package parser

import (
	"errors"
	"strings"
)

type BooleanSet bool

func (b *BooleanSet) Capture(values []string) error {
	*b = len(strings.TrimSpace(values[0])) > 0
	return nil
}

type Variable struct {
	Name string `parser:"@Variable"`
}

type String struct {
	Value string `parser:"@String"`
}

type Text struct {
	Value string `parser:"@(String | Ident | Hostname | Filename)"`
}

type Number struct {
	Value int `parser:"@Number"`
}

type Value[T any] struct {
	Direct      *T        `parser:"@@"`
	Variable    *Variable `parser:"| @@"`
	Parentheses *Value[T] `parser:"| '(' @@ ')'"`
}

type Parentheses[T any] struct {
	Value T `parser:"('(' @@ ')') | @@"`
}

type ValueOrBraceList[T any] struct {
	Values []Value[T] `parser:"@@ | ('{' @@ (',' @@)* '}')"`
}

type ValueOrRawList[T any] struct {
	Values []Value[T] `parser:"@@ (','? @@)*"`
}

type Comment string

func (c *Comment) Capture(values []string) error {
	if len(values) == 0 {
		return errors.New("expecting value for comment")
	}

	raw := values[0]
	*c = Comment(strings.TrimSpace(raw[1:]))
	return nil
}
