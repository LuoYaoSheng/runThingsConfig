package config

type InfluxdbConf struct {
	Addr      string
	Username  string
	Password  string
	Database  string
	Precision string
	Prefix    string
}

type RabbitMqConf struct {
	Mqurl     string
	QueueName string
	Exchange  string
	Key       string
}

type RedisConf struct {
	Addr     string
	Password string
	DB       int
}

type MqttConf struct {
	Broker   string
	Username string
	Password string
	Topics   []string
	Qos      byte
}

type EmqxConf struct {
	Url  string
	User string
	Pass string
}

type OneNetConf struct {
	Key string
}

type AepConf struct {
	AppKey    string
	AppSecret string
	ProductId int64
	MasterKey string
}
