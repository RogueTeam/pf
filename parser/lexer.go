package parser

import (
	"github.com/alecthomas/participle/v2/lexer"
)

const (
	IPv4Expr    = "(" + `\d{1,3}(\.\d{1,3}){3}` + ")"
	IPv6Expr    = "(" + "::::" + ")"
	AddressExpr = "(" + IPv4Expr + "|" + IPv6Expr + ")"
	IPRange     = "(" + AddressExpr + "-" + AddressExpr + ")"
	CIDR        = "(" + AddressExpr + `/\d{1,3})`
)

var lex = lexer.MustStateful(lexer.Rules{
	"Root": []lexer.Rule{
		// Matching
		{Name: "EOL", Pattern: `(\n|\r)+`},
		{Name: "IPRange", Pattern: IPRange},
		{Name: "CIDR", Pattern: CIDR},
		{Name: "Address", Pattern: AddressExpr},
		{Name: "Variable", Pattern: `\$[a-zA-Z_](\w|-)*`},
		{Name: "Ident", Pattern: `[a-zA-Z_](\w|-)*`},
		{Name: "Hexnumber", Pattern: `0x[0-9a-fA-F]+`},
		{Name: "Number", Pattern: `\d+`},
		{Name: "Punct", Pattern: `[-{}()=>!<:,/]+`},
		{Name: "Comment", Pattern: `#[^\n]*`},
		{Name: "String", Pattern: `"(\\"|[^"])*"`},
		{Name: "Hostname", Pattern: `[a-zA-Z_]+([a-zA-Z_-]*[a-zA-Z_])?(\.[a-zA-Z_]+([a-zA-Z_-]*[a-zA-Z_]))*`},
		{Name: "Filename", Pattern: `/?[a-zA-Z0-9_\./-]+`},
		// Not matching
		{Name: "whitespace", Pattern: `(\\\n| |\t)+`},
	},
})
