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
	Address struct {
		IP       *IP     `parser:"@@"`
		Hostname *string `parser:"| @Hostname"`
		Other    *string `parser:"| @Ident"`
	}
	Host struct {
		Negate   BooleanSet `parser:"@('!')?"`
		Address  *Address   `parser:"( ( @@"`
		Weight   *int       `parser:"('weight' @Number)? )"`
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
	OsOption struct {
		Value string `parser:"@(Ident | String)"`
	}
	Os struct {
		Selected ValueOrBraceList[OsOption] `parser:"'os' @@"`
	}
	HostsTarget struct {
		Any     BooleanSet `parser:"( @('any')"`
		NoRoute BooleanSet `parser:"| @('no-route')"`
		Self    BooleanSet `parser:"| @('self')"`
		Route   *string    `parser:"| ('route' @(String | Ident))"`
		Host    *Host      `parser:"| @@ )"`
	}
	HostsFromTo struct {
		From     ValueOrBraceList[HostsTarget] `parser:"'from' @@"`
		FromPort *Port                         `parser:"@@?"`
		FromOs   *Os                           `parser:"@@?"`
		To       ValueOrBraceList[HostsTarget] `parser:"'to' @@"`
		ToPort   *Port                         `parser:"@@?"`
	}
	Hosts struct {
		All         BooleanSet   `parser:"@('all')"`
		HostsFromTo *HostsFromTo `parser:"| @@"`
	}
	User struct {
		Selected ValueOrBraceList[Operation] `parser:"'user' @@"`
	}
	Group struct {
		Selected ValueOrBraceList[Operation] `parser:"'group' @@"`
	}
	Flags struct {
		Left  []string   `parser:"'flags' @('F' | 'S' | 'R' | 'P' | 'A' | 'U' | 'E' | 'W')?"`
		Right string     `parser:"'/' (@('F' | 'S' | 'R' | 'P' | 'A' | 'U' | 'E' | 'W')"`
		Any   BooleanSet `parser:"@('any') )"`
	}
	Tos struct {
		Selected string `parser:"@('lowdelay' | 'throughput' | 'reliability')"`
		Number   int    `parser:"@Hexnumber"`
	}
	Label struct {
		Text string `parser:"'label' @(String | Ident)"`
	}
	Tag struct {
		Text string `parser:"'tag' @(String | Ident)"`
	}
	Tagged struct {
		Negate BooleanSet `parser:"@('!')?"`
		Text   string     `parser:"'tagged' @(String | Ident)"`
	}
	ScrubOption struct {
		NoDf          BooleanSet `parser:"@('no-df')"`
		MinTtl        *int       `parser:"| ('min-ttl' @Number)"`
		MaxMss        *int       `parser:"| ('max-mss' @Number)"`
		ReassembleTcp BooleanSet `parser:"| @('reassemble' 'tcp')"`
		RandomId      BooleanSet `parser:"| @('random-id')"`
	}
	ScrubOptions struct {
		Options ValueOrRawList[ScrubOption] `parser:"@@"`
	}
	FilterOption struct {
		User  *User  `parser:"@@"`
		Group *Group `parser:"| @@"`
		Flags *Flags `parser:"| @@"`
		// TODO: IcmpType     *IcmpType                    `parser:"| @@"`
		// TODO: IcmpType6    *IcmpType6                   `parser:"| @@"`
		Tos          *Tos                         `parser:"| ('tos' @@)"`
		State        *string                      `parser:"| (@('no' | 'keep' | 'modulate' | 'synproxy') 'state')"`
		StateOptions *ValueOrRawList[StateOption] `parser:"| ('(' @@ ')')"`
		ScrubOption  *ScrubOptions                `parser:"| ('scrub' '(' @@ ')')"`
		// TODO: FRAGMENT
		// TODO: ALLOWOPTIONS
		// TODO: once
		// TODO: Divert-packet
		// TODO: divert-reply
		// TODO: divert-to port
		Label  *Label  `parser:"| @@"`
		Tag    *Tag    `parser:"| @@"`
		Tagged *Tagged `parser:"| @@"`
		// max-pkt-rate number "/" seconds: Specifies the maximum packet rate in packets per second.
		// "set delay" number: Sets a delay for packets.
		// "set prio" ( number | "(" number [ [ "," ] number ] ")" ): Sets the priority of packets. Can be a single number or a range.
		// "set queue" ( string | "(" string [ [ "," ] string ] ")" ): Assigns packets to a specific queue. Can be a single queue name or a list of queue names.
		// "rtable" number: Specifies the routing table to use.
		// "probability" number"%": Sets the probability of the rule matching.
		// "prio" number: Sets the priority of the rule itself.
		// "af-to" af "from" ( redirhost | "{" redirhost-list "}" ) [ "to" ( redirhost | "{" redirhost-list "}" ) ]: Redirects traffic to an address family from a specified host or list of hosts, optionally to another host or list of hosts.
		// "binat-to" ( redirhost | "{" redirhost-list "}" ) [ portspec ] [ pooltype ]: Performs bidirectional NAT to a specified host or list of hosts, with optional port specification and pool type.
		// "rdr-to" ( redirhost | "{" redirhost-list "}" ) [ portspec ] [ pooltype ]: Redirects traffic to a specified host or list of hosts, with optional port specification and pool type.
		// "nat-to" ( redirhost | "{" redirhost-list "}" ) [ portspec ] [ pooltype ] [ "static-port" ]: Performs NAT to a specified host or list of hosts, with optional port specification, pool type, and static-port flag.
		// [ route ]: Specifies a routing action.
		// [ "set tos" tos ]: Sets the Type of Service (TOS) field in the IP header.
		// [ [ "!" ] "received-on" ( interface-name | interface-group ) ]: Matches packets received on a specific interface or interface group, optionally negated.
	}
	PfRule struct {
		Action        Action                        `parser:"@@"`
		Direction     *string                       `parser:"@('in' | 'out')?"`
		Log           *Log                          `parser:"@@?"`
		Quick         BooleanSet                    `parser:"@('quick')?"`
		On            *PfRuleOn                     `parser:"@@?"`
		AddressFamily *AddressFamily                `parser:"@@?"`
		ProtoSpec     *ProtoSpec                    `parser:"@@?"`
		Hosts         *Hosts                        `parser:"@@?"`
		FilterOptions *ValueOrRawList[FilterOption] `parser:"@@?"`
	}
	AntiSpoofRule struct {
		Log           *Log           `parser:"'antispoof' @@?"`
		Quick         BooleanSet     `parser:"@('quick')?"`
		IfSpec        IfSpec         `parser:"'for' @@"`
		AddressFamily *AddressFamily `parser:"@@?"`
		Label         *Label         `parser:"@@?"`
	}
	Literal struct {
		Address Address `parser:"@@"`
		String  string  `parser:"| @String"`
		Number  int     `parser:"| @Number"`
	}
	Assignment struct {
		Variable string                    `parser:"@Ident"`
		Value    ValueOrBraceList[Literal] `parser:"'=' @@"`
	}
	Line struct {
		Option        *Option        `parser:"@@"`
		PfRule        *PfRule        `parser:"| @@"`
		Comment       *Comment       `parser:"| @Comment"`
		AntiSpoofRule *AntiSpoofRule `parser:"| @@"`
		Assignment    *Assignment    `parser:"| @@"`
		// QueueRule *QueueRule     `parser:"| @@"`
		// AnchorRule *AnchorRule    `parser:"| @@"`
		// AnchorClose *AnchorClose   `parser:"| @@"`
		// LoadAnchor *LoadAnchor    `parser:"| @@"`
		// TableRule *TableRule     `parser:"| @@"`
		// Inclue *Inclue        `parser:"| @@"`
	}
	Configuration struct {
		Line []*Line `parser:"@@*"`
	}
)
