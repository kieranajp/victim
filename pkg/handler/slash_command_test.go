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

func TestExcludingUsers(t *testing.T) {
	text := "hello <@W123> and <@W345> but not !<@W456> or !<@W678> how goes it"

	users := handler.ExtractExclusions(text)

	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}

	if users[0] != "<@W456>" {
		t.Errorf("Expected user <@W456>, got %s", users[0])
	}
}

func TestResolvingExclusions(t *testing.T) {
	text := "hello <@W123> and <@W345> but actually not !<@W345> or !<@W678> how goes it"

	users := handler.ExtractUsers(text)
	exclusions := handler.ExtractExclusions(text)
	resolved := handler.ResolveExclusions(users, exclusions)

	if len(resolved) != 1 {
		t.Errorf("Expected 1 user, got %d", len(resolved))
	}

	if resolved[0] != "<@W123>" {
		t.Errorf("Expected user <@W123>, got %s", resolved[0])
	}
}
