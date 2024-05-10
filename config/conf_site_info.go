package config

type SiteInfo struct {
	CreatedAt string `yaml:"created_at" json:"created_at"`
	Name      string `yaml:"name" json:"name"`
	Title     string `yaml:"title" json:"title"`
	Version   string `yaml:"version" json:"version"`
	Job       string `yaml:"job" json:"job"`
	Addr      string `yaml:"addr" json:"addr"`
}
