import { App, Plugin, PluginManifest, TFile } from "obsidian";
import { TabulaSettings, DEFAULT_SETTINGS } from "./types";
import { TabulaSettingTab } from "./settings";
import { extractChunks, processChunks, outputChunks } from "./processor";
import { Executer } from "./executer";

export default class TabulaPlugin extends Plugin {
  settings: TabulaSettings = DEFAULT_SETTINGS;
  private updatingFiles = new Set<string>();

  constructor(app: App, manifest: PluginManifest) {
    super(app, manifest);
  }

  async onload() {
    await this.loadSettings();

    this.registerMarkdownCodeBlockProcessor("csv", (source, el, _ctx) => {
      const rows = source.split("\n").filter((row) => row.length > 0);

      const table = el.createEl("table");
      const body = table.createEl("tbody");

      for (let i = 0; i < rows.length; i++) {
        const cols = rows[i].split(",");

        const row = body.createEl("tr");

        for (let j = 0; j < cols.length; j++) {
          row.createEl("td", { text: cols[j] });
        }
      }
    });

    this.registerMarkdownCodeBlockProcessor("tabula", (_source, el, _ctx) => {
      const container = el.createEl("div");
      container.className = "cm-comment tabula-code";
      const button = container.createEl("div");
      button.appendText("⚙ Tabula");
    });

    // Listen for markdown file modifications
    this.registerEvent(
      this.app.vault.on("modify", async (file) => {
        if (!this.settings.autoExecute) {
          return;
        }

        if (this.updatingFiles.has(file.path)) return;

        // Only process markdown files
        if (file instanceof TFile && file.extension === "md") {
          const content = await this.app.vault.read(file);
          const chunks = extractChunks(content);

          const fileBasePath = file.path.substring(
            0,
            file.path.lastIndexOf("/") + 1,
          );
          const executer = new Executer(
            this.settings,
            // @ts-expect-error undocumented
            this.app.vault.adapter.basePath + fileBasePath,
          );
          const processed = await processChunks(executer, chunks);
          const output = outputChunks(processed);

          this.updatingFiles.add(file.path);
          try {
            await file.vault.modify(file, output);
          } finally {
            this.updatingFiles.delete(file.path);
          }
        }
      }),
    );

    // Add settings tab
    this.addSettingTab(new TabulaSettingTab(this.app, this));
  }

  async loadSettings() {
    this.settings = Object.assign({}, DEFAULT_SETTINGS, await this.loadData());
  }

  async saveSettings() {
    await this.saveData(this.settings);
  }
}
