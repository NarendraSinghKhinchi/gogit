package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Suit int

const (
	SPADES Suit = iota
	HEART
	DIAMOND
	CLUB
)

type Rank int

const (
	ACE Rank = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

type Card struct {
	Suit Suit
	Rank Rank
}

func (c *Card) String() string {
	rank := ""
	suit := ""
	switch c.Rank {
	case ACE:
		rank = "Ace"
	case TWO:
		rank = "Two"
	case THREE:
		rank = "Three"
	case FOUR:
		rank = "Four"
	case FIVE:
		rank = "Five"
	case SIX:
		rank = "Six"
	case SEVEN:
		rank = "Seven"
	case EIGHT:
		rank = "Eight"
	case NINE:
		rank = "Nine"
	case TEN:
		rank = "Ten"
	case JACK:
		rank = "Jack"
	case QUEEN:
		rank = "Queen"
	case KING:
		rank = "King"
	}
	switch c.Suit {
	case SPADES:
		suit = "Spades"
	case CLUB:
		suit = "Club"
	case DIAMOND:
		suit = "Diamond"
	case HEART:
		suit = "Heart"
	}
	return fmt.Sprintf("%s of %s", rank, suit)
}

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	deck := &Deck{
		Cards: make([]Card, 0, 52),
	}

	for _, suit := range []Suit{SPADES, CLUB, DIAMOND, HEART} {
		for _, rank := range []Rank{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING} {
			deck.Cards = append(deck.Cards, Card{
				Suit: suit,
				Rank: rank,
			})
		}
	}
	deck.Shuffle()
	return deck
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Deal() (Card, bool) {
	if len(d.Cards) == 0 {
		return Card{}, false
	}
	c := d.Cards[0]
	d.Cards = d.Cards[1:]
	return c, true
}

type Hand interface {
	AddCard(c Card)
	Score() int
}

type BlackjackHand struct {
	cards []Card
}

func (h *BlackjackHand) AddCard(c Card) {
	h.cards = append(h.cards, c)
}

func (h *BlackjackHand) Score() int {
	score := 0
	aces := 0
	for _, card := range h.cards {
		switch card.Rank {
		case ACE:
			aces++
			score += 11
		case JACK, QUEEN, KING:
			score += 10
		default:
			score += int(card.Rank) + 1
		}
	}

	for score > 21 && aces > 0 {
		score -= 10
		aces--
	}
	return score
}

type Player struct {
	Hand Hand
}

func NewPlayer(h Hand) *Player {
	return &Player{
		Hand: h,
	}
}

// GAME: BACCARAT ---
type BaccaratHand struct {
	cards []Card
}

func (h *BaccaratHand) AddCard(c Card) {
	h.cards = append(h.cards, c)
}

func (h *BaccaratHand) Score() int {
	score := 0
	for _, card := range h.cards {
		val := 0
		switch card.Rank {
		case ACE:
			val = 1
		case TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE:
			val = int(card.Rank) + 1
		case TEN, JACK, QUEEN, KING:
			val = 0 // Face cards and 10s are worth 0 in Baccarat
		}
		score += val
	}
	return score % 10 // Baccarat score is modulo 10
}

func main() {
	deck := NewDeck()
	player1 := NewPlayer(&BlackjackHand{})
	player2 := NewPlayer(&BlackjackHand{})
	for i := range 52 {
		card, ok := deck.Deal()
		if !ok {
			fmt.Println("Deck is empty")
			break
		}
		if i%2 == 0 {
			player1.Hand.AddCard(card)
		} else {
			player2.Hand.AddCard(card)
		}
		// fmt.Println(card.String())
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("Player 1 score: ", player1.Hand.Score())
	fmt.Println("Player 2 score: ", player2.Hand.Score())

	if player1.Hand.Score() > player2.Hand.Score() {
		fmt.Println("The Winner is: Player 1")
	} else {
		fmt.Println("The Winner is: Player 2")
	}
}
