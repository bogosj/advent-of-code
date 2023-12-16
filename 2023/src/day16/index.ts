import run from "aocrunner";

const parseInput = (rawInput: string): Array<Array<string>> => {
  return rawInput.split('\n').map(line => line.split(''));
};

interface Beam {
  x: number,
  y: number,
  direction: string
}

const onBoard = (board: Array<Array<string>>, beam: Beam): boolean => {
  return (
    beam.x >= 0 &&
    beam.y >= 0 &&
    beam.y < board.length &&
    beam.x < board[0].length
  )
};

const moves = {
  'U': { y: -1, x: 0 },
  'D': { y: 1, x: 0 },
  'L': { x: -1, y: 0 },
  'R': { x: 1, y: 0 },
}

const moveBeam = (beam: Beam, activeBeams: Beam[], visited: string[]) => {
  beam.x += moves[beam.direction].x;
  beam.y += moves[beam.direction].y;
  const newState = `${beam.x},${beam.y},${beam.direction}`;
  if (!visited.includes(newState)) {
    visited.push(newState);
    activeBeams.push(beam);
  }
}

const cloneBeam = (beam: Beam, newDirection: string): Beam => {
  return {
    x: beam.x,
    y: beam.y,
    direction: newDirection
  }
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const energizedTiles = {};
  const visited: Array<string> = [];
  const activeBeams: Array<Beam> = [
    { x: 0, y: 0, direction: 'R' }
  ];
  for (let i = 0; i < 10000000 && activeBeams.length > 0; i++) {
    const beam = activeBeams.shift();
    if (!onBoard(input, beam)) {
      continue;
    }
    energizedTiles[`${beam.x},${beam.y}`] = true;
    const currentTile = input[beam.y][beam.x];
    if (currentTile == '.') {
      moveBeam(beam, activeBeams, visited);
    }
    if (currentTile == '|') {
      if ('UD'.includes(beam.direction)) {
        moveBeam(beam, activeBeams, visited);
      } else {
        moveBeam(cloneBeam(beam, 'U'), activeBeams, visited);
        moveBeam(cloneBeam(beam, 'D'), activeBeams, visited);
      }
    }
    if (currentTile == '-') {
      if ('LR'.includes(beam.direction)) {
        moveBeam(beam, activeBeams, visited);
      } else {
        moveBeam(cloneBeam(beam, 'L'), activeBeams, visited);
        moveBeam(cloneBeam(beam, 'R'), activeBeams, visited);
      }
    }
    if (currentTile == '\\') {
      if (beam.direction == 'R') {
        beam.direction = 'D';
      } else if (beam.direction == 'U') {
        beam.direction = 'L';
      } else if (beam.direction == 'L') {
        beam.direction = 'U';
      } else if (beam.direction == 'D') {
        beam.direction = 'R';
      }
      moveBeam(beam, activeBeams, visited);
    }
    if (currentTile == '/') {
      if (beam.direction == 'R') {
        beam.direction = 'U';
      } else if (beam.direction == 'U') {
        beam.direction = 'R';
      } else if (beam.direction == 'L') {
        beam.direction = 'D';
      } else if (beam.direction == 'D') {
        beam.direction = 'L';
      }
      moveBeam(beam, activeBeams, visited);
    }
  }
  return Object.keys(energizedTiles).length;
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
        .|...\\....
        |.-.\\.....
        .....|-...
        ........|.
        ..........
        .........\\
        ..../.\\\\..
        .-.-/..|..
        .|....-|.\\
        ..//.|....
        `,
        expected: 46,
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
