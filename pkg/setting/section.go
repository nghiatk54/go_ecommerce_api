package setting

type Config struct {
	MySqlSetting MySqlSetting `mapstructure:"mysql"`
	LoggerSetting LoggerSetting `mapstructure:"logger"`
	ServerSetting ServerSetting `mapstructure:"server"`
	RedisSetting RedisSetting `mapstructure:"redis"`
}

type MySqlSetting struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName string `mapstructure:"dbname"`
	MaxIdleConns int `mapstructure:"maxIdleConns"`
	MaxOpenConns int `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	Level string `mapstructure:"level"`
	Filename string `mapstructure:"filename"`
	MaxSize int `mapstructure:"maxSize"`
	MaxBackups int `mapstructure:"maxBackups"`
	MaxAge int `mapstructure:"maxAge"`
	Compress bool `mapstructure:"compress"`
}

type ServerSetting struct {
	Port string `mapstructure:"port"`
}

type RedisSetting struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db int `mapstructure:"db"`
	PoolSize int `mapstructure:"poolSize"`
}