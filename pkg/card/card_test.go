package card

import (
	"fmt"

	"github.com/MSHE97/types/pkg/types"
)

func ExampleWitdraw_success() {
	card := &types.Card{
		Balance:  100,
		Activity: types.Onn,
	}
	Withdraw(card)

	fmt.Printf("%v", card.Balance)

	// Output:
	// 0
}
