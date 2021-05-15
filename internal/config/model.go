package config

// AntConfig is the structure of the config used by RedisAnt program
type AntConfig struct {
	Database        string `yaml:"database"`
	Collection      string `yaml:"collection"`
	Fields 					[]string `yaml:"fields"`
}
