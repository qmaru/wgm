package models

// CommonModel 公共结构
type CommonModel struct {
	ID        int    `json:"id" db:"INTEGER;PRIMARY KEY"`
	CreatedAt int    `json:"created_at" db:"INTEGER;DEFAULT 0"`
	UpdatedAt int    `json:"updated_at" db:"INTEGER;DEFAULT 0"`
	Remark    string `json:"remark" db:"TEXT;DEFAULT ''"`
	State     int    `json:"state" db:"TINYINT;DEFAULT 1"`
}

// Users 所有节点
type Users struct {
	CommonModel
	Username string `json:"username" db:"TEXT;DEFAULT ''"`
	PriKey   string `json:"prikey" db:"TEXT;NOT NULL"`
	Pubkey   string `json:"pubkey" db:"TEXT;NOT NULL"`
}

// Routes 路由表
type Routes struct {
	CommonModel
	CIDR string `json:"cidr" db:"TEXT;NOT NULL"`
}

// Peers 节点信息
type Peers struct {
	CommonModel
	UserID              int    `json:"user_id" db:"INTEGER;NOT NULL"`
	PublicAddr          string `json:"public_addr" db:"TEXT"`
	PrivateAddr         string `json:"private_addr" db:"TEXT;NOT NULL"`
	Port                int    `json:"port" db:"INTEGER;DEFAULT 0"`
	AllowedIPs          string `json:"allowed_ips" db:"TEXT;NOT NULL"`
	MTU                 int    `json:"mtu" db:"INTEGER;DEFAULT ''"`
	DNS                 string `json:"dns" db:"TEXT;DEFAULT ''"`
	PersistentKeepalive int    `json:"keepalive" db:"INTEGER;DEFAULT 25"`
}
