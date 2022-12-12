package templates

type CFG []struct {
	Country        string `json:"country"`
	Region         string `json:"region"`
	Location       string `json:"location"`
	ConnectionName string `json:"connectionName"`
	IP             []string
}

type CONFIG struct {
	URL       string `yaml:"url"`
	SecureDNS bool   `yaml:"secureDNS"`
}
