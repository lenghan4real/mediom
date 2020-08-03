package models

import (
	"testing"
)

func TestSearchMentionLogins(t *testing.T) {
	body := `@lenghan4real 你好啊 @monster @lenghan4real`
	logins := searchMentionLogins(body)
	if logins[0] != "lenghan4real" && logins[1] != "monster" {
		t.Error("not match result:", logins)
	}
}
