import { TextFileView, Notice, IconName, TFile, WorkspaceLeaf } from "obsidian";
import * as child_process from "child_process";
import * as path from "path";
import { parseCSV, unparseCSV } from "../utils/csv-utils";
import { FileUtils } from "../utils/file-utils";
import { normalizeTableData } from "../utils/util";

import { renderTable } from "../components/table";
import { setupContextMenu } from "../components/context-menu";
import { Settings, TableComments, TableData } from "../types";

export const VIEW_TYPE = "csv-view";

export class TableView extends TextFileView {
  public file: TFile | null;
  private headerEl: HTMLElement;
  private interval: NodeJS.Timeout;

  private tableData: TableData = [];
  private tableComments: TableComments = {};
  private tableEl: HTMLElement;
  private headerContextMenuCleanup: (() => void) | null = null;

  constructor(
    leaf: WorkspaceLeaf,
    private settings: Settings,
  ) {
    super(leaf);
    this.setupSafeSave();
  }

  getIcon(): IconName {
    return "table";
  }
  getViewData() {
    return unparseCSV(this.settings, this.tableData, this.tableComments);
  }

  // We need to create a wrapper for the original requestSave
  private originalRequestSave: () => void;

  /**
   * Setup safe save method with retry logic
   */
  private setupSafeSave() {
    // Store the original requestSave function
    this.originalRequestSave = this.requestSave;

    // Replace with our version that includes retry logic
    this.requestSave = async () => {
      try {
        // Use our retry utility to handle file busy errors
        await FileUtils.withRetry(async () => {
          // Call the original requestSave method
          this.originalRequestSave();
          // Return a resolved promise to satisfy the async function
          return Promise.resolve();
        }).then(async () => {
          globalThis.clearInterval(this.interval);
          await new Promise((r) => {
            this.interval = globalThis.setTimeout(r, 1000);
          });
          if (this.file) {
            const fullPath = path.join(
              //@ts-ignore
              this.file.vault.adapter.basePath,
              this.file.path,
            );
            await this.execute(fullPath);
          }
        });
      } catch (error) {
        new Notice(
          `Failed to save file: ${error.message}. The file might be open in another program.`,
        );
      }
    };
  }

  execute(file: string) {
    return new Promise<void>((resolve, reject) => {
      new Notice("executing");
      const child = child_process.spawn(
        this.settings.tabula,
        ["-a", "-u", file],
        {
          env: process.env,
          shell: true,
        },
      );

      child.stderr.on("data", (data) => {
        new Notice("Err:" + data.toString());
        reject();
      });

      child.stdout.on("data", (data) => {
        new Notice("Out:" + data.toString());
        reject();
      });

      child.on("close", () => {
        new Notice("executing");

        resolve();
      });
    });
  }

  setViewData(data: string, clear: boolean) {
    // if (clear) {
    //   this.tableData = [];
    //   this.refresh();
    //   return;
    // }

    const { data: tableData, comments: tableComments } = parseCSV(
      this.settings,
      data,
    );
    this.tableData = normalizeTableData(tableData || []);
    this.tableComments = tableComments;

    this.refresh();
  }

  //private reparseAndRefresh() {
  //  this.setViewData(this.data, false);
  //}

  refresh() {
    if (!this.contentEl) return;

    this.contentEl
      .querySelectorAll(".csv-source-mode")
      .forEach((el) => el.remove());

    // Safety check: ensure tableData is initialized
    if (!this.tableData || !Array.isArray(this.tableData)) {
      console.warn("Table data not properly initialized, setting default");
      this.tableData = [];
    }

    renderTable(this.tableEl, this.tableData);

    if (this.headerContextMenuCleanup) {
      this.headerContextMenuCleanup();
      this.headerContextMenuCleanup = null;
    }
    this.headerContextMenuCleanup = setupContextMenu(
      this.tableEl,
      () => {
        this.refresh();
        this.requestSave();
      },
      this.tableData,
    );
  }

  clear() {
    this.tableData = [];
    this.refresh();
  }

  getViewType() {
    return VIEW_TYPE;
  }

  async onOpen() {
    try {
      const actionsEl = this.headerEl?.querySelector?.(".view-actions");
      if (actionsEl && !actionsEl.querySelector(".csv-switch-source")) {
        const btn = document.createElement("button");
        btn.className = "clickable-icon csv-switch-source";
        btn.setAttribute("aria-label", "Switch to source mode");
        btn.innerHTML = `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-file-code"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><polyline points="10 13 8 15 10 17"/><polyline points="14 13 16 15 14 17"/></svg>`;
        btn.onclick = async () => {
          const file = this.file;
          if (!file) return;
          const leaves = this.app.workspace.getLeavesOfType("csv-source-view");
          let found = false;
          for (const leaf of leaves) {
            if (
              leaf.view &&
              (leaf.view as any).file &&
              (leaf.view as any).file.path === file.path
            ) {
              this.app.workspace.setActiveLeaf(leaf, true, true);
              found = true;
              break;
            }
          }
          if (!found) {
            const newLeaf = this.app.workspace.getLeaf(true);
            await newLeaf.openFile(file, {
              active: true,
              state: { mode: "source" },
            });
            await newLeaf.setViewState({
              type: "csv-source-view",
              active: true,
              state: { file: file.path },
            });
            this.app.workspace.setActiveLeaf(newLeaf, true, true);
          }
        };
        actionsEl.appendChild(btn);
      }

      this.contentEl.empty();

      const tableWrapper = this.contentEl.createEl("div", {
        cls: "table-wrapper",
      });
      const tableContainer = tableWrapper.createEl("div", {
        cls: "table-container main-scroll",
      });
      this.tableEl = tableContainer.createEl("table", {
        cls: "table",
      });

      this.tableEl.addEventListener("input", (ev: Event) => {
        if (ev.target instanceof HTMLInputElement) {
          const data: DOMStringMap = (ev.target as HTMLInputElement).dataset;
          const x = parseInt(String(data?.x));
          const y = parseInt(String(data?.y));

          if (isNaN(x) || isNaN(y)) return;

          this.tableData[x][y] = ev.target.value;

          this.requestSave();
        }
      });

      // Ensure tableData is initialized before refreshing
      if (!this.tableData || !Array.isArray(this.tableData)) {
        this.tableData = [];
      }

      this.refresh();
    } catch (error) {
      console.error("Error in onOpen:", error);
      new Notice(`Failed to open CSV view: ${error.message}`);

      // Try to recover with minimal UI
      this.contentEl.empty();
      const errorDiv = this.contentEl.createEl("div", {
        cls: "csv-error",
      });
      errorDiv.createEl("h3", { text: "Error opening CSV file" });
      errorDiv.createEl("p", { text: error.message });

      this.tableData = [];
      this.tableEl = this.contentEl.createEl("table");
      this.refresh();
    }
  }

  setTableContent(content: string[][]) {
    this.tableData = content;
    this.refresh();
  }

  getTableContent(): string[][] {
    return this.tableData;
  }

  async onClose() {
    if (this.headerContextMenuCleanup) {
      this.headerContextMenuCleanup();
      this.headerContextMenuCleanup = null;
    }
    this.contentEl.empty();
  }
}
