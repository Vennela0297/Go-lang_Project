package main

import (
	"fmt"
)

type bank struct {
	cus []customer
}

func (ba *bank) numofcustomers() int {
	return len(ba.cus)
}

func (ba *bank) printdata() {
	fmt.Println((ba.cus))

	for _, cusdetail := range ba.cus {
		fmt.Println("customer  name", cusdetail.cname)
		fmt.Println("customer accountnumber", cusdetail.caccno)
		fmt.Println("customer ssn", cusdetail.cssn)
		fmt.Println("customer amount", cusdetail.amount)
	}
}
func (ba *bank) addcustomer() {
	newcus := &customer{}
	fmt.Print("enter customer name :")
	fmt.Scan(&newcus.cname)
	fmt.Print("enter customer ssn :")
	fmt.Scan(&newcus.cssn)
	fmt.Println("enter customer deposit amount :")
	fmt.Scan(&newcus.amount)
	newaccnoadd := ba.numofcustomers()
	newcus.caccno = newaccnoadd + 1
	fmt.Println("account number for the customer is :", newcus.caccno)
	ba.cus = append(ba.cus, *newcus)
	fmt.Println(len(ba.cus))

}

func (ba *bank) deletecustomer(delaccnum int) {

	for index, delcus := range ba.cus {
		if delcus.caccno == delaccnum {
			ba.cus = append(ba.cus[:index], ba.cus[index:]...)
		}
	}
}

type customer struct {
	cname  string
	caccno int
	cssn   int
	amount int
}

func (cus *customer) withdraw(withdrawamount int) {
	if cus.amount >= withdrawamount {
		cus.amount = cus.amount - withdrawamount
		fmt.Println("remaing amount", cus.amount)
	} else {
		fmt.Println("not enough balance to withdraw")
	}
	fmt.Println("in withdraw :", cus)
}
func (cus *customer) depositmoney(dmoney int) {
	cus.amount += dmoney
	fmt.Println("total balance in the account", cus.amount)
}
func custfind(tbank *bank, accnum int) *customer {
	var selectedcustomer customer

	for index, customerdata := range tbank.cus {
		if customerdata.caccno == accnum {
			return &tbank.cus[index]
		}
	}
	return &selectedcustomer
}

func main() {

	abank := &bank{}
	customer1 := &customer{cname: "customer1", caccno: 0001, cssn: 0011, amount: 1000}
	customer2 := &customer{cname: "customer3", caccno: 0002, cssn: 0012, amount: 1001}
	customer3 := &customer{cname: "customer3", caccno: 0003, cssn: 0013, amount: 1002}
	abank.cus = append(abank.cus, *customer1, *customer2, *customer3)
	var (
		cuscheck, serviceType, accnum, amount, exitcheck int
	)
	fmt.Print("if new customer press 1 else enter 0")
	fmt.Scan(&cuscheck)
	for {
		if cuscheck == 1 {
			fmt.Println("HEY")
			abank.addcustomer()
		} else if cuscheck == 0 {
			fmt.Println("Enter the service number you require\n1.WithdrawMoney\n2. DepositMoney\n3.NumberofCustomers\n4.printAllCustomers\n5. PrintBalance\n6. DeleteCustomer")
			fmt.Scan(&serviceType)
			switch serviceType {
			case 1:
				fmt.Println("enter account number")
				fmt.Scan(&accnum)
				fmt.Println("Enter amount to withdraw")
				fmt.Scan(&amount)
				(*custfind(abank, accnum)).withdraw(amount)
				abank.printdata()
				amount = 0
			case 2:
				fmt.Println("enter account number")
				fmt.Scan(&accnum)
				fmt.Println("Enter amount to deposit")
				fmt.Scan(&amount)
				(*custfind(abank, accnum)).depositmoney(amount)

				amount = 0
			case 3:
				fmt.Println("Number of customers in bank are", abank.numofcustomers())
			case 4:
				abank.printdata()
			case 5:
				fmt.Println("enter account number")
				fmt.Scan(&accnum)
				fmt.Println("balance in the account is", (*custfind(abank, accnum)).amount)
			case 6:
				fmt.Println("enter account number")
				fmt.Scan(&accnum)
				abank.deletecustomer(accnum)
				fmt.Println("Account deleted successfully")
			default:
				fmt.Println("choose among given")
			}
		}
		fmt.Println("if you want to exit , enter 1\n to add another customer  enter 0\n to view  customer service  enter 3")
		fmt.Scan(&exitcheck)
		if exitcheck == 0 {
			cuscheck = 1
		} else if exitcheck == 3 {
			cuscheck = 0

		}
		if exitcheck == 1 {
			return
		}
	}
}
