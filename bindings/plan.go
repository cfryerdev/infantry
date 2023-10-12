package bindings

type Plan struct {
	Setup     Setup      `yaml:"setup,omitempty"`
	Protocol  Protocol   `yaml:"protocol,omitempty"`
	Proposals []Proposal `yaml:"proposal,omitempty"`
}

type Setup struct {
	Host   string  `yaml:"host,omitempty"`
	Stages []Stage `yaml:"stages,omitempty"`
}

type Stage struct {
	Iterations            int `yaml:"iterations,omitempty"`
	AddUsersPerIterations int `yaml:"addUsersPerIterations,omitempty"`
	MaxUsersAtOnce        int `yaml:"maxUsersAtOnce,omitempty"`
}

type Protocol struct {
	Http   ProtocolHttp   `yaml:"http,omitempty"`
	Socket ProtocolSocket `yaml:"socket,omitempty"`
}

type ProtocolHttp struct {
	Tls     Tls      `yaml:"tls,omitempty"`
	Timeout int      `yaml:"timeout,omitempty"`
	Headers []Header `yaml:"headers,omitempty"`
}

type Tls struct {
	IgnoreSslErrors bool `yaml:"ignoreSslErrors,omitempty"`
}

type Header struct {
	Key   string `yaml:"key,omitempty"`
	Value string `yaml:"value,omitempty"`
}

type ProtocolHttpHeaders struct {
	Ssl bool `yaml:"ssl,omitempty"`
}

type ProtocolSocket struct {
}

type Proposal struct {
	Name   string `yaml:"name,omitempty"`
	Method string `yaml:"method,omitempty"`
	Url    string `yaml:"url,omitempty"`
}

type Executor interface {
	Initialize(protocol Protocol)
	Execute(proposal Proposal) (Test, error)
}
