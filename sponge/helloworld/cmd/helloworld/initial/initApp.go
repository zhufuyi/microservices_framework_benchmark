// Package initial is the package that starts the service to initialize the service, including
// the initialization configuration, service configuration, connecting to the database, and
// resource release needed when shutting down the service.
package initial

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/jinzhu/copier"

	"github.com/zhufuyi/sponge/pkg/conf"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/nacoscli"
	"github.com/zhufuyi/sponge/pkg/stat"
	"github.com/zhufuyi/sponge/pkg/tracer"

	"helloworld/configs"
	"helloworld/internal/config"
	//"helloworld/internal/model"
)

var (
	version            string
	configFile         string
	enableConfigCenter bool
)

// InitApp initial app configuration
func InitApp() {
	initConfig()
	cfg := config.Get()

	// initializing log
	_, err := logger.Init(
		logger.WithLevel(cfg.Logger.Level),
		logger.WithFormat(cfg.Logger.Format),
		logger.WithSave(cfg.Logger.IsSave),
	)
	if err != nil {
		panic(err)
	}
	logger.Debug(config.Show())
	logger.Info("init logger succeeded")

	// initializing database
	//model.InitDB()
	//logger.Infof("init %s succeeded", cfg.Database.Driver)
	//model.InitCache(cfg.App.CacheType)

	// initializing tracing
	if cfg.App.EnableTrace {
		tracer.InitWithConfig(
			cfg.App.Name,
			cfg.App.Env,
			cfg.App.Version,
			cfg.Jaeger.AgentHost,
			strconv.Itoa(cfg.Jaeger.AgentPort),
			cfg.App.TracingSamplingRate,
		)
		logger.Info("init tracer succeeded")
	}

	// initializing the print system and process resources
	if cfg.App.EnableStat {
		stat.Init(
			stat.WithLog(logger.Get()),
			stat.WithAlarm(), // invalid if it is windows, the default threshold for cpu and memory is 0.8, you can modify them
		)
		logger.Info("init statistics succeeded")
	}
}

func initConfig() {
	flag.StringVar(&version, "version", "", "service Version Number")
	flag.BoolVar(&enableConfigCenter, "enable-cc", false, "whether to get from the configuration center, "+
		"if true, the '-c' parameter indicates the configuration center")
	flag.StringVar(&configFile, "c", "", "configuration file")
	flag.Parse()

	if enableConfigCenter {
		// get the configuration from the configuration center (first get the nacos configuration,
		// then read the service configuration according to the nacos configuration center)
		if configFile == "" {
			configFile = configs.Path("helloworld_cc.yml")
		}
		nacosConfig, err := config.NewCenter(configFile)
		if err != nil {
			panic(err)
		}
		appConfig := &config.Config{}
		params := &nacoscli.Params{}
		_ = copier.Copy(params, &nacosConfig.Nacos)
		format, data, err := nacoscli.GetConfig(params)
		if err != nil {
			panic(fmt.Sprintf("connect to configuration center err, %v", err))
		}
		err = conf.ParseConfigData(data, format, appConfig)
		if err != nil {
			panic(fmt.Sprintf("parse configuration data err, %v", err))
		}
		if appConfig.App.Name == "" {
			panic("read the config from center error, config data is empty")
		}
		config.Set(appConfig)
	} else {
		// get configuration from local configuration file
		if configFile == "" {
			configFile = configs.Path("helloworld.yml")
		}
		err := config.Init(configFile)
		if err != nil {
			panic("init config error: " + err.Error())
		}
	}

	if version != "" {
		config.Get().App.Version = version
	}
}
