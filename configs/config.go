package configs

type ServerConfig struct {
	Name         string             `mapstructure:"name"`
	Mode         string             `mapstructure:"mode"`
	TimeLocation string             `mapstructure:"time_location"`
	ConfigPath   string             `mapstructure:"config_path"`
	MySQLConfig  *MySQLConf         `mapstructure:"mysql"`
	ConsulConfig *ConsulConf        `mapstructure:"consul"`
	RedisConfig  RedisConf          `mapstructure:"redis"`
	EmailConfig  EmailConf          `mapstructure:"email"`
	LarkConfig   map[uint]*LarkConf `mapstructure:"lark"`
}

type MySQLConf struct {
	DataSourceName  string `mapstructure:"dsn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}

type ConsulConf struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type RedisConf struct {
	Addr string `mapstructure:"addr"`
}

type EmailConf struct {
	Host           string `mapstructure:"host" json:"host"`
	Port           int    `mapstructure:"port" json:"port"`
	Password       string `mapstructure:"password" json:"password"`
	SendPersonName string `mapstructure:"send_person_name"`
	SendPersonAddr string `mapstructure:"send_person_addr"`
}

type LarkConf struct {
	AppID     string `mapstructure:"app_id"`
	AppSecret string `mapstructure:"app_secret"`
	Type      int    `mapstructure:"type"`
}
