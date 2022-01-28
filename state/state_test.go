package state

import (
	"testing"
)

// Bruker testing pakken for å teste feilmeldinger.
func TestViewState(t *testing.T) {
	wanted := "[kylling rev korn mann ---vest  \\___/ ___________øst---]"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func TestPutinBoat(t *testing.T) {
	wanted := "[rev korn ---vest \\_mann_kylling_/ ___________øst---]"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func TestCrossRiver(t *testing.T) {
	wanted := "[rev korn ---vest \\___/ ___________øst--- mann kylling]"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
