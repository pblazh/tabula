import { App, PluginSettingTab, Setting } from "obsidian";
import Tabula from "../main";
import { validateCSVSettings } from "../services/csv-service";
import { checkTabulaAvailable } from "../services/process-service";
import { showErrorNotice, showWarningNotice, showSuccessNotice } from "../utils/error-utils";

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
    validationEl.createDiv({ cls: "setting-item-name", text: "Settings Validation" });
    const validationDesc = validationEl.createDiv({ cls: "setting-item-description" });
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

  private async updateValidationStatus(validationEl: HTMLElement): Promise<void> {
    validationEl.empty();
    
    // Validate CSV settings
    const csvErrors = validateCSVSettings(this.plugin.settings);
    
    // Check tabula availability
    let tabulaAvailable = false;
    try {
      tabulaAvailable = await checkTabulaAvailable(this.plugin.settings.tabula);
    } catch (error) {
      // Ignore errors for now
    }
    
    if (csvErrors.length === 0 && tabulaAvailable) {
      validationEl.createSpan({ 
        text: "✓ All settings are valid", 
        cls: "setting-validation-success" 
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
        cls: "setting-validation-error" 
      });
    }
    
    // Add CSS for validation styling
    this.addValidationStyles();
  }

  private addValidationStyles(): void {
    if (!document.getElementById("tabula-validation-styles")) {
      const style = document.createElement("style");
      style.id = "tabula-validation-styles";
      style.textContent = `
        .setting-validation-success {
          color: var(--text-success);
          font-weight: 500;
        }
        .setting-validation-error {
          color: var(--text-error);
          font-weight: 500;
        }
        .setting-item-info {
          margin-bottom: 1rem;
          padding: 0.5rem;
          background: var(--background-secondary);
          border-radius: var(--radius-s);
        }
      `;
      document.head.appendChild(style);
    }
  }
}
