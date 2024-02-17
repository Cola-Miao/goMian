package utils

import "golang.org/x/crypto/bcrypt"

func Encrypt(s []byte) ([]byte, error) {
	encode, err := bcrypt.GenerateFromPassword(s, bcrypt.DefaultCost)
	return encode, err
}

func CompareEncode(hash, str []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, str)
	return err
}
