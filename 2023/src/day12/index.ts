import run from "aocrunner";

interface springCondition {
  template: string,
  runLengths: Array<number>,
  validPattern: string
}

const parseInput = (rawInput: string): Array<springCondition> => {
  const retVal = [];
  rawInput.split('\n').forEach(line => {
    const s = {
      template: line.split(' ')[0],
      runLengths: line.split(' ')[1].split(',').map(v => parseInt(v, 10)),
      validPattern: ''
    };

    const validPattern: Array<string> = [];
    while (s.runLengths.length) {
      validPattern.push('#'.repeat(s.runLengths.shift()));
    }
    s.validPattern = validPattern.join(',');
    retVal.push(s);
  });
  return retVal; 
};

const arrangementChecker = () => {
  const cache = {};
  return (s: springCondition, c): number => {
    if (s.template.includes('?')) {
      return (
        c({ template: s.template.replace('?', '#'), runLengths: [...s.runLengths], validPattern: s.validPattern }, c) +
        c({ template: s.template.replace('?', '.'), runLengths: [...s.runLengths], validPattern: s.validPattern }, c)
      )
    }
    const cacheKey = `${s.template}|${s.runLengths}`;
    if (cache[cacheKey]) {
      return cache[cacheKey];
    }  
    
    let pattern = s.template.replace(/^\.+/, '').replace(/\.+$/, '').replace(/\.+/g, ',');
    let count = 0;
    if (pattern == s.validPattern) { 
      count = 1;
    } 
    cache[cacheKey] = count;
    return count;
  }
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  let arrangements = 0;
  for (let i=0; i<input.length; i++) {
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
