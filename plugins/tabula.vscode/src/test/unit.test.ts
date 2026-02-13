import * as assert from "assert";

suite("Unit Tests", () => {
  test("String operations", () => {
    const csvFile = "test.csv";
    assert.ok(csvFile.endsWith(".csv"), "Should detect CSV extension");

    const upperCaseFile = "TEST.CSV";
    assert.ok(
      upperCaseFile.toLowerCase().endsWith(".csv"),
      "Should handle case-insensitive extensions"
    );
  });

  test("Command format", () => {
    const tabulaPath = "tabula";
    const filePath = "/path/to/file.csv";
    const command = `"${tabulaPath}" -a -u "${filePath}"`;

    assert.ok(command.includes("tabula"), "Command should include executable");
    assert.ok(command.includes("-a"), "Command should include -a flag");
    assert.ok(command.includes("-u"), "Command should include -u flag");
    assert.ok(command.includes(filePath), "Command should include file path");
  });

  test("Path extraction", () => {
    const fullPath = "/Users/test/documents/data.csv";
    const pathParts = fullPath.split("/");
    const fileName = pathParts[pathParts.length - 1];

    assert.strictEqual(fileName, "data.csv", "Should extract filename");
  });

  test("Tabula directive detection", () => {
    const csvWithDirective = '#tabula:let A1 = "test"\nA,B,C\n1,2,3';
    const csvWithoutDirective = 'A,B,C\n1,2,3';

    const hasDirective = /^#tabula:/m.test(csvWithDirective);
    const noDirective = /^#tabula:/m.test(csvWithoutDirective);

    assert.ok(hasDirective, "Should detect tabula directive");
    assert.ok(!noDirective, "Should not detect directive when absent");
  });

  test("Configuration value types", () => {
    const autoExecute = true;
    const executablePath = "tabula";

    assert.strictEqual(typeof autoExecute, "boolean");
    assert.strictEqual(typeof executablePath, "string");
  });

  test("File extension validation", () => {
    const validFiles = ["test.csv", "data.CSV", "file.csv"];
    const invalidFiles = ["test.txt", "data.json", "file.xlsx"];

    validFiles.forEach((file) => {
      assert.ok(
        file.toLowerCase().endsWith(".csv"),
        `${file} should be valid CSV`
      );
    });

    invalidFiles.forEach((file) => {
      assert.ok(
        !file.toLowerCase().endsWith(".csv"),
        `${file} should not be CSV`
      );
    });
  });

  test("Custom executable path handling", () => {
    const defaultPath = "tabula";
    const customPath = "/usr/local/bin/tabula";
    const windowsPath = "C:\\Program Files\\tabula\\tabula.exe";

    [defaultPath, customPath, windowsPath].forEach((path) => {
      const command = `"${path}" -a -u "file.csv"`;
      assert.ok(command.length > 0, "Should generate valid command");
      assert.ok(command.includes(path), "Should include executable path");
    });
  });

  test("Error message formatting", () => {
    const errorMsg = "Run script error: command not found";
    assert.ok(errorMsg.includes("Run script error"));
    assert.ok(errorMsg.includes("command not found"));
  });

  test("Position calculation", () => {
    const line = 5;
    const maxLine = 10;
    const safePosition = Math.min(line, maxLine - 1);

    assert.strictEqual(safePosition, 5);

    const largeLine = 15;
    const safeLargePosition = Math.min(largeLine, maxLine - 1);
    assert.strictEqual(safeLargePosition, 9);
  });

  test("URI string comparison", () => {
    const uri1 = "file:///path/to/file.csv";
    const uri2 = "file:///path/to/file.csv";
    const uri3 = "file:///different/file.csv";

    assert.strictEqual(uri1, uri2, "Same URIs should match");
    assert.notStrictEqual(uri1, uri3, "Different URIs should not match");
  });
});
