package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) (res int) {
	switch card {
        case "ace":
        	res = 11
        case "two":
        	res = 2
        case "three":
        	res = 3
        case "four":
        	res = 4
        case "five":
        	res = 5
        case "six":
        	res = 6
        case "seven":
        	res = 7
        case "eight":
        	res = 8
        case "nine":
        	res = 9
        case "ten", "jack", "queen", "king":
        	res = 10
        default:
        	res = 0
    }
    return
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) (res string) {
    sumCard := ParseCard(card1) + ParseCard(card2)
    switch {
        case sumCard == 21 && ParseCard(dealerCard) != 10 && ParseCard(dealerCard) != 11:
        	res = "W"
        case sumCard == 21 && (ParseCard(dealerCard) == 10 || ParseCard(dealerCard) == 11):
        	res = "S"
        case sumCard == 22:
        	res = "P"
        case sumCard >= 17 && sumCard <= 20:
        	res = "S"
        case sumCard >= 12 && sumCard <= 16 && ParseCard(dealerCard) < 7:
        	res = "S"
        case sumCard >= 12 && sumCard <= 16 && ParseCard(dealerCard) >= 7:
        	res = "H"
        case sumCard <= 11:
        	res = "H"
    }
    return
}
