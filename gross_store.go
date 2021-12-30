package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	//panic("Please implement the Units() function")
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	//panic("Please implement the NewBill() function")
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	//panic("Please implement the AddItem() function")
	if u, ok := units[unit]; ok {
		if _, ok = bill[item]; ok {
			bill[item] += u
		} else {
			bill[item] = u
		}
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	//panic("Please implement the RemoveItem() function")
	if u, ok := units[unit]; ok {
		var i int
		if i, ok = GetItem(bill, item); ok {
			t := i - u
			switch {
			case t < 0:
				return false
			case t == 0:
				delete(bill, item)
				return true
			default:
				bill[item] -= u
				return true
			}
		}
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	//panic("Please implement the GetItem() function")
	i, ok := bill[item]
	return i, ok
}
