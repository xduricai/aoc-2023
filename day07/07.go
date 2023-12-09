package day07

import "github.com/xduricai/aoc-2023/util"

var cardValueMap = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type Hand struct {
	cards string
	bid   int
}

func (hand *Hand) Compare(input *Hand) int {
	for idx := range (*input).cards {
		a := rune((*hand).cards[idx])
		b := rune((*input).cards[idx])

		if cardValueMap[a] > cardValueMap[b] {
			return 1
		} else if cardValueMap[a] < cardValueMap[b] {
			return -1
		}
	}
	return 0
}

func newHand(cards *string, value int, bid int) *Hand {
	return &Hand{
		cards: *cards,
		bid:   bid,
	}
}

func RankHands(jokers bool) (int, error) {
	id := "07"
	lines, err := util.ReadLines(id)

	if err != nil {
		return *new(int), err
	}
	if jokers {
		cardValueMap['J'] = 1
	} else {
		cardValueMap['J'] = 11
	}

	handSets := [7][]Hand{}

	for _, line := range lines {
		var value int
		cards := line[:5]
		bid := line[6:]

		if jokers {
			value = getJokerHandValue(&cards)
		} else {
			value = getHandValue(&cards)
		}

		hand := newHand(&cards, value, util.ParseIntFromString(&bid))
		handSets[value] = append(handSets[value], *hand)
	}

	for _, set := range handSets {
		quickSort(set)
	}

	var sum int
	multiplier := 1

	for _, set := range handSets {
		for _, hand := range set {
			sum += hand.bid * multiplier
			multiplier++
		}
	}

	return sum, nil
}

func getHandValue(hand *string) int {
	counts := map[rune]int{}
	best := 1
	secondary := 1

	for _, char := range *hand {
		counts[char]++
		count := counts[char]

		if count > best {
			best = count
		} else if count > secondary {
			secondary = count
		}
	}

	switch {
	case best == 5:
		return 6
	case best == 4:
		return 5
	case best == 3 && secondary == 2:
		return 4
	case best == 3:
		return 3
	case best == 2 && secondary == 2:
		return 2
	case best == 2:
		return 1
	default:
		return 0
	}
}

func getJokerHandValue(hand *string) int {
	counts := map[rune]int{}
	best := 1
	secondary := 1
	jokers := 0

	for _, char := range *hand {
		if char == 'J' {
			jokers++
			continue
		}

		counts[char]++
		count := counts[char]

		if count > best {
			best = count
		} else if count > secondary {
			secondary = count
		}
	}

	best += jokers

	switch {
	case best == 5 || jokers == 5:
		return 6
	case best == 4:
		return 5
	case best == 3 && secondary == 2:
		return 4
	case best == 3:
		return 3
	case best == 2 && secondary == 2:
		return 2
	case best == 2:
		return 1
	default:
		return 0
	}
}
