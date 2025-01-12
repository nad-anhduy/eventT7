package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type JsonStruct struct {
	Name string    `json:"name"`
	Env  EnvStruct `json:"env"`
}

type EnvStruct struct {
	PGConnRO string `mapstructure:"pgConnRO" json:"PGConnRO"`
	PGConnRW string `mapstructure:"pgConnRW" json:"PGConnRW"`
}

func Load(envIs string) (result JsonStruct, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(envIs)
	viper.SetConfigType(`json`)

	if err = viper.ReadInConfig(); err != nil {
		return
	}
	tmp := viper.Get(`env`).(map[string]interface{})
	if err = mapstructure.Decode(tmp, &result.Env); err != nil {
		return
	}
	return
}
