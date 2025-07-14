package parser

import (
	"net/netip"

	"go4.org/netipx"
)

type (
	Value struct {
		Address    *netip.Addr     `parser:"@Address"`
		CIDR       *netip.Prefix   `parser:"| @CIDR"`
		Range      *netipx.IPRange `parser:"| @IPRange"`
		Identifier *string         `parser:"| @(Ident | String)"`
	}
	Values[T any] struct {
		Values []T `parser:"@@ | ('{' @@ (',' @@)* '}')"`
	}
	Action struct {
		Block   *string     `parser:"( 'block' @('drop' | 'return' | 'return-icmp' | 'return-icmp6' | 'return-rst')?"`
		Pass    *BooleanSet `parser:"| 'pass'"`
		Match   *BooleanSet `parser:"| 'match' )"`
		Options *Options    `parser:"@@?"`
	}
	Options struct {
		Direction *Direction `parser:"@@?"`
		Quick     BooleanSet `parser:"@('quick')?"`
		Log       *Log       `parser:"@@?"`
		// on interface | any
		// on rdomain number
		Family   *AddressFamily `parser:"@@?"`
		Protocol *Protocol      `parser:"@@?"`
		Target   *Target        `parser:"@@?"`
		// allow-opts
		// divert-packet port port
		// divert-reply
		// divert-To
		// flags a/b | any
		// group group
		// icmp-type type [code code]
		// icmp6-type type [code code]
		Label *string `parser:"('label' @String)?"`
		// max-pkt-rate number|seconds
		// once
		// probability number%
		// prio number
		// !received-on interface
		// rtable number
		// set delay milliseconds
		// set prio priority | (priority, priority)
		// set queue queue | (queue, queue)
		// set tos string | number
		// tag string
		// [!]tagged string
		// tos string | number
		// user user
	}
	Log struct {
		IsAll       BooleanSet `parser:"'log' ( @('all')"`
		IsMatches   BooleanSet `parser:"| @('matches')"`
		IsUser      BooleanSet `parser:"@('user')"`
		ToInterface *string    `parser:"| ('to' @(Ident | String)))?"`
	}
	Direction struct {
		Direction string `parser:"@('in' | 'out')"`
	}
	AddressFamily struct {
		Family string `parser:"@('inet' | 'inet6')"`
	}
	Protocol struct {
		Selected string `parser:"'proto' @('icmp' | 'icmp6' | 'tcp' | 'udp')"`
	}
	Port struct {
		Operator *string `parser:"'port' @('!=' | '<=' | '>=' | '<' | '>' | '=')?"`
		Lhs      *string `parser:"@(Number | Ident)"`
		Op       *string `parser:"( @( '<>' | '><' | ':' )"`
		Rhs      *string `parser:"@(Number | Ident) )?"`
	}
	Source struct {
		Any        *BooleanSet     `parser:"( @('any')"`
		NoRoute    *BooleanSet     `parser:"| @('no-route')"`
		Self       *BooleanSet     `parser:"| @('self')"`
		UrpfFailed *BooleanSet     `parser:"| @('urpf-failed')"`
		Route      *string         `parser:"| ('route' @(String | Ident))"`
		Address    *netip.Addr     `parser:"| @Address"`
		CIDR       *netip.Prefix   `parser:"| @CIDR"`
		Range      *netipx.IPRange `parser:"| @IPRange"`
		Identifier *string         `parser:"| @(Ident | String) )"`
		Port       *Port           `parser:"@@?"`
	}
	FromTo struct {
		From *Source `parser:"'from' @@"`
		To   *Source `parser:"'to' @@"`
	}
	Target struct {
		FromTo *FromTo     `parser:"@@"`
		All    *BooleanSet `parser:"| @('all')"`
	}
	OptionBlockPolicy struct {
		Policy string `parser:"'block-policy' @('drop' | 'return')"`
	}
	OptionDebugLevel struct {
		Level string `parser:"'debug' @('emerg' | 'alert' | 'crit' | 'err' | 'warning' | 'notice' | 'info' | 'debug')"`
	}
	OptionTimeout struct {
		Variable string `parser:"'timeout' @Ident"`
		Value    string `parser:""`
	}
	Option struct {
		BlockPolicy *OptionBlockPolicy `parser:"'set' ( @@"`
		Debug       *OptionDebugLevel  `parser:"| @@"`
		// TODO: Others
		Timeout *OptionTimeout `parser:"| @@ )"`
	}
	Assignment[T any] struct {
		Identifier string `parser:"@Ident"`
		Value      T      `parser:"'=' @@"`
	}
	Entry struct {
		Comment    *Comment                   `parser:"@Comment"`
		Assignment *Assignment[Values[Value]] `parser:"| @@"`
		Action     *Action                    `parser:"| @@"`
		Option     *Option                    `parser:"| @@"`
	}
	Configuration struct {
		Entries []*Entry `parser:"@@*"`
	}
)
