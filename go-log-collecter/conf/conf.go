package conf

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var C = InitConfig()

type Config struct {
	viper *viper.Viper
	Kafka *Kafka
	Es    *Es
}

type Kafka struct {
	Address     string
	ChanMaxSize int
}

type Es struct {
	Address     string
	ChanMaxSize int
	Workers     int
}

func InitConfig() *Config {
	conf := &Config{viper: viper.New()}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath("/etc/ms_project/user")
	conf.viper.AddConfigPath(workDir + "/config")
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	conf.InitKafkaConfig()
	conf.InitEsConfig()
	return conf
}

func (c *Config) InitKafkaConfig() {
	kf := &Kafka{
		Address:     c.viper.GetString("kafka.address"),
		ChanMaxSize: c.viper.GetInt("kafka.chan_max_size"),
	}
	c.Kafka = kf
}

func (c *Config) InitEsConfig() {
	e := &Es{
		Address:     c.viper.GetString("es.address"),
		ChanMaxSize: c.viper.GetInt("es.chan_max_size"),
		Workers:     c.viper.GetInt("es.workers"),
	}
	c.Es = e
}
