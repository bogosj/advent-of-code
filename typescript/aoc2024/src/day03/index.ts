import run from "aocrunner";

const parseInput = (rawInput: string) => rawInput;

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const re = /mul\((\d+),(\d+)\)/g;
  const matches = [...input.matchAll(re)];
  let sum = 0;
  matches.forEach(match => {
    sum += Number(match[1]) * Number(match[2]);
  });
  return sum;
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const re = /mul\((\d+),(\d+)\)|do\(\)|don't\(\)/g;
  const matches = [...input.matchAll(re)];
  let sum = 0;
  let enabled = true;
  matches.forEach(match => {
    if (match[0] == "do()") {
      enabled = true;
    } else if (match[0] == "don't()") {
      enabled = false;
    } else {
      if (enabled) {
        sum += Number(match[1]) * Number(match[2]);
      }
    }
  });
  return sum;
};

run({
  part1: {
    tests: [
      {
        input: `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`,
        expected: 161,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`,
        expected: 48,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
