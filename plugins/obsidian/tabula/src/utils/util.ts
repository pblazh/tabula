import { TableData } from "../types";

export const appendRow = (rowIdx: number, tableData: TableData) => {
  tableData.splice(rowIdx + 1, 0, Array(tableData[0].length).fill(""));
};

export const prependRow = (rowIdx: number, tableData: TableData) => {
  tableData.splice(rowIdx, 0, Array(tableData[0].length).fill(""));
};

export const appendColumn = (colIdx: number, tableData: TableData) => {
  tableData.forEach((row) => row.splice(colIdx + 1, 0, ""));
};

export const preppendColumn = (colIdx: number, tableData: TableData) => {
  tableData.forEach((row) => row.splice(colIdx, 0, ""));
};

export const deleteRow = (rowIdx: number, tableData: TableData) => {
  if (tableData.length <= 1) return;
  tableData.splice(rowIdx, 1);
};

export const deleteColumn = (colIdx: number, tableData: TableData) => {
  if (tableData[0].length < 1) return;
  tableData.forEach((row) => row.splice(colIdx, 1));
};

export const moveRow = (
  fromIndex: number,
  toIndex: number,
  tableData: TableData,
) => {
  if (isOffBounds(fromIndex, toIndex, tableData.length)) return;

  const row = tableData.splice(fromIndex, 1)[0];
  tableData.splice(toIndex, 0, row);
};

export const moveColumn = (
  fromIndex: number,
  toIndex: number,
  tableData: TableData,
) => {
  if (isOffBounds(fromIndex, toIndex, tableData[0].length)) return;

  tableData.forEach((row) => {
    const col = row.splice(fromIndex, 1)[0];
    row.splice(toIndex, 0, col);
  });
};

const isOffBounds = (from: number, to: number, bounds: number): boolean =>
  from < 0 || to < 0 || from >= bounds || to >= bounds;

export const getColumnLabel = (idx: number): string => {
  let result = "";
  do {
    result = String.fromCharCode(65 + (idx % 26)) + result;
    idx = Math.floor(idx / 26) - 1;
  } while (idx >= 0);
  return result;
};

export const normalizeTableData = (tableData: string[][]): string[][] => {
  if (!tableData || tableData.length === 0) return [[""]];

  // Find maximum number of columns
  let maxCols = 0;
  for (const row of tableData) {
    if (row) {
      maxCols = Math.max(maxCols, row.length);
    }
  }

  // Ensure each row has the same number of columns
  const normalizedData = tableData.map((row) => {
    const newRow = row ? [...row] : [];
    while (newRow.length < maxCols) {
      newRow.push("");
    }
    return newRow;
  });

  return normalizedData;
};
