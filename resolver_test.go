package consul_go

import (
	"testing"
)

func Test_consulWatcher_Next(t *testing.T) {
	resolver := NewConsulResolver("127.0.0.1:8500", "grpc-go")
	resolve, err := resolver.Resolve("")
	if err != nil {
		t.Fatal(err)
	}
	next, err := resolve.Next()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("grpc-go instances is %d", len(next))
	for _, n := range next {
		t.Log(n.Addr)
	}
}
