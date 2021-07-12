package conf

// LogTransfer 全局配置
//type name struct {
//
//}

type LogTransferCfg struct {
	KafkaCfg `ini:"kafka"`
	EsCfg    `ini:"es"`
}

// kafka ...
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

// EsCfg ...
type EsCfg struct {
	Address  string `ini:"address"`
	ChanSize int    `ini:"chan_size"`
	Nums     int    `ini:"nums"`
}
