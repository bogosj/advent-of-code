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
  return values.reduce((total, val) => total + val);
};

const convertWordsToNumbers = (rawInput: string): string => {
  const nums = {
    'one': '1',
    'two': '2',
    'three': '3',
    'four': '4',
    'five': '5',
    'six': '6',
    'seven': '7',
    'eight': '8',
    'nine': '9',
  }
  return rawInput.split('\n').map(line => {
    let newLine = '';
    while (line) {
      let found = false;
      Object.entries(nums).forEach(([key, val]) => {
        if (line.startsWith(key)) {
          newLine += val;
          found = true;
        }
      });
      if (!found) {
        newLine += line[0];
      }
      line = line.slice(1);
    }
    return newLine;
  }).join('\n');
}

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  return part1(convertWordsToNumbers(input));
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
      {
        input: `two1nine
        eightwothree
        abcone2threexyz
        xtwone3four
        4nineeightseven2
        zoneight234
        7pqrstsixteen`,
        expected: 281,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
