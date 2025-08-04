package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	//"strconv"
	//"strings"
)

var numberplayers = 6
var player [6]players
var topofdiscardpile card
var decksize int          // Tracks the size of the current deck
var activePlayerCount int // Tracks the number of inactive players
var move string
var valid bool
var roundactive = true
var gameactive = true
var currentdeck deck

// card represents a playing card with it's name and value
type card struct {
	cardvalue int
	cardname  string
}

// Deck represents a collection of cards.
type deck []card

type players struct {
	name          string
	human         bool
	playing       bool
	whitecounters int
	blackcounters int
	score         int
	hand          deck
	playorder     int
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

	// Create the player array
	for i := 0; i < numberplayers; i++ {
		player[i] = NewPlayer(i)
	}

	player[0].human = true // set the first player as human

	for gameactive {

		currentdeck = NewDeck()
		decksize = len(currentdeck)

		fmt.Println("Now Shuffling the cards:")
		currentdeck.Shuffle()
		fmt.Println("Now Dealing the cards:")
		// clear out players hands first and reset their state
		for i := 0; i < numberplayers; i++ {
			player[i].hand = []card{}
			player[i].playing = true
		}
		activePlayerCount = numberplayers // Reset active player count at the start of a round
		// deal the 6 cards to each player
		for j := 0; j < 6; j++ {
			for i := 0; i < numberplayers; i++ {
				player[i].hand = append(player[i].hand, currentdeck.DealCard())
			}
		}

		// turn over top card onto the discard pile
		topofdiscardpile = currentdeck.DealCard()

		// playing a round the main loop
		for roundactive {

			for i := 0; i < numberplayers; i++ {
				if player[i].playing {
					Display() // display the round status
					DoVaildMoves(i, CheckVaildMoves(i))

					hand := player[i].hand // Get the current player's hand
					if len(hand) == 0 {    // If the player has no cards left, they have won the round, end the round
						fmt.Println(player[i].name, "is all out, and has won the round")
						roundactive = false
						break
					}
					if activePlayerCount == 0 { // If no active players left, end the round
						roundactive = false
						break
					}
				}
			}
		}
		fmt.Println("----------------------------------------------------")
		fmt.Println("Round Over")
		EndofRoundScore()
		fmt.Println("----------------------------------------------------")

		if CheckGameEnd() {
			fmt.Println("Press return to play the next round")
			waitkey := ""
			fmt.Scanln(&waitkey)
			roundactive = true
		}
		gameactive = CheckGameEnd()
	}
	DisplayGameEnd()
	fmt.Println("Thanks for playing Fuji-Llama!")
}

// create a new player with a name and default values
func NewPlayer(index int) players {
	playernames := []string{"Simon", "Lorenzo", "Jeff", "John", "Terry", "Eric", "Graham"}
	newPlayer := players{
		name:          playernames[index],
		human:         false,
		playing:       false,
		whitecounters: 0,
		blackcounters: 0,
		score:         0,
		playorder:     index,
	}
	return newPlayer
}

// NewDeck creates a new Llama 56-card deck.

func NewDeck() deck {
	cardNames := []string{"One", "Two", "Three", "Four", "Five", "Six", "Llama"}
	newDeck := make(deck, 0, 56) // Pre-allocate slice with capacity for 56 cards

	for i := 0; i < 8; i++ { // Repeat the cards 8 times to create a full deck
		for value, name := range cardNames {
			newDeck = append(newDeck, card{cardvalue: value + 1, cardname: name})
		}
	}
	return newDeck
}

// Shuffle shuffles the deck using the Fisher-Yates algorithm.
func (d *deck) Shuffle() {
	for i := len(*d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}

// DealCard deals a single card from the top of the deck.
func (d *deck) DealCard() card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

// display the current status
func Display() {
	fmt.Println("Current round status:")
	fmt.Println("----------------------------------------------------")

	for i := 0; i < numberplayers; i++ {
		var statetext string
		if player[i].playing {
			statetext = "playing"
		} else {
			statetext = "folded"
		}

		fmt.Println(player[i].name, "has", len(player[i].hand), "Cards\t:", statetext,
			":", player[i].whitecounters, "white counters and",
			player[i].blackcounters, "black counters and a score of",
			player[i].score)
		// fmt.Println(DisplayHand(i)) // Display the player's hand for debugging

	}
	fmt.Println("----------------------------------------------------")

}

// DisplayHand displays the cards in a player's hand.
// It returns a string representation of the hand.
// If the player index is invalid, it returns an error message.
// If the player has no cards in hand, it returns a message indicating that.
func DisplayHand(playerIndex int) string {
	if playerIndex < 0 || playerIndex >= numberplayers {
		return "Invalid player index"
	}

	hand := player[playerIndex].hand
	if len(hand) == 0 {
		return "No cards in hand"
	}

	handDisplay := "[ "
	for _, card := range hand {
		handDisplay += card.cardname + " "
	}
	handDisplay += "]"
	return handDisplay
}

// CheckVaildMoves checks the player's hand for valid moves
// The valid moves are determined based on them holding a card that matches the top card of the discard pile or next card in the sequence.
// then checks they can draw a card and if not they must or fold (they can fold if they don't want to fold).
// It returns a string indicating the valid moves available.
func CheckVaildMoves(index int) string {
	validmove := ""
	nextcard := 0
	currentcardcount := 0
	nextcardcount := 0
	decksize = len(currentdeck)
	if topofdiscardpile.cardvalue == 7 {
		nextcard = 1
	} else {
		nextcard = topofdiscardpile.cardvalue + 1
	}

	for _, s := range player[index].hand { // check each card in the players hand for valid cards
		if s.cardvalue == topofdiscardpile.cardvalue {
			currentcardcount++ // Count how many of the current card the player has
		}
		if s.cardvalue == nextcard {
			nextcardcount++ // Count how many of the current card the player has
		}
	}
	switch {
	case currentcardcount > 0 && nextcardcount > 0 && activePlayerCount > 1 && decksize > 0: // If they have both current and next card
		validmove = "CNcnDFdf" // Can play current or next card, draw a card or fold
	case currentcardcount > 0 && activePlayerCount > 1 && decksize > 0: // If they have just the current card
		validmove = "CcDFdf" // Can play current card, draw a card or fold
	case nextcardcount > 0 && activePlayerCount > 1 && decksize > 0: // If they have just the next card, draw a card or fold
		validmove = "NnDFdf"
	case currentcardcount > 0 && nextcardcount > 0 && (activePlayerCount == 1 || decksize < 1): // If they have both current and next card but your the only player left or deck is depleted
		validmove = "CNcnFf" // Can play current or next card or fold
	case currentcardcount > 0 && (activePlayerCount == 1 || decksize < 1): // If they have just the current but your the only player left or deck is depleted
		validmove = "CcFf" // Can play current card  or fold
	case nextcardcount > 0 && (activePlayerCount == 1 || decksize < 1): // If they have just the next card but your the only player left or deck is depleted
		validmove = "NnFf" // Can play next card or fold
	case decksize > 0: // if there are cards left in the deck you can draw or fold
		validmove = "DFdf"
	default: // fold is then only option left
		validmove = "Ff"
	}
	return validmove
}

// DoVaildMoves asks the play what moves they want to make and checks they are valid, also current does simple AI player moves.
func DoVaildMoves(index int, validmove string) {

	validcheck := false
	if player[index].human {
		fmt.Println("The top card showing is:", topofdiscardpile.cardname, "and there are", len(currentdeck), "cards remaining")
		fmt.Println("your hand is:", DisplayHand(index)) // Display the player's hand
		for !validcheck {
			fmt.Println(player[index].name, "your play options are:")
			switch {
			case validmove == "CNcnDFdf":
				fmt.Println("Play current card (C), Play next card (N), Draw card (D) or Fold (F)")
			case validmove == "CcDFdf":
				fmt.Println("Play current card (C), Draw card (D) or Fold (F)")
			case validmove == "NnDFdf":
				fmt.Println("Play next card (N), Draw card (D) or Fold (F)")
			case validmove == "CNcnFf":
				fmt.Println("Play current card (C), Play next card (N), or Fold (F)")
			case validmove == "CcFf":
				fmt.Println("Play current card (C) or Fold (F)")
			case validmove == "NnFf":
				fmt.Println("Play Next card (C) or Fold (F)")
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
		fmt.Println(player[index].name, "played a", topofdiscardpile.cardname, "on to the discard pile")
		RemoveCard(index, topofdiscardpile.cardvalue)

	case "n", "N":
		// need to set card to 1 if current card is Llama (Im sure there is a better way to do this )
		if topofdiscardpile.cardvalue == 7 {
			topofdiscardpile.cardvalue = 1
			topofdiscardpile.cardname = "One"
		} else {
			topofdiscardpile.cardvalue = topofdiscardpile.cardvalue + 1
			cardNames := []string{"One", "Two", "Three", "Four", "Five", "Six", "Llama"}
			topofdiscardpile.cardname = cardNames[topofdiscardpile.cardvalue-1]
		}
		fmt.Println(player[index].name, "played a", topofdiscardpile.cardname, "on to the discard pile")
		RemoveCard(index, topofdiscardpile.cardvalue)

	case "d", "D":
		fmt.Println(player[index].name, "drew a new card from the deck")
		player[index].hand = append(player[index].hand, currentdeck.DealCard())
	case "f", "F":
		fmt.Println(player[index].name, "folded")
		player[index].playing = false
		activePlayerCount-- // decrement active player count
	}
}

// Remove the played card from players hand
func RemoveCard(index, cardvalue int) {

	for i, s := range player[index].hand { // find the card in the player's hand and then remove it
		if s.cardvalue == cardvalue {
			// Remove the card by slicing out the matched card
			player[index].hand = append(player[index].hand[:i], player[index].hand[i+1:]...)
			break // Exit the loop after removing the first matching card
		}
	}
}

// End of round scoreing
func EndofRoundScore() {
	fmt.Println("------------- End of round summary ------------------")
	for i := 0; i < numberplayers; i++ {
		a := 0
		b := 0
		c := 0
		for ii := 0; ii < len(player[i].hand); ii++ {
			switch {
			case player[i].hand[ii].cardvalue == 0:
				// do nothing
			case player[i].hand[ii].cardvalue == 7:
				b = 1 // Llama is worth 1 black counter but only get 1 black counter no matter how many you have
			default:
				a = a + player[i].hand[ii].cardvalue
			}
		}

		c = a / 10
		b = b + c
		a = a - (c * 10)
		if (a + (b * 10)) == 0 {
			switch {
			case player[i].blackcounters > 0:
				fmt.Println(player[i].name, "finshed with no cards, so scores zero points and is returning one black token")
				player[1].blackcounters = player[i].blackcounters - 1
			case player[i].whitecounters > 0:
				fmt.Println(player[i].name, "finshed with no cards, so scores zero points and is returning one white token")
				player[i].whitecounters = player[i].whitecounters - 1
			default:
				fmt.Println(player[i].name, "finshed with no cards, so scores zero points")
			}
		} else {
			fmt.Println(player[i].name, "cards are", DisplayHand(i), "and gains", a, "white counters and", b, "black counters, scoring", a+(b*10), "points")
			player[i].whitecounters = player[i].whitecounters + a
			player[i].blackcounters = player[i].blackcounters + b
			player[i].score = player[i].score + a + (b * 10)
		}

	}

	// Sort by score, usign dummy slice to avoid modifying the original player array
	// This is useful for displaying the scores without stuffign up the player Array which craps out the code if I try to use a slice :-(
	var dummyplayerSlice []players
	for i := 0; i < numberplayers; i++ {
		if player[i].name != "" {
			dummyplayerSlice = append(dummyplayerSlice, player[i])
		}
	}

	sort.SliceStable(dummyplayerSlice[:], func(i, j int) bool {
		return dummyplayerSlice[i].score < dummyplayerSlice[j].score
	})

	fmt.Println("------------- End of round totals  ------------------")
	for i := 0; i < numberplayers; i++ {
		fmt.Println(dummyplayerSlice[i].name, "has", dummyplayerSlice[i].whitecounters, "white counters and",
			dummyplayerSlice[i].blackcounters, "black counters and a total score of",
			dummyplayerSlice[i].score)
	}
	// work out who was last to quit or emptied their hand
	//fmt.Println("The last player to quit or empty their hand was:", player[numberplayers-1].name, "with a score of", player[numberplayers-1].score)

	// Reset the playorder for the next round
	/*

			for i := 0; i < numberplayers; i++ {
				player[i].playorder = i // Reset playorder to the original order
			}
			 Change the playorder of players
			player[0].playorder = 6
			player[1].playorder = 5
			player[2].playorder = 2
			player[3].playorder = 1
			player[4].playorder = 3
			player[5].playorder = 0

		// Sort by playorder based on the playorder field
		// This will ensure that the players are sorted in the order they played in the last round
		// This is useful for determining the order of play in the next round.
		sort.SliceStable(player[:], func(i, j int) bool {
			return player[i].playorder < player[j].playorder
		})
	*/
}

// Check for game end
func CheckGameEnd() bool {
	var scorenotmax = true
	for i := 0; i < numberplayers; i++ {
		if player[i].score >= 40 {
			scorenotmax = false
		}
	}
	return scorenotmax
}

// show game over summary and the winner etc
func DisplayGameEnd() {
	fmt.Println("The Game is Over")
	// Sort by score, usign dummy slice to avoid modifying the original player array
	// This is useful for displaying the scores without stuffign up the player Array which craps out the code if I try to use a slice :-(
	var dummyplayerSlice []players
	for i := 0; i < numberplayers; i++ {
		if player[i].name != "" {
			dummyplayerSlice = append(dummyplayerSlice, player[i])
		}
	}

	sort.SliceStable(dummyplayerSlice[:], func(i, j int) bool {
		return dummyplayerSlice[i].score < dummyplayerSlice[j].score
	})

	// Sort by score with the dummy slice
	sort.SliceStable(dummyplayerSlice[:], func(i, j int) bool {
		return dummyplayerSlice[i].score < dummyplayerSlice[j].score
	})
	fmt.Println("------------- Final scores  ------------------")
	for i := 0; i < numberplayers; i++ {
		fmt.Println(dummyplayerSlice[i].name, "has", dummyplayerSlice[i].whitecounters, "white counters and",
			dummyplayerSlice[i].blackcounters, "black counters and a total score of",
			dummyplayerSlice[i].score)
	}
	fmt.Println("The Winner is:", dummyplayerSlice[0].name, "with a score of", dummyplayerSlice[0].score)
}
