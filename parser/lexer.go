package parser

import (
	"errors"
	"strings"

	"github.com/alecthomas/participle/v2/lexer"
)

type Comment string

func (c *Comment) Capture(values []string) error {
	if len(values) == 0 {
		return errors.New("expecting value for comment")
	}

	raw := values[0]
	*c = Comment(strings.TrimSpace(raw[1:]))
	return nil
}

const (
	IPv4Expr    = "(" + `\d{1,3}(\.\d{1,3}){3}` + ")"
	IPv6Expr    = "(" + "::" + ")"
	AddressExpr = "(" + IPv4Expr + "|" + IPv6Expr + ")"
	CidrExpr    = AddressExpr + `/\d{1,3}`
	IPRange     = AddressExpr + `-` + AddressExpr
)

var lex = lexer.MustSimple([]lexer.SimpleRule{
	{Name: "CIDR", Pattern: CidrExpr},
	{Name: "IPRange", Pattern: IPRange},
	{Name: "Address", Pattern: AddressExpr},
	{Name: "Comment", Pattern: `#[^\n]*`},
	{Name: "String", Pattern: `"(\\"|[^"])*"`},
	{Name: "Number", Pattern: `[-+]?(\d*\.)?\d+`},
	{Name: "Ident", Pattern: `[a-zA-Z_](\w|-)*`},
	{Name: "Punct", Pattern: `[-=:<>!]+`},
	{Name: "eol", Pattern: `[\n\r]+`},
	{Name: "whitespace", Pattern: `[ \t]+`},
})
