package defective

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/Kwynto/mech-exp/internal/intypes"
	"github.com/Kwynto/mech-exp/pkg/incolor"
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

func spaceSimbol(k int) string {
	if (k / 10) >= 1 {
		return ""
	}
	return " "
}

func preAnalize(slStInput []intypes.TStGame, iGame int) int {
	var (
		slWork []intypes.TStGame
		// iAllNumbers int = 0
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
		// iWinCount := len(stGame.Wins)
		iWrongCount := len(stGame.Wrong)
		// iAllNumbers = iAllNumbers + iWinCount + iWrongCount
		iAllWrong = iAllWrong + iWrongCount
	}

	iBorder = iAllWrong / NUMBERS_IN_GAME
	if (iAllWrong % NUMBERS_IN_GAME) > 0 {
		iBorder = iBorder + ADD_TO_BORER
	}
	iBorder = iBorder + ADD_TO_BORER

	return iBorder
}

func startAnalize(slStInput []intypes.TStGame, iGame, iBorder int) {
	var slWork []intypes.TStGame

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
		fmt.Println(incolor.StringBlueH("Анализ всех тиражей."))
		slWork = slStPrepear
	} else {
		sMsg1 := fmt.Sprintf("Анализ %d тиражей.", iGame)
		fmt.Println(incolor.StringBlueH(sMsg1))
		slWork = slStPrepear[0:iGame]
	}

	fmt.Println("")

	// for _, v := range slWork {
	// 	fmt.Println(v.Game)
	// }

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

	fmt.Println(incolor.StringGreenH("Премиальные номера:"))
	for i := range NUMBERS_IN_GAME {
		k1 := i + 1
		for k, v := range mStatNumbers {
			if k == k1 {
				if v.Wrong < iBorder {
					if v.PremiumWin > (iBorder * 4) {
						sMsg := fmt.Sprintf("Номер %s%s premium = %d", spaceSimbol(k), incolor.StringGreen("%d", k), v.PremiumWin)
						fmt.Println(sMsg)
					}
				}
			}
		}
	}

	fmt.Println("")

	fmt.Println(incolor.StringRedH("Номера зоны риска:"))
	for i := range NUMBERS_IN_GAME {
		k1 := i + 1
		for k, v := range mStatNumbers {
			if k == k1 {
				if v.Wrong >= iBorder {
					sMsg := fmt.Sprintf("Номер %s%s wrong = %d", spaceSimbol(k), incolor.StringRed("%d", k), v.Wrong)
					fmt.Println(sMsg)
				}
			}
		}
	}

}

func Start(slStInput []intypes.TStGame) {
	var sGame, sBorder string

	SlStGames = slStInput

	fmt.Println(incolor.StringBlue("Дефектный анализ:"))

	fmt.Print(incolor.StringMagenta("Кол-во последних тиражей для анализа (0 для всех тиражей) > "))
	fmt.Scanf("%v\n", &sGame)
	iGame, err1 := strconv.Atoi(sGame)
	if err1 != nil {
		fmt.Println("Conversion failed:", err1)
	}

	iPreBorder := preAnalize(slStInput, iGame)

	fmt.Print(incolor.StringMagenta("Граница повторений (%d или более) (сейчас рекомендация %d повторений) > ", DEFAULT_BORDER, iPreBorder))
	fmt.Scanf("%v\n", &sBorder)
	iBorder, err2 := strconv.Atoi(sBorder)
	if err2 != nil {
		fmt.Println("Conversion failed:", err2)
	}

	if iBorder < DEFAULT_BORDER {
		iBorder = DEFAULT_BORDER
	}

	fmt.Println("")

	startAnalize(slStInput, iGame, iBorder)
}
