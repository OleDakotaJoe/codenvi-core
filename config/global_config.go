package config

type GlobalConfig struct {
	HomeDirName string
}

func Global() *GlobalConfig {
	return &GlobalConfig{
		HomeDirName: ".codenvi",
	}
}
