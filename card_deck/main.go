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

func (c *Card) Score() int {
	switch c.Rank {
	case ACE:
		return 11
	case JACK, QUEEN, KING:
		return 10
	default:
		return int(c.Rank) + 1
	}
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

type Hand struct {
	Cards []Card
}

func (h *Hand) AddCard(c Card) {
	h.Cards = append(h.Cards, c)
}

func (h *Hand) Score() int {
	score := 0
	for _, card := range h.Cards {
		score += card.Score()
	}
	return score
}

type Player struct {
	Hand *Hand
}

func NewPlayer() *Player {
	return &Player{
		Hand: &Hand{
			Cards: make([]Card, 0),
		},
	}
}

func main() {
	deck := NewDeck()
	player1 := NewPlayer()
	player2 := NewPlayer()
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
