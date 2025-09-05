import { Plugin, WorkspaceLeaf, moment } from "obsidian";
import { TableView, VIEW_TYPE } from "./views/table-view";
import { SourceView, VIEW_TYPE as VIEW_TYPE_SRC } from "./views/source-view";
import { i18n } from "./i18n";
import { SettingTab } from "./views/settings";
import { Settings } from "./types";
import { validateSettings } from "./services/validateSettings";
import { showWarningNotice } from "./utils/error-utils";
import { InvalidSettingsError, PluginFaildedToStartError } from "./errors";

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

      const validationErrors = validateSettings(this.settings);
      if (validationErrors.length > 0) {
        showWarningNotice(InvalidSettingsError(validationErrors));
      }

      this.addSettingTab(new SettingTab(this.app, this));

      const obsidianLang = moment.locale();
      i18n.setLocale(obsidianLang);

      this.registerView(
        VIEW_TYPE,
        (leaf: WorkspaceLeaf) => new TableView(leaf, this.settings),
      );

      this.registerView(
        VIEW_TYPE_SRC,
        (leaf: WorkspaceLeaf) => new SourceView(leaf),
      );

      this.registerExtensions(["csv"], VIEW_TYPE);
    } catch (error) {
      console.error("Failed to load Tabula plugin:", error);
      showWarningNotice(PluginFaildedToStartError(error));
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
