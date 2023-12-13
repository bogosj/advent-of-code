import run from "aocrunner";

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

const arrangementChecker = () => {
  const worker = (s: springCondition, c: any): number => {
    // No more groups of broken springs
    if (s.runLengths.length == 0) {
      // Pattern shows there's more
      if (s.pattern.includes('#')) {
        return 0;
      } 
      return  1;
    }

    // More broken springs to find, but pattern exhausted.
    if (s.pattern.length == 0) {
      return 0;
    }

    const nextChar = s.pattern[0];
    const nextRun = s.runLengths[0];

    const isDot = (): number => {
      return c(
        {
          pattern: s.pattern.slice(1),
          runLengths: [...s.runLengths]
        },
        c
      )
    };

    const isHash = (): number => {
      // If the first is a hash, the next N have to be a hash
      const nextN = s.pattern.slice(0, nextRun).replace(/\?/g, '#');
      if (nextN != '#'.repeat(nextRun)) {
        return 0;
      } 

      if (s.pattern.length==nextRun) {
        if (s.runLengths.length==1) {
          return 1;
        }  
        return 0;
      }

      if ('?.'.includes(s.pattern[nextRun])) {
        return c(
          {
            pattern: s.pattern.slice(nextRun+1),
            runLengths: s.runLengths.slice(1)
          },
          c
        )
      }

      return 0; 
    };

    if (nextChar == '.') {
      return isDot();
    }
    if (nextChar == '#') {
      return isHash();
    }
    if (nextChar == '?') {
      return isDot() + isHash();
    }
  }

  const cache = {};
  return (s: springCondition, c: any): number => {
    const k = `${s.pattern}|${s.runLengths}`; 
    if (cache[k]) {
      return cache[k];
    }
    cache[k] = worker(s, c);
    return cache[k];
  }
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  let arrangements = 0;
  for (let i = 0; i < input.length; i++) {
    const checker = arrangementChecker();
    arrangements += checker(input[i], checker);
  }

  return arrangements;
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
