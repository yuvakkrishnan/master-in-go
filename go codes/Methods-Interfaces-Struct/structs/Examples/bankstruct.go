package main

import (
	"fmt"
)

type bankDatas struct {
	id        int
	name      string
	accountno int
	balance   float64
}

type maintanceofAc struct {
	deposit  float64
	withdraw float64
}

func (holder bankDatas) accountHolder() string {
	return holder.name + " Welocome to BitCoin Bank"
}
func (details bankDatas) accountDetails() float64 {
	result := details.balance
	return result
}
func (cash maintanceofAc) cashDeposit(bankDatas bankDatas) float64 {
	DepositAmt := bankDatas.balance + cash.deposit
	return DepositAmt
}

func (cash maintanceofAc) cashWithdraw(bankDatas bankDatas) float64 {
	withdrawAmt := bankDatas.balance - cash.withdraw
	return withdrawAmt
}
func main() {
	var bankData bankDatas
	var cashMaintain maintanceofAc
	bankData = bankDatas{id: 1, name: "Yuvakkrishnan", accountno: 203018717134, balance: 0}

	cashMaintain = maintanceofAc{deposit: 10000, withdraw: 5000}
	fmt.Println("Account-Holder", bankData.accountHolder())
	fmt.Println("Total Balance", bankData.accountDetails())
	fmt.Println("Deposit amount:", cashMaintain.cashDeposit(bankData))
	fmt.Println("Withdraw amount:", cashMaintain.cashWithdraw(bankData))

}
