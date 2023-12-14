import run from "aocrunner";
import { group } from "console";

interface springCondition {
  pattern: string,
  runLengths: Array<number>
}

const parseInput = (rawInput: string): Array<springCondition> => {
  return rawInput.split('\n').map(line => {
    return {
      pattern: line.split(' ')[0],
      runLengths: line.split(' ')[1].split(',').map(v => parseInt(v, 10))
    }
  });
};

interface stringToNum {
  [key: string]: number;
}

const solver = (s: springCondition): number => {
  let validPermutations: stringToNum = {};
  validPermutations['0,0'] = 1;

  s.pattern.split('').forEach(c => {
    const nextStates = [];
    Object.entries(validPermutations).forEach(([key, count]) => {
      const [groupIdx, groupCount] = key.split(',').map(i => parseInt(i, 10));

      // Either a # or a ?, treat as a #
      if (c != '.') {
        if (groupIdx < s.runLengths.length && groupCount < s.runLengths[groupIdx]) {
          nextStates.push([groupIdx, groupCount + 1, count]);
        }
      }

      // Either a . or a ?, treat as a .
      if (c != '#') {
        if (groupCount == 0) {
          nextStates.push([groupIdx, groupCount, count]);
        } else if (groupCount == s.runLengths[groupIdx]) {
          nextStates.push([groupIdx + 1, 0, count]);
        }
      }
      validPermutations = {};

      nextStates.forEach(([groupIdx, groupCount, count]) => {
        const key = `${groupIdx},${groupCount}`;
        if (!validPermutations[key]) {
          validPermutations[key] = 0;
        }
        validPermutations[key] += count;
      })
    });
  });

  return Object.entries(validPermutations).filter(([k, v]) => {
    const [groupIdx, groupCount] = k.split(',').map(i => parseInt(i, 10));
    return (
      groupIdx == s.runLengths.length ||
      (groupIdx == s.runLengths.length - 1 && groupCount == s.runLengths[groupIdx])
    )
  }).map(v => v[1]).reduce((a, b) => a + b);
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  let arrangements = 0;
  for (let i = 0; i < input.length; i++) {
    arrangements += solver(input[i]);
  }

  return arrangements;
};

const unfold = (s: springCondition): springCondition => {
  return {
    pattern: Array(5).fill(s.pattern).join('?'),
    runLengths: Array(5).fill(s.runLengths).flat(),
  };
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  let arrangements = 0;
  for (let i = 0; i < input.length; i++) {
    arrangements += solver(unfold(input[i]));
  }
  return arrangements;
};

run({
  part1: {
    tests: [
      {
        input: `
        ???.### 1,1,3
        .??..??...?##. 1,1,3
        ?#?#?#?#?#?#?#? 1,3,1,6
        ????.#...#... 4,1,1
        ????.######..#####. 1,6,5
        ?###???????? 3,2,1
        `,
        expected: 21,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `
        ???.### 1,1,3
        .??..??...?##. 1,1,3
        ?#?#?#?#?#?#?#? 1,3,1,6
        ????.#...#... 4,1,1
        ????.######..#####. 1,6,5
        ?###???????? 3,2,1
        `,
        expected: 525152,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
