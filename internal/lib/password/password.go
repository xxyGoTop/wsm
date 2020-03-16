package password

import "golang.org/x/crypto/bcrypt"

const (
	prefix = "ped&13()%"
	suffix = "d;'^3@!#"
)

// 生成密码
func Generate(plaintext string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(prefix+plaintext+suffix), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func Verify(plaintext, hash string) bool  {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(prefix+plaintext+suffix))

	if err != nil {
		return false
	} else {
		return true
	}
}