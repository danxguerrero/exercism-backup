package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
    if prepTimePerLayer == 0 {
        prepTimePerLayer = 2
    }

    return len(layers) * prepTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    noodleCount := 0
    sauceCount := 0

	for _, layer := range layers {
        if layer == "noodles" {
            noodleCount++
        } else if layer == "sauce" {
            sauceCount++
        }
    }

    noodles = noodleCount * 50
    sauce = float64(sauceCount) * 0.2

    return noodles, sauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
    secretIngredient := friendsList[len(friendsList) - 1]
    myList[len(myList) -1] = secretIngredient
    return myList
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int)[]float64{
    scaledQuantities := make([]float64, len(quantities))
    for i, quantity := range quantities {
        singlePortion := quantity / 2.0
        scaledQuantities[i] = singlePortion * float64(portions)
    }
    return scaledQuantities
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
