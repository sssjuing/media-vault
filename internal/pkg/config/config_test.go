package config

import "testing"

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
