package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) (res int) {
	for _, i := range cb[file] {
        if i {
            res++
        }
    }
    return
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (res int) {
	if rank < 1 || rank > 8 {
        return
    }
    
	for _, i := range cb {
        if i[rank-1] {
            res++
        }
    }
    return
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (res int) {
	for _,i := range cb {
        res += len(i)
    }
    return
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (res int) {
	for _,i := range cb {
        for _,j := range i {
        	if j {
                res ++
            }
    	}
    }
    return
}
