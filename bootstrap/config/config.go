package config

type Config struct {
	Cors     Cors     `yaml:"cors"`
	Mysql    Mysql    `yaml:"mysql"`
	Redis    Redis    `yaml:"redis"`
	Password Password `yaml:"password"`
	Jwt      Jwt      `yaml:"jwt"`
	Server   Server   `yaml:"server"`
	Logs     Logs     `yaml:"logs"`
}

type Server struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	RunMode string `yaml:"runMode"`
}

type Logs struct {
	MaxAge     int    `yaml:"maxAge"`
	Compress   bool   `yaml:"compress"`
	Level      int    `yaml:"level"`
	Path       string `yaml:"path"`
	MaxSize    int    `yaml:"maxSize"`
	MaxBackups int    `yaml:"maxBackups"`
}

type Cors struct {
	AllowOrigins string `yaml:"allowOrigins"`
}

type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
	Query    string `yaml:"query"`
	LogMode  bool   `yaml:"logMode"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type Redis struct {
	ReadTimeout        int    `yaml:"readTimeout"`
	WriteTimeout       int    `yaml:"writeTimeout"`
	PoolSize           int    `yaml:"poolSize"`
	PoolTimeout        int    `yaml:"poolTimeout"`
	IdleCheckFrequency int    `yaml:"idleCheckFrequency"`
	Host               string `yaml:"host"`
	DialTimeout        int    `yaml:"dialTimeout"`
	Db                 int    `yaml:"db"`
	Port               int    `yaml:"port"`
	Password           string `yaml:"password"`
}

type Password struct {
	IncludeUppercase bool `yaml:"includeUppercase"`
	IncludeLowercase bool `yaml:"includeLowercase"`
	IncludeChars     bool `yaml:"includeChars"`
	IncludeDigits    bool `yaml:"includeDigits"`
	MinLength        int  `yaml:"minLength"`
	MaxLength        int  `yaml:"maxLength"`
}

type Jwt struct {
	AccessTokenExpireDuration  int    `yaml:"accessTokenExpireDuration"`
	RefreshTokenExpireDuration int    `yaml:"refreshTokenExpireDuration"`
	Secret                     string `yaml:"secret"`
	RefreshSecret              string `yaml:"refreshSecret"`
}
