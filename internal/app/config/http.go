package config

import "github.com/xxyGoTop/wsm/internal/lib/dotenv"

type TLS struct {
	Cert string `json:"cert"` // 证书文件
	Key  string `json:"key"`  // Key 文件
}

type http struct {
	Domain string `json:"domain"` // 用户端 API 绑定的域名, 例如 https://example.com
	Port   string `json:"port"`   // 用户端 API 监听端口
	Secret string `json:"secret"` // 用户端密钥，用于加密/解密 token
}

var Http http

func init() {
	Http.Port = dotenv.GetByDefault("PORT", "8080")
	Http.Secret = dotenv.GetByDefault("JWT_TOKEN", "some_str_for_jwt_token")
}