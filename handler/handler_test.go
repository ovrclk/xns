package handler

import "testing"

func TestResolve(t *testing.T) {
	h := &Handler{RootDomain: "example.com"}
	tests := []struct {
		input  string
		output string
	}{
		{"127.0.0.1.example.com.", "127.0.0.1"},
		{"foo.127.0.0.1.example.com.", "127.0.0.1"},
		{"foo.boo.127.0.0.1.example.com.", "127.0.0.1"},
		{"bar.foo.boo.127.0.0.1.example.com.", "127.0.0.1"},
		{"foo.127-0-0-1.example.com.", "127.0.0.1"},
		{"bar.foo.127-0-0-1.example.com.", "127.0.0.1"},
		{"bar.foo.baz.127-0-0-1.example.com.", "127.0.0.1"},
		// error cases
		{"foo.boo.example.com.", ""},
		{"foo.127.0.0.1.boo.example.com.", ""},
		{"foo.127.0.0.1.bad.com.", ""},
	}
	for _, v := range tests {
		if addr := h.Resolve(v.input); addr != v.output {
			t.Fatalf("Testing for: %s | Expected: %s, Got: %s", v.input, v.output, addr)
		}
	}
}
