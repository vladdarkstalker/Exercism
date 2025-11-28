package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) (res bool) {
	if kind == "car" || kind == "truck" {
        res = true 
    } else if kind == "bike" {
        return
    }
    return
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    res := ""
    if option1 < option2 {
        res = option1
    } else {
        res = option2
    }
    return res + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) (res float64) {
	if age < 3.0 {
        res = originalPrice * 0.8
    } else if age >= 10.0 {
        res = originalPrice * 0.5
    } else if age >= 3.0 && age <= 10.0 {
        res = originalPrice * 0.7
    }
    return
}
