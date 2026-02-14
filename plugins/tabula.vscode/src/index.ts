import * as vscode from "vscode";
import { exec } from "child_process";

export function activate(context: vscode.ExtensionContext) {
  // Auto-execute Tabula on CSV file save
  const autoExecuteDisposable = vscode.workspace.onDidSaveTextDocument(
    async (document: vscode.TextDocument) => {
      // Check configuration
      const config = vscode.workspace.getConfiguration("tabula");
      const autoExecute = config.get<boolean>("autoExecute", true);

      if (!autoExecute) {
        return;
      }

      // Check if it's a CSV file and if it's the active document
      const editor = vscode.window.activeTextEditor;
      if (
        editor &&
        editor.document.uri.toString() === document.uri.toString() &&
        document.languageId === "csv"
      ) {
        // Execute the tabula.execute command
        await vscode.commands.executeCommand("tabula.execute");
      }
    },
  );

  context.subscriptions.push(autoExecuteDisposable);

  // Manual command to toggle auto-execution
  const toggleCommand = vscode.commands.registerCommand(
    "tabula.toggleAutoExecute",
    () => {
      const config = vscode.workspace.getConfiguration("tabula");
      const autoExecute = !config.get<boolean>("autoExecute", true);
      config.update(
        "autoExecute",
        autoExecute,
        vscode.ConfigurationTarget.Global,
      );
      vscode.window.showInformationMessage(
        `Tabula auto-execute ${autoExecute ? "enabled" : "disabled"}`,
      );
    },
  );

  context.subscriptions.push(toggleCommand);

  const executeCommand = vscode.commands.registerCommand(
    "tabula.execute",
    async () => {
      const editor = vscode.window.activeTextEditor;
      if (!editor) {
        vscode.window.showErrorMessage("No active editor found");
        return;
      }

      const document = editor.document;

      // Check if it's a CSV file
      // if (document.languageId !== "csv" && !isCsvFile(document.fileName)) {
      if (document.languageId !== "csv") {
        vscode.window.showErrorMessage(
          "Current file is not a CSV file. Tabula can only be executed on CSV files.",
        );
        return;
      }

      // Save the document if it has unsaved changes
      if (document.isDirty) {
        await document.save();
      }

      try {
        // Save cursor position
        const cursorPosition = editor.selection.active;

        await runScript(document.uri);

        // Reload the document from disk to show changes
        await reloadDocument(document.uri);

        // Restore cursor position
        const newPosition = new vscode.Position(
          Math.min(cursorPosition.line, editor.document.lineCount - 1),
          cursorPosition.character,
        );
        editor.selection = new vscode.Selection(newPosition, newPosition);
        editor.revealRange(
          new vscode.Range(newPosition, newPosition),
          vscode.TextEditorRevealType.InCenterIfOutsideViewport,
        );
      } catch (error) {
        vscode.window.showErrorMessage(
          `Tabula execution failed: ${error instanceof Error ? error.message : String(error)}`,
        );
      }
    },
  );

  context.subscriptions.push(executeCommand);
}

// This method is called when your extension is deactivated
export function deactivate() {}

const runScript = (path: vscode.Uri): Promise<void> =>
  new Promise((resolve, reject) => {
    const pathParts = path.path.split("/");
    const fileName = pathParts[pathParts.length - 1];

    // Get configuration settings
    const config = vscode.workspace.getConfiguration("tabula");
    const tabulaPath = config.get<string>("executablePath", "tabula");
    const autoFormat = config.get<boolean>("autoFormat", true);

    // Build command with optional -a flag
    const autoFormatFlag = autoFormat ? "-a" : "";
    const command = `"${tabulaPath}" ${autoFormatFlag} -u "${path.path}"`;

    exec(command, (error, _stdout, stderr) => {
      if (error) {
        vscode.window.showErrorMessage(
          `Run script error: ${error.message}. Check that tabula is installed and the path is correct in settings.`,
        );
        console.error(`exec error: ${error}`);
        return reject(error);
      }
      if (stderr) {
        console.error(`stderr: ${stderr}`);
      }
      vscode.window.showInformationMessage(`Tabula updated ${fileName}.`);
      resolve();
    });
  });

const reloadDocument = async (uri: vscode.Uri): Promise<void> => {
  // Find all text editors showing this document
  const editors = vscode.window.visibleTextEditors.filter(
    (editor) => editor.document.uri.toString() === uri.toString(),
  );

  if (editors.length === 0) {
    return;
  }

  // Close the document
  await vscode.commands.executeCommand("workbench.action.closeActiveEditor");

  // Small delay to ensure file system changes are visible
  await new Promise((resolve) => setTimeout(resolve, 100));

  // Reopen the document
  const document = await vscode.workspace.openTextDocument(uri);
  await vscode.window.showTextDocument(document, {
    preview: false,
    preserveFocus: false,
  });
};
