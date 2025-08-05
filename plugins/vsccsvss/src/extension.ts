// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import { parseString, format } from "fast-csv";
import * as vscode from "vscode";
type CsvRowType = {
  [key: string]: string;
};
// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {
  const webViewCommand = vscode.commands.registerCommand(
    "vsccsvss.start",
    async () => {
      // Create and show a new webview
      const panel = vscode.window.createWebviewPanel(
        "csvEditing", // Identifies the type of the webview. Used internally
        "File Content to Edit", // Title of the panel displayed to the user
        vscode.ViewColumn.One, // Editor column to show the new webview panel in.
        {
          enableScripts: true,
          localResourceRoots: [
            vscode.Uri.joinPath(context.extensionUri, "media"),
            vscode.Uri.joinPath(context.extensionUri, "src", "js"),
          ],
        }
      );

      const onDiskPathCsv = vscode.Uri.joinPath(
        context.extensionUri,
        "media",
        "test.csv"
      );

      const scriptPath = vscode.Uri.joinPath(
        context.extensionUri,
        "src",
        "js",
        "script.js"
      );
      const scriptUri = panel.webview.asWebviewUri(scriptPath);
      const cspSource = panel.webview.cspSource;

      try {
        const csvContent = await readFileContent(onDiskPathCsv);

        panel.webview.html = getWebviewContent(
          csvContent,
          scriptUri,
          cspSource
        );

        panel.webview.onDidReceiveMessage(
          async (message) => {
            switch (message.command) {
              case "updateCell":
                csvContent[message.rowIndex][message.columnIndex] =
                  message.value;
                await saveFileContent(onDiskPathCsv, csvContent);
                return;
            }
          },
          undefined,
          context.subscriptions
        );
      } catch (error) {
        console.log(error);
      }
    }
  );

  context.subscriptions.push(webViewCommand);
}

function getWebviewContent(
  csvRows: CsvRowType[] | undefined,
  scriptPath: vscode.Uri,
  cspSource: string
) {
  if (!csvRows) {
    return "No file to edit";
  }

  const columnNames = Object.keys(csvRows[0]);
  const rows = csvRows.reduce<{
    content: typeof csvRows;
    comments: typeof csvRows;
  }>(
    (acc, row) => {
      const updatedAcc = { content: acc.content, comments: acc.comments };
      if (Object.values(row)[0].includes("#")) {
        updatedAcc.comments.push(row);
      } else {
        updatedAcc.content.push(row);
      }
      return updatedAcc;
    },
    { content: [], comments: [] }
  );
  return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="Content-Security-Policy" content="default-src 'none'; script-src ${cspSource}; style-src 'unsafe-inline';">
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
       (row, indexRow) =>
         `<tr>${columnNames
           .map(
             (column, indexColumn) =>
               `<td>
                  <input
                    value="${row[column]}"
                    data-row-index="${indexRow}"
                    data-column-index="${indexColumn}"
                  />
                </td>`
           )
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

<script src="${scriptPath}"></script>
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
      .on("end", () => {
        resolve(results);
      });
  });
}

async function saveFileContent(
  fileUri: vscode.Uri,
  data: CsvRowType[]
): Promise<void> {
  try {
    const csvString = await new Promise<string>((resolve, reject) => {
      const chunks: Buffer[] = [];

      const stringifyStream = format({ headers: true });

      stringifyStream.on("data", (chunk: Buffer) => {
        chunks.push(chunk);
      });

      stringifyStream.on("error", (err) => {
        reject(err);
      });

      stringifyStream.on("end", () => {
        resolve(Buffer.concat(chunks).toString("utf-8"));
      });

      data.forEach((row) => stringifyStream.write(row));

      stringifyStream.end();
    });

    const writeData: Uint8Array = new TextEncoder().encode(csvString);

    await vscode.workspace.fs.writeFile(fileUri, writeData);

    vscode.window.showInformationMessage(`File was saved: ${fileUri.fsPath}`);
  } catch (error) {
    console.error("Save file by fast-csv Error:", error);
    vscode.window.showErrorMessage(`Saving file failure`);
    return Promise.reject(error);
  }
}

// This method is called when your extension is deactivated
export function deactivate() {}
