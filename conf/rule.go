package conf

// Rule the replacer rule
type Rule struct {
	Name string            `yaml:"name"`
	Old  []string          `yaml:"old"`
	New  map[string]string `yaml:"new"`
}
