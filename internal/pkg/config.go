package pkg

type ServerHTTP struct {
	ServiceName  string `toml:"service_name"`
	BindAddr     string `toml:"bind_addr_http"`
	ReadTimeout  int    `toml:"read_timeout"`
	WriteTimeout int    `toml:"write_timeout"`
	Protocol     string `toml:"protocol"`
}

type Logger struct {
	LogLevel string `toml:"log_level"`
	LogAddr  string `toml:"log_path"`
}

type DatabaseParams struct {
	MaxOpenCons int `toml:"max_open_cons"`
}

type Config struct {
	ServerHTTPMain ServerHTTP     `toml:"server_http_main"`
	Logger         Logger         `toml:"logger"`
	DatabaseParams DatabaseParams `toml:"database_params"`
}

func NewConfig() *Config {
	return &Config{}
}
