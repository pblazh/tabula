import { App, PluginSettingTab, Setting } from "obsidian";
import Tabula from "../main";
import { validateSettings } from "../services/validateSettings";
import { checkTabulaExecutableAvailable } from "../services/process-service";
import { showErrorNotice } from "../utils/error-utils";

export class SettingTab extends PluginSettingTab {
  plugin: Tabula;
  private validationTimeout: NodeJS.Timeout | null = null;

  constructor(app: App, plugin: Tabula) {
    super(app, plugin);
    this.plugin = plugin;
  }

  display(): void {
    const { containerEl } = this;

    containerEl.empty();

    // Add validation info section
    const validationEl = containerEl.createDiv({ cls: "setting-item-info" });
    validationEl.createDiv({
      cls: "setting-item-name",
      text: "Settings Validation",
    });
    const validationDesc = validationEl.createDiv({
      cls: "setting-item-description",
    });
    this.updateValidationStatus(validationDesc);

    new Setting(containerEl)
      .setName("Tabula executable path")
      .setDesc(
        "Absolute path to the tabula executable or just 'tabula' if in $PATH",
      )
      .addText((text) =>
        text
          .setPlaceholder("tabula")
          .setValue(this.plugin.settings.tabula)
          .onChange(async (value) => {
            this.plugin.settings.tabula = value.trim();
            await this.plugin.saveSettings();
            this.scheduleValidation(validationDesc);
          }),
      );

    new Setting(containerEl)
      .setName("CSV comments prefix")
      .setDesc(
        "If a line in a CSV starts with this value it's ignored and treated as a comment",
      )
      .addText((text) =>
        text
          .setPlaceholder("#")
          .setValue(this.plugin.settings.comment)
          .onChange(async (value) => {
            this.plugin.settings.comment = value;
            await this.plugin.saveSettings();
            this.scheduleValidation(validationDesc);
          }),
      );

    new Setting(containerEl)
      .setName("CSV delimiter")
      .setDesc("Delimiter used to separate fields in CSV")
      .addText((text) =>
        text
          .setPlaceholder(",")
          .setValue(this.plugin.settings.delimiter)
          .onChange(async (value) => {
            // Ensure delimiter is not empty
            if (!value.trim()) {
              showErrorNotice("Delimiter cannot be empty");
              return;
            }
            this.plugin.settings.delimiter = value;
            await this.plugin.saveSettings();
            this.scheduleValidation(validationDesc);
          }),
      );
  }

  private scheduleValidation(validationEl: HTMLElement): void {
    if (this.validationTimeout) {
      clearTimeout(this.validationTimeout);
    }

    this.validationTimeout = setTimeout(() => {
      this.updateValidationStatus(validationEl);
    }, 500);
  }

  private async updateValidationStatus(
    validationEl: HTMLElement,
  ): Promise<void> {
    validationEl.empty();

    // Validate CSV settings
    const csvErrors = validateSettings(this.plugin.settings);

    // Check tabula availability
    let tabulaAvailable = false;
    try {
      tabulaAvailable = await checkTabulaExecutableAvailable(
        this.plugin.settings.tabula,
      );
    } catch (error) {
      // Ignore errors for now
    }

    if (csvErrors.length === 0 && tabulaAvailable) {
      validationEl.createSpan({
        text: "✓ All settings are valid",
        cls: "setting-validation-success",
      });
    } else {
      const issues: string[] = [];

      if (csvErrors.length > 0) {
        issues.push(...csvErrors);
      }

      if (!tabulaAvailable) {
        issues.push("Tabula executable not found or not accessible");
      }

      const issueText = issues.length === 1 ? "Issue" : "Issues";
      validationEl.createSpan({
        text: `⚠ ${issueText}: ${issues.join(", ")}`,
        cls: "setting-validation-error",
      });
    }
  }
}
