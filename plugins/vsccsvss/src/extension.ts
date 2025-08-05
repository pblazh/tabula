// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import { parseString, format } from "fast-csv";
import * as vscode from "vscode";
import * as tableTemplate from "./templates/table";

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

      const csvPath = vscode.Uri.joinPath(
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

      const template = tableTemplate.createTablePage(
        scriptUri.toString(),
        cspSource
      );

      try {
        const csvContent = await readFileContent(csvPath);

        const { head, body, foot } = parseTable(csvContent);

        const thead = tableTemplate.createTableHead(head);

        const tbody = tableTemplate.createTableBody(body);

        const tfoot = tableTemplate.createTableFooter(foot);

        panel.webview.html = template(thead, tbody, tfoot);

        panel.webview.onDidReceiveMessage(
          (message) => {
            switch (message.command) {
              case "updateCell":
                csvContent[message.rowIndex][message.columnIndex] =
                  message.value;
                saveFileContent(csvPath, csvContent)
                  .then(showSavingResult(csvPath))
                  .catch(showSavingError);
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

function parseTable(table: string[][]): {
  head: string[];
  body: string[][];
  foot: string[][];
} {
  const head = Object.keys(table[0]);
  const { body, foot } = table.reduce<{
    body: typeof table;
    foot: typeof table;
  }>(
    (acc, row) => {
      const updatedAcc = { body: acc.body, foot: acc.foot };
      if (Object.values(row)[0].includes("#")) {
        updatedAcc.foot.push(row);
      } else {
        updatedAcc.body.push(row);
      }
      return updatedAcc;
    },
    { body: [], foot: [] }
  );

  return { head, body, foot };
}

//TODO create functions: getWebviewTable and getWebviewNotTable

async function readFileContent(fileUri: vscode.Uri) {
  const readData: Uint8Array = await vscode.workspace.fs.readFile(fileUri);
  const fileContent: string = new TextDecoder("utf-8").decode(readData);

  return new Promise<string[][]>((resolve, reject) => {
    const results: string[][] = [];

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
  data: string[][]
): Promise<void> {
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

  vscode.workspace.fs.writeFile(fileUri, writeData);
}

// This method is called when your extension is deactivated
export function deactivate() {}

//TODO try to check is it dev mode now
function showSavingError(error: unknown): void {
  vscode.window.showErrorMessage(
    `Saving file failure: ${
      error instanceof Error ? error.message : String(error)
    }`
  );
}

const showSavingResult = (path: vscode.Uri) => () => {
  vscode.window.showInformationMessage(`File was saved: ${path}`);
};
