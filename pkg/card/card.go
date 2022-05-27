package card

import (
	"github.com/MSHE97/types/pkg/types"
)

func Issue() types.Card {
	return types.Card{
		Balance:  0,
		Currency: types.Currency("TJS"),
	}
}

func Withdraw(card types.Card) types.Card {
	card.Balance -= 100
	return card
}
