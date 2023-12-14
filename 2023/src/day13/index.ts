import run from "aocrunner";

const parseInput = (rawInput: string): Array<Array<Array<string>>> => {
  const patterns = [];
  let pattern = [];
  rawInput.split('\n').forEach(line => {
    if (line == '') {
      patterns.push([...pattern]);
      pattern = [];
    } else {
      pattern.push(line.split(''));
    }
  });
  patterns.push([...pattern]);
  return patterns;
};

const columnsToLeft = (pattern: Array<Array<string>>): number => {
  for (let mirrorLine = 1; mirrorLine < pattern[0].length; mirrorLine++) {
    let l = mirrorLine - 1;
    let r = mirrorLine;
    let found = true;

    for (; found && l >= 0 && r < pattern[0].length; l--, r++) {
      for (let y = 0; y < pattern.length; y++) {
        if (pattern[y][l] != pattern[y][r]) {
          found = false;
          break;
        }
      }
    }
    if (found) {
      return mirrorLine;
    }
  }
  return 0;
}

const rowsAbove = (pattern: Array<Array<string>>): number => {
  for (let mirrorLine = 1; mirrorLine < pattern.length; mirrorLine++) {
    let a = mirrorLine - 1;
    let b = mirrorLine;
    let found = true;

    for (; found && a >= 0 && b < pattern.length; a--, b++) {
      for (let x = 0; x < pattern[0].length; x++) {
        if (pattern[a][x] != pattern[b][x]) {
          found = false;
          break;
        }
      }
    }
    if (found) {
      return mirrorLine;
    }
  }
  return 0;
}

const scorePattern = (pattern: Array<Array<string>>): number => {
  return columnsToLeft(pattern) + (100 * rowsAbove(pattern));
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  return input.map(scorePattern).reduce((a, b) => a + b);
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
        #.##..##.
        ..#.##.#.
        ##......#
        ##......#
        ..#.##.#.
        ..##..##.
        #.#.##.#.
        
        #...##..#
        #....#..#
        ..##..###
        #####.##.
        #####.##.
        ..##..###
        #....#..#
        `,
        expected: 405,
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
