package stats

import 	"github.com/YusKom/bank/pkg/bank/types"

//Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	countPayments := types.Money(len(payments))
	sumPaymenys := types.Money(0)
	for _, payment := range payments {
		moneyPayments := payment.Amount
		sumPaymenys += moneyPayments
	}
	return sumPaymenys / countPayments
}
