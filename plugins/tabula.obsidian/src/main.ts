import { App, Plugin, PluginManifest, TFile, WorkspaceLeaf } from 'obsidian'
import { TabulaSettings, DEFAULT_SETTINGS } from './types'
import { TabulaSettingTab } from './settings'
import { Executer } from './executer'
import { renderTable } from './renderTable'
import { renderCode } from './renderCode'

function highlight(self: Plugin) {
  // The defineSimpleMode function is not immediately available during
  // onload, so continue to try and define the language until it is.
  const setupInterval = setInterval(() => {
    //@ts-expect-error Obsidian built-in
    if (CodeMirror && CodeMirror.defineSimpleMode) {
      const mode = {
        start: [
          { regex: /[-+*:=/<>]/, token: 'operator' },
          { regex: /include\b/i, token: 'keyword' },
          { regex: /\b#include\b/i, token: 'keyword' },
          { regex: /\b(fmt|let)\b/i, token: 'keyword' },
          { regex: /\b"[^"]*"\b/, token: 'string' },
          { regex: /\b0x[0-9a-f]+\b/i, token: 'number' },
          { regex: /\b-?\d+\b/, token: 'number' },
          { regex: /\b\s[0-9.]+\b/, token: 'number' },
          { regex: /\b([A-z]+[0-9]+:[A-z]+[0-9]+)\b/i, token: 'variable' },
          { regex: /\b([A-z]+[0-9]+)\b/i, token: 'variable' },
          { regex: /\b(".+")\b/i, token: 'string' },
          {
            regex:
              /@?(?:("|')(?:(?!\1)[^\n\\]|\\[\s\S])*\1(?!"|')|"""(?:[^\\]|\\[\s\S])*?""")/,
            token: 'string',
          },
          {
            regex:
              /\b(SUM|AVERAGE|MIN|MAX|ABS|ROUND|SQRT|MOD|POWER|FLOOR|CEIL|TRUNC|INT|SIGN|RAND|RANDBETWEEN|GCD|LCM|CONCATENATE|LEFT|RIGHT|MID|UPPER|LOWER|TRIM|LEN|SUBSTITUTE|FIND|REPLACE|TEXT|VALUE|SPLIT|JOIN|DATE|DATEVALUE|YEAR|MONTH|DAY|HOUR|MINUTE|SECOND|NOW|TODAY|DATEDIF|WEEKDAY|WORKDAY|NETWORKDAYS|TIME|TIMEVALUE|IF|AND|OR|NOT|TRUE|FALSE|ISNUMBER|ISTEXT|ISLOGICAL|ISBLANK|ISERROR|COLUMN|ROW|COLUMNS|ROWS|ADDRESS|REF|RANGE|INDEX|MATCH|VLOOKUP|HLOOKUP|EXEC|COUNT|COUNTA|COUNTBLANK|SUMIF|AVERAGEIF|MAXIFS|MINIFS)\b/,
            token: 'keyword',
          },
          { regex: /\b(?:false|true)\b/, token: 'number' },
        ],
        var_type: [{ regex: /(\w+)/, token: 'attribute', pop: true }],
        definition: [{ regex: /(\w+)/, token: 'attribute', pop: true }],
      }

      // @ts-expect-error wrong types
      CodeMirror.defineSimpleMode('tabula', mode)

      self.app.workspace.iterateAllLeaves((leaf: WorkspaceLeaf) => {
        // @ts-expect-error wrong types
        leaf?.rebuildView?.()
      })

      clearInterval(setupInterval)
    }
  }, 100)
}

export default class TabulaPlugin extends Plugin {
  settings: TabulaSettings = DEFAULT_SETTINGS
  private updatingFiles = new Set<string>()

  constructor(app: App, manifest: PluginManifest) {
    super(app, manifest)
  }

  invalidate() {}

  async onunload() {
    // @ts-expect-error wrong types
    delete CodeMirror.modes['tabula']
  }

  async onload() {
    await this.loadSettings()
    highlight(this)

    this.addCommand({
      id: 'tabula-execute',
      name: 'execute',
      callback: () => {
        const file = this.app.workspace.getActiveFile()
        if (file instanceof TFile && file.extension === 'md') {
          this.executeOnFile(file)
        }
      },
    })

    this.addCommand({
      id: 'tabula-toggle-auto-execution',
      name: 'toggle auto execution',
      callback: () => {
        this.settings.autoExecution = !this.settings.autoExecution
        this.saveSettings()
      },
    })

    this.addCommand({
      id: 'tabula-toggle-table-index',
      name: 'toggles rows and columns names visibility',
      callback: () => {
        this.settings.tableIndex = !this.settings.tableIndex
        this.saveSettings()
      },
    })

    this.registerMarkdownCodeBlockProcessor('csv', (source, el, _ctx) => {
      renderTable(this.settings, source, el)
    })
    this.registerMarkdownCodeBlockProcessor('tabula', (source, el, _ctx) => {
      renderCode(this.settings, source, el)
    })

    // Listen for markdown file modifications
    this.registerEvent(
      this.app.vault.on('modify', async (file) => {
        if (!this.settings.autoExecution) {
          return
        }

        // Only process markdown files
        if (file instanceof TFile && file.extension === 'md') {
          this.executeOnFile(file)
        }
      }),
    )

    // Add settings tab
    this.addSettingTab(new TabulaSettingTab(this.app, this))
  }

  private async executeOnFile(file: TFile) {
    if (file.extension.toLowerCase() !== 'md') {
      console.error('file type is not supported')
    }
    if (this.updatingFiles.has(file.path)) return

    // Only process markdown files
    if (file.extension === 'md') {
      const content = await this.app.vault.read(file)
      const executer = new Executer(
        this.settings,
        this.app.vault.adapter,
        file.parent?.path || '',
      )
      const processed = await executer.execute(content)

      try {
        this.updatingFiles.add(file.path)
        await file.vault.modify(file, processed)
      } finally {
        this.updatingFiles.delete(file.path)
      }
    }
  }

  async loadSettings() {
    this.settings = Object.assign({}, DEFAULT_SETTINGS, await this.loadData())
  }

  async saveSettings() {
    await this.saveData(this.settings)

    this.app.workspace.iterateAllLeaves((leaf: WorkspaceLeaf) => {
      // @ts-expect-error wrong types
      leaf?.rebuildView?.()
    })
  }
}
