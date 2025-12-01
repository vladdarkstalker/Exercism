package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) (res float32) {
    switch {
        case balance < float64(0):
        	res = 3.213
        case balance >= float64(0) && balance < float64(1000):
        	res = 0.5
        case balance >= float64(1000) && balance < float64(5000):
        	res = 1.621
        case balance >= float64(5000):
        	res = 2.475
    }
    return
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	p := float64(InterestRate(balance)) / 100
    return balance * p
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) (years int) {
    for balance < targetBalance {
        balance += Interest(balance)
        years++
    }
	return 
}
