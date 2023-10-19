package config

type Config struct {
	Logs     Logs     `yaml:"logs"`
	Cors     Cors     `yaml:"cors"`
	Mysql    Mysql    `yaml:"mysql"`
	Redis    Redis    `yaml:"redis"`
	Password Password `yaml:"password"`
	Jwt      Jwt      `yaml:"jwt"`
	Server   Server   `yaml:"server"`
}
type Redis struct {
	Host               string `yaml:"host"`
	DialTimeout        int    `yaml:"dialTimeout"`
	PoolSize           int    `yaml:"poolSize"`
	Port               int    `yaml:"port"`
	Password           string `yaml:"password"`
	Db                 int    `yaml:"db"`
	ReadTimeout        int    `yaml:"readTimeout"`
	WriteTimeout       int    `yaml:"writeTimeout"`
	PoolTimeout        int    `yaml:"poolTimeout"`
	IdleCheckFrequency int    `yaml:"idleCheckFrequency"`
}

type Password struct {
	IncludeChars     bool `yaml:"includeChars"`
	IncludeDigits    bool `yaml:"includeDigits"`
	MinLength        int  `yaml:"minLength"`
	MaxLength        int  `yaml:"maxLength"`
	IncludeUppercase bool `yaml:"includeUppercase"`
	IncludeLowercase bool `yaml:"includeLowercase"`
}

type Jwt struct {
	Secret                     string `yaml:"secret"`
	RefreshSecret              string `yaml:"refreshSecret"`
	AccessTokenExpireDuration  int    `yaml:"accessTokenExpireDuration"`
	RefreshTokenExpireDuration int    `yaml:"refreshTokenExpireDuration"`
}

type Server struct {
	Host    string `yaml:"host"`
	RunMode string `yaml:"runMode"`
	Port    int    `yaml:"port"`
}

type Logs struct {
	Compress   bool   `yaml:"compress"`
	Level      int    `yaml:"level"`
	Path       string `yaml:"path"`
	MaxSize    int    `yaml:"maxSize"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxAge     int    `yaml:"maxAge"`
}

type Cors struct {
	AllowOrigins string `yaml:"allowOrigins"`
}

type Mysql struct {
	Password        string `yaml:"password"`
	DbName          string `yaml:"dbName"`
	SslMode         string `yaml:"sslMode"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
	Port            int    `yaml:"port"`
	User            string `yaml:"user"`
	Host            string `yaml:"host"`
}
