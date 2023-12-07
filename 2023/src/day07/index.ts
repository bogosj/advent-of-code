import run from "aocrunner";

interface handOfCards {
  cards: Array<string>,
  bid: number
}

const cardToRank = {};
'23456789TJQKA'.split('').forEach((card, idx) => {
  cardToRank[card] = idx;
})

const parseInput = (rawInput: string): Array<handOfCards> => {
  return rawInput.split('\n').map(line => {
    const tokens = line.split(/ +/);
    return {
      cards: tokens[0].split(''),
      bid: parseInt(tokens[1], 10)
    }
  })
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
  return typeVals[typeArray.reduce((p, c) => (p * 10 + c))];
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
    if (cardToRank[a.cards[i]] < cardToRank[b.cards[i]]) {
      return -1;
    }
    if (cardToRank[b.cards[i]] < cardToRank[a.cards[i]]) {
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

  return;
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
      // {
      //   input: ``,
      //   expected: "",
      // },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
