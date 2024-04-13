package utils

import "testing"

func TestPassword(t *testing.T) {
	password := Password("e10adc3949ba59abbe56e057f20f883e", "123")
	t.Log(password)
}
