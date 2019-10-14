package users

import "golang.org/x/crypto/bcrypt"

// passwordHash ハッシュ化したパスワードの作成
func passwordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}
