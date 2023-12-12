import run from "aocrunner";

const parseInput = (rawInput: string) => {
  return rawInput.split('\n').map(line => line.split(''));
};

const expandUniverse = (universe: Array<Array<string>>) => {
  const expandedY: Array<Array<string>> = [];
  universe.forEach(line => {
    if (line.every(char => char == '.')) {
      expandedY.push(line);
    }
    expandedY.push(line);
  });

  const columnsToExpand = [];
  for (let x = 0; x < expandedY[0].length; x++) {
    if (expandedY.map(line => line[x]).every(char => char == '.')) {
      columnsToExpand.push(x);
    }
  }
  const expandedX = [];
  expandedY.forEach(line => {
    const newLine: Array<string> = [];
    line.forEach((char, idx) => {
      if (columnsToExpand.includes(idx)) {
        newLine.push('.');
      }
      newLine.push(char);
    })
    expandedX.push(newLine);
  });
  return expandedX;
};

const manhattanDistance = (a: Array<number>, b: Array<number>): number => {
  return Math.abs(a[0] - b[0]) + Math.abs(a[1] - b[1]);
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const galaxyCoords = [];
  expandUniverse(input).forEach((line, y) => {
    return line.forEach((char, x) => {
      if (char == '#') {
        galaxyCoords.push([y, x]);
      }
    });
  });
  const distances = [];
  for (let i = 0; i < galaxyCoords.length; i++) {
    for (let j = i + 1; j < galaxyCoords.length; j++) {
      distances.push(manhattanDistance(galaxyCoords[i], galaxyCoords[j]));
    }
  }
  return distances.reduce((a, b) => a + b);
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
