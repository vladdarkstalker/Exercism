package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
    return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) (res string) {
	res = fmt.Sprintf("Welcome to my party, %s!\n", name) +
    		func (res int) (rest string) {
                if res < 10 {
                    rest = fmt.Sprintf("You have been assigned to table 00%d.", res)
                } else if res < 100 {
                    rest = fmt.Sprintf("You have been assigned to table 0%d.", res)
                } else {
                    rest = fmt.Sprintf("You have been assigned to table %d.", res)
                }
                return
            } (table) + 
    		fmt.Sprintf(" Your table is %s, exactly %0.1f meters from here.\n", direction, distance) +
    		fmt.Sprintf("You will be sitting next to %s.", neighbor)
    return
}
