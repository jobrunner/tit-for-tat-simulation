package main

type Decider interface {
	// Decides whether C (cooperation) or D (defection) is selected
	Decide(opponentLastChoice string) string
	// The type of strategy
	Kind() string
}

// Always cooperate
type AlwaysCooperate struct{}

func (ac AlwaysCooperate) Decide(opponentLastChoice string) string {
	return "C"
}

func (ac AlwaysCooperate) Kind() string {
	return "AC"
}

// Always defect
type AlwaysDefect struct{}

func (ad AlwaysDefect) Decide(opponentLastChoice string) string {
	return "D"
}

func (ad AlwaysDefect) Kind() string {
	return "AD"
}

// Tit-for-Tat
type TitForTat struct{}

func (tft TitForTat) Decide(opponentLastChoice string) string {
	// If the opponent betrays, the agent also chooses defection (D).
	// Otherwise he cooperates (C).
	if opponentLastChoice == "D" {
		return "D"
	}

	return "C"
}

func (tft *TitForTat) Kind() string {
	return "TfT"
}
