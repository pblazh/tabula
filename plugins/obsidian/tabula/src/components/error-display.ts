export interface ErrorDisplayOptions {
  title?: string;
  collapsible?: boolean;
  maxHeight?: string;
}

export class ErrorDisplay {
  private container: HTMLElement;
  private contentEl: HTMLElement;
  private isCollapsed: boolean = true;

  constructor(
    private parentEl: HTMLElement,
    private options: ErrorDisplayOptions = {}
  ) {
    this.createContainer();
  }

  private createContainer(): void {
    this.container = this.parentEl.createEl("div", {
      cls: "tabula-error-display",
    });

    const header = this.container.createEl("div", {
      cls: "tabula-error-header",
    });

    if (this.options.collapsible) {
      header.addClass("tabula-error-header-collapsible");
      header.addEventListener("click", () => this.toggle());
    }

    const titleEl = header.createEl("span", {
      cls: "tabula-error-title",
      text: this.options.title || "Error Details",
    });

    if (this.options.collapsible) {
      const toggleIcon = header.createEl("span", {
        cls: "tabula-error-toggle",
        text: "▶",
      });
      toggleIcon.setAttribute("aria-label", "Toggle error details");
    }

    this.contentEl = this.container.createEl("div", {
      cls: "tabula-error-content",
    });

    if (this.options.collapsible) {
      this.contentEl.style.display = "none";
    }

    if (this.options.maxHeight) {
      this.contentEl.style.maxHeight = this.options.maxHeight;
      this.contentEl.style.overflowY = "auto";
    }

    this.addStyles();
  }

  private addStyles(): void {
    if (document.querySelector("#tabula-error-styles")) return;

    const style = document.createElement("style");
    style.id = "tabula-error-styles";
    style.textContent = `
      .tabula-error-display {
        margin: 10px 0;
        border: 1px solid var(--background-modifier-error);
        border-radius: 6px;
        background: var(--background-primary-alt);
      }

      .tabula-error-header {
        padding: 8px 12px;
        background: var(--background-modifier-error);
        color: var(--text-on-accent);
        font-weight: 600;
        display: flex;
        justify-content: space-between;
        align-items: center;
        border-radius: 6px 6px 0 0;
      }

      .tabula-error-header-collapsible {
        cursor: pointer;
        user-select: none;
      }

      .tabula-error-header-collapsible:hover {
        background: var(--background-modifier-error-hover);
      }

      .tabula-error-toggle {
        font-size: 12px;
        transition: transform 0.2s;
      }

      .tabula-error-toggle.expanded {
        transform: rotate(90deg);
      }

      .tabula-error-content {
        padding: 12px;
        font-family: var(--font-monospace);
        font-size: 13px;
        line-height: 1.4;
        white-space: pre-wrap;
        word-break: break-word;
        color: var(--text-muted);
      }

      .tabula-error-content:empty {
        display: none;
      }
    `;
    document.head.appendChild(style);
  }

  public show(errorMessage: string): void {
    this.contentEl.textContent = errorMessage.trim();
    this.container.style.display = "block";
  }

  public hide(): void {
    this.container.style.display = "none";
  }

  public clear(): void {
    this.contentEl.textContent = "";
    this.hide();
  }

  private toggle(): void {
    if (!this.options.collapsible) return;

    this.isCollapsed = !this.isCollapsed;
    const toggleIcon = this.container.querySelector(".tabula-error-toggle") as HTMLElement;
    
    if (this.isCollapsed) {
      this.contentEl.style.display = "none";
      if (toggleIcon) {
        toggleIcon.textContent = "▶";
        toggleIcon.classList.remove("expanded");
      }
    } else {
      this.contentEl.style.display = "block";
      if (toggleIcon) {
        toggleIcon.textContent = "▼";
        toggleIcon.classList.add("expanded");
      }
    }
  }

  public destroy(): void {
    this.container.remove();
  }
}