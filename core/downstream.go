package core

type DownstreamConfig struct {
	URIs            []string `toml:",omitempty"`
	Exchange        string   `toml:",omitempty"`
	RoutingKey      string   `toml:",omitempty"`
	RetryInterval   int      `toml:",omitempty"` //in millsecond
	TimeoutInterval int      `toml:",omitempty"` //in millsecond
	Compress        bool     `toml:",omitempty"`
}
