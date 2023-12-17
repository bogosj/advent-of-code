import run from "aocrunner";

const parseInput = (rawInput: string): Array<any> => {
  const ret = [];
  let pair = [];
  const lines = rawInput.split('\n');
  for (let i = 0; i < lines.length; i++) {
    if (lines[i] == '') {
      ret.push(pair);
      pair = [];
    } else {
      pair.push(eval(lines[i]));
    }
  }
  ret.push(pair);
  return ret;
};

const pairInOrder = (left: any, right: any): boolean => {
  if (typeof left === 'number' && typeof right === 'number') {
    return left > right ? false : left < right ? true : undefined;
  } else if (Array.isArray(left) !== Array.isArray(right)) {
    return pairInOrder(Array.isArray(left) ? left : [left], Array.isArray(right) ? right : [right]);
  }

  for (let i = 0, end = Math.max(left.length, right.length); i < end; i++) {
    if (left[i] === undefined) return true;
    if (right[i] === undefined) return false;
    const result = pairInOrder(left[i], right[i]);
    if (result !== undefined) return result;
  }
  return undefined;
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  return input.map((pair, idx) => {
    if (pairInOrder(pair[0], pair[1])) {
      return idx + 1;
    }
    return 0;
  }).reduce((a, b) => a + b);
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
        [1,1,3,1,1]
        [1,1,5,1,1]

        [[1],[2,3,4]]
        [[1],4]

        [9]
        [[8,7,6]]

        [[4,4],4,4]
        [[4,4],4,4,4]

        [7,7,7,7]
        [7,7,7]

        []
        [3]

        [[[]]]
        [[]]

        [1,[2,[3,[4,[5,6,7]]]],8,9]
        [1,[2,[3,[4,[5,6,0]]]],8,9]
        `,
        expected: 13,
      },
      {
        input: `
        [[1],[2,3,4]]
        [[1],4]
        `,
        expected: 1
      },
      {
        input: `
        [[1],4]
        [[1],[2,3,4]]
        `,
        expected: 0
      }
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
