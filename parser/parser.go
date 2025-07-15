package parser

import (
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/alecthomas/participle/v2"
)

var parser = participle.MustBuild[Configuration](
	participle.Lexer(lex),
	participle.Unquote("String"),
)

func ParseReader(r io.Reader) (conf *Configuration, err error) {
	return parser.Parse(":reader:", r, participle.Trace(os.Stdout))
}

func ParseContent[T string | ~[]byte](b T) (conf *Configuration, err error) {
	asAny := any(b)

	var r io.Reader
	if _, ok := asAny.(string); ok {
		r = strings.NewReader(asAny.(string))
	} else {
		r = bytes.NewReader(asAny.([]byte))
	}

	return ParseReader(r)
}
