package utils

import "fmt"

type FamilyAccount struct {
	//总金额
	Balance float64
	//每次的金额
	Money   float64
	//说明
	Note    string

	IsFirst bool
	Key     string
	ExitTag bool
	Details string
}

func (f *FamilyAccount) showDetails()  {
	fmt.Println("----------当前收支明细记录----------")
	if !f.IsFirst{
		fmt.Println(f.Details)
	}else {
		fmt.Println("当前没有收支明细, 请来一笔")
	}
}

func (f *FamilyAccount) inCome()  {
	if f.IsFirst{
		f.IsFirst = false
	}
	fmt.Println("本次收入金额:")
	fmt.Scanln(&f.Money)
	f.Balance += f.Money
	fmt.Println("本次收入说明")
	fmt.Scanln(&f.Note)

	f.Details += fmt.Sprintf("\n收入\t%v\t%v\t%v",
		f.Balance, f.Money, f.Note)
}

func (f *FamilyAccount) outCome()  {
	if f.IsFirst{
		f.IsFirst = false
	}
	fmt.Println("本次支出金额")
	if fmt.Scanln(&f.Money); f.Money > f.Balance {
		fmt.Println("no money left enough")
		return
	} else {
		f.Balance -= f.Money
	}
	fmt.Println("本次支出说明")
	fmt.Scanln(&f.Note)

	f.Details += fmt.Sprintf("\n支出\t%v\t%v\t%v",
		f.Balance, f.Money, f.Note)
}

func (f *FamilyAccount) showChoice()  {

	fmt.Println("记账")
	fmt.Println("1: 明细")
	fmt.Println("2: 收入")
	fmt.Println("3: 支出")
	fmt.Println("4: 退出")
}

func (f *FamilyAccount) MainMenu() {

	for {
		//fmt.Println(details)
		f.showChoice()

		fmt.Scanln(&f.Key)
		switch f.Key {
		case "1":
			f.showDetails()
		case "2":
			f.inCome()
		case "3":
			f.outCome()
		case "4":
			f.ExitTag = true
		default:
			fmt.Println()
		}
		if f.ExitTag {
			break
		}
	}
}