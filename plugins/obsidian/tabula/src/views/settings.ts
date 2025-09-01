import { App, PluginSettingTab, Setting } from "obsidian";
import Tabula from "../main";

export class SettingTab extends PluginSettingTab {
  plugin: Tabula;

  constructor(app: App, plugin: Tabula) {
    super(app, plugin);
    this.plugin = plugin;
  }

  display(): void {
    let { containerEl } = this;

    containerEl.empty();

    new Setting(containerEl)
      .setName("Tabula executable path")
      .setDesc(
        "absolute path to the tabula executable on just tabula if in $PATH",
      )
      .addText((text) =>
        text
          .setPlaceholder("tabula")
          .setValue(this.plugin.settings.tabula)
          .onChange(async (value) => {
            this.plugin.settings.tabula = value;
            await this.plugin.saveSettings();
          }),
      );

    new Setting(containerEl)
      .setName("CSV comments prefix")
      .setDesc(
        "If a line in a CSV staarts with this value it's ignored and treated as a comment",
      )
      .addText((text) =>
        text
          .setPlaceholder("comment")
          .setValue(this.plugin.settings.comment)
          .onChange(async (value) => {
            this.plugin.settings.comment = value;
            await this.plugin.saveSettings();
          }),
      );

    new Setting(containerEl)
      .setName("CSV delimieter")
      .setDesc("delimieter used to separate filelds in csv")
      .addText((text) =>
        text
          .setPlaceholder(",")
          .setValue(this.plugin.settings.delimiter)
          .onChange(async (value) => {
            this.plugin.settings.delimiter = value;
            await this.plugin.saveSettings();
          }),
      );
  }
}
