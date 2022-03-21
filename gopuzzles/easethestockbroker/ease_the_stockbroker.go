package easethestockbroker

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func BalanceStatement(lst string) string {
	if len(lst) == 0 {
		return fmt.Sprintf("Buy: %d Sell: %d", 0, 0)
	}

	orders := strings.Split(lst, ",")
	badlyFormedOrders := []string{}
	boughtStock := 0.0
	soldStock := 0.0

	for _, simpleOrder := range orders {
		order := strings.Trim(simpleOrder, " ")
		orderParts := strings.Fields(order)

		if len(orderParts) != 4 {
			badlyFormedOrders = append(badlyFormedOrders, order)
			continue
		}

		qty, price, status := orderParts[1], orderParts[2], orderParts[3]

		q, err := strconv.Atoi(qty)
		if err != nil {
			badlyFormedOrders = append(badlyFormedOrders, order)
			continue
		}

		if !strings.Contains(price, ".") {
			badlyFormedOrders = append(badlyFormedOrders, order)
			continue
		}

		if status != "B" && status != "S" {
			badlyFormedOrders = append(badlyFormedOrders, order)
			continue
		}

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			badlyFormedOrders = append(badlyFormedOrders, order)
			continue
		}

		if status == "B" {
			boughtStock += float64(q) * p
		} else {
			soldStock += float64(q) * p
		}
	}

	result := fmt.Sprintf("Buy: %.0f Sell: %.0f", boughtStock, soldStock)

	if len(badlyFormedOrders) != 0 {
		result += fmt.Sprintf("; Badly formed %d: %s ;", len(badlyFormedOrders), strings.Join(badlyFormedOrders, " ;"))
	}

	return result
}

func BalanceStatementV2(lst string) string {
	if len(lst) == 0 {
		return "Buy: 0 Sell: 0"
	}
	var badform = []string{}
	prices := map[string]float64{"B": 0.0, "S": 0.0}
	arr := strings.Split(lst, ", ")
	var valid = regexp.MustCompile(`\S+ \d+ \d*\.\d+ (B|S)`)
	for _, v := range arr {
		if !valid.MatchString(v) {
			badform = append(badform, v+" ;")
		} else {
			uu := strings.Split(v, " ")
			if len(uu) != 4 {
				badform = append(badform, v+" ;")
			} else {
				_, quantity, price, status := uu[0], uu[1], uu[2], uu[3]
				q, _ := strconv.Atoi(quantity)
				p, _ := strconv.ParseFloat(price, 64)
				newprice := prices[status] + float64(q)*p
				prices[status] = newprice
			}
		}
	}
	res := fmt.Sprintf("Buy: %.0f Sell: %.0f", prices["B"], prices["S"])
	if len(badform) != 0 {
		res += fmt.Sprintf("; Badly formed %d: %s", len(badform), strings.Join(badform, ""))
	}
	return res
}
