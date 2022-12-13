package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"minepin/com/constvar"
	"minepin/com/utils"
	"strings"
)

type config struct {
	Name string
}

type cfg struct {
	name  string
	def   interface{}
	typ   string
	value interface{}
}

var configMap = make(map[string]*cfg)

func (c cfg) RegisterCfg() {
	_, ok := configMap[c.name]
	if ok {
		panic(fmt.Sprintf("%s is already registered", c.name))
	}
	c.value = nil
	configMap[c.name] = &c
}

func Get(n string, noErr bool) interface{} {
	var v interface{}

	n = strings.ToLower(n)

	c, ok := configMap[n]
	if !ok {
		errMsg := fmt.Sprintf("cfg key [%s] is not registered", n)
		if noErr {
			return nil
		} else {
			panic(errMsg)
		}
	}

	if c.value != nil {
		return c.value
	}

	if c.typ == "string" {
		v = strings.TrimSpace(viper.GetString(n))
	} else if c.typ == "int" {
		v = viper.GetInt(n)
	} else if c.typ == "int64" {
		// 默认获取 64 位 int，避免跨平台带来的问题
		v = viper.GetInt64(n)
	} else if c.typ == "bool" {
		v = viper.GetBool(n)
	} else {
		v = viper.Get(n)
	}

	c.value = v
	return v
}
func (c *config) initConfig() error {
	var cfgFile string

	if c.Name != "" {
		viper.SetConfigFile(c.Name)
		fmt.Printf("run with abstract config %s\n", c.Name)
		cfgFile = c.Name
	} else if isExist, _ := utils.PathExists(constvar.DefaultCfgFile); isExist {
		fmt.Printf("run with default config %s\n", constvar.DefaultCfgFile)
		viper.AddConfigPath(constvar.DefaultCfgPath)
		viper.SetConfigName(constvar.DefaultCfgName)
		viper.SetConfigType(constvar.DefaultCfgType)
		cfgFile = constvar.DefaultCfgFile
	}

	for k, v := range configMap {
		viper.SetDefault(k, v.def)
	}

	viper.AutomaticEnv() // 读取匹配的环境变量，环境变量优先级最高

	if constvar.DefaultCfgEnvPrefix == "" {
		viper.AllowEmptyEnv(true) // 允许空环境变量
	} else {
		viper.SetEnvPrefix(constvar.DefaultCfgEnvPrefix) // 读取环境变量的前缀
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if isExist, _ := utils.PathExists(cfgFile); !isExist {
		return nil
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func SetString(k string, v string) {
	if _, ok := configMap[k]; !ok {
		// log.Panicf("cfg key [%s] is not registered", k)
		return
	}
	if configMap[k].typ != "string" {
		// log.Panicf("cfg key [%s] is not string type", k)
		return
	}
	configMap[k].value = v
}

func GetString(n string) string { return Get(n, false).(string) }
func GetInt(n string) int       { return Get(n, false).(int) }
func GetInt64(n string) int64   { return Get(n, false).(int64) }
func GetBool(n string) bool     { return Get(n, false).(bool) }

// RegisterCfg
// 注：必须在 Init 之前完成所有配置的注册，否则默认配置不会生效。
// 注：内部变量返回指针，会发生逃逸，编译器自动在堆上分配内存，
// 考虑到该变量伴随程序整个周期，可以接受堆栈内存性能损失。
func RegisterCfg(k string, d interface{}, t string) {
	cfg{
		name:  strings.ToLower(k),
		def:   d,
		typ:   t,
		value: nil,
	}.RegisterCfg()
}

// Init
// 完成配置文件的读取、默认配置的注册。
func Init(cfg string) error {
	c := config{
		Name: cfg,
	}

	if err := c.initConfig(); err != nil {
		return err
	}

	// c.watchConfig() // 热加载当前无意义，且会在配置文件不存在时阻塞，暂时移除

	return nil
}
