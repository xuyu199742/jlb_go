package config

type Local struct {
	Path   string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件路径
	ImgUrl string `mapstructure:"img-url" json:"imgUrl" yaml:"img-url"`
}
