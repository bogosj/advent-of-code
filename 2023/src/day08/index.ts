import run from "aocrunner";

import lcm from 'compute-lcm';

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
  const nodes = Object.keys(input.nodes).filter(x => x.endsWith('A'));
  const nodeMoves = Array.from(Array(nodes.length), () => 0);
  const nodeSolved = Array.from(Array(nodes.length), () => []);
  while (true) {
    for (let i = 0; i < input.instructions.length; i++) {
      for (let j = 0; j < nodes.length; j++) {
        let currentNode = nodes[j];
        const move = input.instructions[i];
        if (move == 'L') {
          currentNode = input.nodes[currentNode][0];
        } else {
          currentNode = input.nodes[currentNode][1];
        }
        nodeMoves[j]++;
        if (currentNode.endsWith('Z')) {
          nodeSolved[j].push(nodeMoves[j]);
        }
        nodes[j] = currentNode;
      }
      if (nodeSolved.every(v => v.length > 0)) {
        return lcm(nodeSolved.map(v => v[0]));
      };
    }
  }
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
      {
        input: `
        LR

        11A = (11B, XXX)
        11B = (XXX, 11Z)
        11Z = (11B, XXX)
        22A = (22B, XXX)
        22B = (22C, 22C)
        22C = (22Z, 22Z)
        22Z = (22B, 22B)
        XXX = (XXX, XXX)
        `,
        expected: 6,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
