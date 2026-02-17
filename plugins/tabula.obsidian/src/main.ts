import { App, Plugin, PluginManifest, TFile } from 'obsidian'
import { TabulaSettings, DEFAULT_SETTINGS } from './types'
import { TabulaSettingTab } from './settings'
import { extractChunks, processChunks, outputChunks } from './processor'
import { Executer } from './executer'
import { renderTable } from './renderTable'
import { renderCode } from './renderCode'

export default class TabulaPlugin extends Plugin {
  settings: TabulaSettings = DEFAULT_SETTINGS
  private updatingFiles = new Set<string>()

  constructor(app: App, manifest: PluginManifest) {
    super(app, manifest)
  }

  async onload() {
    await this.loadSettings()

    this.registerMarkdownCodeBlockProcessor('csv', (source, el, _ctx) => {
      renderTable(this.settings, source, el)
    })

    this.registerMarkdownCodeBlockProcessor('tabula', (source, el, _ctx) => {
      renderCode(this.settings, source, el)
    })

    // Listen for markdown file modifications
    this.registerEvent(
      this.app.vault.on('modify', async (file) => {
        if (!this.settings.autoExecute) {
          return
        }

        if (this.updatingFiles.has(file.path)) return

        // Only process markdown files
        if (file instanceof TFile && file.extension === 'md') {
          const content = await this.app.vault.read(file)
          const chunks = extractChunks(content)

          const executer = new Executer(
            this.settings,
            this.app.vault.adapter,
            file.parent?.path || '',
          )
          const processed = await processChunks(executer, chunks)
          const output = outputChunks(processed)

          this.updatingFiles.add(file.path)
          try {
            await file.vault.modify(file, output)
          } finally {
            this.updatingFiles.delete(file.path)
          }
        }
      }),
    )

    // Add settings tab
    this.addSettingTab(new TabulaSettingTab(this.app, this))
  }

  async loadSettings() {
    this.settings = Object.assign({}, DEFAULT_SETTINGS, await this.loadData())
  }

  async saveSettings() {
    await this.saveData(this.settings)
  }
}
