#!/usr/bin/env ts-node

/**
 * Test runner utility for running specific test suites
 */

import { execSync } from 'child_process';

const testSuites = {
  unit: [
    'src/__tests__/services/',
    'src/__tests__/utils/',
    'src/__tests__/components/'
  ],
  integration: [
    'src/__tests__/integration/'
  ],
  all: [
    'src/__tests__/'
  ]
};

function runTests(suites: string[], options: { coverage?: boolean; watch?: boolean } = {}) {
  const jestArgs: string[] = [];
  
  // Add test paths
  jestArgs.push(...suites);
  
  // Add coverage if requested
  if (options.coverage) {
    jestArgs.push('--coverage');
  }
  
  // Add watch mode if requested  
  if (options.watch) {
    jestArgs.push('--watch');
  }
  
  // Add other useful flags
  jestArgs.push('--verbose');
  jestArgs.push('--passWithNoTests');
  
  const command = `npx jest ${jestArgs.join(' ')}`;
  console.log(`Running: ${command}`);
  
  try {
    execSync(command, { stdio: 'inherit' });
  } catch (error) {
    console.error('Tests failed');
    process.exit(1);
  }
}

// CLI interface
function main() {
  const args = process.argv.slice(2);
  const suiteType = args[0] || 'all';
  
  const options = {
    coverage: args.includes('--coverage'),
    watch: args.includes('--watch')
  };
  
  if (!testSuites[suiteType as keyof typeof testSuites]) {
    console.error(`Unknown test suite: ${suiteType}`);
    console.error(`Available suites: ${Object.keys(testSuites).join(', ')}`);
    process.exit(1);
  }
  
  const suites = testSuites[suiteType as keyof typeof testSuites];
  runTests(suites, options);
}

// Run if called directly
if (require.main === module) {
  main();
}

export { runTests, testSuites };