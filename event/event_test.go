package event

import "testing"

func TestPut(t *testing.T) {
	wanted := "[rev korn mann ---\\ \\_kylling__/ ___________/---]"
	got := Put("kylling")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}
func TestCross(t *testing.T) {
	wanted := "[rev korn ---\\ \\___/ ___________/--- mann_kylling]"
	got := Cross("mann_kylling")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}

}
