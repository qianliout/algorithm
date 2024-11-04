package main

func main() {

}

func accountBalanceAfterPurchase(purchaseAmount int) int {
	n1 := purchaseAmount / 10
	n2 := (purchaseAmount + 9) / 10

	if purchaseAmount-n1*10 >= n2*10-purchaseAmount {
		return 100 - n2*10
	}
	return 100 - n1*10
}
