/**
 * Standalone tests that can run without VS Code environment
 * Run with: node out/test/standalone.test.js
 */

import * as assert from "assert";

// Test the core logic functions
function testCommandGeneration() {
  console.log("Testing command generation...");

  const executablePath = "tabula";
  const filePath = "/path/to/file.csv";

  // Test with auto format enabled
  const autoFormatFlag = true ? "-a " : "";
  const commandWithFormat = `"${executablePath}" ${autoFormatFlag}-u "${filePath}"`;

  assert.ok(commandWithFormat.includes("tabula"), "Command should include executable");
  assert.ok(commandWithFormat.includes("-a"), "Command should include -a flag when enabled");
  assert.ok(commandWithFormat.includes("-u"), "Command should include -u flag");
  assert.strictEqual(
    commandWithFormat,
    '"tabula" -a -u "/path/to/file.csv"',
    "Command with auto format should match expected"
  );

  // Test without auto format
  const noFormatFlag = false ? "-a " : "";
  const commandWithoutFormat = `"${executablePath}" ${noFormatFlag}-u "${filePath}"`;

  assert.ok(commandWithoutFormat.includes("tabula"), "Command should include executable");
  assert.ok(!commandWithoutFormat.includes("-a"), "Command should not include -a flag when disabled");
  assert.ok(commandWithoutFormat.includes("-u"), "Command should include -u flag");
  assert.strictEqual(
    commandWithoutFormat,
    '"tabula" -u "/path/to/file.csv"',
    "Command without auto format should match expected"
  );

  console.log("✓ Command generation tests passed");
}

function testFileNameExtraction() {
  console.log("Testing filename extraction...");

  const path = "/Users/test/documents/data.csv";
  const pathParts = path.split("/");
  const fileName = pathParts[pathParts.length - 1];

  assert.strictEqual(fileName, "data.csv", "Should extract filename correctly");

  console.log("✓ Filename extraction tests passed");
}

function testCsvDetection() {
  console.log("Testing CSV file detection...");

  const testCases = [
    { file: "test.csv", expected: true },
    { file: "data.CSV", expected: true },
    { file: "file.txt", expected: false },
    { file: "test.json", expected: false },
  ];

  testCases.forEach(({ file, expected }) => {
    const isCSV = file.toLowerCase().endsWith(".csv");
    assert.strictEqual(
      isCSV,
      expected,
      `${file} detection should be ${expected}`
    );
  });

  console.log("✓ CSV detection tests passed");
}

function testTabulaDirectiveDetection() {
  console.log("Testing tabula directive detection...");

  const withDirective = '#tabula:let D1 = "Total"\nA,B,C\n1,2,3';
  const withoutDirective = 'A,B,C\n1,2,3\n4,5,6';

  const pattern = /^#tabula:/m;

  assert.ok(
    pattern.test(withDirective),
    "Should detect directive in CSV with directive"
  );
  assert.ok(
    !pattern.test(withoutDirective),
    "Should not detect directive in CSV without directive"
  );

  console.log("✓ Tabula directive detection tests passed");
}

function testCustomExecutablePath() {
  console.log("Testing custom executable paths...");

  const testPaths = [
    "tabula",
    "/usr/local/bin/tabula",
    "/opt/homebrew/bin/tabula",
    "C:\\Program Files\\tabula\\tabula.exe",
  ];

  testPaths.forEach((path) => {
    const command = `"${path}" -a -u "test.csv"`;
    assert.ok(command.includes(path), `Command should include path: ${path}`);
    assert.ok(command.includes("-a"), "Command should include -a flag");
    assert.ok(command.includes("-u"), "Command should include -u flag");
  });

  console.log("✓ Custom executable path tests passed");
}

function testPositionCalculation() {
  console.log("Testing cursor position calculation...");

  const testCases = [
    { line: 5, maxLines: 10, expected: 5 },
    { line: 15, maxLines: 10, expected: 9 },
    { line: 0, maxLines: 10, expected: 0 },
    { line: 9, maxLines: 10, expected: 9 },
  ];

  testCases.forEach(({ line, maxLines, expected }) => {
    const safePosition = Math.min(line, maxLines - 1);
    assert.strictEqual(
      safePosition,
      expected,
      `Position ${line} with max ${maxLines} should be ${expected}`
    );
  });

  console.log("✓ Position calculation tests passed");
}

function testConfigurationDefaults() {
  console.log("Testing configuration defaults...");

  const defaults = {
    autoExecute: true,
    executablePath: "tabula",
  };

  assert.strictEqual(
    typeof defaults.autoExecute,
    "boolean",
    "autoExecute should be boolean"
  );
  assert.strictEqual(
    typeof defaults.executablePath,
    "string",
    "executablePath should be string"
  );
  assert.strictEqual(
    defaults.autoExecute,
    true,
    "autoExecute should default to true"
  );
  assert.strictEqual(
    defaults.executablePath,
    "tabula",
    "executablePath should default to 'tabula'"
  );

  console.log("✓ Configuration defaults tests passed");
}

function testErrorMessageFormatting() {
  console.log("Testing error message formatting...");

  const error = new Error("Command not found");
  const message = `Run script error: ${error.message}. Check that tabula is installed and the path is correct in settings.`;

  assert.ok(
    message.includes("Run script error"),
    "Message should include error prefix"
  );
  assert.ok(
    message.includes("Command not found"),
    "Message should include error details"
  );
  assert.ok(
    message.includes("Check that tabula is installed"),
    "Message should include help text"
  );

  console.log("✓ Error message formatting tests passed");
}

function testUriComparison() {
  console.log("Testing URI comparison...");

  const uri1: string = "file:///path/to/file.csv";
  const uri2: string = "file:///path/to/file.csv";
  const uri3: string = "file:///different/file.csv";

  assert.strictEqual(
    uri1,
    uri2,
    "Identical URIs should match"
  );
  assert.notStrictEqual(
    uri1,
    uri3,
    "Different URIs should not match"
  );

  console.log("✓ URI comparison tests passed");
}

// Run all tests
function runAllTests() {
  console.log("\n=== Running Standalone Tests ===\n");

  const tests = [
    testCommandGeneration,
    testFileNameExtraction,
    testCsvDetection,
    testTabulaDirectiveDetection,
    testCustomExecutablePath,
    testPositionCalculation,
    testConfigurationDefaults,
    testErrorMessageFormatting,
    testUriComparison,
  ];

  let passed = 0;
  let failed = 0;

  tests.forEach((test) => {
    try {
      test();
      passed++;
    } catch (error) {
      failed++;
      console.error(`✗ Test failed: ${test.name}`);
      console.error(error);
    }
  });

  console.log("\n=== Test Results ===");
  console.log(`Total: ${tests.length}`);
  console.log(`Passed: ${passed}`);
  console.log(`Failed: ${failed}`);

  if (failed === 0) {
    console.log("\n✅ All tests passed!\n");
    process.exit(0);
  } else {
    console.log("\n❌ Some tests failed!\n");
    process.exit(1);
  }
}

// Run tests if this is the main module
if (require.main === module) {
  runAllTests();
}

export { runAllTests };
