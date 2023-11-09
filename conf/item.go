package conf

// Item the replacer config item
type Item struct {
	Name     string   `yaml:"name"`
	Disabled bool     `yaml:"disabled"`
	Paths    []string `yaml:"paths"`
	Rules    []Rule   `yaml:"rules"`
}
