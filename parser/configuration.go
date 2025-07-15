package parser

import "net/netip"

type (
	Timeout struct {
		Variable string  `parser:"@('tcp.first' | 'tcp.opening' | 'tcp.established' | 'tcp.closing' | 'tcp.finwait' | 'tcp.closed' | 'tcp.tsdiff' | 'udp.first' | 'udp.single' | 'udp.multiple' | 'icmp.first' | 'icmp.error' | 'other.first' | 'other.single' | 'other.multiple' | 'frag' | 'interval' | 'src.track' | 'adaptive.start' | 'adaptive.end')"`
		Value    float64 `parser:"@Number"`
	}
	TimeoutOption struct {
		ValueOrBraceList[Timeout] `parser:"'timeout' @@"`
	}
	RulesetOptimizationOption struct {
		Value string `parser:"'ruleset-optimization' @('none' | 'basic' | 'profile')"`
	}
	OptimizatioOption struct {
		Value string `parser:"'optimization' @('default' | 'normal' | 'high-latency' | 'satellite' | 'aggressive' | 'convervative')"`
	}
	LimitItem struct {
		Variable string `parser:"@('states' | 'frags' | 'src-nodes' | 'tables' | 'table-entries')"`
		Value    int    `parser:"@Number"`
	}
	LimitOption struct {
		ValueOrBraceList[LimitItem] `parser:"'limit' @@"`
	}
	LoginInterfaceOption struct {
		None      BooleanSet `parser:"'logininterface' ( @('none')"`
		Interface string     `parser:"| @Ident)"`
	}
	BlockPolicyOption struct {
		Policy string `parser:"'block-policy' @('drop' | 'return')"`
	}
	StatePolicyOption struct {
		Policy string `parser:"'state-policy' @('if-bound' | 'floating')"`
	}
	FlushStateOverloadEntry struct {
		Global BooleanSet `parser:"'flush' @('global')?"`
	}
	StateOverloadEntry struct {
		Value string                   `parser:"'overload' '<' @String '>'"`
		Flush *FlushStateOverloadEntry `parser:"@@?"`
	}
	StateOption struct {
		Max            *int       `parser:"('max' @Number)"`
		NoSync         BooleanSet `parser:"| @('no-sync')"`
		*Timeout       `parser:"| @@"`
		Sloppy         BooleanSet          `parser:"| @('sloppy')"`
		Pflow          BooleanSet          `parser:"| @('pflow')"`
		SourceTrack    string              `parser:"| ('source-track' @('rule' | 'global'))"`
		MaxSrcNodes    *int                `parser:"| ('max-src-nodes' @Number)"`
		MaxSrcStates   *int                `parser:"| ('max-src-states' @Number)"`
		MaxSrcConn     *int                `parser:"| ('max-src-conn' @Number)"`
		MaxSrcConnRage *[2]int             `parser:"| ('max-src-conn-rate' @Number '/' @Number)"`
		Overload       *StateOverloadEntry `parser:"| @@"`
		IfFloating     BooleanSet          `parser:"@('if-floating')"`
		Floating       BooleanSet          `parser:"@('floating')"`
	}
	StateDefaultsOption struct {
		ValueOrRawList[StateOption] `parser:"'state-defaults' @@"`
	}
	FingerPrintsOption struct {
		Filename string `parser:"'fingerprints' @(Ident | String | Filename)"`
	}
	IfSpecEntry struct {
		Negate                    BooleanSet `parser:"'!'?"`
		InterfaceOrInterfaceGroup string     `parser:"@Ident"`
	}
	IfSpec       ValueOrBraceList[IfSpecEntry]
	SkipOnOption struct {
		IfSpec `parser:"'skip' 'on' @@"`
	}
	DebugOption struct {
		Level string `parser:"'debug' @('emerg' | 'alert' | 'crit' | 'err' | 'warning' | 'notice' | 'info' | 'debug')"`
	}
	ReassembleOption struct {
		Reassemble BooleanSet `parser:"'reassemble' (@('yes') | 'no')"`
		NoDf       BooleanSet `parser:"@('no-df')?"`
	}
	Option struct {
		*TimeoutOption             `parser:"'set' (@@"`
		*RulesetOptimizationOption `parser:"| @@"`
		*OptimizatioOption         `parser:"| @@"`
		*LimitOption               `parser:"| @@"`
		*BlockPolicyOption         `parser:"| @@"`
		*StatePolicyOption         `parser:"| @@"`
		*StateDefaultsOption       `parser:"| @@"`
		*FingerPrintsOption        `parser:"| @@"`
		*SkipOnOption              `parser:"| @@"`
		*DebugOption               `parser:"| @@"`
		*ReassembleOption          `parser:"| @@)"`
	}
	ActionBlockReturn struct {
		Return string `parser:"@('return' | 'drop')"`
	}
	ActionBlock struct {
		*ActionBlockReturn `parser:"'block' @@?"`
	}
	Action struct {
		Pass  BooleanSet   `parser:"@('pass')"`
		Match BooleanSet   `parser:"| @('match')"`
		Block *ActionBlock `parser:"| @@"`
	}
	LogOption struct {
		All     BooleanSet `parser:"@('all'"`
		Matches BooleanSet `parser:"| @('matches')"`
		User    BooleanSet `parser:"| @('user')"`
		To      *string    `parser:"| ('to' @Ident))"`
	}
	Log struct {
		Options *ValueOrRawList[LogOption] `parser:"'log' ('(' @@ ')')?"`
	}
	PfRuleOn struct {
		*IfSpec `parser:"'on' ( @@"`
		Rdomain *string `parser:"| ('rdomain' @Number))"`
	}
	AddressFamily struct {
		Is4 BooleanSet `parser:"@('inet') | 'inet6'"`
	}
	Protocol struct {
		Name   *string `parser:"@Ident"`
		Number *string `parser:"| @Number"`
	}
	ProtoSpec struct {
		ValueOrBraceList[Protocol] `parser:"'proto' @@"`
	}
	Address struct {
		Hostname                  *string     `parser:"@Hostname"`
		InterfaceOrInterfaceGroup *string     `parser:"| (@Ident | ( '(' @Ident ')' ))"`
		IP                        *netip.Addr `parser:"| @Address"`
	}
	Host struct {
		Negate   BooleanSet `parser:"@('!')?"`
		Address  *Address   `parser:"( ( @@"`
		Mask     *int       `parser:"('/' @Number)?"`
		Weight   *int       `parser:"('weight' @Number)? )"`
		AsString *string    `parser:"| ('<' @String '>') )"`
	}
	Unary struct {
		Operator string  `parser:"@('=' | '!=' | '<' | '<=' | '>' | '>=')?"`
		Name     *string `parser:"(@Ident"`
		Number   *int    `parser:"| @Number)"`
	}
	Binary struct {
		Lhs      int    `parser:"@Number"`
		Operator string `parser:"@('<>' | '><' | ':')"`
		Rhs      int    `parser:"@Number"`
	}
	PortOp struct {
		*Unary  `parser:"@@"`
		*Binary `parser:"| @@"`
	}
	Port struct {
		ValueOrBraceList[PortOp] `parser:"'port' @@"`
	}
	HostsTarget struct {
		Any     BooleanSet `parser:"( @('any')"`
		NoRoute BooleanSet `parser:"| @('no-route')"`
		Self    BooleanSet `parser:"| @('self')"`
		Route   *string    `parser:"| ('route' @String)"`
		Host    *Host      `parser:"| @@ )"`
		*Port   `parser:"@@?"`
	}
	HostsFromTo struct {
		From HostsTarget `parser:"'from' @@"`
		To   HostsTarget `parser:"'to' @@"`
	}
	Hosts struct {
		All          BooleanSet `parser:"@('all')"`
		*HostsFromTo `parser:"| @@"`
	}
	FilterOption struct {
		// *User        `parser:"@@"`
		// *Group       `parser:"| @@"`
		// *Flags       `parser:"| @@"`
		// *IcmpType    `parser:"| @@"`
		// *IcmpType6   `parser:"| @@"`
		// *Tos         `parser:"| ('tos' @@)"`
		State        *string                      `parser:"(@('no' | 'keep' | 'modulate' | 'synproxy') 'state')"`
		StateOptions *ValueOrRawList[StateOption] `parser:"| ('(' @@ ')')"`
	}
	PfRule struct {
		Action        `parser:"@@"`
		Direction     *string `parser:"@('in' | 'out')?"`
		*Log          `parser:"@@?"`
		Quick         BooleanSet                    `parser:"@('quick')?"`
		On            *PfRuleOn                     `parser:"@@?"`
		AddressFamily *AddressFamily                `parser:"@@?"`
		ProtoSpec     *ProtoSpec                    `parser:"@@?"`
		Hosts         *Hosts                        `parser:"@@?"`
		FilterOptions *ValueOrRawList[FilterOption] `parser:"@@?"`
	}
	Line struct {
		*Comment `parser:"@Comment"`
		*Option  `parser:"| @@"`
		*PfRule  `parser:"| @@"`
		// *AntiSpoofRule `parser:"| @@"`
		// *QueueRule     `parser:"| @@"`
		// *AnchorRule    `parser:"| @@"`
		// *AnchorClose   `parser:"| @@"`
		// *LoadAnchor    `parser:"| @@"`
		// *TableRule     `parser:"| @@"`
		// *Inclue        `parser:"| @@"`
	}
	Configuration struct {
		Line []*Line `parser:"@@*"`
	}
)
