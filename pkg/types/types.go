package types

// Money представляет собой денежную сумму в минималных единиц 
type Money int64

// PaymentCategory представляет 
type PaymentCategory string

// PaymentStatus статус плятежей.
type PaymentStatus string

// Предопределеные статусы платежа.
const (
	PaymentStatusOk       PaymentStatus = "OK"
	PaymentStatusFail     PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

// Payment представляет игформацию о платиже.
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

// Phone представляет информацию о  телефона
type Phone string

// Account  о счёте пользлователя.
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}

type Favorite struct {
	ID        string
	AccountID int64
	Name      string
	Amount    Money
	Category  PaymentCategory
}