package setting

import (
	"Txray/core"
	"Txray/core/setting/key"
	"Txray/log"
	"os"

	"github.com/spf13/viper"
)

func init() {
	// 配置文件不存在则创建
	if _, err := os.Stat(core.SettingFile); os.IsNotExist(err) {
		file, _ := os.Create(core.SettingFile)
		_ = file.Close()
	}
	viper.SetConfigName("setting")
	viper.SetConfigType("toml")
	viper.AddConfigPath(core.GetConfigDir())
	// 设置默认值
	viper.SetDefault(key.Socks, 1080)
	viper.SetDefault(key.Http, 0)
	viper.SetDefault(key.UDP, true)
	viper.SetDefault(key.Sniffing, true)
	viper.SetDefault(key.FromLanConn, false)
	viper.SetDefault(key.Mux, false)

	viper.SetDefault(key.RoutingStrategy, "IPIfNonMatch") //路由策略
	viper.SetDefault(key.RoutingBypass, true)             // 绕过局域网和大陆

	viper.SetDefault(key.DNSPort, 1350)
	viper.SetDefault(key.DNSForeign, "1.1.1.1")
	viper.SetDefault(key.DNSDomestic, "119.29.29.29")
	viper.SetDefault(key.DNSBackup, "114.114.114.114")

	viper.SetDefault(key.TestURL, "https://www.youtube.com")
	viper.SetDefault(key.TestTimeout, 5)
	viper.SetDefault(key.RunBefore, "")

	viper.SetDefault(key.PID, 0)
	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Error(err)
	}
	// 监听配置文件
	viper.WatchConfig()
}