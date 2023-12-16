import run from "aocrunner";

const parseInput = (rawInput: string): Array<Array<string>> => {
  return rawInput.split('\n').map(line => line.split(''));
};

const directionShifts = {
  'N': [-1, 0],
  'S': [1, 0],
  'W': [0, -1],
  'E': [0, 1],
};

const newPositionIsValid = (rocks: Array<Array<string>>, position: Array<number>): boolean => {
  return (
    position[0] >= 0 &&
    position[1] >= 0 &&
    position[0] < rocks.length &&
    position[1] < rocks[0].length
  )
}

const shiftRocks = (rocks: Array<Array<string>>, direction: string) => {
  let moves = 0;
  const shift = directionShifts[direction];

  while (true) {
    for (let y = 0; y < rocks.length; y++) {
      for (let x = 0; x < rocks[0].length; x++) {
        const curr = rocks[y][x];
        if (curr == 'O') {
          const newPosition = [y + shift[0], x + shift[1]];
          if (newPositionIsValid(rocks, newPosition)) {
            if (rocks[newPosition[0]][newPosition[1]] == '.') {
              rocks[y][x] = '.';
              rocks[newPosition[0]][newPosition[1]] = 'O';
              moves++;
            }
          }
        }
      }
    }
    if (moves == 0) {
      break;
    }
    moves = 0;
  }
}

const calculateLoad = (rocks: Array<Array<string>>): number => {
  const height = rocks.length;
  return rocks.map((line, idx) => {
    return line.map(c => {
      if (c != 'O') {
        return 0;
      }
      return height - idx;
    }).reduce((a, b) => a + b)
  }).reduce((a, b) => a + b);
}


const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  shiftRocks(input, 'N');

  return calculateLoad(input);
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  const rotations = ['N', 'W', 'S', 'E'];
  const formationsToPeriod = {};

  for (let i = 0; ; i++) {
    for (let j = 0; j < 4; j++) {
      shiftRocks(input, rotations[j]);
    }
    const formation = input.map(line => line.join('')).join('\n');
    if (!formationsToPeriod[formation]) {
      formationsToPeriod[formation] = [];
    }
    formationsToPeriod[formation].push(i + 1);
    if (formationsToPeriod[formation].length == 3) {
      break;
    }
  }

  const repeats = Object.values(formationsToPeriod).filter((v: Array<number>) => {
    return v.length > 1;
  });
  const periodLength = repeats.length;
  const cycleStart = Math.min(...repeats.map((v: Array<number>) => { return v[0] }));

  const cycleToPick = (1_000_000_000 - cycleStart) % periodLength + cycleStart;
  let answer = 0;
  Object.entries(formationsToPeriod).forEach(([formation, cycles]) => {
    if (cycles[0] == cycleToPick) {
      answer = calculateLoad(parseInput(formation));
    }
  });
  return answer;
};

run({
  part1: {
    tests: [
      {
        input: `
        O....#....
        O.OO#....#
        .....##...
        OO.#O....O
        .O.....O#.
        O.#..O.#.#
        ..O..#O..O
        .......O..
        #....###..
        #OO..#....
        `,
        expected: 136,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `
        O....#....
        O.OO#....#
        .....##...
        OO.#O....O
        .O.....O#.
        O.#..O.#.#
        ..O..#O..O
        .......O..
        #....###..
        #OO..#....
        `,
        expected: 64,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
