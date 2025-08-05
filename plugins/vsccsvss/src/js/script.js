const vscode = acquireVsCodeApi();

function handleInputChange(event) {
  const input = event.target;
  const newValue = input.value;
  const rowIndex = input.dataset.rowIndex;
  const columnIndex = input.dataset.columnIndex;

  vscode.postMessage({
    command: "updateCell",
    rowIndex,
    columnIndex,
    value: newValue,
  });
}

const allInputs = document.querySelectorAll("input");

allInputs.forEach((input) => {
  input.addEventListener("change", handleInputChange);
});
