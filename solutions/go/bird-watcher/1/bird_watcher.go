package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) (res int) {
    for _, el := range birdsPerDay {
        res += el
    }
    return
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) (sum int) {
    el := week * 7
    for _, i := range birdsPerDay[el-7:el] {
        sum += i
    }
	return 
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i+=2 {
        birdsPerDay[i] = birdsPerDay[i] + 1
    }
    return birdsPerDay
}
