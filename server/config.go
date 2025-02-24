package server

type restyConfig struct {
	Host                string
	Port                uint16
	ServerHeaderEnabled bool
}

var RestyConfig = restyConfig{
	Host:                "localhost",
	Port:                3001,
	ServerHeaderEnabled: true,
}
