package state

import (
	"testing"
)

// Bruker testing pakken for å teste feilmeldinger.
func TestViewState(t *testing.T) {
	wanted := "[kylling rev korn mann ---\\ \\___/ ___________/---]"
	state := ViewState("")
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func TestPutinBoat(t *testing.T) {
	wanted := "[rev korn ---\\ \\_mann_kylling_/ ___________/---]"
	state := PutinBoat("mann_kylling")
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func TestCrossRiver(t *testing.T) {
	wanted := "[rev korn ---\\ \\___/ ___________/--- mann_kylling]"
	state := CrossRiver("mann_kylling")
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
