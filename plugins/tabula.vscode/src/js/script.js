//TODO: TypeScript
const vscode = acquireVsCodeApi();

function getCellIndexesFromId(id) {
  const splittedId = id.split("-");
  return {
    columnIndex: splittedId[0],
    rowIndex: splittedId[1],
  };
}

function handleInputChange(event) {
  const input = event.target;
  const newValue = input.value;
  const id = input.id;

  const { columnIndex, rowIndex } = getCellIndexesFromId(id);

  vscode.postMessage({
    command: "updateCell",
    rowIndex,
    columnIndex,
    value: newValue,
  });
}

const table = document.querySelector("table");

table.addEventListener("change", handleInputChange);
