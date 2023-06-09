package config

type mysql struct {
	Db         string `yaml:"db"`
	DbHost     string `yaml:"dbHost"`
	DbPort     string `yaml:"dbPort"`
	DbUser     string `yaml:"dbUser"`
	DbPassWord string `yaml:"dbPassWord"`
	DbName     string `yaml:"dbName"`
}
