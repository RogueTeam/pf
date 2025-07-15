package parser

import (
	"github.com/alecthomas/participle/v2/lexer"
)

const (
	IPv4Expr    = "(" + `\d{1,3}(\.\d{1,3}){3}` + ")"
	IPv6Expr    = "(" + "::" + ")"
	AddressExpr = "(" + IPv4Expr + "|" + IPv6Expr + ")"
)

var lex = lexer.MustSimple([]lexer.SimpleRule{
	{Name: "Address", Pattern: AddressExpr},
	{Name: "Comment", Pattern: `#[^\n]*`},
	{Name: "String", Pattern: `"(\\"|[^"])*"`},
	{Name: "Ident", Pattern: `\$?[a-zA-Z_](\w|-)*`},
	{Name: "Hostname", Pattern: `[a-zA-Z_]+([a-zA-Z_-]*[a-zA-Z_])?(\.[a-zA-Z_]+([a-zA-Z_-]*[a-zA-Z_]))*`},
	{Name: "Filename", Pattern: `/?[a-zA-Z0-9_./-]+`},
	{Name: "Number", Pattern: `[-+]?(\d*\.)?\d+`},
	{Name: "Punct", Pattern: `[-=:<>!]+`},
	{Name: "eol", Pattern: `(\n|\r)+`},
	{Name: "whitespace", Pattern: `(\\|\s|\t)+`},
})
