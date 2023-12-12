import run from "aocrunner";

const parseInput = (rawInput: string) => {
  return rawInput.split('\n').map(line => line.split(''));
};

const expandUniverse = (galaxyCoords: Array<Array<number>>, universe: Array<Array<string>>, expansionRate: number) => {
  const linesToExpand = [];
  universe.forEach((line, y) => {
    if (line.every(char => char == '.')) {
      linesToExpand.push(y)
    }
  });

  const columnsToExpand = [];
  for (let x = 0; x < universe[0].length; x++) {
    if (universe.map(line => line[x]).every(char => char == '.')) {
      columnsToExpand.push(x);
    }
  }
  const newCoords = [];
  galaxyCoords.forEach(coord => {
    const colsBefore = columnsToExpand.filter(x => x < coord[1]);
    const linesAbove = linesToExpand.filter(y => y < coord[0]);
    newCoords.push(
      [
        coord[0] + (linesAbove.length * (expansionRate - 1)),
        coord[1] + (colsBefore.length * (expansionRate - 1)),
      ]);
  })
  return newCoords;
};

const manhattanDistance = (a: Array<number>, b: Array<number>): number => {
  return Math.abs(a[0] - b[0]) + Math.abs(a[1] - b[1]);
};

const solver = (rawInput: string, expansionRate: number): number => {
  const input = parseInput(rawInput);
  const galaxyCoords = [];
  input.forEach((line, y) => {
    return line.forEach((char, x) => {
      if (char == '#') {
        galaxyCoords.push([y, x]);
      }
    });
  });
  const expandedGalaxyCoords = expandUniverse(galaxyCoords, input, expansionRate);
  const distances = [];
  for (let i = 0; i < expandedGalaxyCoords.length; i++) {
    for (let j = i + 1; j < expandedGalaxyCoords.length; j++) {
      distances.push(manhattanDistance(expandedGalaxyCoords[i], expandedGalaxyCoords[j]));
    }
  }
  return distances.reduce((a, b) => a + b);
}

const part1 = (rawInput: string) => {
  return solver(rawInput, 2);
};

const part2 = (rawInput: string) => {
  return solver(rawInput, 1_000_000);
};

run({
  part1: {
    tests: [
      {
        input: `
        ...#......
        .......#..
        #.........
        ..........
        ......#...
        .#........
        .........#
        ..........
        .......#..
        #...#.....
        `,
        expected: 374,
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
