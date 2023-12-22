import run from "aocrunner";

enum ModuleType {
  FlipFlop = '%',
  Conjunction = '&',
  Boadcaster = 'b',
  Untyped = 'U'
}

enum ModuleState {
  Off = 0,
  On
}

enum PulseType {
  Low = 0,
  High
}

interface Module {
  type: ModuleType,
  state: ModuleState,
  memory: Record<string, PulseType>,
  targets: string[]
}

const parseInput = (rawInput: string): Record<string, Module> => {
  const allModules = {};
  rawInput.split('\n').forEach(line => {
    let modName = line.split(' ')[0];
    let module = {
      type: modName.substring(0, 1),
      state: ModuleState.Off,
      memory: {},
      targets: line.split(' ').slice(2).join('').split(',')
    };
    if (module.type == ModuleType.FlipFlop || module.type == ModuleType.Conjunction) {
      modName = modName.substring(1);
    }
    allModules[modName] = module;
  });
  Object.keys(allModules).forEach(moduleName => {
    const module: Module = allModules[moduleName];
    module.targets.forEach(targetName => {
      let targetModule: Module = allModules[targetName];
      if (targetModule == undefined) {
        targetModule = {
          type: ModuleType.Untyped,
          state: ModuleState.Off,
          memory: {},
          targets: []
        };
        allModules[targetName] = targetModule;
      }
      targetModule.memory[moduleName] = PulseType.Low;
    })
  });
  return allModules;
}

interface Pulse {
  source: string,
  target: string,
  type: PulseType
}

const part1 = (rawInput: string) => {
  const allModules = parseInput(rawInput);
  const pulseCount = [0, 0];
  for (let i = 0; i < 1000; i++) {
    const pulsesToSend: Pulse[] = [{ source: 'button', target: 'broadcaster', type: PulseType.Low }];
    while (pulsesToSend.length > 0) {
      const nextPulse = pulsesToSend.shift();
      pulseCount[nextPulse.type] += 1;
      const targetModule = allModules[nextPulse.target];
      if (targetModule.type == ModuleType.Boadcaster) {
        targetModule.targets.forEach(targetModuleName => {
          pulsesToSend.push({
            source: nextPulse.target,
            target: targetModuleName,
            type: nextPulse.type
          });
        });
      }
      if (targetModule.type == ModuleType.FlipFlop) {
        let pulseToSend: PulseType = undefined;
        if (nextPulse.type == PulseType.High) {
          continue;
        }
        if (targetModule.state == ModuleState.Off) {
          targetModule.state = ModuleState.On;
          pulseToSend = PulseType.High
        } else {
          targetModule.state = ModuleState.Off;
          pulseToSend = PulseType.Low;
        }
        targetModule.targets.forEach(targetModuleName => {
          pulsesToSend.push({
            source: nextPulse.target,
            target: targetModuleName,
            type: pulseToSend
          });
        });
      }
      if (targetModule.type == ModuleType.Conjunction) {
        // Update memory
        targetModule.memory[nextPulse.source] = nextPulse.type;

        // If all memory is high send a low, otherwise send a high
        if (Object.values(targetModule.memory).every(pt => pt == PulseType.High)) {
          targetModule.targets.forEach(targetModuleName => {
            pulsesToSend.push({
              source: nextPulse.target,
              target: targetModuleName,
              type: PulseType.Low
            });
          });
        } else {
          targetModule.targets.forEach(targetModuleName => {
            pulsesToSend.push({
              source: nextPulse.target,
              target: targetModuleName,
              type: PulseType.High
            });
          });
        }
      }
    }
  }
  return pulseCount[0] * pulseCount[1];
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
        broadcaster -> a, b, c
        %a -> b
        %b -> c
        %c -> inv
        &inv -> a
        `,
        expected: 32000000,
      },
      {
        input: `
        broadcaster -> a
        %a -> inv, con
        &inv -> b
        %b -> con
        &con -> output
        `,
        expected: 11687500,
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
