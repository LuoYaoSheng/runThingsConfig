package model

// Rule 存储规则结构体
type Rule struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Code    string `json:"code"`
	Sn      string `json:"sn"`
	Content string `json:"content"` // 此处不直接使用 RuleContent
	//Content []RuleContent `json:"content"`
}

// RuleContent 规则具体内容
type RuleContent struct {
	Property  string      `json:"property"`
	Condition int         `json:"condition"`
	Value     interface{} `json:"value"` // 范围示例： [12.3,30.8] // 最小值 ~ 最大值
}
