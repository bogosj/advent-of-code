import run from "aocrunner";

interface pipeMap {
  start: Array<number>,
  pipes: Array<Array<string>>
}

const parseInput = (rawInput: string): pipeMap => {
  const ret = {
    start: [0, 0],
    pipes: []
  };
  rawInput.split('\n').forEach((line, y) => {
    const pipeLine = [];
    line.split('').forEach((char, x) => {
      if (char == 'S') {
        ret.start = [y, x];
      }
      pipeLine.push(char);
    });
    ret.pipes.push(pipeLine);
  });
  return ret
};

function* neighbors(point: Array<number>) {
  yield {
    pos: [point[0] - 1, point[1]],
    dir: 'N'
  }
  yield {
    pos: [point[0] + 1, point[1]],
    dir: 'S'
  }
  yield {
    pos: [point[0], point[1] - 1],
    dir: 'W'
  }
  yield {
    pos: [point[0], point[1] + 1],
    dir: 'E'
  }
};

interface walkState {
  position: Array<number>,
  distance: number,
  visited: Array<string>
}

// Going this direction I can enter these pipes.
const validPipes = {
  'N': '|7F'.split(''),
  'S': '|JL'.split(''),
  'W': '-LF'.split(''),
  'E': '-J7'.split(''),
};

// I can leave this pipe in this direction
const validDirections = {
  '|': 'NS'.split(''),
  '-': 'EW'.split(''),
  'L': 'NE'.split(''),
  'J': 'NW'.split(''),
  '7': 'WS'.split(''),
  'F': 'ES'.split(''),
  'S': 'NSWE'.split('')
}

const firstValidNeighbor = (state: walkState, pipes: pipeMap): Array<number> => {
  for (const neighbor of neighbors(state.position)) {
    if (state.visited.includes(positionKey(neighbor.pos))) {
      continue;
    }
    try {
      const currentPipe = pipes.pipes[state.position[0]][state.position[1]]
      const neighborPipe = pipes.pipes[neighbor.pos[0]][neighbor.pos[1]];
      if (validPipes[neighbor.dir].includes(neighborPipe)) {
        if (validDirections[currentPipe].includes(neighbor.dir)) {
          return neighbor.pos;
        }
      }
    } catch (e) { continue }
  }
};

const positionKey = (p: Array<number>): string => {
  return `${p[0]}|${p[1]}`;
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  const state = {
    position: input.start,
    distance: 0,
    visited: []
  }
  while (true) {
    const nextPosition = firstValidNeighbor(state, input);
    state.distance += 1;
    if (!nextPosition) {
      break;
    }
    state.visited.push(positionKey(state.position));
    state.position = nextPosition;
  }

  return state.distance / 2;
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
        .....
        .S-7.
        .|.|.
        .L-J.
        .....
        `,
        expected: 4,
      },
      {
        input: `
        ..F7.
        .FJ|.
        SJ.L7
        |F--J
        LJ...
        `,
        expected: 8,
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
