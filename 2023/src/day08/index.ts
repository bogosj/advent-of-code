import run from "aocrunner";

interface stringDict {
  [key: string]: string[]
}

interface network {
  instructions: string[],
  nodes: stringDict
}

const parseInput = (rawInput: string): network => {
  const net = {
    instructions: rawInput.split('\n')[0].split(''),
    nodes: {}
  };

  rawInput.split('\n').slice(2).forEach(line => {
    const tokens = line.split(/ +/);
    net.nodes[tokens[0]] = [
      tokens[2].replace('(', '').replace(',', ''),
      tokens[3].replace(')', '')
    ];
  });

  return net
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  let currentNode = 'AAA';
  let moveCount = 0;
  while (true) {
    for (let i = 0; i < input.instructions.length; i++) {
      const move = input.instructions[i];
      if (move == 'L') {
        currentNode = input.nodes[currentNode][0];
      } else {
        currentNode = input.nodes[currentNode][1];
      }
      moveCount++;
      if (currentNode == 'ZZZ') {
        return moveCount;
      }
    }
  }
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
        LLR

        AAA = (BBB, BBB)
        BBB = (AAA, ZZZ)
        ZZZ = (ZZZ, ZZZ)
        `,
        expected: 6,
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
