package config

type SiteInfo struct {
	Name        string `yaml:"name" json:"name"`
	Logo        string `yaml:"logo" json:"logo"`
	Title       string `yaml:"title" json:"title"`
	Email       string `yaml:"email" json:"email"`
	Addr        string `yaml:"addr" json:"addr"`
	Version     string `yaml:"version" json:"version"`
	WechatImage string `yaml:"wechat_image" json:"wechat_image"`
	Job         string `yaml:"job" json:"job"`
	GithubUrl   string `yaml:"github_url" json:"github_url"`
	GiteeUrl    string `yaml:"gitee_url" json:"gitee_url"`
}
