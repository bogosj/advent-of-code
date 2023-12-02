import run from "aocrunner";

const parseInput = (rawInput: string) => {
  const retVal = {};
  rawInput.split('\n').map(line => {
    const game = line.split(':');
    const gameId = game[0].split(' ')[1];
    const pulls = game[1].split(';');

    let pullResults = [];
    pulls.forEach(pull => {
      const pullObj = {};
      const dice = pull.split(',');
      dice.forEach(dies => {
        const count = dies.trim().split(' ');
        pullObj[count[1]] = parseInt(count[0], 10);
      })
      pullResults.push(pullObj);
    });
    retVal[gameId] = pullResults;
  });
  return retVal;
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  let result = 0;
  Object.entries(input).forEach(([game, results]) => {
    if (results.every(result => {
      const red = result['red'] || 0;
      const green = result['green'] || 0;
      const blue = result['blue'] || 0;
      return red <= 12 && green <= 13 && blue <= 14;
    })) {
      result += parseInt(game, 10);
    }
  });
  return result; 
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  let result = 0;
  Object.entries(input).forEach(([game, results]) => {
    const reds = [];
    const blues = [];
    const greens = [];
    results.forEach(result => {
      reds.push(result.red || 0);
      blues.push(result.blue || 0);
      greens.push(result.green || 0);
    });
    result += Math.max(...reds) * Math.max(...blues) * Math.max(...greens);
  });
  return result;
};

run({
  part1: {
    tests: [
      {
        input: `
        Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
        Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
        Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
        Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
        Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
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
        Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
        Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
        Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
        Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
        Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
        `,
        expected: 2286,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
