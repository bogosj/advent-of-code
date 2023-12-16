import run from "aocrunner";

const parseInput = (rawInput: string): Array<string> => {
  return rawInput.split(',');
};

const hashString = (s: string): number => {
  let current = 0;
  s.split('').forEach(c => {
    current += c.charCodeAt(0);
    current *= 17;
    current %= 256;
  })
  return current;
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  return input.map(hashString).reduce((a, b) => a + b);
};

interface lense {
  label: string,
  focalLength: number
}

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const boxes: { [key: number]: Array<lense> } = {};
  for (let i = 0; i < 256; i++) {
    boxes[i] = [];
  }
  input.forEach(v => {
    if (v.includes('=')) {
      const [label, focalLength] = v.split('=');
      const hash = hashString(label);
      const lense = { label: label, focalLength: parseInt(focalLength, 10) }
      let replaced = false;
      boxes[hash].map(l => {
        if (l.label == lense.label) {
          l.focalLength = lense.focalLength;
          replaced = true;
          return lense;
        }
        return l;
      });
      if (!replaced) {
        boxes[hash].push(lense);
      }
    } else {
      const [label, focalLength] = v.split('-');
      const hash = hashString(label);
      const lense = { label: label, focalLength: parseInt(focalLength, 10) };
      boxes[hash] = boxes[hash].filter(l => {
        if (l.label != lense.label) {
          return l;
        }
      });
    }
  });

  return Object.entries(boxes).map(([k, v]) => {
    const mag = parseInt(k, 10) + 1;
    if (v.length == 0) {
      return 0;
    }
    return v.map((lense, idx) => {
      return mag * (idx + 1) * lense.focalLength;
    }).reduce((a, b) => a + b);
  }).reduce((a, b) => a + b);
};

run({
  part1: {
    tests: [
      {
        input: `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`,
        expected: 1320,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`,
        expected: 145,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
