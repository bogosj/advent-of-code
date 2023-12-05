import run from "aocrunner";

const mapOrder = [
  'seed-to-soil',
  'soil-to-fertilizer',
  'fertilizer-to-water',
  'water-to-light',
  'light-to-temperature',
  'temperature-to-humidity',
  'humidity-to-location'
];

const reverseMapOrder = [...mapOrder].reverse();

interface mapping {
  source: number,
  destination: number,
  rangeLength: number,
}

const parseInput = (rawInput: string) => {
  const lines = rawInput.split('\n');
  const retVal = {};
  retVal['seeds'] = lines[0].split(' ').slice(1).map(x => parseInt(x, 10));

  let key = '';
  lines.slice(2).forEach(line => {
    if (line.endsWith(':')) {
      key = line.split(' ')[0];
      retVal[key] = [];
      return;
    }
    if (line == '') {
      return;
    }
    retVal[key].push({
      'destination': parseInt(line.split(' ')[0], 10),
      'source': parseInt(line.split(' ')[1], 10),
      'rangeLength': parseInt(line.split(' ')[2], 10),
    });
  })

  return retVal;
};

const seedToLocation = (seed: number, allMaps: Object): number => {
  let location = seed;
  for (let i = 0; i < mapOrder.length; i++) {
    const maps: Array<mapping> = Object.values(allMaps[mapOrder[i]]);
    for (let j = 0; j < maps.length; j++) {
      if (location >= maps[j].source && location <= maps[j].source + maps[j].rangeLength) {
        location = location - (maps[j].source - maps[j].destination);
        break;
      }
    }
  }
  return location;
};

const locationToSeed = (loc: number, allMaps: Object): number => {
  for (let i = 0; i < reverseMapOrder.length; i++) {
    const maps: Array<mapping> = Object.values(allMaps[reverseMapOrder[i]]);
    for (let j=0; j<maps.length; j++) {
      if (maps[j].destination <= loc && maps[j].destination + maps[j].rangeLength > loc) {
        loc = maps[j].source + loc - maps[j].destination;
        break;
      }
    }
  } 
  return loc;
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  const seedLocations = input['seeds'].map(seed => {
    return seedToLocation(seed, input);
  });

  return Math.min(...seedLocations);
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  for (let i = 0; ; i++) {
    const loc = locationToSeed(i, input);
    const s = input['seeds'];
    for (let j = 0; j < s.length; j += 2) {
      if (loc >= s[j] && loc < s[j] + s[j + 1]) {
        return i;
      }
    }
  }
};

run({
  part1: {
    tests: [
      {
        input: `
        seeds: 79 14 55 13

        seed-to-soil map:
        50 98 2
        52 50 48

        soil-to-fertilizer map:
        0 15 37
        37 52 2
        39 0 15

        fertilizer-to-water map:
        49 53 8
        0 11 42
        42 0 7
        57 7 4

        water-to-light map:
        88 18 7
        18 25 70

        light-to-temperature map:
        45 77 23
        81 45 19
        68 64 13

        temperature-to-humidity map:
        0 69 1
        1 0 69

        humidity-to-location map:
        60 56 37
        56 93 4
        `,
        expected: 35,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `
        seeds: 79 14 55 13

        seed-to-soil map:
        50 98 2
        52 50 48

        soil-to-fertilizer map:
        0 15 37
        37 52 2
        39 0 15

        fertilizer-to-water map:
        49 53 8
        0 11 42
        42 0 7
        57 7 4

        water-to-light map:
        88 18 7
        18 25 70

        light-to-temperature map:
        45 77 23
        81 45 19
        68 64 13

        temperature-to-humidity map:
        0 69 1
        1 0 69

        humidity-to-location map:
        60 56 37
        56 93 4
        `,
        expected: 46,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
