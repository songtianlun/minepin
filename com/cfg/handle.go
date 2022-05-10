package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"minepin/com/constvar"
	"minepin/com/utils"
	"strings"
)

type Config struct {
	Name string
}

type Cfg struct {
	name  string
	def   interface{}
	typ   string
	value interface{}
}

var ConfigMap = make(map[string]*Cfg)

func (c Cfg) RegisterCfg() {
	_, ok := ConfigMap[c.name]
	if ok {
		panic(fmt.Sprintf("%s is already registered", c.name))
	}
	c.value = nil
	ConfigMap[c.name] = &c
}

func Get(n string, noErr bool) interface{} {
	var v interface{}

	n = strings.ToLower(n)

	c, ok := ConfigMap[n]
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

func GetString(n string) string { return Get(n, false).(string) }
func GetInt(n string) int       { return Get(n, false).(int) }
func GetInt64(n string) int64   { return Get(n, false).(int64) }
func GetBool(n string) bool     { return Get(n, false).(bool) }

func (c *Config) initConfig() error {
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

	for k, v := range ConfigMap {
		viper.SetDefault(k, v.def)
	}

	viper.AutomaticEnv()                             // 读取匹配的环境变量，环境变量优先级最高
	viper.SetEnvPrefix(constvar.DefaultCfgEnvPrefix) // 读取环境变量的前缀为 MINEGIN
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

// RegisterCfg
// 注：必须在 Init 之前完成所有配置的注册，否则默认配置不会生效。
// 注：内部变量返回指针，会发生逃逸，编译器自动在堆上分配内存，
// 考虑到该变量伴随程序整个周期，可以接受堆栈内存性能损失。
func RegisterCfg(k string, d interface{}, t string) {
	Cfg{
		name:  strings.ToLower(k),
		def:   d,
		typ:   t,
		value: nil,
	}.RegisterCfg()
}

// Init
// 完成配置文件的读取、默认配置的注册。
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	if err := c.initConfig(); err != nil {
		return err
	}

	//c.watchConfig() // 热加载当前无意义，且会在配置文件不存在时阻塞，暂时移除

	return nil
}
