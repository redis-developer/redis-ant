package config

// AntConfig is the structure of the config used by RedisAnt program
type AntConfig struct {
	Database   string   `yaml:"database"`
	Collection string   `yaml:"collection"`
	KeyField   string   `yaml:"key_field"`
	Fields     []string `yaml:"fields"`

	GlobalRedisURL  string `envconfig:"GLOBAL_REDIS_URL"`
	GlobalRedisPass string `envconfig:"GLOBAL_REDIS_PASS"`
	GlobalRedisPort string `envconfig:"GLOBAL_REDIS_PORT"`

	ClientRedisURL  string `envconfig:"CLIENT_REDIS_URL"`
	ClientRedisPass string `envconfig:"CLIENT_REDIS_PASS"`
	ClientRedisPort string `envconfig:"CLIENT_REDIS_PORT"`

	MongoURI  string `envconfig:"MONGO_URI"`
	MongoUser string `envconfig:"MONGO_USER"`
	MongoPass string `envconfig:"MONGO_PASS"`
}
