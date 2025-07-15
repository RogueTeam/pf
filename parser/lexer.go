package parser

import (
	"github.com/alecthomas/participle/v2/lexer"
)

const (
	IPv4Expr    = "(" + `\d{1,3}(\.\d{1,3}){3}` + ")"
	IPv6Expr    = "(" + "::::" + ")"
	AddressExpr = "(" + IPv4Expr + ")" // "|" + IPv6Expr + ")"
	IPRange     = "(" + AddressExpr + "-" + AddressExpr + ")"
	CIDR        = "(" + AddressExpr + `/\d{1,3})`
)

var lex = lexer.MustStateful(lexer.Rules{
	"Root": []lexer.Rule{
		{Name: "IPRange", Pattern: IPRange},
		{Name: "CIDR", Pattern: CIDR},
		{Name: "Address", Pattern: AddressExpr},
		{Name: "Ident", Pattern: `\$?[a-zA-Z_](\w|-)*`},
		{Name: "Punct", Pattern: `[-{}()=>!<:,]+`},
		{Name: "Number", Pattern: `\d+`},
		{Name: "Comment", Pattern: `#[^\n]*`},
		{Name: "String", Pattern: `"(\\"|[^"])*"`},
		{Name: "Hostname", Pattern: `[a-zA-Z_]+([a-zA-Z_-]*[a-zA-Z_])?(\.[a-zA-Z_]+([a-zA-Z_-]*[a-zA-Z_]))*`},
		{Name: "Filename", Pattern: `/?[a-zA-Z0-9_\./-]+`},
		{Name: "eol", Pattern: `(\n|\r)+`},
		{Name: "whitespace", Pattern: `(\\|\s|\t)+`},
	},
})
