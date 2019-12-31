package structdemo

// Config struct
type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		DbHost string `yaml:"dbhost"`
		DbPort int    `yaml:"dbport"`
		DbUser string `yaml:"dbuser"`
		DbPass string `yaml:"dbpass"`
		DbName string `yaml:"dbname"`
	} `yaml:"database"`
}
