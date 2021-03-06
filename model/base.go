package model

// Eq2MqLog 日志
type Eq2MqLog struct {
	Sn       string                 `json:"sn"`       // 设备唯一编码
	Product  string                 `json:"product"`  // 产品key
	Protocol int64                  `json:"protocol"` // 协议
	Status   int64                  `json:"status"`   // 设备状态
	Content  map[string]interface{} `json:"content"`  // 带具体参数
	Title    string                 `json:"title"`    // 推送等标题
	Link     bool                   `json:"link"`     // 连接
}

// Eq2MqHeartbeat 心跳
type Eq2MqHeartbeat struct {
	Sn        string `json:"sn"`        // 设备唯一编码
	Product   string `json:"product"`   // 产品key
	Heartbeat int64  `json:"heartbeat"` // 心跳，单位[秒]
}

// Eq2MqCmd 下发指令
type Eq2MqCmd struct {
	Sn      string                 `json:"sn"`      // 设备唯一编码
	Content map[string]interface{} `json:"content"` // 带具体参数
}

// Eq2MqThreshold 设置阈值
type Eq2MqThreshold struct {
	Operate uint `json:"type"`    // 类别;0-新增;1-更新;2-删除
	Content Rule `json:"content"` // 带具体参数
}

// Eq2MqSynchroMQ 同步信息
type Eq2MqSynchroMQ struct {
	Product string                 `json:"product"` // 产品key
	Content map[string]interface{} `json:"content"` // 带具体参数
}
