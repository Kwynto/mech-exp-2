package intypes

type TStGame struct {
	Game  int
	Wins  []int
	Wrong []int
}

type TStStatNumber struct {
	PremiumWin  int
	OrdinariWin int
	Wrong       int
}

type TMapNembers map[int]TStStatNumber

type TCharacteristicNumder struct {
	Name int
	Sum  int
}

type TPremiumNumber []TCharacteristicNumder
type TRiskNumbers []TCharacteristicNumder
