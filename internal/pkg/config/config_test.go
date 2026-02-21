package config

import "testing"

func TestGetPostgresDsn(t *testing.T) {
	dsn := GetPostgresDsn()
	t.Logf("dsn: %s", dsn)
}

func TestGetPostgresReplicas(t *testing.T) {
	replicas := GetPostgresReplicas()
	for _, r := range replicas {
		t.Logf("replica: %s", r)
	}
}

func TestGetMinioPublicUrl(t *testing.T) {
	url := GetMinioPublicUrl()
	t.Logf("url: %s", url)
}

func TestGetUsers(t *testing.T) {
	users, _ := GetUsers()
	for _, u := range users {
		t.Logf("username: %s, password: %s", u.Username, u.Password)
	}
}
