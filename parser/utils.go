package parser

import "strings"

type BooleanSet bool

func (b *BooleanSet) Capture(values []string) error {
	*b = len(strings.TrimSpace(values[0])) > 0
	return nil
}
