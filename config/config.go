package config

type App struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Local  Local  `mapstructure:"local" json:"local" yaml:"local"`
	Wx     Wx     `mapstructure:"wx" json:"wx" yaml:"wx"`
}
