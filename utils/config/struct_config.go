package config

type config struct {
	LogLevel     uint32   `yaml:"logLevel"`
	MySqlDsn     string   `yaml:"mySqlDsn"`
	MySqlMigrate bool     `yaml:"mySqlMigrate"`
	ServerHeader string   `yaml:"serverHeader"`
	BackAddress  string   `yaml:"backAddress"`
	FrontAddress string   `yaml:"frontAddress"`
	Cors         []string `yaml:"cors"`
}
