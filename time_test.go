package time

import (
	"fmt"
	"testing"
	"time"
)

type Money float64

type UserTransactions struct {
	HourRate     Money
	Transactions []Money
}

type MoneyInTime struct {
	Money Money
	Time  string
}

func TestUserPayment(t *testing.T) {

	userTxs := UserTransactions{
		HourRate:     30,
		Transactions: []Money{10, 20, 60, -60, 100, -200},
	}

	userTransactions := userTransactionsInTime(userTxs)
	userBalance := userBalance(userTxs)

	fmt.Println("user hour rate is: ", userTxs.HourRate)
	fmt.Printf("user transactions: %+v\n", userTransactions)
	fmt.Printf("user balance: %+v\n", userBalance)
}

func userTransactionsInTime(txs UserTransactions) []MoneyInTime {
	var transactionsInTime []MoneyInTime
	for i := 0; i < len(txs.Transactions); i++ {
		rateInHours := txs.Transactions[i] / txs.HourRate
		duration := time.Duration(rateInHours * 3600 * 1e9)
		transactionsInTime = append(transactionsInTime, MoneyInTime{
			Money: txs.Transactions[i],
			Time:  duration.String(),
		})
	}
	return transactionsInTime
}

func userBalance(txs UserTransactions) MoneyInTime {
	var total MoneyInTime
	userTransactionsWithTime := userTransactionsInTime(txs)
	for i := 0; i < len(userTransactionsWithTime); i++ {
		currentDuration, _ := time.ParseDuration(total.Time)
		newDuration, _ := time.ParseDuration(userTransactionsWithTime[i].Time)
		durationResult := currentDuration + newDuration
		total = MoneyInTime{
			Money: total.Money + userTransactionsWithTime[i].Money,
			Time:  durationResult.String(),
		}
	}
	return total
}
