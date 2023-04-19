package config

type config struct {
	LogLevel     uint32 `yaml:"logLevel"`
	MySqlDsn     string `yaml:"mySqlDsn"`
	MySqlMigrate bool   `yaml:"mySqlMigrate"`
}
