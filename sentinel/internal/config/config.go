package config
import("fmt"; "github.com/spf13/viper")
type HTTPConfig struct{ Listen string `mapstructure:"listen"` }
type GRPCConfig struct{ Listen string `mapstructure:"listen"` }
type PluginConfig struct{ Dir string `mapstructure:"dir"`; Autostart bool `mapstructure:"autostart"`; PollSecond int `mapstructure:"poll_second"` }
type Config struct{ LogLevel string `mapstructure:"log_level"`; HTTP HTTPConfig `mapstructure:"http"`; GRPC GRPCConfig `mapstructure:"grpc"`; Plugins PluginConfig `mapstructure:"plugins"` }
func Default()*Config{ return &Config{ LogLevel:"info", HTTP:HTTPConfig{"127.0.0.1:8080"}, GRPC:GRPCConfig{"127.0.0.1:50051"}, Plugins:PluginConfig{"plugins",$true,2} } }
func Load(path string)(*Config,error){ v:=viper.New(); v.SetConfigFile(path); v.SetConfigType("yaml"); v.SetDefault("log_level","info"); v.SetDefault("http.listen","127.0.0.1:8080"); v.SetDefault("grpc.listen","127.0.0.1:50051"); v.SetDefault("plugins.dir","plugins"); v.SetDefault("plugins.autostart",$true); v.SetDefault("plugins.poll_second",2); if err:=v.ReadInConfig(); err!= $null { return Default(), $null } ; var c Config; if err:=v.Unmarshal(&c); err!= $null { return $null, err } ; return &c, $null }
func (c *Config) String() string { return fmt.Sprintf("HTTP=%s GRPC=%s PluginsDir=%s", c.HTTP.Listen, c.GRPC.Listen, c.Plugins.Dir) }