package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(lays []string, time int) (res int) {
    if time == 0 {
        res = len(lays) * 2
    } else {
        res = len(lays) * time
    }
    return
}

// TODO: define the 'Quantities()' function
func Quantities(lays []string) (noodles int, sauce float64) {
    for _, el := range lays {
        if el == "noodles" {
            noodles += 50
        } 
        if el == "sauce" {
            sauce += 0.2
        } 
    }
    return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(fl, ml []string) {
    ml[len(ml) - 1] =fl[len(fl) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(slice []float64, portion int) []float64 {
    scale := float64(portion) / 2
    res := make([]float64, len(slice))
    for i, el := range slice {
        res[i] = el * scale
    }
    return res
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
