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
    for (let i = 1; i < r.time; i++) {
      if (i * (r.time - i) > r.distance) {
        winCount++;
      }
    }
    return winCount;
  });

  return wins.reduce((p, c) => p * c)
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  let time = '';
  let distance = '';

  for (let i = 0; i < input.length; i++) {
    time += input[i].time.toString();
    distance += input[i].distance.toString();
  }

  const r = {
    time: parseInt(time, 10),
    distance: parseInt(distance, 10)
  }

  let winCount = 0;
  for (let i = 1; i < r.time; i++) {
    if (i * (r.time - i) > r.distance) {
      winCount++;
    }
  }
  return winCount;
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
      {
        input: `
        Time:      7  15   30
        Distance:  9  40  200
        `,
        expected: 71503,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
