package tests

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPasswordLen(t *testing.T) {
	str := []byte("$2a$04$Px3Oum35Ln0ddVmeBI2Miu5hnVDaYGWlfa83UQY0zR2L2zfLeO2Aa")
	pass := []byte("123456")
	err := bcrypt.CompareHashAndPassword(str, pass)
	if err != nil {
		t.Error(err)
	}
}
