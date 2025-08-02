package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

var numberplayers = 6
var playername = []string{"Simon", "Lorenzo", "Jeff", "John", "Terry", "Eric", "Graham"}
var cardnames = []string{"Null", "One", "Two", "Three", "Four", "Five", "Six", "Llama"}
var playerstate = []bool{true, true, true, true, true, true}
var playerhand [6][56]int
var playerwhitecounters = []int{0, 0, 0, 0, 0, 0}
var playerblackcounters = []int{0, 0, 0, 0, 0, 0}
var playerscore = []int{0, 0, 0, 0, 0, 0}
var topofdiscardpile int
var deck Deck
var tempcard int
var move string
var valid bool
var validmoves string
var roundactive = true
var gameactive = true

// card represents a playing card with it value
type card struct {
	cardvalue int
}

// Deck represents a collection of cards.
type Deck []card

type players struct {
	name  string
	score int
}

func main() {

	fmt.Println("Welcome to Fuji-Llama!")
	// get the number of AI players with validate
	for !valid {
		fmt.Print("How many AI Bot's ? (1-5)")
		_, err := fmt.Scan(&numberplayers)
		if err != nil {
			fmt.Println("enter a number only, please try again")
		} else {
			valid = true
		}
		if numberplayers >= 1 && numberplayers <= 5 {
			valid = true
		} else {
			valid = false
			fmt.Println("1 to 5 only, please try again")
		}
	}
	numberplayers++ // add in the human player

	for gameactive {
		deck = NewDeck()
		fmt.Println("Now Shuffling the cards:")
		deck.Shuffle()

		fmt.Println("Now Dealing the cards:")
		// clear out players hands first
		for i := 0; i < 56; i++ {
			for ii := 0; ii < numberplayers; ii++ {
				playerhand[ii][i] = 0
			}
		}
		// deal the new cards
		for i := 0; i < 6; i++ {
			for ii := 0; ii < numberplayers; ii++ {
				if card, ok := deck.DealCard(); ok {
					playerhand[ii][i] = card.cardvalue
					playerstate[ii] = true
				}
			}
		}
		// turn over top card
		if card, ok := deck.DealCard(); ok {
			topofdiscardpile = card.cardvalue
		}
		/*
			// cheat to test ending round on 0
			for i := 0; i < 6; i++ {

				playerhand[0][i] = 0

			}
			playerhand[0][0] = topofdiscardpile
			playerwhitecounters[0] = 0
			playerblackcounters[0] = 0
			playerscore[0] = 0
			//*/
		// playing a round
		for roundactive {

			Display() // display the status
			notactivecount := 0

			for i := 0; i < numberplayers; i++ {
				if playerstate[i] {
					validmoves = CheckVaildMoves(i)
					DoVaildMoves(i, validmoves)
				} else {
					notactivecount++
				}
				if CountPlayersCards(i) == 0 {
					fmt.Println(playername[i], "is all out, and has won the round")
					roundactive = false
					i = numberplayers
				}

			}
			if notactivecount == numberplayers {
				roundactive = false
			}
		}

		fmt.Println("round over")
		EndofRoundScore()
		fmt.Println("Press return to play the next round")
		waitkey := ""
		fmt.Scanln(&waitkey)
		roundactive = true
		gameactive = CheckGameEnd()
	}
	DisplayGameEnd()
}

// NewDeck creates a new Llama 56-card deck.
func NewDeck() Deck {
	cardvalues := []int{1, 2, 4, 4, 5, 6, 7}
	deck := make(Deck, 0, 55)
	for i := 0; i < 8; i++ {
		for _, cardvalue := range cardvalues {
			deck = append(deck, card{cardvalue: cardvalue})
		}
	}
	return deck
}

// Shuffle shuffles the deck using the Fisher-Yates algorithm.
func (d *Deck) Shuffle() {
	for i := len(*d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}

// DealCard deals a single card from the top of the deck.
// Returns the card and a boolean indicating if a card was available.
func (d *Deck) DealCard() (card, bool) {
	if len(*d) == 0 {
		return card{}, false
	}
	card := (*d)[0]
	*d = (*d)[1:]
	return card, true
}

// count the deck and return the number of cards remaining
func Deckcount() int {
	remainingcards := 0
	for _, cards := range deck {
		if cards.cardvalue != 0 {
			remainingcards++
		}
	}
	return remainingcards
}

// display the current status
func Display() {
	fmt.Println("Current round status:")
	remainingcards := Deckcount()

	fmt.Println("The top card showing is:", cardnames[topofdiscardpile], "and there are", remainingcards, "cards remaining")

	for i := 1; i < numberplayers; i++ {

		playerremainingcards := 0
		for _, cards := range playerhand[i] {
			if cards != 0 {
				playerremainingcards++
			}

		}

		fmt.Println(playername[i], "has", CountPlayersCards(i), "Cards and is",
			playerstate[i], ":", playerwhitecounters[i], "white counters and",
			playerblackcounters[i], "black counters and a score of",
			playerscore[i])

	}

	fmt.Println(playername[0], "cards are", PlayerHandDisplay(0), "and you are",
		playerstate[0], ":", playerwhitecounters[0], "white counters and",
		playerblackcounters[0], "black counters and a score of",
		playerscore[0])
}

// Check for vaild moves
func CheckVaildMoves(player int) string {
	validmove := ""
	currentcards := 0
	nextcards := 0
	tempcard = 8 // when the Llama is topcard 1 is also a vaild card, so this makes the number roll etc
	remainingcards := Deckcount()
	for _, s := range playerhand[player] {
		if topofdiscardpile == 7 {
			tempcard = 1
		}
		switch s {
		case topofdiscardpile:
			currentcards++
		case topofdiscardpile + 1:
			nextcards++
		case tempcard:
			nextcards++
		}
	}

	switch {
	case currentcards > 0 && nextcards > 0:
		validmove = "CNcn"
	case currentcards > 0:
		validmove = "Cc"
	case nextcards > 0:
		validmove = "Nn"
	case AmIalone():
		validmove = "Ff"
	case remainingcards > 0:
		validmove = "DFdf"
	default:
		validmove = "Ff"
	}

	return validmove
}

// do move for vaild moves
func DoVaildMoves(player int, validmove string) {

	validcheck := false
	if player == 0 {
		for !validcheck {
			fmt.Println("Your play options are:")
			switch {
			case validmove == "CNcn":
				fmt.Println("Play current card (C) or Play next card (N)")
			case validmove == "Cc":
				fmt.Println("Play current card (C)")
			case validmove == "Nn":
				fmt.Println("Play next card (N)")
			case validmove == "DFdf":
				fmt.Println("Draw card (D) or Fold (F)")
			default:
				fmt.Println("Fold (F)")
			}
			fmt.Scan(&move)
			if strings.Contains(validmove, move) {
				validcheck = true
			} else {
				fmt.Println("not a valid move, please try again")
				validcheck = false
			}
		}
	} else {
		// do AI stuff just do first vaild option for now
		move = validmove[0:1]
	}

	switch move {
	case "c", "C":
		fmt.Println(playername[player], "is playing a", cardnames[topofdiscardpile], "on to the discard pile")
		RemoveCard(player, topofdiscardpile)
	case "n", "N":
		// need to set card to 1 if current card is Llama (Im sure there is a better way to do this )
		if topofdiscardpile == 7 {
			topofdiscardpile = 1
		} else {
			topofdiscardpile++
		}
		fmt.Println(playername[player], "is playing a", cardnames[topofdiscardpile], "on to the discard pile")
		RemoveCard(player, topofdiscardpile)

	case "d", "D":
		fmt.Println(playername[player], "is drawing a new card from the deck")
		if card, ok := deck.DealCard(); ok {
			for i := 0; i < 56; i++ {
				if playerhand[player][i] == 0 {
					playerhand[player][i] = card.cardvalue
					i = 56
				}
			}
		}

	case "f", "F":
		fmt.Println(playername[player], "is folding")
		playerstate[player] = false
	}
}

// remove played card from players hand
func RemoveCard(player, cardvalue int) {
	for i := 0; i < 56; i++ {
		if playerhand[player][i] == cardvalue {
			playerhand[player][i] = 0
			i = 56
		}
	}
}

// count the cards in a players hand
func CountPlayersCards(player int) int {
	playerremainingcards := 0
	for _, cards := range playerhand[player] {
		if cards != 0 {
			playerremainingcards++
		}
	}
	return playerremainingcards
}

// Get the cards in a players hand and return as a string
func PlayerHandDisplay(player int) string {
	displayhand := "[ "
	for i := 0; i < 56; i++ {
		if playerhand[player][i] != 0 {

			displayhand = displayhand + cardnames[playerhand[player][i]] + " "
		}
	}
	displayhand = displayhand + "]"
	return displayhand
}

// End of round scoreing
func EndofRoundScore() {
	for i := 0; i < numberplayers; i++ {
		a := 0
		b := 0
		c := 0
		for ii := 0; ii < 56; ii++ {
			switch {
			case playerhand[i][ii] == 0:
				// do nothing
			case playerhand[i][ii] == 7 && b != 1:
				b = 1
			case playerhand[i][ii] != 7:
				a = a + playerhand[i][ii]
			}

		}
		c = a / 10
		b = b + c
		a = a - (c * 10)
		if (a + (b * 10)) == 0 {
			switch {
			case playerblackcounters[i] > 0:
				fmt.Println(playername[i], "finshed with no cards,so scores zero points and is returning one black token")
			case playerwhitecounters[i] > 0:
				fmt.Println(playername[i], "finshed with no cards,so scores zero points and is returning one white token")
			default:
				fmt.Println(playername[i], "finshed with no cards,so scores zero points")
			}
		} else {
			fmt.Println(playername[i], "has", PlayerHandDisplay(i), "gaining", a, "white counters and", b, "black counters, scoreing", a+(b*10), "points")
			playerwhitecounters[i] = playerwhitecounters[i] + a
			playerblackcounters[i] = playerblackcounters[i] + b
			playerscore[i] = playerscore[i] + a + (b * 10)
		}
	}
}

// Check for game end
func CheckGameEnd() bool {
	var scorenotmax = true
	for i := 0; i < numberplayers; i++ {
		if playerscore[i] >= 40 {
			scorenotmax = false
		}
	}
	return scorenotmax
}

// show game over summary and the winner etc
func DisplayGameEnd() {
	fmt.Println("The Game is Over")
	players := [6]players{}
	for i := 0; i < numberplayers; i++ {
		fmt.Println(playername[i], "has", playerwhitecounters[i], "white counters and",
			playerblackcounters[i], "black counters and a total score of",
			playerscore[i])
		players[i].name = playername[i]
		players[i].score = playerscore[i]
	}
	fmt.Println("not sorted Sorted by score:", players)
	s := players[:]
	sort.Slice(s, func(i, j int) bool { return s[i].score < s[j].score })
	//sort.Slice(players, func(i, j int) bool { return players[i].score < players[j].score })
	fmt.Println("Sorted by score:", players)
	fmt.Println("The winner is", players[0])
}

// check count of players active
func AmIalone() bool {
	notactivecount := 0

	for i := 0; i < numberplayers; i++ {
		if !playerstate[i] {
			notactivecount++
		}
	}
	return notactivecount >= (numberplayers - 1)
}
