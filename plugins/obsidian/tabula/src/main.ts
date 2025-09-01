import { Plugin, WorkspaceLeaf, moment } from "obsidian";
import { TableView, VIEW_TYPE } from "./views/table-view";
import { SourceView, VIEW_TYPE as VIEW_TYPE_SRC } from "./views/source-view";
import { i18n } from "./i18n";
import { SettingTab } from "./views/settings";
import { Settings } from "./types";

const DEFAULT_SETTINGS: Settings = {
  tabula: "tabula",
  delimiter: ",",
  comment: "#",
  quote: '"',
};

export default class Tabula extends Plugin {
  settings: Settings;

  async onload() {
    await this.loadSettings();
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
  }

  onunload() {
    // Remove views
  }

  async loadSettings() {
    this.settings = Object.assign({}, DEFAULT_SETTINGS, await this.loadData());
  }

  async saveSettings() {
    await this.saveData(this.settings);
  }
}
