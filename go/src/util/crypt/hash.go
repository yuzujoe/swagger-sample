package crypt

import "golang.org/x/crypto/bcrypt"

// PasswordHash ハッシュ化したパスワードの作成
func PasswordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}
