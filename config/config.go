package config

var Def *Config

func init() {
	Def = &Config{
		ClusterName: "nexus",
		NodeName: "testNode",
		Level:"info",
	}
}

type Config struct {
	ClusterName string `yaml:"cluster_name"`
	NodeName    string `yaml:"node_name"`
	Level       string `yaml:"level"`
}
