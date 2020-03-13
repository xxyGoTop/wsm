package config

import (
	"crypto/aes"

	"github.com/xxyGoTop/wsm/internal/lib/dotenv"
)

var (
	ModeProduction = "production"
)

type common struct {
	MachineId int64  `json:"machine_id"` // 机器ID 分布式部署
	Mode      string `json:"mode"`       //环境变量 区分生产环境和开发环境
	Exiting   bool   `json:"exiting"`    // 进程是否退出，用于优雅退出进程
	Secret    string `json:"secret"`     // 秘钥，加密
}

var Common *common

func init() {
	Common = &common{}
	Common.Mode = dotenv.GetByDefault("GO_MODE", ModeProduction)
	Common.MachineId = dotenv.GetInt64ByDefault("MACHINE_ID", 1)
	Common.Secret = dotenv.GetByDefault("SECRET", "astaxie12798akljzmknm.ahkjkljl;k")

	k := len(Common.Secret)

	switch k {
	case 32:
		break
	default:
		panic(aes.KeySizeError(k))
	}
}
