package config

type Jwt struct {
	SecretKey string `yaml:"secret_key" json:"secret_key"` // 密钥
	Expire    int    `yaml:"expire" json:"expire"`         // 过期时间
	Issuer    string `yaml:"issuer" json:"issuer"`         // 签发者
}
