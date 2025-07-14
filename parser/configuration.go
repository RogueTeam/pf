package parser

import (
	"net/netip"

	"go4.org/netipx"
)

type (
	Block struct {
		PolicyOption *string `parser:"'block' @('drop' | 'return' | 'return-icmp' | 'return-icmp6' | 'return-rst')?"`
		Target       *Target `parser:"@@?"`
	}
	Port struct {
		Operator *string `parser:"'port' @('!=' | '<=' | '>=' | '<' | '>' | '=')?"`
		Lhs      *string `parser:"@(Number | Ident)"`
		Op       *string `parser:"( @( '<>' | '><' | ':' )"`
		Rhs      *string `parser:"@(Number | Ident) )?"`
	}
	Host struct {
		Any        *BooleanSet     `parser:"( @('any')"`
		NoRoute    *BooleanSet     `parser:"| @('no-route')"`
		Self       *BooleanSet     `parser:"| @('self')"`
		UrpfFailed *BooleanSet     `parser:"| @('urpf-failed')"`
		Route      *string         `parser:"| ('route' @(String | Ident))"`
		Address    *netip.Addr     `parser:"| @Address"`
		CIDR       *netip.Prefix   `parser:"| @CIDR"`
		Range      *netipx.IPRange `parser:"| @IPRange )"`
		Port       *Port           `parser:"@@?"`
	}
	Target struct {
		From *Host       `parser:"('from' @@"`
		To   *Host       `parser:" 'to' @@)"`
		All  *BooleanSet `parser:"| @('all')"`
	}
	Entry struct {
		Comment *Comment `parser:"@Comment"`
		Block   *Block   `parser:"| @@"`
	}
	Configuration struct {
		Entries []*Entry `parser:"@@*"`
	}
)
