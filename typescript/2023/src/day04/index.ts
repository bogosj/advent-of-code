import run from "aocrunner";

const parseInput = (rawInput: string) => {
  return rawInput.split('\n').map(line => {
    return line.split(/ +/);
  })
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  let score = 0;
  input.map(card => {
    const winningNumbers = [];
    const myNumbers = [];
    card = card.slice(2);
    let next = card.shift();
    while (next != '|') {
      winningNumbers.push(parseInt(next, 10));
      next = card.shift();
    }
    next = card.shift();
    while (next) {
      myNumbers.push(parseInt(next, 10));
      next = card.shift();
    }
    let cardScore = 0;
    winningNumbers.map(num => {
      if (myNumbers.includes(num)) {
        if (cardScore == 0) {
          cardScore = 1;
        } else {
          cardScore *= 2;
        }
      }
    });
    score += cardScore;
  })
  return score;
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  let cardScores = [];
  input.map(card => {
    const winningNumbers = [];
    const myNumbers = [];
    card = card.slice(2);
    let next = card.shift();
    while (next != '|') {
      winningNumbers.push(parseInt(next, 10));
      next = card.shift();
    }
    next = card.shift();
    while (next) {
      myNumbers.push(parseInt(next, 10));
      next = card.shift();
    }
    let cardScore = 0;
    winningNumbers.map(num => {
      if (myNumbers.includes(num)) {
        cardScore += 1;
      }
    });
    cardScores.push(cardScore);
  });
  const cardCounts = Array(cardScores.length).fill(1);
  for (let i = 0; i < cardScores.length; i++) {
    for (let j = 0; j < cardScores[i]; j++) {
      cardCounts[i + j + 1] += cardCounts[i];
    } 
  } 
  return cardCounts.reduce((p, c) => p + c);
};

run({
  part1: {
    tests: [
      {
        input: `
        Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
        Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
        Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
        Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
        Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
        Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
        `,
        expected: 13,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `
        Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
        Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
        Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
        Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
        Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
        Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
        `,
        expected: 30,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
