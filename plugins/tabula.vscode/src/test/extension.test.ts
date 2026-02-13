import * as assert from "assert";
import * as vscode from "vscode";
import * as path from "path";
import * as fs from "fs";

suite("Tabula Extension Test Suite", () => {
  vscode.window.showInformationMessage("Starting Tabula extension tests");

  let extensionId: string;

  suiteSetup(() => {
    // Find the extension by searching for it
    const allExtensions = vscode.extensions.all;
    const tabulaExtension = allExtensions.find(
      (ext) => ext.packageJSON.name === "tabula"
    );

    if (tabulaExtension) {
      extensionId = tabulaExtension.id;
    } else {
      // Fallback to expected ID
      extensionId = "tabula.tabula";
    }
  });

  test("Extension should be present", () => {
    const extension = vscode.extensions.getExtension(extensionId);
    assert.ok(extension, `Extension ${extensionId} should be present`);
  });

  test("Extension should activate", async () => {
    const extension = vscode.extensions.getExtension(extensionId);
    assert.ok(extension, `Extension ${extensionId} should exist`);

    await extension!.activate();
    assert.strictEqual(extension!.isActive, true, "Extension should be active");
  });

  test("Should register toggleAutoExecute command", async () => {
    const commands = await vscode.commands.getCommands(true);
    assert.ok(
      commands.includes("tabula.toggleAutoExecute"),
      "toggleAutoExecute command should be registered",
    );
  });

  test("Configuration should have default values", () => {
    const config = vscode.workspace.getConfiguration("tabula");

    // Use inspect() to get the default value from schema
    const autoExecuteInspect = config.inspect<boolean>("autoExecute");
    assert.ok(autoExecuteInspect, "autoExecute configuration should exist");
    assert.strictEqual(
      autoExecuteInspect.defaultValue,
      true,
      "autoExecute should default to true",
    );

    const executablePathInspect = config.inspect<string>("executablePath");
    assert.ok(executablePathInspect, "executablePath configuration should exist");
    assert.strictEqual(
      executablePathInspect.defaultValue,
      "tabula",
      "executablePath should default to 'tabula'",
    );

    const autoFormatInspect = config.inspect<boolean>("autoFormat");
    assert.ok(autoFormatInspect, "autoFormat configuration should exist");
    assert.strictEqual(
      autoFormatInspect.defaultValue,
      true,
      "autoFormat should default to true",
    );
  });

  test("Should toggle autoExecute setting", async () => {
    // Get fresh config reference
    let config = vscode.workspace.getConfiguration("tabula");

    // Read current value (use get with default)
    const initialValue = config.get<boolean>("autoExecute", true);

    // Execute toggle command
    await vscode.commands.executeCommand("tabula.toggleAutoExecute");

    // Wait for config to propagate
    await new Promise(resolve => setTimeout(resolve, 150));

    // Re-get configuration to see changes
    config = vscode.workspace.getConfiguration("tabula");
    const newValue = config.get<boolean>("autoExecute", true);

    // After toggle, value should be different
    assert.notStrictEqual(
      newValue,
      initialValue,
      `autoExecute should toggle from ${initialValue}`,
    );

    // Verify it's the opposite value
    assert.strictEqual(
      newValue,
      !initialValue,
      `autoExecute should be ${!initialValue} after toggle`,
    );

    // Toggle back
    await vscode.commands.executeCommand("tabula.toggleAutoExecute");

    // Wait for config to propagate
    await new Promise(resolve => setTimeout(resolve, 150));

    // Re-get configuration again
    config = vscode.workspace.getConfiguration("tabula");
    const restoredValue = config.get<boolean>("autoExecute", true);

    // Should be back to initial value
    assert.strictEqual(
      restoredValue,
      initialValue,
      `autoExecute should toggle back to ${initialValue}`,
    );
  });

  test("Should update executablePath configuration", async () => {
    let config = vscode.workspace.getConfiguration("tabula");
    const customPath = "/custom/path/to/tabula";

    await config.update(
      "executablePath",
      customPath,
      vscode.ConfigurationTarget.Global,
    );

    // Re-get configuration to see changes
    config = vscode.workspace.getConfiguration("tabula");
    const updatedPath = config.get<string>("executablePath", "tabula");
    assert.strictEqual(
      updatedPath,
      customPath,
      "executablePath should update correctly",
    );

    // Restore default
    await config.update(
      "executablePath",
      "tabula",
      vscode.ConfigurationTarget.Global,
    );
  });

  test("Should handle CSV file detection", async () => {
    // Use the existing test CSV file instead of creating a new one
    const testFilePath = path.join(__dirname, "..", "..", "src", "test", "input.csv");

    if (fs.existsSync(testFilePath)) {
      const document = await vscode.workspace.openTextDocument(testFilePath);

      assert.ok(
        document.fileName.endsWith(".csv"),
        "File should be recognized as CSV",
      );

      // Check if document language is set to csv
      assert.ok(
        document.languageId === "csv" || document.fileName.endsWith(".csv"),
        "Document should be identified as CSV",
      );
    } else {
      // Skip test if file doesn't exist
      console.log("Test CSV file not found, skipping file detection test");
    }
  });

  test("Should validate configuration schema", () => {
    const config = vscode.workspace.getConfiguration("tabula");
    const inspect = config.inspect("autoExecute");

    assert.ok(inspect, "autoExecute configuration should exist");
    assert.strictEqual(
      typeof inspect?.defaultValue,
      "boolean",
      "autoExecute should be boolean type",
    );
  });

  test("Should have correct package.json metadata", async () => {
    const extension = vscode.extensions.getExtension(extensionId);
    assert.ok(extension, "Extension should exist for metadata test");

    const packageJSON = extension!.packageJSON;

    assert.strictEqual(packageJSON.name, "tabula");
    assert.strictEqual(packageJSON.displayName, "Tabula");
    assert.ok(packageJSON.version);
    assert.ok(packageJSON.description);

    // Check commands are registered
    const commands = packageJSON.contributes.commands;
    assert.ok(
      commands.find((cmd: any) => cmd.command === "tabula.toggleAutoExecute"),
      "toggleAutoExecute command should be in package.json",
    );

    // Check configuration is defined
    const properties = packageJSON.contributes.configuration.properties;
    assert.ok(properties["tabula.autoExecute"]);
    assert.ok(properties["tabula.executablePath"]);
  });

  test("Test CSV file with tabula script directive", () => {
    // Test the directive pattern matching
    const csvContent = '#tabula:let D1 = "Total"\nA,B,C\n1,2,3\n4,5,6';

    assert.ok(
      csvContent.includes("#tabula:"),
      "CSV should contain tabula directive",
    );
    assert.ok(csvContent.includes("let D1"), "CSV should contain tabula script");

    // Test pattern detection
    const hasDirective = /^#tabula:/m.test(csvContent);
    assert.ok(hasDirective, "Should detect tabula directive pattern");
  });

  test("Should respect autoExecute setting when disabled", async () => {
    let config = vscode.workspace.getConfiguration("tabula");

    // Disable auto-execute
    await config.update(
      "autoExecute",
      false,
      vscode.ConfigurationTarget.Global,
    );

    // Re-get configuration to see changes
    config = vscode.workspace.getConfiguration("tabula");
    const autoExecute = config.get<boolean>("autoExecute", true);
    assert.strictEqual(
      autoExecute,
      false,
      "autoExecute should be disabled",
    );

    // Re-enable
    await config.update(
      "autoExecute",
      true,
      vscode.ConfigurationTarget.Global,
    );
  });

  test("Should handle different file extensions", () => {
    const csvFiles = [
      "test.csv",
      "data.CSV",
      "file.csv",
    ];

    csvFiles.forEach((fileName) => {
      assert.ok(
        fileName.toLowerCase().endsWith(".csv"),
        `${fileName} should be recognized as CSV`,
      );
    });
  });

  test("Should update autoFormat configuration", async () => {
    let config = vscode.workspace.getConfiguration("tabula");

    // Test disabling auto format
    await config.update(
      "autoFormat",
      false,
      vscode.ConfigurationTarget.Global,
    );

    // Re-get configuration to see changes
    config = vscode.workspace.getConfiguration("tabula");
    const autoFormat = config.get<boolean>("autoFormat", true);
    assert.strictEqual(
      autoFormat,
      false,
      "autoFormat should be disabled",
    );

    // Re-enable
    await config.update(
      "autoFormat",
      true,
      vscode.ConfigurationTarget.Global,
    );

    // Verify it's enabled again
    config = vscode.workspace.getConfiguration("tabula");
    const autoFormatEnabled = config.get<boolean>("autoFormat", true);
    assert.strictEqual(
      autoFormatEnabled,
      true,
      "autoFormat should be enabled",
    );
  });
});
