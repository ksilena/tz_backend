package message

import "fmt"

type Message struct {
	Price    int
	Quantity int
	Amount   int
	Object   int
	Method   int
}

func (m Message) String() string {
	return fmt.Sprintf("price: %d\n quantity: %d\n amount: %d\n object: %d\n method: %d\n",
		m.Price, m.Quantity, m.Amount, m.Object, m.Method)
}
