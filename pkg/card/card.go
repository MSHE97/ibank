package card

import (
	"github.com/MSHE97/types/pkg/types"
)

func Issue() types.Card {
	return types.Card{
		Activity: types.Onn,
		Balance:  0,
		Currency: types.Currency("TJS"),
	}
}

func Withdraw(card *types.Card) {
	card.Balance -= 100
}
