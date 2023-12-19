import run from "aocrunner";

interface Instruction {
  dir: string,
  len: number,
  color: string,
}

const moves = {
  'U': { y: -1, x: 0 },
  'D': { y: 1, x: 0 },
  'L': { x: -1, y: 0 },
  'R': { x: 1, y: 0 },
}

const parseInput = (rawInput: string): Instruction[] => {
  return rawInput.split('\n').map(line => {
    const tokens = line.split(' ');
    return {
      dir: tokens[0],
      len: parseInt(tokens[1], 10),
      color: tokens[2].substring(2, tokens[2].length - 1)
    }
  });
};

interface Point {
  x: number,
  y: number,
}

const getArea = (verticies: Point[]): number => {
  let area = 0;

  let j = verticies.length - 1;
  for (let i = 0; i < verticies.length; i++) {
    area += (verticies[j].x + verticies[i].x) * (verticies[j].y - verticies[i].y);
    j = i;
  }

  return Math.abs(area) / 2;
};

const trenchArea = (instructions: Instruction[]): number => {
  const pos = { x: 0, y: 0 };
  const verticies = [];
  let perimeter = 0;
  instructions.forEach((instruction) => {
    pos.x += moves[instruction.dir].x * instruction.len;
    pos.y += moves[instruction.dir].y * instruction.len;
    perimeter += instruction.len;
    verticies.push({ x: pos.x, y: pos.y });
  });

  return getArea(verticies) + perimeter / 2 + 1;
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  return trenchArea(input);
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
        R 6 (#70c710)
        D 5 (#0dc571)
        L 2 (#5713f0)
        D 2 (#d2c081)
        R 2 (#59c680)
        D 2 (#411b91)
        L 5 (#8ceee2)
        U 2 (#caa173)
        L 1 (#1b58a2)
        U 2 (#caa171)
        R 2 (#7807d2)
        U 3 (#a77fa3)
        L 2 (#015232)
        U 2 (#7a21e3)
        `,
        expected: 62,
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

