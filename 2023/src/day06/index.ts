import run from "aocrunner";

interface race {
  time: number
  distance: number
}

const parseInput = (rawInput: string): race[] => {
  const times = rawInput.split('\n')[0].split(/ +/).slice(1).map(c => parseInt(c, 10));
  const distances = rawInput.split('\n')[1].split(/ +/).slice(1).map(c => parseInt(c, 10));

  const races: race[] = [];

  for (let i = 0; i < times.length; i++) {
    races.push({ time: times[i], distance: distances[i] });
  }

  return races;
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  const wins = input.map(r => {
    let winCount = 0;
    for (let i=1; i<r.time; i++) {
      if (i*(r.time - i) > r.distance) {
        winCount++;
      }
    }
    return winCount;
  });

  return wins.reduce((p, c) => p * c)
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
        Time:      7  15   30
        Distance:  9  40  200
        `,
        expected: 288,
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
