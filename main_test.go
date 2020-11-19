package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	pb "github.com/3ventic/twirphelloworld/rpc"
)

// InitTestServer initializes a test server for HelloWorld and returns its address
func InitTestServer() string {
	handler := pb.NewHelloWorldServer(&HelloWorldServer{})
	server := httptest.NewServer(handler)
	return server.URL
}

func TestHello(t *testing.T) {
	url := InitTestServer()
	clients := map[string]pb.HelloWorld{
		"json": pb.NewHelloWorldJSONClient(url, http.DefaultClient),
		"pb":   pb.NewHelloWorldProtobufClient(url, http.DefaultClient),
	}

	for typ, client := range clients {
		t.Run(typ, func(t *testing.T) {
			ctx := context.Background()
			result, err := client.Hello(ctx, &pb.HelloReq{
				Subject: "test",
			})
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			if result.Text != "Hello test" {
				t.Errorf("result didn't match 'Hello test', was '%s'", result.Text)
			}
		})
	}
}
