package chance
import "math/rand"
import "time"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	rand.Seed(time.Now().UnixNano())
    return  1 + rand.Intn(20)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	source := rand.NewSource(time.Now().UnixNano())
    generator := rand.New(source)
    return generator.Float64() * 12.0
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

    rand.Shuffle(len(animals), func(i, j int) {
        animals[i], animals[j] = animals[j], animals[i]
    }) // Shuffle contents
    return animals
}
