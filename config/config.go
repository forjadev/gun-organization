package config

import (
	"fmt"
)

func Init() error {
	var err error

	if err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}

	return nil
}
