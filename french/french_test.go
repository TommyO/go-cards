package french

import(
	"fmt"
)

func Example() {
	cards := NewDecks(2, 4)

	fmt.Println(len(cards))

	// Output: 108
}
