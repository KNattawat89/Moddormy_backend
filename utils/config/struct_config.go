package config

type config struct {
	Path          string   `yaml:"path"`
	URL           string   `yaml:"url"`
	LogLevel      uint32   `yaml:"logLevel"`
	MySqlDsn      string   `yaml:"mySqlDsn"`
	MySqlMigrate  bool     `yaml:"mySqlMigrate"`
	ServerHeader  string   `yaml:"serverHeader"`
	BackAddress   string   `yaml:"backAddress"`
	Cors          []string `yaml:"cors"`
	ProductionURL string   `yaml:"productionURL"`
}
