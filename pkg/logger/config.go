package logger

type Env string

func (e Env) String() string {
	return string(e)
}

func (e Env) IsValid() bool {
	switch e {
	case Dev, Prod, Local:
		return true
	default:
		return false
	}
}

const (
	Local Env = "local"
	Dev   Env = "dev"
	Prod  Env = "prod"
)

type Config struct {
	Env    Env
	Folder string
}

func NewConfig(env Env, folder string) *Config {
	return &Config{
		Env:    env,
		Folder: folder,
	}
}
