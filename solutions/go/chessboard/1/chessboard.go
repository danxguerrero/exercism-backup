package chessboard
//import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    count := 0
	for k, v := range cb {
        if k == file {
            for x := range v {
				if v[x] == true {
                    count += 1
                }
            }
        }
    }
    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank > 8 || rank < 0 {
        return 0
    }
    
	count := 0
	for _, v := range cb {
        for x := range v {
            if x == rank - 1 && v[x] == true {
                count += 1
            }
        }
    }
    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    count := 0
	for _, v := range cb {
        count += len(v)
    }
    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
    for _, v := range cb {
        for x := range v {
            if v[x] == true {
                count += 1
            }
        }
    }
    return count
}
