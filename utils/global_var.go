package utils

import (
	"github.com/spf13/viper"
)

const SecretPassword string = "abc&1*~#^2^#s0^=)^^7%b34"

type GlobalVar struct {
	EncryptByte string `mapstructure:"SECRET"`
}

func CallGlobalVar() (globalVar GlobalVar) {
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	if err = viper.Unmarshal(&globalVar); err != nil {
		return
	}
	return
}
