package types

// Money прелставляет собой денежную сумму в минимальных единицах (центыб копейки, дирамы и т.д.).
type Money int64

// Currency представляет код валюты.
type Currency string

// PAN представляет номер карты.
type PAN string

// Category прелставляет собой категорию, в которой был совершён платёж (авто, аптеки, рестораны и т.д.).
type Category string

// Status прелставляет собой статус платежей.
type Status string

// Предопределённые статусы платежей.
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Коды валюты.
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// Payment предоставляет информаницию о платеже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

// Card представляет информацию о платёжной карте.
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}
