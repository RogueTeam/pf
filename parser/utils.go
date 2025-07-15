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

type ValueOrBraceList[T any] struct {
	Values []T `parser:"@@ | ('{' @@ (',' @@)* '}')"`
}

type ValueOrRawList[T any] struct {
	Values []T `parser:"@@ | ( @@ (',' @@)* )"`
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
