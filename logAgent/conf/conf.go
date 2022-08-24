package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}

type EtcdConf struct {
	Addr    string `ini:"addr"`
	Key     string `ini:"log_agent_conf"`
	Timeout int    `ini:"timeout"`
}

type KafkaConf struct {
	Addr    string `ini:"addr"`
	MaxSize int    `ini:"chan_max_size"`
	//Topic string `ini:"topic"`
}

type TaillogConf struct {
	FileName string `ini:"fileName"`
}
