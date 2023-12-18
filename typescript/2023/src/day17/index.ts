import run from "aocrunner";
import {
  MinPriorityQueue,
  IGetCompareValue,
} from "@datastructures-js/priority-queue";

const parseInput = (rawInput: string): Array<Array<number>> => {
  return rawInput.split('\n').map(line => {
    return line.split('').map(c => parseInt(c, 10));
  })
};

interface WalkState {
  x: number,
  y: number,
  heatLoss: number,
  dir: string,
  stepsInDir: number
}

const getHeatLoss: IGetCompareValue<WalkState> = (state) => state.heatLoss;

const moves = {
  'U': { y: -1, x: 0 },
  'D': { y: 1, x: 0 },
  'L': { x: -1, y: 0 },
  'R': { x: 1, y: 0 },
}

const nextDirection = {
  'U': 'LRU'.split(''),
  'D': 'LRD'.split(''),
  'L': 'UDL'.split(''),
  'R': 'UDR'.split('')
}

const inCity = (state: WalkState, city: Array<Array<number>>): boolean => {
  return (
    state.x >= 0 &&
    state.y >= 0 &&
    state.x < city[0].length &&
    state.y < city.length
  )
}

const atGoal = (state: WalkState, city: Array<Array<number>>): boolean => {
  return (
    state.x == city[0].length - 1 &&
    state.y == city.length - 1
  )
}

const visitedKey = (s: WalkState): string => { return `${s.y},${s.x},${s.dir},${s.stepsInDir}`; }

const shortestPath = (city: Array<Array<number>>, minBeforeTurn: number, maxBeforeTurn: number): number => {
  const visited = {};
  const statesQueue = new MinPriorityQueue<WalkState>(getHeatLoss);
  statesQueue.enqueue({ x: 1, y: 0, heatLoss: 0, dir: 'R', stepsInDir: 1 });
  statesQueue.enqueue({ x: 0, y: 1, heatLoss: 0, dir: 'D', stepsInDir: 1 });

  for (let i = 0; i < 10_000_000 && !statesQueue.isEmpty(); i++) {
    const next = statesQueue.dequeue();
    // Are we out of bounds?
    if (!inCity(next, city)) {
      continue
    }

    // Have we been here before, entering from this direciton?
    // If so, check if the prior way was better.
    const last = visited[visitedKey(next)];
    if (last) {
      continue;
    }
    visited[visitedKey(next)] = next;

    next.heatLoss += city[next.y][next.x];
    if (atGoal(next, city)) {
      if (next.stepsInDir < minBeforeTurn) {
        continue;
      }
      return next.heatLoss;
    }

    nextDirection[next.dir].forEach(dir => {
      const newNext: WalkState = {
        x: next.x + moves[dir].x,
        y: next.y + moves[dir].y,
        heatLoss: next.heatLoss,
        dir: dir,
        stepsInDir: next.stepsInDir + 1
      }
      if (dir != next.dir) { // Turning
        if (next.stepsInDir < minBeforeTurn) {
          return;
        }
        newNext.stepsInDir = 1;
      } else { // Going straight
        // Too too many forward steps
        if (newNext.stepsInDir > maxBeforeTurn) {
          return;
        }
      }
      statesQueue.enqueue(newNext);
    });
  }

  return Number.MAX_VALUE;
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  return shortestPath(input, 1, 3);
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  return shortestPath(input, 4, 10);
};

run({
  part1: {
    tests: [
      {
        input: `
        2413432311323
        3215453535623
        3255245654254
        3446585845452
        4546657867536
        1438598798454
        4457876987766
        3637877979653
        4654967986887
        4564679986453
        1224686865563
        2546548887735
        4322674655533
        `,
        expected: 102,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `
        2413432311323
        3215453535623
        3255245654254
        3446585845452
        4546657867536
        1438598798454
        4457876987766
        3637877979653
        4654967986887
        4564679986453
        1224686865563
        2546548887735
        4322674655533
        `,
        expected: 94,
      },
      {
        input: `
        111111111111
        999999999991
        999999999991
        999999999991
        999999999991
        `,
        expected: 71,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
