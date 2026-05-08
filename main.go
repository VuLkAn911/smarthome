package main

import "fmt"

func main() {
	const BaseTariff = 0.45
	const HightLoadTax = 0.15
	const NightDiscount = 0.30

	fmt.Println("Выберите девайс:")
	for {
		fmt.Print("Введите название прибора (или 'done'): ")
		var name string
		fmt.Scan(&name)
		if name == "done" {
			break
		}
		fmt.Print("Введите мощность ватт: ")
		var Power float64
		fmt.Scan(&Power)

		fmt.Print("Время работы девайса: ")
		var HoursOfWorking float64
		fmt.Scan(&HoursOfWorking)

		fmt.Print("Ночной режим:")
		var NightMode bool
		fmt.Scan(&NightMode)

		Power = Power * HoursOfWorking / 1000
		Total := Power * BaseTariff

		if NightMode == true {
			Total -= NightDiscount
		}
		if Power > 10 {
			Power += HightLoadTax
		}

		switch {
		case Power < 100:
			fmt.Println("Экономный")
		case Power > 100 && Power < 1000:
			fmt.Println("Стандартный")
		case Power > 1000:
			fmt.Println("Мощный")

		}
	}

	fmt.Println("Расчет завершен")
}
