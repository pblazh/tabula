import { Plugin, WorkspaceLeaf, moment } from "obsidian";
import { TableView, VIEW_TYPE } from "./views/table-view";
import { SourceView, VIEW_TYPE as VIEW_TYPE_SRC } from "./views/source-view";
import { i18n } from "./i18n";
import { SettingTab } from "./views/settings";
import { Settings } from "./types";
import { validateCSVSettings } from "./services/csv-service";
import { showWarningNotice } from "./utils/error-utils";

const DEFAULT_SETTINGS: Settings = {
  tabula: "tabula",
  delimiter: ",",
  comment: "#",
  quote: '"',
};

export default class Tabula extends Plugin {
  settings: Settings;

  async onload() {
    try {
      await this.loadSettings();
      
      // Validate settings on startup
      const validationErrors = validateCSVSettings(this.settings);
      if (validationErrors.length > 0) {
        showWarningNotice(`Settings validation warnings: ${validationErrors.join(", ")}`);
      }
      
      this.addSettingTab(new SettingTab(this.app, this));

      const obsidianLang = moment.locale();
      i18n.setLocale(obsidianLang);

      // Register CSV view type
      this.registerView(
        VIEW_TYPE,
        (leaf: WorkspaceLeaf) => new TableView(leaf, this.settings),
      );

      // Register source view type
      this.registerView(
        VIEW_TYPE_SRC,
        (leaf: WorkspaceLeaf) => new SourceView(leaf),
      );

      // Bind .csv file extension with view type
      this.registerExtensions(["csv"], VIEW_TYPE);
    } catch (error) {
      console.error("Failed to load Tabula plugin:", error);
      showWarningNotice(`Failed to load Tabula plugin: ${error instanceof Error ? error.message : 'Unknown error'}`);
    }
  }

  onunload() {
    // Remove views
  }

  async loadSettings() {
    try {
      const loadedData = await this.loadData();
      this.settings = Object.assign({}, DEFAULT_SETTINGS, loadedData);
    } catch (error) {
      console.error("Failed to load settings, using defaults:", error);
      this.settings = { ...DEFAULT_SETTINGS };
    }
  }

  async saveSettings() {
    try {
      await this.saveData(this.settings);
    } catch (error) {
      console.error("Failed to save settings:", error);
      throw error;
    }
  }
}
