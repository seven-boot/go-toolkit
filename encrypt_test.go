package toolkit

import (
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	password, err := GeneratePassword("123456")
	if err != nil {
		t.Error(err)
	}
	t.Log(password)
}

func TestComparePassword(t *testing.T) {
	equal, err := ComparePassword("123456", "$2a$10$WuZuSOAWgcH3bPNC9QdF3uVEvvL72EeyV9vq0qTk2ckgxOq6X.CC.")
	if equal {
		t.Log("密码正确")
		return
	}
	t.Error(err)
}
