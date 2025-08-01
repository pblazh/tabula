// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import { parseString } from "fast-csv";
import * as vscode from "vscode";
type CsvRowType = {
  [key: string]: string;
};
// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {
  // Use the console to output diagnostic information (console.log) and errors (console.error)
  // This line of code will only be executed once when your extension is activated
  console.log('Congratulations, your extension "vsccsvss" is now active!');

  // The command has been defined in the package.json file
  // Now provide the implementation of the command with registerCommand
  // The commandId parameter must match the command field in package.json
  const disposable = vscode.commands.registerCommand(
    "vsccsvss.helloWorld",
    () => {
      // The code you place here will be executed every time your command is executed
      // Display a message box to the user
      vscode.window.showInformationMessage("Hello World from VS CODE!");
    }
  );

  const webViewCommand = vscode.commands.registerCommand(
    "vsccsvss.start",
    async () => {
      // Create and show a new webview
      const panel = vscode.window.createWebviewPanel(
        "csvEditing", // Identifies the type of the webview. Used internally
        "File Content to Edit", // Title of the panel displayed to the user
        vscode.ViewColumn.One, // Editor column to show the new webview panel in.
        {} // Webview options. More on these later.
      );

      const onDiskPathCsv = vscode.Uri.joinPath(
        context.extensionUri,
        "media",
        "test.csv"
      );
      try {
        const csvContent = await readFileContent(onDiskPathCsv);

        panel.webview.html = getWebviewContent(csvContent);
      } catch (error) {
        console.log(error);
      }
    }
  );

  context.subscriptions.push(disposable, webViewCommand);
}

function getWebviewContent(csvRows: CsvRowType[] | undefined) {
  if (!csvRows) return "No file to edit";

  const columnNames = Object.keys(csvRows[0]);
  const rows = csvRows.reduce<{
    content: typeof csvRows;
    comments: typeof csvRows;
  }>(
    (acc, row) => {
      const updatedAcc = { content: acc.content, comments: acc.comments };
      if (Object.values(row)[0].includes("#")) updatedAcc.comments.push(row);
      else updatedAcc.content.push(row);
      return updatedAcc;
    },
    { content: [], comments: [] }
  );
  return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>File Content</title>
</head>
<body>
<section>
<h2>Table View of CSV</h2>
<table>
  <thead>
    <tr>
	${columnNames?.map((item) => `<td>${item}</td>`).join("")}
    </tr>
  </thead>
  <tbody>
   ${rows.content
     ?.map(
       (row) =>
         `<tr>${columnNames
           .map((column) => `<td><input value="${row[column]}" /></td>`)
           .join("")}</tr>`
     )
     .join("")}
  </tbody>
<!--  <tfoot>
     ${rows.comments
       ?.map(
         (row) =>
           `<tr colspan="${columnNames.length}"><td>${
             Object.values(row)[0]
           }</td></tr>`
       )
       .join("")}
  </tfoot> -->
</table>
</section>
<section>
<h2>Parsed Rows of CSV</h2>
    ${csvRows?.map((item) => `<div>${JSON.stringify(item)}</div><br>`).join("")}
</section>
	</body>
</html>`;
}

async function readFileContent(fileUri: vscode.Uri) {
  const readData: Uint8Array = await vscode.workspace.fs.readFile(fileUri);
  const fileContent: string = new TextDecoder("utf-8").decode(readData);

  return new Promise<CsvRowType[]>((resolve, reject) => {
    const results: CsvRowType[] = [];

    parseString(fileContent)
      .on("error", (error) => {
        reject(error);
      })
      .on("data", (row) => {
        results.push(row);
      })
      .on("end", (rowCount: number) => {
        console.log(`Parsing ended. Number of rows: ${rowCount}`);
        resolve(results);
      });
  });
}

// This method is called when your extension is deactivated
export function deactivate() {}
