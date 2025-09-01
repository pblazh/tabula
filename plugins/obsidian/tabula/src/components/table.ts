import { getColumnLabel } from "../utils/util";

export function renderTable(tableEl: HTMLElement, tableData: string[][]): void {
  tableEl.empty();

  const fragment = document.createDocumentFragment();

  // Create column number row (at top of table) - Fix: render header first
  const headerRow = fragment.createEl("thead").createEl("tr");

  // Create column number row
  if (tableData[0]) {
    Array(tableData[0].length + 1)
      .fill(null)
      .forEach((_, x) => {
        const th = headerRow.createEl("th", {
          cls: x
            ? "csv-column-number"
            : "csv-column-number csv-column-number-first",
          attr: { "data-x": x - 1 },
        });
        th.textContent = x > 0 ? getColumnLabel(x - 1) : "";
      });
  }

  const tableBody = fragment.createEl("tbody");

  for (let y = 0; y < tableData.length; y++) {
    const row = tableData[y];
    const tableRow = tableBody.createEl("tr");
    const rowNumberCell = tableRow.createEl("td", {
      cls: "csv-row-number",
      attr: { "data-y": y },
    });
    rowNumberCell.textContent = (y + 1).toString();

    row.forEach((cell, x) => {
      const td = tableRow.createEl("td");
      td.createEl("input", {
        cls: "csv-cell-input",
        attr: { value: cell, "data-y": y, "data-x": x },
      });
    });
  }

  tableEl.appendChild(fragment);
}
