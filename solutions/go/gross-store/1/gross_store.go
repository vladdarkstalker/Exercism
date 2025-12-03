package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    res := map[string]int{"quarter_of_a_dozen": 3, "half_of_a_dozen": 6, "dozen": 12, "small_gross": 120, "gross": 144, "great_gross": 1728}
	return res
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    value_u, is_u := units[unit]
    if !is_u {
        return false
    }
    
    _, is_b := bill[item]
    if is_b {
        bill[item] += value_u
    } else {
        bill[item] = value_u
    }
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value_b, is_b := bill[item]
    if !is_b {
        return false
    }
    value_u, is_u := units[unit]
    if !is_u {
        return false
    }

    if value_b - value_u < 0 {
        return false
    } 
    bill[item] = value_b - value_u
	if bill[item] == 0 {
        delete(bill, item)
    }
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value_b, is_b := bill[item]
    if !is_b {
        return 0, false
    } else {
        return value_b, true
    }
}
