package cleaner

import (
	"strings"
	"testing"
)

func TestSurgicalTargetsIntegrity(t *testing.T) {
	// Verify that none of the targets contain 'globalStorage' or 'workspaceStorage'
	// to guarantee user chats and project contexts are strictly preserved.
	forbidden := []string{"globalstorage", "workspacestorage", "history", "conversations"}

	targets := []string{
		"Cookies",
		"Cookies-journal",
		"DIPS",
		"TransportSecurity",
		"blob_storage",
		"Session Storage",
	}

	for _, target := range targets {
		lower := strings.ToLower(target)
		for _, f := range forbidden {
			if strings.Contains(lower, f) {
				t.Fatalf("Target %s violates chat preservation rule by touching %s", target, f)
			}
		}
	}
}
