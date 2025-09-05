import { TableData } from "../../types";
import {
  appendRow,
  prependRow,
  deleteColumn,
  deleteRow,
  moveColumn,
  moveRow,
  prependColumn,
  appendColumn,
} from "../../utils/util";

type Action = {
  label: string;
  action: VoidFunction;
};

export const createRowOptions = (
  refresh: VoidFunction,
  idx: number,
  tableData: TableData,
): Action[] => [
  {
    label: "contextMenu.insertBefore",
    action: () => {
      prependRow(idx, tableData);
      refresh();
    },
  },
  {
    label: "contextMenu.insertAfter",
    action: () => {
      appendRow(idx, tableData);
      refresh();
    },
  },
  {
    label: "contextMenu.moveUp",
    action: () => {
      moveRow(idx, idx - 1, tableData);
      refresh();
    },
  },
  {
    label: "contextMenu.moveDown",
    action: () => {
      moveRow(idx, idx + 1, tableData);
      refresh();
    },
  },
  {
    label: "contextMenu.delete",
    action: () => {
      deleteRow(idx, tableData);
      refresh();
    },
  },
];

export const createColumnOptions = (
  refresh: VoidFunction,
  idx: number,
  tableData: TableData,
): Action[] => [
  {
    label: "contextMenu.delete",
    action: () => {
      deleteColumn(idx, tableData);
      refresh();
    },
  },
  {
    label: "contextMenu.insertBefore",
    action: () => {
      prependColumn(idx, tableData);
      refresh();
    },
  },
  {
    label: "contextMenu.insertAfter",
    action: () => {
      appendColumn(idx, tableData);
      refresh();
    },
  },
  {
    label: "contextMenu.moveLeft",
    action: () => {
      moveColumn(idx, idx - 1, tableData);
      refresh();
    },
  },
  {
    label: "contextMenu.moveRight",
    action: () => {
      moveColumn(idx, idx + 1, tableData);
      refresh();
    },
  },
];
