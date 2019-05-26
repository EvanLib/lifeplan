package conf

type Config struct {
	Enforcer Enforcer
}

type Enforcer struct {
	Host string `envconfig:"HOST" required:"false" default:"127.0.0.1"`
	Port int    `envconfig:"PORT" required:"false" default:"6379"`
}
