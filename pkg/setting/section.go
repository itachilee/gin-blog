package setting

import "time"

type App struct {
	PageSize        int
	MaxPageSize     int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}
type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Port        string
	Name        string
	TablePrefix string
}

type Redis struct {
	Addr     string
	Username string
	Password string
	Db       int
}

type Mqtt struct {
	Host     string
	Port     int
	Username string
	Password string
}

type JWT struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type Email struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string `delim:"|" `
}
