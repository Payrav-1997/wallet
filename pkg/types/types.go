package types

// Money представляет собой денежную сумму в минималных единиц 
type Money int64

// PaymentCategory представляет 
type PaymentCategory string

// PaymentStatus статус плятежей.
type PaymentStatus string

// Предопределеные статусы платежа.
const (
	StatusOk       PaymentStatus = "OK"
	StatusFail     PaymentStatus = "FAIL"
	StatusProgress PaymentStatus = "INPROGRESS"
)

// Payment представляет игформацию о платиже.
type Payment struct {
	ID       string
	Amount   Money
	Category PaymentCategory
	Status   PaymentStatus
}

// Phone представляет информацию о  телефона
type Phone string

// Account  о счёте пользлователя.
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}