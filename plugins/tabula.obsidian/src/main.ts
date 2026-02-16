import { Plugin, TFile } from "obsidian";
import { TabulaSettings, DEFAULT_SETTINGS } from "./types";
import { TabulaSettingTab } from "./settings";
import { extractChunks, processChunks, outputChunks } from "./processor";
import { Executer } from "./executer";

export default class TabulaPlugin extends Plugin {
  settings: TabulaSettings;
  executer: Executer;
  private updatingFiles = new Set<string>();

  async onload() {
    await this.loadSettings();

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
          this.executer = new Executer(
            this.settings,
            // @ts-expect-error undocumented
            this.app.vault.adapter.basePath + fileBasePath,
          );
          const processed = await processChunks(this.executer, chunks);
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
