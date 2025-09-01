import { TableData } from "types";
import { createColumnOptions, createRowOptions } from "./options";
import { render } from "./render";

export function setupContextMenu(
  tableEl: HTMLElement,
  refresh: VoidFunction,
  tableData: TableData,
): () => void {
  let closeMenu = () => {};

  const handler = (ev: MouseEvent) => {
    const target = ev.target as HTMLElement;
    const isRowNumber = target.classList.contains("csv-row-number");
    const isColumnNumber = target.classList.contains("csv-column-number");

    if (!isRowNumber && !isColumnNumber) return () => {};

    ev.preventDefault();

    const data: DOMStringMap = target.dataset;

    if (isRowNumber) {
      const y = parseInt(String(data?.y));
      if (isNaN(y)) return;

      const items = createRowOptions(refresh, y, tableData);
      closeMenu = render(items, ev.pageX, ev.pageY);
    }

    if (isColumnNumber) {
      const x = parseInt(String(data?.x));
      if (isNaN(x)) return;

      const items = createColumnOptions(refresh, x, tableData);
      closeMenu = render(items, ev.pageX, ev.pageY);
    }
  };
  tableEl.addEventListener("contextmenu", handler);

  return () => {
    tableEl.removeEventListener("contextmenu", handler);
    closeMenu();
  };
}
