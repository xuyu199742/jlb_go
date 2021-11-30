package config

type Wx struct {
	Appid     string `mapstructure:"appid" json:"appID" yaml:"appid"`
	Secret    string `mapstructure:"secret" json:"secret" yaml:"secret"`
	SfaAppid  string `mapstructure:"sfa-appid" json:"sfaAppid" yaml:"sfa-appid"`
	SfaSecret string `mapstructure:"sfa-secret" json:"sfaSecret" yaml:"sfa-secret"`
}
