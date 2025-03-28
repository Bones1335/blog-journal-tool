package config

type Config struct {
	Editor       string `yaml:"editor"`
	SaveLocation string `yaml:"save_location"`
}

func Read() Config {
	config := Config{
		Editor:       "nvim",
		SaveLocation: "/Documents/blog",
	}

	return config
}
