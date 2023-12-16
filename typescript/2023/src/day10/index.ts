import run from "aocrunner";
import { on } from "events";

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

interface neighborT {
  pos: Array<number>,
  dir: string
}

const neighbors = (point: Array<number>): Array<neighborT> => {
  return [
    {
      pos: [point[0] - 1, point[1]],
      dir: 'N'
    },
    {
      pos: [point[0] + 1, point[1]],
      dir: 'S'
    },
    {
      pos: [point[0], point[1] - 1],
      dir: 'W'
    },
    {
      pos: [point[0], point[1] + 1],
      dir: 'E'
    }
  ];
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

const firstValidNeighbor = (state: walkState, pipes: pipeMap): neighborT => {
  for (const neighbor of neighbors(state.position)) {
    if (state.visited.includes(positionKey(neighbor.pos))) {
      continue;
    }
    try {
      const currentPipe = pipes.pipes[state.position[0]][state.position[1]]
      const neighborPipe = pipes.pipes[neighbor.pos[0]][neighbor.pos[1]];
      if (validPipes[neighbor.dir].includes(neighborPipe)) {
        if (validDirections[currentPipe].includes(neighbor.dir)) {
          return neighbor;
        }
      }
    } catch (e) { continue }
  }
};

const positionKey = (p: Array<number>): string => {
  return `${p[0]}|${p[1]}`;
}

interface pipeState {
  position: Array<number>,
  distance: number,
  visited: Array<string>
}

const startPositions = {
  'SW': 'F',
  'SS': '|',
  'SE': '7',
  'WS': 'L',
  'WW': '-',
  'WN': 'F',
  'NW': 'L',
  'NE': 'J',
  'NN': '|',
  'ES': '7',
  'EE': '-',
  'EN': 'L'
}

const findPipeLoop = (input: pipeMap): pipeState => {
  const state = {
    position: input.start,
    distance: 0,
    visited: [],
    exitEnterStart: []
  };
  let lastNeighbor: neighborT = undefined;
  while (true) {
    const neighbor = firstValidNeighbor(state, input);
    if (!neighbor) {
      let delta = [
        input.start[0] - lastNeighbor.pos[0],
        input.start[1] - lastNeighbor.pos[1]
      ]
      let enter = '';
      if (delta[0] == 0 && delta[1] == -1) {
        enter = 'W';
      } else if (delta[0] == 0 && delta[1] == 1) {
        enter = 'E';
      } else if (delta[0] == -1 && delta[1] == 0) {
        enter = 'S';
      } else if (delta[0] == 1 && delta[1] == 0) {
        enter = 'N';
      }
      state.exitEnterStart.push(enter);
    }
    if (state.distance == 0) {
      state.exitEnterStart.push(neighbor.dir);
    }
    state.distance += 1;
    state.visited.push(positionKey(state.position));
    if (!neighbor) {
      break;
    }
    state.position = neighbor.pos;
    lastNeighbor = neighbor;
  }

  let startPos = state.exitEnterStart.join('');
  input.pipes[input.start[0]][input.start[1]] = startPositions[startPos]

  return state;
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const state = findPipeLoop(input);
  return state.distance / 2;
};

const markAllNonPipes = (pipes: Array<Array<string>>, state: pipeState) => {
  pipes.forEach((line, y) => {
    line.forEach((c, x) => {
      if (!state.visited.includes(`${y}|${x}`)) {
        pipes[y][x] = '.';
      }
    })
  })
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const state = findPipeLoop(input);
  markAllNonPipes(input.pipes, state);
  let interiorPointCount = 0;
  input.pipes.forEach((row, idx) => {
    let inside = false;
    row.forEach(c => {
      if (c == '.' && inside) {
        interiorPointCount++;
      }
      if ('F7|'.split('').includes(c)) {
        inside = !inside;
      }
    });
  });
  return interiorPointCount;
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
      {
        input: `
        ...........
        .S-------7.
        .|F-----7|.
        .||.....||.
        .||.....||.
        .|L-7.F-J|.
        .|..|.|..|.
        .L--J.L--J.
        ...........
        `,
        expected: 4,
      },
      {
        input: `
        .F----7F7F7F7F-7....
        .|F--7||||||||FJ....
        .||.FJ||||||||L7....
        FJL7L7LJLJ||LJ.L-7..
        L--J.L7...LJS7F-7L7.
        ....F-J..F7FJ|L7L7L7
        ....L7.F7||L7|.L7L7|
        .....|FJLJ|FJ|F7|.LJ
        ....FJL-7.||.||||...
        ....L---J.LJ.LJLJ...
        `,
        expected: 8,
      },
      {
        input: `
        FF7FSF7F7F7F7F7F---7
        L|LJ||||||||||||F--J
        FL-7LJLJ||||||LJL-77
        F--JF--7||LJLJ7F7FJ-
        L---JF-JLJ.||-FJLJJ7
        |F|F-JF---7F7-L7L|7|
        |FFJF7L7F-JF7|JL---7
        7-L-JL7||F7|L7F-7F7|
        L.L7LFJ|||||FJL7||LJ
        L7JLJL-JLJLJL--JLJ.L
        `,
        expected: 10,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
