package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type QQType struct {
	Enable   bool   `json:"enable"` // 是否启用
	AppID    string `json:"app_id"`
	Key      string `json:"key"`
	Redirect string `json:"redirect"` // 登录之后的回调地址
	WebPath  string `json:"webPath"`  // 点击跳转的路径
}

func (q QQType) GetPath() string {
	if q.Key == "" || q.AppID == "" || q.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_uri=%s", q.AppID, q.Redirect)
}

// Scan 取出来的时候的数据
func (c *QQType) Scan(val interface{}) error {
	err := json.Unmarshal(val.([]byte), c)
	if err != nil {
		return err
	}
	return nil
}

// Value 入库的数据
func (c QQType) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
