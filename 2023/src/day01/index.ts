import run from "aocrunner";

const parseInput = (rawInput: string) => rawInput;

const isCharNumber = (c: string) => {
  return c >= '0' && c <= '9';
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const lines = input.split('\n');
  const digitsOnly = lines.map((line) => {
    return line.split('').filter((c) => isCharNumber(c));
  });
  const values = digitsOnly.map((digits) => {
    let val = parseInt(digits[0], 10) * 10;
    val += parseInt(digits.pop(), 10);
    return val;
  });
  console.log(values);
  return values.reduce((total, val) => total + val);
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  return;
};

run({
  part1: {
    tests: [
      {
        input: `1abc2
         pqr3stu8vwx
         a1b2c3d4e5f
         treb7uchet`,
        expected: 142,
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
