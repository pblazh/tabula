import {
  App,
  Notice,
  Plugin,
  PluginManifest,
  TFile,
  WorkspaceLeaf,
} from 'obsidian'

import { TabulaSettings, DEFAULT_SETTINGS } from './types'
import { TabulaSettingTab } from './settings'
import { Executer } from './executer'
import { renderTable } from './renderTable'
import { renderCode } from './renderCode'
import { highlightSyntax } from './highlightSyntax'

export default class TabulaPlugin extends Plugin {
  settings: TabulaSettings = DEFAULT_SETTINGS
  private updatingFiles = new Set<string>()
  private lastErrorNotice: Notice | null = null

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
    highlightSyntax(this)

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
      id: 'tabula-toggle',
      name: 'toggle auto execution',
      callback: () => {
        this.settings.autoExecution = !this.settings.autoExecution
        this.saveSettings()
      },
    })

    this.addCommand({
      id: 'tabula-index',
      name: 'toggle rows and columns names visibility',
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
      let processed: string
      try {
        processed = await executer.execute(content)
      } catch (err) {
        const error = err instanceof Error ? err : new Error(String(err))
        const isNotFound = (error as NodeJS.ErrnoException).code === 'ENOENT'
        const message = isNotFound
          ? `Tabula: executable not found at "${this.settings.executablePath}". Check the path in settings.`
          : `Tabula: ${error.message}`
        this.lastErrorNotice?.hide()
        this.lastErrorNotice = new Notice(message, 0)
        return
      }

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
