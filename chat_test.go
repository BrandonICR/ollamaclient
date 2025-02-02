package ollamaclient_test

import (
	"strings"
	"testing"

	"github.com/BrandonICR/ollamaclient"
)

func TestCutThink(t *testing.T) {
	content := `<think>Okay, so I had this really weird dream last night, and I'm trying to figure out what it means.
			</think>Your dream of walking through a forest that shifts from peaceful to eerie can be interpreted as follows:`
	expected := "Your dream of walking through a forest that shifts from peaceful to eerie can be interpreted as follows:"

	got := ollamaclient.CutThink(content)

	if strings.Compare(expected, got) != 0 {
		t.Fatalf("expected [%s] but got [%s]", expected, got)
	}
}

func TestCutThink_CloseTagThinkNotPresent(t *testing.T) {
	content := "Your dream of walking through a forest that shifts from peaceful to eerie can be interpreted as follows:"
	expected := "Your dream of walking through a forest that shifts from peaceful to eerie can be interpreted as follows:"

	got := ollamaclient.CutThink(content)

	if strings.Compare(expected, got) != 0 {
		t.Fatalf("expected [%s] but got [%s]", expected, got)
	}
}

func TestCutThink_JustCloseTagThinkPresent(t *testing.T) {
	content := "<think>Okay, so I had this really weird dream last night, and I'm trying to figure out what it means.</think>"
	var expected string

	got := ollamaclient.CutThink(content)

	if strings.Compare(expected, got) != 0 {
		t.Fatalf("expected [%s] but got [%s]", expected, got)
	}
}
