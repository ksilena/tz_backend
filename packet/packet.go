package packet

import "fmt"

// Packet тестовый пакет, получаемый от клиента
type Packet struct {
	Price    int
	Quantity int
	Amount   int
	Object   int
	Method   int
}

func (p Packet) String() string {
	return fmt.Sprintf("price: %d, quantity: %d, amount: %d, object: %d, method: %d\n",
		p.Price, p.Quantity, p.Amount, p.Object, p.Method)
}
