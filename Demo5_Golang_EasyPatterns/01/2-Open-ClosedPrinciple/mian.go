package main

import "fmt"

type AbstractBanker interface {
	DoBusiness()
}

// DepositBanker 存款业务
type DepositBanker struct{}

func (db *DepositBanker) DoBusiness() {
	fmt.Println("存款业务")
}

// TransferAccountsBanker 转账业务
type TransferAccountsBanker struct{}

func (tab *TransferAccountsBanker) DoBusiness() {
	fmt.Println("转账业务")
}

// BankerBusiness 实现一个架构层（基于抽象层进行业务封装 - 针对interface接口进行封装）
func BankerBusiness(ab AbstractBanker) {
	// 通过接口向下来调用（多态的现象）
	ab.DoBusiness()
}

func main() {
	BankerBusiness(&DepositBanker{})
	BankerBusiness(&TransferAccountsBanker{})
}
