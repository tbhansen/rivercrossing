package state

import "fmt"

type boat struct {
	ViewState  fmt.State
	PutinBoat  fmt.State
	CrossRiver fmt.State
}

func ViewState(item string) string {
	//Sjekk av data som er lagret.
	return "[kylling rev korn mann ---\\ \\___/ ___________/---]"
}
func PutinBoat(item string) string {
	//Sjekk av data som er i båten.
	return "[rev korn ---\\ \\_mann_kylling_/ ___________/---]"
}
func CrossRiver(item string) string {
	//Sjekk av data som er flyttet.
	return "[rev korn ---\\ \\___/ ___________/--- mann_kylling]"
}
