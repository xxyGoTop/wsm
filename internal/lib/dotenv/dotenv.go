package dotenv

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/xxyGoTop/wsm/internal/lib/fs"
)

var (
	RootDir string // 当前运行的二进制所在的目录
	loaded  bool   // 是否已初始化过
)

func init() {
	if err := Load(); err != nil {
		log.Panic(err)
	}
}

// Load env file
func Load() (err error) {
	if loaded == true {
		return
	}

	defer func() {
		if err == nil {
			loaded = true
		}
	}()

	ex, err := os.Executable()

	if err != nil {
		log.Panicln(err)
	}

	exPathDir := filepath.Dir(ex)
	RootDir = exPathDir

	dotEnvFilePath := path.Join(RootDir, ".env")

	if exists, err := fs.PathExists(dotEnvFilePath); err != nil {
		return err
	} else if !exists {
		return nil
	}

	fmt.Println(fmt.Sprintf("加载环境变量文件: `%s`", color.GreenString(dotEnvFilePath)))

	err = godotenv.Load(dotEnvFilePath)

	return
}

func Get(key string) string {
	if loaded == false {
		_ = Load()
	}
	return os.Getenv(key)
}

func GetIntByDefault(key string, defaultValue int) int {
	val := GetByDefault(key, fmt.Sprintf("%d", defaultValue))

	result, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func GetInt64ByDefault(key string, defaultValue int64) int64 {
	val := GetByDefault(key, fmt.Sprintf("%d", defaultValue))

	result, err := strconv.ParseInt(val, 0, 10)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func GetArrayByDefault(key string, defaultValue []string) []string {
	val := GetByDefault(key, fmt.Sprintf("%s", strings.Join(defaultValue, ",")))

	arr := strings.Split(val, ",")

	var result []string

	for _, val := range arr {
		result = append(result, strings.TrimSpace(val))
	}

	return result
}

func GetByDefault(key string, defaultValue string) string {
	if loaded == false {
		_ = Load()
	}

	result := os.Getenv(key)

	if result == "" {
		return defaultValue
	} else {
		return result
	}
}
