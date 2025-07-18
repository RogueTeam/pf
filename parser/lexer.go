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
		{Name: "eol", Pattern: `(\n|\r)+`},
		{Name: "whitespace", Pattern: `(\\|\s|\t)+`},
		// {Name: "Keyword", Pattern: `ruleset-optimization|max-src-conn-rate|tcp\.established|adaptive\.start|other\.multiple|state-defaults|max-src-states|sticky-address|other\.single|table-entries|udp\.multiple|adaptive\.end|divert-packet|max-src-nodes|optimization|tcp\.closing|tcp\.opening|other\.first|max-pkt-rate|max-src-conn|least-states|divert-reply|tcp\.finwait|return-icmp6|state-policy|source-track|high-latency|conservative|block-policy|loginterface|fingerprints|tcp\.tsdiff|source-hash|icmp\.first|round-robin|udp\.single|reliability|received-on|tcp\.closed|probability|icmp\.error|urpf-failed|return-icmp|static-port|return-rst|reassemble|src\.track|aggressive|allow-opts|icmp6-type|udp\.first|throughput|tcp\.first|satellite|divert-to|bandwidth|icmp-type|antispoof|src-nodes|no-route|counters|overload|modulate|floating|fragment|lowdelay|if-bound|interval|synproxy|binat-to|include|no-sync|matches|persist|profile|warning|rdomain|default|bitmask|quantum|timeout|tagged|normal|nat-to|anchor|weight|random|sloppy|states|rdr-to|parent|tables|return|global|rtable|notice|qlimit|debug|proto|pflow|inet6|label|route|flows|burst|match|flags|quick|no-df|state|group|basic|const|emerg|flush|block|queue|scrub|alert|af-to|limit|delay|table|frags|rule|skip|inet|drop|user|prio|none|frag|self|crit|from|once|file|port|pass|load|code|keep|info|min|tag|set|err|yes|log|max|any|ttl|out|for|all|tos|no|in|ms|on|os|to|S|P|F|E|U|R|A|W`},
		{Name: "IPRange", Pattern: IPRange},
		{Name: "CIDR", Pattern: CIDR},
		{Name: "Address", Pattern: AddressExpr},
		{Name: "Ident", Pattern: `\$?[a-zA-Z_](\w|-)*`},
		{Name: "Punct", Pattern: `[-{}()=>!<:,]+`},
		{Name: "Hexnumber", Pattern: `0x[0-9a-fA-F]+`},
		{Name: "Number", Pattern: `\d+`},
		{Name: "Comment", Pattern: `#[^\n]*`},
		{Name: "String", Pattern: `"(\\"|[^"])*"`},
		{Name: "Hostname", Pattern: `[a-zA-Z_]+([a-zA-Z_-]*[a-zA-Z_])?(\.[a-zA-Z_]+([a-zA-Z_-]*[a-zA-Z_]))*`},
		{Name: "Filename", Pattern: `/?[a-zA-Z0-9_\./-]+`},
	},
})
