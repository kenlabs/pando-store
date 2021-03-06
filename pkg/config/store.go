package config

import "github.com/mitchellh/go-homedir"

const (
	DefaultStoreDir  = "PandoStore"
	DefaultCacheSize = 1024 * 1024 * 10
)

type StoreConfig struct {
	Type             string
	StoreRoot        string
	Dir              string
	SnapShotInterval string
	CacheSize        int
}

func DefaultConfig() *StoreConfig {
	homeRoot, _ := homedir.Expand("~/.pando")

	return &StoreConfig{
		Type:             "levelds",
		StoreRoot:        homeRoot,
		Dir:              DefaultStoreDir,
		SnapShotInterval: "60m",
		CacheSize:        DefaultCacheSize,
	}
}
