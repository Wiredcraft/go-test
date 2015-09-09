package httpLogin

import "testing"

func TestLogin(t *testing.T) {
	cases := []struct {
		url, name, password string
	}{
		{"http://localhost:3000/api/login", "bob", "some_wrong_pwd"},
		{"http://localhost:3000/api/login", "alice", "secret"},
	}
	for i, c := range cases {
		result := Login(c.url, c.name, c.password)
		if i == 0 && result != false {
			t.Errorf("Expected false , got %t", result)
		}
		if i == 1 && result != true {
			t.Errorf("Expected true , got %t", result)
		}
	}
}
