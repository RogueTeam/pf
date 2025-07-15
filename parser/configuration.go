package parser

import (
	"net/netip"

	"go4.org/netipx"
)

type (
	Timeout struct {
		Variable string  `parser:"@('tcp.first' | 'tcp.opening' | 'tcp.established' | 'tcp.closing' | 'tcp.finwait' | 'tcp.closed' | 'tcp.tsdiff' | 'udp.first' | 'udp.single' | 'udp.multiple' | 'icmp.first' | 'icmp.error' | 'other.first' | 'other.single' | 'other.multiple' | 'frag' | 'interval' | 'src.track' | 'adaptive.start' | 'adaptive.end')"`
		Value    float64 `parser:"@Number"`
	}
	TimeoutOption struct {
		Timeout ValueOrBraceList[Timeout] `parser:"'timeout' @@"`
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
		Limit ValueOrBraceList[LimitItem] `parser:"'limit' @@"`
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
		Max            *int                `parser:"('max' @Number)"`
		NoSync         BooleanSet          `parser:"| @('no-sync')"`
		Timeout        *Timeout            `parser:"| @@"`
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
		Defaults ValueOrRawList[StateOption] `parser:"'state-defaults' @@"`
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
		IfSpec IfSpec `parser:"'skip' 'on' @@"`
	}
	DebugOption struct {
		Level string `parser:"'debug' @('emerg' | 'alert' | 'crit' | 'err' | 'warning' | 'notice' | 'info' | 'debug')"`
	}
	ReassembleOption struct {
		Reassemble BooleanSet `parser:"'reassemble' (@('yes') | 'no')"`
		NoDf       BooleanSet `parser:"@('no-df')?"`
	}
	Option struct {
		Timeout             *TimeoutOption             `parser:"'set' (@@"`
		RulesetOptimization *RulesetOptimizationOption `parser:"| @@"`
		Optimizatio         *OptimizatioOption         `parser:"| @@"`
		Limit               *LimitOption               `parser:"| @@"`
		BlockPolicy         *BlockPolicyOption         `parser:"| @@"`
		StatePolicy         *StatePolicyOption         `parser:"| @@"`
		StateDefaults       *StateDefaultsOption       `parser:"| @@"`
		FingerPrints        *FingerPrintsOption        `parser:"| @@"`
		SkipOn              *SkipOnOption              `parser:"| @@"`
		Debug               *DebugOption               `parser:"| @@"`
		Reassemble          *ReassembleOption          `parser:"| @@)"`
	}
	ActionBlockReturn struct {
		Return string `parser:"@('return' | 'drop')"`
	}
	ActionBlock struct {
		Return *string `parser:"'block' @('return' | 'return-icmp' | 'return-icmp6' | 'return-rst' | 'drop')?"`
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
		Options ValueOrRawList[LogOption] `parser:"'log' ('(' @@ ')')?"`
	}
	PfRuleOn struct {
		IfSpec  *IfSpec `parser:"'on' ( @@"`
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
		Protocol ValueOrBraceList[Protocol] `parser:"'proto' @@"`
	}
	IP struct {
		Mask    *netipx.IPRange `parser:"@IPRange"`
		CIDR    *netip.Prefix   `parser:"| @CIDR"`
		Address *netip.Addr     `parser:"| @Address"`
	}
	Host struct {
		Negate   BooleanSet `parser:"@('!')?"`
		IP       *IP        `parser:"( ( @@"`
		Weight   *int       `parser:"('weight' @Number)? )"`
		Other    *string    `parser:"| @Ident"`
		AsString *string    `parser:"| ('<' @String '>') )"`
	}
	Unary struct {
		Operator string  `parser:"@('=' | '!=' | '<' | '<=' | '>' | '>=')?"`
		Number   *int    `parser:"( @Number "`
		Name     *string `parser:"| @Ident )"`
	}
	Binary struct {
		Lhs      int    `parser:"@Number"`
		Operator string `parser:"@(':' | '<>' | '><')"`
		Rhs      int    `parser:"@Number"`
	}
	Operation struct {
		Binary *Binary `parser:"@@"`
		Unary  *Unary  `parser:"| @@"`
	}
	Port struct {
		Ports ValueOrBraceList[Operation] `parser:"'port' @@"`
	}
	HostsTarget struct {
		Any     BooleanSet `parser:"( @('any')"`
		NoRoute BooleanSet `parser:"| @('no-route')"`
		Self    BooleanSet `parser:"| @('self')"`
		Route   *string    `parser:"| ('route' @(String | Ident))"`
		Host    *Host      `parser:"| @@ )"`
		Port    *Port      `parser:"@@?"`
	}
	HostsFromTo struct {
		From HostsTarget `parser:"'from' @@"`
		To   HostsTarget `parser:"'to' @@"`
	}
	Hosts struct {
		All         BooleanSet   `parser:"@('all')"`
		HostsFromTo *HostsFromTo `parser:"| @@"`
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
		Action    Action     `parser:"@@"`
		Direction *string    `parser:"@('in' | 'out')?"`
		Log       *Log       `parser:"@@?"`
		Quick     BooleanSet `parser:"@('quick')?"`
		// On            *PfRuleOn      `parser:"@@?"`
		AddressFamily *AddressFamily `parser:"@@?"`
		ProtoSpec     *ProtoSpec     `parser:"@@?"`
		Hosts         *Hosts         `parser:"@@?"`
		// FilterOptions *ValueOrRawList[FilterOption] `parser:"@@?"`
	}
	Line struct {
		Option  *Option  `parser:"@@"`
		PfRule  *PfRule  `parser:"| @@"`
		Comment *Comment `parser:"| @Comment"`
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
