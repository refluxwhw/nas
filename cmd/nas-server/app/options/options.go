package options

type NasServerOption struct {
	ConfigFile string
}

func NewNasServerOption() *NasServerOption {
	return &NasServerOption{
		ConfigFile: "conf.yaml",
	}
}
