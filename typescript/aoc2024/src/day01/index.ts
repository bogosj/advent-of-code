import run from "aocrunner";

const parseInput = (rawInput: string) => rawInput;

function getLists(input: string) {
  const left = [];
  const right = [];
  input.split('\n').forEach(line => {
    const data = line.split(/\s+/);
    left.push(parseInt(data[0], 10));
    right.push(parseInt(data[1], 10));
  });
  return { left, right };
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const { left, right } = getLists(input);
  left.sort();
  right.sort();

  let total = 0;
  for (let i = 0; i < left.length; i++) {
    total += Math.abs(left[i] - right[i]);
  }

  return total;
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const { left, right } = getLists(input);

  return left.reduce(
    (acc, cur) => {
      return acc + cur * right.reduce(
        (acc2, cur2) => {
          if (cur == cur2) {
            return acc2 + 1;
          } else {
            return acc2 + 0;
          }
        },
        0
      );
    },
    0
  );
};

run({
  part1: {
    tests: [
      {
        input: `3   4
4   3
2   5
1   3
3   9
3   3`,
        expected: 11,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `3   4
4   3
2   5
1   3
3   9
3   3`,
        expected: 31,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
