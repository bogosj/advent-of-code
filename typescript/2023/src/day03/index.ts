import run from "aocrunner";

const parseInput = (rawInput: string) => {
  return rawInput.split('\n').map(line => {
    return line.split('');
  });
};

const isNumber = (c: string): boolean => {
  return c >= '0' && c <= '9';
};

const isSymbol = (c: string): boolean => {
  if (!c) { return false; }
  return !(isNumber(c) || c == '.');
};

const touchesSymbol = (input: string[][], y: number, x: number): boolean => {
  for (let y2 = y - 1; y2 <= y + 1; y2++) {
    for (let x2 = x - 1; x2 <= x + 1; x2++) {
      try {
        if (isSymbol(input[y2][x2])) {
          return true;
        }
      } catch (error) { }
    }
  }
  return false
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const goodNumbers = [];
  let accumulator = 0;
  let isGoodNumber = false;
  for (let y = 0; y < input.length; y++) {
    for (let x = 0; x < input[y].length; x++) {
      if (isNumber(input[y][x])) {
        accumulator = accumulator * 10 + parseInt(input[y][x], 10);
        isGoodNumber = isGoodNumber || touchesSymbol(input, y, x);
      } else {
        if (isGoodNumber) {
          goodNumbers.push(accumulator);
        }
        accumulator = 0;
        isGoodNumber = false;
      }
    }
  }
  return goodNumbers.reduce((prev, curr) => prev + curr);
};

const touchesStar = (input: string[][], y: number, x: number): { x: number, y: number } => {
  for (let y2 = y - 1; y2 <= y + 1; y2++) {
    for (let x2 = x - 1; x2 <= x + 1; x2++) {
      try {
        if (input[y2][x2] == '*') {
          return { x: x2, y: y2 };
        }
      } catch (error) { }
    }
  }
  return undefined;
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  const gears = {};

  let accumulator = 0;
  let starPosition = undefined;
  for (let y = 0; y < input.length; y++) {
    for (let x = 0; x < input[y].length; x++) {
      if (isNumber(input[y][x])) {
        accumulator = accumulator * 10 + parseInt(input[y][x], 10);
        starPosition = starPosition || touchesStar(input, y, x);
      } else {
        if (starPosition) {
          const key = `${starPosition.x}|${starPosition.y}`;
          if (!gears[key]) {
            gears[key] = [];
          }
          gears[key].push(accumulator);
        }
        accumulator = 0; 
        starPosition = undefined;
      }
    }
  }

  let retVal = 0;
  Object.values(gears).forEach(parts => {
    if (parts.length == 2) {
      retVal += parts[0] * parts[1];
    }
  });
  return retVal;
};

run({
  part1: {
    tests: [
      {
        input: `
        467..114..
        ...*......
        ..35..633.
        ......#...
        617*......
        .....+.58.
        ..592.....
        ......755.
        ...$.*....
        .664.598..
`,
        expected: 4361,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `
        467..114..
        ...*......
        ..35..633.
        ......#...
        617*......
        .....+.58.
        ..592.....
        ......755.
        ...$.*....
        .664.598..
`,
        expected: 467835,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
