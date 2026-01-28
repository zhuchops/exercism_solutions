package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	measurements := map[string]int {
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen": 12,
		"small_gross": 120,
		"gross": 144,
		"great_gross": 1728,
	}
	return measurements
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	int_unit, exists := units[unit]
	if !exists {
		return false
	}
	_, exists = bill[item]
	if exists {
		bill[item] += int_unit
	} else {
		bill[item] = int_unit
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unit_count, exists := units[unit]
	if !exists {
		return false
	}
	current_count, exists := bill[item]
	if !exists {
		return false
	}
	if current_count - unit_count < 0 {
		return false
	}
	if current_count - unit_count == 0 {
		delete(bill, item)
		return true
	}
	bill[item] -= unit_count
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	current_count, exists := bill[item]
	if !exists {
		return 0, false
	}
	return current_count, true
}
