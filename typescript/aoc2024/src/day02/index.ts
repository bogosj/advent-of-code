import run from "aocrunner";

const parseInput = (rawInput: string) => rawInput;

const inOrder = (l: number[]): boolean => {
  const sorted = [...l].sort((a, b) => a - b);
  if (JSON.stringify(l) == JSON.stringify(sorted)) {
    return true;
  }
  sorted.reverse();
  if (JSON.stringify(l) == JSON.stringify(sorted)) {
    return true;
  }
  return false;
};

const isSafe = (vals: number[]): number => {
  if (!inOrder(vals)) {
    return 0;
  }
  for (let i = 0; i < vals.length - 1; i++) {
    const delta = Math.abs(vals[i] - vals[i + 1]);
    if (delta < 1 || delta > 3) {
      return 0;
    }
  }
  return 1;
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  let safeCount = 0;
  input.split('\n').forEach(line => {
    const vals = line.split(' ').map(Number);
    safeCount += isSafe(vals);
  });

  return safeCount;
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  let safeCount = 0;
  input.split('\n').forEach(line => {
    const vals = line.split(' ').map(Number);
    for (let i = 0; i < vals.length; i++) {
      const newVals = [...vals];
      newVals.splice(i, 1);
      const innerSafe = isSafe(newVals);
      if (innerSafe > 0) {
        safeCount += innerSafe;
        return;
      }
    }
  });

  return safeCount;
};

run({
  part1: {
    tests: [
      {
        input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
        expected: 2,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
        expected: 4,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
