package conf

// Conf the replacer config
type Conf struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Items   []Item `yaml:"items"`
}
