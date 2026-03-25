package defective

import (
	"slices"

	"github.com/Kwynto/mech-exp-2/internal/intypes"
)

const (
	MAX_NUMBER   = 40
	PREMIUM_WIN  = 16
	ORDINARI_WIN = 38

	DEFAULT_BORDER  = 2
	NUMBERS_IN_GAME = 40
	ADD_TO_BORER    = 1
)

var SlStGames []intypes.TStGame

func initMapNumbers() intypes.TMapNembers {
	initMap := make(intypes.TMapNembers, 40)
	iCount := MAX_NUMBER + 1
	for i := 1; i < iCount; i++ {
		initMap[i] = intypes.TStStatNumber{
			PremiumWin:  0,
			OrdinariWin: 0,
			Wrong:       0,
		}
	}

	return initMap
}

func PreAnalize(slStInput []intypes.TStGame, iGame int) int {
	var (
		slWork    []intypes.TStGame
		iAllWrong int = 0
		iBorder   int = 0
	)

	slStPrepear := slices.Clone(slStInput)
	slices.Reverse(slStPrepear)

	if iGame < 0 {
		iGame = 0
	}

	iLenSlIn := len(slStPrepear)
	if iLenSlIn < iGame {
		iGame = iLenSlIn
	}

	if iGame == 0 {
		slWork = slStPrepear
	} else {
		slWork = slStPrepear[0:iGame]
	}

	for _, stGame := range slWork {
		iWrongCount := len(stGame.Wrong)
		iAllWrong = iAllWrong + iWrongCount
	}

	iBorder = iAllWrong / NUMBERS_IN_GAME
	if (iAllWrong % NUMBERS_IN_GAME) > 0 {
		iBorder = iBorder + ADD_TO_BORER
	}
	iBorder = iBorder + ADD_TO_BORER

	return iBorder
}

func StartAnalize(slStInput []intypes.TStGame, iGame, iBorder int) (intypes.TPremiumNumber, intypes.TRiskNumbers) {
	var slWork []intypes.TStGame

	var (
		slPremiumNumber intypes.TPremiumNumber
		slRiskNumbers   intypes.TRiskNumbers
	)

	slStPrepear := slices.Clone(slStInput)
	slices.Reverse(slStPrepear)

	if iGame < 0 {
		iGame = 0
	}

	iLenSlIn := len(slStPrepear)
	if iLenSlIn < iGame {
		iGame = iLenSlIn
	}

	if iGame == 0 {
		slWork = slStPrepear
	} else {
		slWork = slStPrepear[0:iGame]
	}

	mStatNumbers := initMapNumbers()

	for _, stGame := range slWork {
		for i01, iWinNum := range stGame.Wins {
			if i01 < PREMIUM_WIN {
				tempMStatNumber := mStatNumbers[iWinNum]
				tempMStatNumber.PremiumWin = tempMStatNumber.PremiumWin + 1
				mStatNumbers[iWinNum] = tempMStatNumber
			} else if (i01 >= PREMIUM_WIN) && (i01 < ORDINARI_WIN) {
				tempMStatNumber := mStatNumbers[iWinNum]
				tempMStatNumber.OrdinariWin = tempMStatNumber.OrdinariWin + 1
				mStatNumbers[iWinNum] = tempMStatNumber
			} else {
				tempMStatNumber := mStatNumbers[iWinNum]
				tempMStatNumber.Wrong = tempMStatNumber.Wrong + 1
				mStatNumbers[iWinNum] = tempMStatNumber
			}
		}

		for _, iWrongNum := range stGame.Wrong {
			tempMStatNumber := mStatNumbers[iWrongNum]
			tempMStatNumber.Wrong = tempMStatNumber.Wrong + 1
			mStatNumbers[iWrongNum] = tempMStatNumber
		}
	}

	for i := range NUMBERS_IN_GAME {
		k1 := i + 1
		for k, v := range mStatNumbers {
			if k == k1 {
				if v.Wrong < iBorder {
					if v.PremiumWin > (iBorder * 4) {
						slPremiumNumber = append(slPremiumNumber, intypes.TCharacteristicNumder{
							Name: k,
							Sum:  v.PremiumWin,
						})
					}
				}
			}
		}
	}

	for i := range NUMBERS_IN_GAME {
		k1 := i + 1
		for k, v := range mStatNumbers {
			if k == k1 {
				if v.Wrong >= iBorder {
					slRiskNumbers = append(slRiskNumbers, intypes.TCharacteristicNumder{
						Name: k,
						Sum:  v.Wrong,
					})
				}
			}
		}
	}

	return slPremiumNumber, slRiskNumbers
}
