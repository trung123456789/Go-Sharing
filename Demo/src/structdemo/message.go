package structdemo

// Message struct
type Message struct {
	Msg struct {
		ConnectErr string `yaml:"ConnectErr"`
		SearchErr  string `yaml:"SearchErr"`
		DatabseErr string `yaml:"DatabseErr"`
		JSONErr    string `yaml:"JSONErr"`
		NotFound   string `yaml:"NotFound"`
		ServerErr  string `yaml:"ServerErr"`
	} `yaml:"Msg"`
	ErrMsg struct {
		ConnectErr string `yaml:"ConnectErr"`
		SearchErr  string `yaml:"SearchErr"`
		DatabseErr string `yaml:"DatabseErr"`
		JSONErr    string `yaml:"JSONErr"`
		NotFound   string `yaml:"NotFound"`
		ServerErr  string `yaml:"ServerErr"`
	} `yaml:"ErrMsg"`
}
