package utils

import (
	"github.com/spf13/viper"
)

const SecretPassword string = "abc&1*~#^2^#s0^=)^^7%b34"

type GlobalVar struct {
	EncryptByte string `mapstructure:"SECRET"`
}

func CallGlobalVar() (globalVar GlobalVar) {
	v := viper.New()
	v.AutomaticEnv()
	v.BindEnv("SECRET")

	if err := v.Unmarshal(&globalVar); err != nil {
		return
	}
	return
}
