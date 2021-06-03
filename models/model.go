package models

const (
	UsersTable   = "users"
	ServersTable = "servers"
	RulesTable   = "rules"
	RulemapTable = "rulemap"
)

// CommonModel 公共结构
type CommonModel struct {
	ID        int `json:"id" db:"INTEGER;PRIMARY KEY"`
	CreatedAt int `json:"created_at" db:"INTEGER;DEFAULT 0"`
	UpdatedAt int `json:"updated_at" db:"INTEGER;DEFAULT 0"`
	Status    int `json:"status" db:"CHAR(1);DEFAULT 1"`
}

// Users 用户表
//	ServerID: 服务器ID
//	Username: 用户名
//	Prikey: 私钥
//	Pubkey: 公钥
//	IP: 分配IP
//	DefaultRule: 默认Rule
//	PersistentKeepalive: 重连间隔时间
type Users struct {
	CommonModel
	ServerID            int    `json:"server_id" db:"INTEGER;NOT NULL"`
	Username            string `json:"username" db:"TEXT;DEFAULT ''"`
	PriKey              string `json:"prikey" db:"TEXT;NOT NULL"`
	Pubkey              string `json:"pubkey" db:"TEXT;NOT NULL"`
	IP                  string `json:"ip" db:"TEXT;NOT NULL"`
	DefaultRule         string `json:"default_rule" db:"TEXT;DEFAULT ''"`
	IsExtra             int    `json:"is_extra" db:"TINYINT;DEFAULT 0"`
	PersistentKeepalive int    `json:"keepalive" db:"INTEGER;DEFAULT 25"`
}

// Servers 中心服务器表
//	Title: 服务器名字
//	Address: 服务器监听地址
//	Port: 服务器监听端口
//	LanIP / LanNetmask: 局域网划分
//	MTU / DNS 局域网组网一般无需设置
type Servers struct {
	CommonModel
	Title      string `json:"title" db:"TEXT;NOT NULL"`
	Address    string `json:"address" db:"TEXT;NOT NULL"`
	Port       int    `json:"port" db:"INTEGER;NOT NULL"`
	LanIP      string `json:"lan_ip" db:"TEXT;NOT NULL"`
	LanNetmask string `json:"lan_netmask" db:"TEXT;NOT NULL"`
	MTU        int    `json:"mtu" db:"INTEGER;DEFAULT ''"`
	DNS        string `json:"dns" db:"TEXT;DEFAULT ''"`
}

// Rules AllowedIP 规则表
type Rules struct {
	CommonModel
	AllowedIP string `json:"allowedip" db:"TEXT;DEFAULT ''"`
}

// RuleMap 额外规则映射
type RuleMap struct {
	CommonModel
	UserID int `json:"user_id" db:"INTEGER;NOT NULL"`
	RuleID int `json:"rule_id" db:"INTEGER;NOT NULL"`
}
