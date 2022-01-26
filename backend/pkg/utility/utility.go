package utility

import "math"

func BankDepositInterset(amount uint32, rate float32, durationYear int) float32 {
	var result float32 = float32(amount)

	//复利
	// for i := 1; i <= duration_year; i++ {
	// 	result = result * (1 + rate/100)
	// }

	//单利
	result = result*(rate/100)*float32(durationYear) + float32(amount)
	return float32(result)
}

func MoneyDepreciation(amount float32, rate float32, durationYear int) float32 {
	var result float32 = amount
	//贬值
	for i := 1; i <= durationYear; i++ {
		result = result * (1 - rate/100)
	}
	return float32(result)
}

func HouseAppreciation(amount uint32, rate float32, durationYear int) float32 {
	var result float32 = float32(amount)
	//增值
	for i := 1; i <= durationYear; i++ {
		result = result * (1 + rate/100)
	}
	return float32(result)
}

//等额本金 （每月偿还的贷款本金一样）
func BankLoanEqualPrincipal(amount uint32, rate float32, durationYear uint32, periodicFee float32, oneTimeFee float32) (float32, float32) {

	monthlyRate := rate / 12 / 100

	periods := durationYear * 12

	var monthlyPrincipal float32 = float32(amount / periods)

	var totalIntersetAndGebyrer float32 = 0

	for i := 0; i < int(periods); i++ {
		periodsInterset := monthlyRate * (float32(amount) - (monthlyPrincipal * float32(i)))
		totalIntersetAndGebyrer = totalIntersetAndGebyrer + periodsInterset + periodicFee
	}

	totalPayment := float32(amount) + totalIntersetAndGebyrer + oneTimeFee

	return totalPayment, totalIntersetAndGebyrer
}

//等额本息 （每月偿还的金额一样）
func BankLoanEqualInstallments(amount uint32, rate float32, durationYear uint32, periodicFee float32, oneTimeFee float32) (float32, float32, float32) {

	periodicRate := rate / 12 / 100

	periods := durationYear * 12

	exponent := float32(math.Pow(float64((1 + periodicRate)), float64(periods)))

	periodocPaymentAmount := float32(amount)*((periodicRate*exponent)/(exponent-1)) + periodicFee

	totalPayment := periodocPaymentAmount*float32(periods) + oneTimeFee

	totalIntersetAndGebyrer := totalPayment - float32(amount)

	return periodocPaymentAmount, totalPayment, totalIntersetAndGebyrer
}

func TotalAnnualSpendAndSale(amount float32, durationYear int) float32 {

	var result float32 = amount*0.019 + 45000 + 3500*12*float32(durationYear) + 15000*float32(durationYear)

	return result
}

func TotalRentalIncome(rent int, durationYear int) float32 {
	return float32(rent*12*durationYear) - float32(rent)*0.2
}
