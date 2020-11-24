package server

// ConfigModel 服务器配置文件模型
type ConfigModel struct {
	ListenAddress string `mapstructure:"listen_address"`
	StoragePath   string `mapstructure:"storage_path"`
	DirList       bool   `mapstructure:"dir_list"`
}
