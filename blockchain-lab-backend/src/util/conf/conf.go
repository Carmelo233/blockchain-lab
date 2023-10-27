package conf

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Connconf struct {
	MspID        string
	CryptoPath   string
	CertPath     string
	KeyPath      string
	TlsCertPath  string
	PeerEndpoint string
	GatewayPeer  string
}

var (
	Peer1 Connconf

	configPath string
)

func InitConf() error {
	pflag.StringVarP(&configPath, "conf", "", "/home/yehh/goProject/blockchain-lab/blockchain-lab-backend/src/config/connection.yaml",
		"config path, eg: --conf config.yaml")
	pflag.Parse()
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		//lgLogger.Logger.Error("read config failed: ", zap.String("err", err.Error()))
		fmt.Println("read config failed: ", zap.String("err", err.Error()))
		panic(err)
	}
	if err := v.Unmarshal(&Peer1); err != nil {
		//lgLogger.Logger.Error("config parse failed: ", zap.String("err", err.Error()))
		fmt.Println("config parse failed: ", zap.String("err", err.Error()))
		return err
	}
	fmt.Println(Peer1)
	return nil
}
