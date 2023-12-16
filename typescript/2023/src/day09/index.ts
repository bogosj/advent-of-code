import run from "aocrunner";

const parseInput = (rawInput: string) => {
  return rawInput.split('\n').map(line => {
    return line.split(/ +/).map(token => parseInt(token, 10));
  });
};

const findNextValue = (history: number[]): number => {
  const lists = [[...history]];
  let lastList = lists[lists.length - 1];
  while (!lastList.every(x => x == 0)) {
    const newList = [];
    for (let i = 1; i < lastList.length; i++) {
      newList.push(lastList[i] - lastList[i - 1]);
    }
    lists.push([...newList]);
    lastList = lists[lists.length - 1];
  }
  while (lists.length > 1) {
    lastList = lists.pop();
    const nextList = lists[lists.length - 1];
    nextList.push(nextList[nextList.length - 1] + lastList.pop());
  } 
  return lists[0].pop();
};

const findPreviousValue = (history: number[]): number => {
  const lists = [[...history]];
  let lastList = lists[lists.length - 1];
  while (!lastList.every(x => x == 0)) {
    const newList = [];
    for (let i = 1; i < lastList.length; i++) {
      newList.push(lastList[i] - lastList[i - 1]);
    }
    lists.push([...newList]);
    lastList = lists[lists.length - 1];
  }
  while (lists.length > 1) {
    lastList = lists.pop();
    const nextList = lists[lists.length - 1];
    nextList.unshift(nextList[0] - lastList.shift());
  } 
  return lists[0].shift();
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  return input.map(findNextValue).reduce((a, b) => a + b);
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  return input.map(findPreviousValue).reduce((a, b) => a + b);
};

run({
  part1: {
    tests: [
      {
        input: `
        0 3 6 9 12 15
        1 3 6 10 15 21
        10 13 16 21 30 45
        `,
        expected: 114,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `
        0 3 6 9 12 15
        1 3 6 10 15 21
        10 13 16 21 30 45
        `,
        expected: 2,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
