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

      // Check if it's a CSV file
      if (document.languageId === "csv" || document.fileName.endsWith(".csv")) {
        try {
          // Save cursor position
          const editor = vscode.window.activeTextEditor;
          const cursorPosition = editor?.selection.active;

          await runScript(document.uri);

          // Reload the document from disk to show changes
          await reloadDocument(document.uri);

          // Restore cursor position
          if (editor && cursorPosition) {
            const newPosition = new vscode.Position(
              Math.min(cursorPosition.line, editor.document.lineCount - 1),
              cursorPosition.character,
            );
            editor.selection = new vscode.Selection(newPosition, newPosition);
            editor.revealRange(
              new vscode.Range(newPosition, newPosition),
              vscode.TextEditorRevealType.InCenterIfOutsideViewport,
            );
          }
        } catch (error) {
          vscode.window.showErrorMessage(
            `Tabula execution failed: ${error instanceof Error ? error.message : String(error)}`,
          );
        }
      }
    },
  );

  context.subscriptions.push(autoExecuteDisposable);

  // Manual command to toggle auto-execution
  const toggleCommand = vscode.commands.registerCommand(
    "tabula.toggleAutoExecute",
    () => {
      const config = vscode.workspace.getConfiguration("tabula");
      const currentValue = config.get<boolean>("autoExecute", true);
      config.update(
        "autoExecute",
        !currentValue,
        vscode.ConfigurationTarget.Global,
      );
      vscode.window.showInformationMessage(
        `Tabula auto-execute ${!currentValue ? "enabled" : "disabled"}`,
      );
    },
  );

  context.subscriptions.push(toggleCommand);
}

// This method is called when your extension is deactivated
export function deactivate() {}

const runScript = (path: vscode.Uri): Promise<void> => {
  return new Promise((resolve, reject) => {
    const pathParts = path.path.split("/");
    const fileName = pathParts[pathParts.length - 1];

    // Get configuration settings
    const config = vscode.workspace.getConfiguration("tabula");
    const tabulaPath = config.get<string>("executablePath", "tabula");
    const autoFormat = config.get<boolean>("autoFormat", true);

    // Build command with optional -a flag
    const autoFormatFlag = autoFormat ? "-a " : "";
    const command = `"${tabulaPath}" ${autoFormatFlag}-u "${path.path}"`;

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
      vscode.window.showInformationMessage(`Script for ${fileName} done.`);
      resolve();
    });
  });
};

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
