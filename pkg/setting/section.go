package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Logger LoggerSetting `mapstructure:"logger"`
	Mysql  MysqlSetting  `mapstructure:"mysql"`
	Redis  RedisSetting  `mapstructure:"redis"`
	Smtp   SmtpSetting   `mapstructure:"smtp"`
	Kafka  KafkaSetting  `mapstructure:"kafka"`
	Jwt    JwtSetting    `mapstructure:"jwt"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type LoggerSetting struct {
	LogLevel   string `mapstructure:"logLevel"`
	FileName   string `mapstructure:"fileName"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
	Compress   bool   `mapstructure:"compress"`
}

type MysqlSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	UserName        string `mapstructure:"userName"`
	Password        string `mapstructure:"password"`
	DbName          string `mapstructure:"dbName"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
	PoolSize int    `mapstructure:"poolSize"`
}

type SmtpSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	From     string `mapstructure:"from"`
}

type KafkaSetting struct {
	Host  string `mapstructure:"host"`
	Port  int    `mapstructure:"port"`
	Topic string `mapstructure:"topic"`
}

type JwtSetting struct {
	TOKEN_HOUR_LIFESPAN int    `mapstructure:"TOKEN_HOUR_LIFESPAN"`
	JWT_EXPIRATION      string `mapstructure:"JWT_EXPIRATION"`
	API_SECRET          string `mapstructure:"API_SECRET"`
}
