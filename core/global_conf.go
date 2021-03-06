package core

import (
	"github.com/spf13/viper"
	"fmt"
)

type GlobalConf struct {
	GeoBlockFileName string
	GeoLocationFileName string
	IpFileReloadInterval int64
	AdFileName string
	AdFileReloadInterval int64
	// log config
	LogLevel int
	AdServerLogFileName string
	SearchLogFileName string
	ImpressionLogFileName string
	ClickLogFileName string
	ConversionLogFileName string
	// track url prefix
	ImpressionTrackUrlPrefix string
	ClickTrackUrlPrefix string
	ConversionTrackUrlPrefix string
	AdServerPort int
}

var GlobalConfObject *GlobalConf

func init()  {
	GlobalConfObject = new(GlobalConf)
}

func LoadGlobalConf(configPath, configFileName string)  {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFileName)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
		panic(-1)
	}
	err := viper.Unmarshal(GlobalConfObject)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
		panic(-1)
	}
	fmt.Printf("GeoBlockFileName=%s\n", GlobalConfObject.GeoBlockFileName)
	fmt.Printf("GeoLocationFileName=%s\n", GlobalConfObject.GeoLocationFileName)
	fmt.Printf("AdFileName=%s\n", GlobalConfObject.AdFileName)
	fmt.Printf("AdFileReloadInterval=%d\n", GlobalConfObject.AdFileReloadInterval)
	fmt.Printf("LogLevel=%d\n", GlobalConfObject.LogLevel)
	fmt.Printf("AdServerLogFileName=%s\n", GlobalConfObject.AdServerLogFileName)
	fmt.Printf("SearchLogFileName=%s\n", GlobalConfObject.SearchLogFileName)
	fmt.Printf("ImpressionLogFileName=%s\n", GlobalConfObject.ImpressionLogFileName)
	fmt.Printf("ClickLogFileName=%s\n", GlobalConfObject.ClickLogFileName)
	fmt.Printf("ConversionLogFileName=%s\n", GlobalConfObject.ConversionLogFileName)
	fmt.Printf("ImpressionTrackUrlPrefix=%s\n", GlobalConfObject.ImpressionTrackUrlPrefix)
	fmt.Printf("ClickTrackUrlPrefix=%s\n", GlobalConfObject.ClickTrackUrlPrefix)
	fmt.Printf("ConversionTrackUrlPrefix=%s\n", GlobalConfObject.ConversionTrackUrlPrefix)
	fmt.Printf("AdServerPort=%d\n", GlobalConfObject.AdServerPort)
}
