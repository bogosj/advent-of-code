import run from "aocrunner";
import { type } from "os";

interface handOfCards {
  cards: Array<string>,
  bid: number,
  jokersWild: boolean,
  cardRanks: Object
}

const cardToRank = {};
'23456789TJQKA'.split('').forEach((card, idx) => {
  cardToRank[card] = idx;
})

const cardToRankPart2 = {};
'J23456789TQKA'.split('').forEach((card, idx) => {
  cardToRankPart2[card] = idx;
})

const parseInput = (rawInput: string): Array<handOfCards> => {
  return rawInput.split('\n').map(line => {
    const tokens = line.split(/ +/);
    return {
      cards: tokens[0].split(''),
      bid: parseInt(tokens[1], 10),
      jokersWild: false,
      cardRanks: cardToRank
    }
  })
};

const strengthenHand = (hand: handOfCards, typeOfHand: number): number => {
  if (!hand.jokersWild || !hand.cards.includes('J')) {
    return typeOfHand;
  }
  const jackCount = hand.cards.filter(card => card == 'J').length;
  if (jackCount == 1) {
    if (typeOfHand == 5 || typeOfHand == 0) {
      return typeOfHand + 1
    }
    return typeOfHand + 2;
  }
  if (jackCount == 2) {
    if (typeOfHand == 2) {
      return typeOfHand + 3;
    }
    return typeOfHand + 2;
  }
  if (jackCount == 3) {
    return typeOfHand + 2;
  }
  return 6;
};

const handToType = (hand: handOfCards): number => {
  const cardCounts = {};
  hand.cards.forEach(card => {
    cardCounts[card] = cardCounts[card] + 1 || 1;
  });
  const typeArray: Array<number> = Object.values(cardCounts);
  typeArray.sort();
  const typeVals = {
    5: 6,
    14: 5,
    23: 4,
    113: 3,
    122: 2,
    1112: 1,
    11111: 0
  };
  const typeOfHand = typeArray.reduce((p, c) => (p * 10 + c));
  return strengthenHand(hand, typeVals[typeOfHand]);
};

const compareCardHands = (a: handOfCards, b: handOfCards): number => {
  const typeA = handToType(a);
  const typeB = handToType(b);
  if (typeA < typeB) {
    return -1;
  }
  if (typeB < typeA) {
    return 1;
  }

  // Same type order by cards
  for (let i = 0; i < 5; i++) {
    if (a.cardRanks[a.cards[i]] < a.cardRanks[b.cards[i]]) {
      return -1;
    }
    if (a.cardRanks[b.cards[i]] < a.cardRanks[a.cards[i]]) {
      return 1;
    }
  }
  return 0;
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  input.sort(compareCardHands);
  return input.map((hand, idx) => {
    return hand.bid * (idx + 1);
  }).reduce((p, c) => p + c);
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  input.forEach(hand => {
    hand.jokersWild = true;
    hand.cardRanks = cardToRankPart2;
  });
  input.sort(compareCardHands);
  return input.map((hand, idx) => {
    return hand.bid * (idx + 1);
  }).reduce((p, c) => p + c);
};

run({
  part1: {
    tests: [
      {
        input: `
        32T3K 765
        T55J5 684
        KK677 28
        KTJJT 220
        QQQJA 483
        `,
        expected: 6440,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `
        32T3K 765
        T55J5 684
        KK677 28
        KTJJT 220
        QQQJA 483
        `,
        expected: 5905,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
