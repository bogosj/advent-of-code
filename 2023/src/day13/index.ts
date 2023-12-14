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

const columnsToLeft = (pattern: Array<Array<string>>, smudge: boolean): number => {
  const errorsAllowed = smudge ? 1 : 0;
  for (let mirrorLine = 1; mirrorLine < pattern[0].length; mirrorLine++) {
    let l = mirrorLine - 1;
    let r = mirrorLine;
    let errors = 0;

    for (; l >= 0 && r < pattern[0].length; l--, r++) {
      for (let y = 0; y < pattern.length; y++) {
        if (pattern[y][l] != pattern[y][r]) {
          errors++
        }
      }
    }
    if (errors == errorsAllowed) {
      return mirrorLine;
    }
  }
  return 0;
}

const rowsAbove = (pattern: Array<Array<string>>, smudge: boolean): number => {
  const errorsAllowed = smudge ? 1 : 0;
  for (let mirrorLine = 1; mirrorLine < pattern.length; mirrorLine++) {
    let a = mirrorLine - 1;
    let b = mirrorLine;
    let errors = 0;

    for (; a >= 0 && b < pattern.length; a--, b++) {
      for (let x = 0; x < pattern[0].length; x++) {
        if (pattern[a][x] != pattern[b][x]) {
          errors++;
        }
      }
    }
    if (errors == errorsAllowed) {
      return mirrorLine;
    }
  }
  return 0;
}

const scorePattern = (pattern: Array<Array<string>>, smudge: boolean): number => {
  return columnsToLeft(pattern, smudge) + (100 * rowsAbove(pattern, smudge));
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const solve = (pattern: Array<Array<string>>): number => {
    return scorePattern(pattern, false);
  }
  return input.map(solve).reduce((a, b) => a + b);
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const solve = (pattern: Array<Array<string>>): number => {
    return scorePattern(pattern, true);
  }
  return input.map(solve).reduce((a, b) => a + b);
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
        expected: 400,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
