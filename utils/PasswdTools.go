package utils

import "golang.org/x/crypto/bcrypt"

func PasswdHash(pwd string) (string, error) {
	if bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func PasswdVerify(pwd, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd)) == nil
}
