package handler_test

import (
	"testing"

	"github.com/kieranajp/victim/pkg/handler"
)

func TestExtractingUsers(t *testing.T) {
	text := "hello <@W123> and <@W345> how goes it"

	users := handler.ExtractUsers(text)

	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}

	if users[0] != "<@W123>" {
		t.Errorf("Expected user <@W123>, got %s", users[0])
	}
}
