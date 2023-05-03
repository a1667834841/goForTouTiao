package config

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	fmt.Print("config init")
	fmt.Print(config.Cookie)
}
