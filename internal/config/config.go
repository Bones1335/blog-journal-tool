package config

import (
	"fmt"
	"time"
)

type Config struct {
	Editor  string
	Journal string
	Blog    string
}

func Read() Config {
	currentYear := time.Now()
	config := Config{
		Editor:  "nvim",
		Journal: fmt.Sprintf("/Documents/journal/%v", currentYear.Year()),
		Blog:    fmt.Sprintf("/programs/read-the-bones/content/posts/%v", currentYear.Year()),
	}

	return config
}
