package socks5

// CredentialStore 用户账号密码方式验证接口
type CredentialStore interface {
	Valid(username, password string) bool
}

// MemoryCredentials 基于内存方式的验证存储
type MemoryCredentials map[string]string

func (m MemoryCredentials) Valid(username, password string) bool {
	pass, ok := m[username]
	if !ok {
		return false
	}
	return pass == password

}
