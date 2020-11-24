package library

import (
	"github.com/spf13/viper"
)

// NewConfig new config
func NewConfig(fpath, fname, ftype string, cfg interface{}) error {
	vpr := viper.New()
	vpr.AddConfigPath(fpath)
	vpr.SetConfigName(fname)
	vpr.SetConfigType(ftype)
	if err := vpr.ReadInConfig(); err != nil {
		return err
	}
	if err := vpr.Unmarshal(cfg); err != nil {
		return err
	}
	return nil
}
