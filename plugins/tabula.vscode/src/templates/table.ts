export const createTablePage =
  (scriptPath: string, cspSource: string) =>
  (head: string, table: string, footer: string) => {
    return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="Content-Security-Policy" content="default-src 'none'; script-src ${cspSource}; style-src 'unsafe-inline';">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>File Content</title>
</head>
<body>
<h2>Table View of CSV</h2>
<table>
<thead>
    ${head}  
</thead>
<tbody>
    ${table}  
</tbody>
<tfoot>
    ${footer}  
</tfoot>

<script src="${scriptPath}"></script>
</body>
</html>`;
  };

export function createTableHead(head: string[]): string {
  return `<tr> ${head
    .map(
      (item) =>
        `<td>${String.fromCharCode("A".charCodeAt(0) + parseInt(item))}</td>`,
    )
    .join("")}
    </tr>`;
}

export function createTableBody(content: string[][]): string {
  return `${content
    ?.map(
      (row, indexRow) =>
        `<tr>${row
          .map(
            (column, indexColumn) =>
              `<td>
                  <input
                    id="${indexColumn}-${indexRow}"
                    value="${column}"
                  />
                </td>`,
          )
          .join("")}</tr>`,
    )
    .join("")}`;
}

export function createTableFooter(foot: string[][]): string {
  const lines: string[] = [];
  for (let tr of foot) {
    lines.push(`<tr>`);
    for (let td of tr) {
      lines.push(`<td>${td}</td>`);
    }
    lines.push(`</tr>`);
  }
  return lines.join("");
}
