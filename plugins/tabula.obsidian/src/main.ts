import { MarkdownView, Notice, Plugin, TFile, WorkspaceLeaf } from 'obsidian'

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

  onunload() {
    // @ts-expect-error wrong types
    delete CodeMirror.modes['tabula']
  }

  async onload() {
    await this.loadSettings()
    highlightSyntax(this)

    this.addCommand({
      id: 'tabula-execute',
      name: 'execute',
      checkCallback: (checking: boolean) => {
        const file = this.app.workspace.getActiveFile()
        if (file instanceof TFile && file.extension === 'md') {
          if (!checking) {
            this.executeOnFile(file)
          }
          return true
        }
        return false
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
    if (this.updatingFiles.has(file.path)) return
    this.updatingFiles.add(file.path)

    try {
      await this.runExecution(file)
    } finally {
      this.updatingFiles.delete(file.path)
    }
  }

  private async runExecution(file: TFile): Promise<void> {
    const executer = new Executer(this.settings, file)
    const originalContent = await this.app.vault.read(file)

    let processed: string
    try {
      processed = await executer.execute()
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

    const activeView = this.app.workspace.getActiveViewOfType(MarkdownView)
    const isActiveFile = activeView?.file?.path === file.path
    const cursor = isActiveFile ? activeView!.editor.getCursor() : null

    // vault.process ensures the write is serialized with other vault operations.
    // If the file was edited during execution, preserve user changes.
    await this.app.vault.process(file, (currentContent) => {
      if (currentContent !== originalContent) return currentContent
      return processed
    })

    if (isActiveFile) {
      requestAnimationFrame(() => {
        const view = this.app.workspace.getActiveViewOfType(MarkdownView)
        if (view?.file?.path !== file.path) return
        if (cursor) view.editor.setCursor(cursor)
        view.editor.focus()
      })
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
