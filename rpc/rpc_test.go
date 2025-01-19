package rpc_test

import (
	"testing"

	"github.com/sagnik395/vanaralsp/rpc"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})

	if expected == actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
