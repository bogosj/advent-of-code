import run from "aocrunner";

const parseInput = (rawInput: string): Array<string> => {
  return rawInput.split(',');
};

const hashString = (s: string): number => {
  let current = 0;
  s.split('').forEach(c => {
    current += c.charCodeAt(0);
    current *= 17;
    current %= 256;
  })
  return current;
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  return input.map(hashString).reduce((a, b) => a + b);
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  return;
};

run({
  part1: {
    tests: [
      {
        input: `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`,
        expected: 1320,
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
