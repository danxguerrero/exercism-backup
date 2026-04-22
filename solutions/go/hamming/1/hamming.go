package hamming
import "errors"

func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("Strands are not the same length")
    }
	var hamming_dist int

    for i , _ := range a {
        if a[i] != b[i] {
            hamming_dist++
        }
    }

    return hamming_dist, nil
}
