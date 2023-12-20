import run from "aocrunner";

const sum = (arr: number[]) => arr.reduce((a, b) => a + b);

interface Part {
  x: number,
  m: number,
  a: number,
  s: number
}

const parseInput = (rawInput: string): [any, Part[]] => {
  const rules = {};
  const parts = [];

  let readingRules = true;
  rawInput.split('\n').forEach(line => {
    if (line == '') {
      readingRules = false;
      return;
    }
    if (readingRules) {
      const name = line.split('{')[0];
      const paths = line.split('{')[1].split('}')[0];
      rules[name] = paths.split(',');
    } else {
      const vals = line.replace('{', '').replace('}', '').split(',')
      parts.push(Object.fromEntries(vals.map(value => {
        return [value.split('=')[0], parseInt(value.split('=')[1], 10)]
      })));
    }
  });
  return [rules, parts];
};

const valuePart = (p: Part): number => {
  return Object.values(p).reduce((a, b) => a + b);
}

const partIsValid = (rules, p: Part): boolean => {
  let currentRule = rules['in'];
  while (true) {
    for (let i = 0; i < currentRule.length; i++) {
      let [test, nextRule] = currentRule[i].split(':');
      if (nextRule == undefined) {
        nextRule = test;
        test = 'x==p.x';
      }
      if (eval(`p.${test}`)) {
        if ('AR'.includes(nextRule)) {
          return nextRule == 'A';
        }
        currentRule = rules[nextRule];
        break;
      }
    }
  }
};

const part1 = (rawInput: string) => {
  const [rules, parts] = parseInput(rawInput);
  const partFilter = (p: Part) => {
    return partIsValid(rules, p);
  };
  return parts.filter(partFilter).map(valuePart).reduce((a, b) => a + b);
};

const valueOfFilters = (filters: string[]): number => {
  const ranges = {
    x: [1, 4000],
    m: [1, 4000],
    a: [1, 4000],
    s: [1, 4000],
  }
  filters.forEach(filter => {
    const letter = filter.substring(0, 1);
    const gt = filter.substring(1, 2) == '>';
    const val = parseInt(filter.substring(2), 10);
    if (gt) {
      if (val > ranges[letter][0]) {
        ranges[letter][0] = val + 1;
      }
    } else {
      if (val < ranges[letter][1]) {
        ranges[letter][1] = val - 1;
      }
    }
  });

  return (
    (ranges.x[1] - ranges.x[0] + 1) *
    (ranges.m[1] - ranges.m[0] + 1) *
    (ranges.a[1] - ranges.a[0] + 1) *
    (ranges.s[1] - ranges.s[0] + 1)
  )
};

const invertRule = (rule: string): string => {
  const gt = rule.charAt(1) == '>';
  const comp = gt ? '<' : '>';
  let val = parseInt(rule.substring(2), 10);
  val = gt ? val + 1 : val - 1;
  return `${rule.charAt(0)}${comp}${val}`;
};

const allValidRules = (rules: string[], rulesToApply: string[], allRules): string[][] => {
  const ruleToEval = rules[0];
  const futureRules = rules.slice(1);
  if (ruleToEval.includes(':')) {
    const [rule, nextRuleName] = ruleToEval.split(':');
    return [  
      allValidRules([...allRules[nextRuleName]], [...rulesToApply, rule], allRules),
      allValidRules(futureRules, [...rulesToApply, invertRule(rule)], allRules)
    ].flat(); 
  } else { // default rule
    if (ruleToEval == 'R') {
      return [];
    }
    if (ruleToEval == 'A') {
      return [
        [...rulesToApply]
      ];
    } 
    return [
      allValidRules([...allRules[ruleToEval]], [...rulesToApply], allRules),
    ].flat();
  }
};

const part2 = (rawInput: string) => {
  const [rules] = parseInput(rawInput);
  rules['A'] = ['A']
  rules['R'] = ['R']
  const allRules = allValidRules(rules['in'], [], rules);
  return sum(allRules.map(valueOfFilters));
};

run({
  part1: {
    tests: [
      {
        input: `
        px{a<2006:qkq,m>2090:A,rfg}
        pv{a>1716:R,A}
        lnx{m>1548:A,A}
        rfg{s<537:gd,x>2440:R,A}
        qs{s>3448:A,lnx}
        qkq{x<1416:A,crn}
        crn{x>2662:A,R}
        in{s<1351:px,qqz}
        qqz{s>2770:qs,m<1801:hdj,R}
        gd{a>3333:R,R}
        hdj{m>838:A,pv}
        
        {x=787,m=2655,a=1222,s=2876}
        {x=1679,m=44,a=2067,s=496}
        {x=2036,m=264,a=79,s=2244}
        {x=2461,m=1339,a=466,s=291}
        {x=2127,m=1623,a=2188,s=1013}
        `,
        expected: 19114,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `
        px{a<2006:qkq,m>2090:A,rfg}
        pv{a>1716:R,A}
        lnx{m>1548:A,A}
        rfg{s<537:gd,x>2440:R,A}
        qs{s>3448:A,lnx}
        qkq{x<1416:A,crn}
        crn{x>2662:A,R}
        in{s<1351:px,qqz}
        qqz{s>2770:qs,m<1801:hdj,R}
        gd{a>3333:R,R}
        hdj{m>838:A,pv}
        
        {x=787,m=2655,a=1222,s=2876}
        {x=1679,m=44,a=2067,s=496}
        {x=2036,m=264,a=79,s=2244}
        {x=2461,m=1339,a=466,s=291}
        {x=2127,m=1623,a=2188,s=1013}
        `,
        expected: 167409079868000,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
