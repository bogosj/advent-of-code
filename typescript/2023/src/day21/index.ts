import run from "aocrunner";

enum MapTile {
  GardenPlot = '.',
  Rock = '#',
  Start = 'S'
}

type Point = {
  x: number,
  y: number
}

const directionDelta = {
  'U': { x: 0, y: -1 },
  'D': { x: 0, y: 1 },
  'L': { x: -1, y: 0 },
  'R': { x: 1, y: 0 }
}

const parseInput = (rawInput: string): string[][] => {
  return rawInput.split('\n').map(line => line.split(''));
};

const findStart = (garden: string[][]): Point[] => {
  for (let y = 0; y < garden.length; y++) {
    for (let x = 0; x < garden[0].length; x++) {
      if (garden[y][x] == 'S') {
        return [{ x: x, y: y }]
      }
    }
  }
}

const safePoint = (garden: string[][], point: Point): boolean => {
  if (
    point.y >= 0 &&
    point.y < garden.length &&
    point.x >= 0 &&
    point.x < garden[0].length
  ) {
    return garden[point.y][point.x] != '#';
  } else {
    return false
  }
}

const walk = (garden: string[][], steps: number): number => {
  let validPoints = findStart(garden);
  for (let step = 0; step < steps; step++) {
    let nextValidPoints = [];
    while (validPoints.length > 0) {
      const nextPoint = validPoints.shift();
      Object.values(directionDelta).forEach((delta) => {
        const p = {
          y: nextPoint.y + delta.y,
          x: nextPoint.x + delta.x
        };
        if (safePoint(garden, p)) {
          nextValidPoints.push(`${p.y}|${p.x}`);
        }
      });
    }
    validPoints = [...new Set(nextValidPoints)].map(s => {
      const t = s.split('|');
      return { y: parseInt(t[0], 10), x: parseInt(t[1], 10) }
    });
  }
  return validPoints.length;
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  if (input.length == 11) { // test case
    return walk(input, 6);
  } else {
    return walk(input, 64);
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
        ...........
        .....###.#.
        .###.##..#.
        ..#.#...#..
        ....#.#....
        .##..S####.
        .##..#...#.
        .......##..
        .##.#.####.
        .##..##.##.
        ...........
        `,
        expected: 16,
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
