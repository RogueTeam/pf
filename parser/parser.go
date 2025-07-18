package parser

import (
	"bytes"
	"io"
	"strings"

	"github.com/alecthomas/participle/v2"
)

var parser = participle.MustBuild[Configuration](
	participle.Lexer(lex),
	participle.Unquote("String"),
)

func init() {
	// fmt.Println(parser.String())
}

func ParseReader(r io.Reader) (conf *Configuration, err error) {
	// file, _ := os.Create(time.Now().String())
	// defer file.Close()
	return parser.Parse(":reader:", r) //	participle.Trace(file),

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
