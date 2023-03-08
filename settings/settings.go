package settings
import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)
//Conf 全局变量，用来保存程序所以配置信息
var Conf = new(AppConfig)
type AppConfig struct{
	Name string`mapstructure:"name"`
	Mode string`mapstructure:"mode"`
	Version string`mapstructure:"version"`
	Port int`mapstructure:"port"`
	StartTime string`mapstructure:"start_time"`
	MachineId int64 `mapstructure:"machine_id"`

	*LogConfig `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`

}
type LogConfig struct {
	Level string `mapstructure:"level"`
	Filename string `mapstructure:"filename"`
	MaxSize int `mapstructure:"max_size"`
	MaxBackups int `mapstructure:"max_backups"`
	MaxAge int `mapstructure:"max_age"`
	}
type MySQLConfig struct {
	Host string `mapstructure:"host"`
	User string `mapstructure:"user"`
	Port int `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DbName string `mapstructure:"dbname"`
	MaxOpenConns int `mapstructure:"max_open_conns"`
	MaxIdleConns int `mapstructure:"max_idle_conns"`

}
type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB int `mapstructure:"db"`
	PoolSize int `mapstructure:"pool_size"`
}
func Init() (err error) {
	viper.SetConfigFile("/Users/liuming/GolandProjects/web_app_2/conf/config.yaml") // 指定配置文件路径
	err = viper.ReadInConfig()                // 读取配置信息
	if err != nil {                           // 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed: %v\n", err)
		return
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//把读取的配置信息反序列化到变量中

	if err := viper.Unmarshal(Conf); err != nil{
		fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
	}

	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了....")
		if err := viper.Unmarshal(Conf); err != nil{
			fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
		}
	})
	return
}
