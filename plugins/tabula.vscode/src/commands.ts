import * as vscode from 'vscode'
import { runScript } from './runner'
import { reloadDocument } from './reload'

export const commandExecute = async (..._: string[]) => {
  const editor = vscode.window.activeTextEditor
  if (!editor) {
    vscode.window.showErrorMessage('No active editor found')
    return
  }

  const document = editor.document
  // Check if it's a CSV or Markdown file
  if (!['csv', 'md'].includes(document.languageId.toLowerCase())) {
    vscode.window.showErrorMessage(
      'Current file is not a CSV nor Markdown file.',
    )
    return
  }

  // Save the document if it has unsaved changes
  if (document.isDirty) {
    await document.save()
  }

  try {
    // Save cursor position
    const cursorPosition = editor.selection.active

    await runScript(document.uri)

    // Reload the document from disk to show changes
    await reloadDocument(document.uri)

    // Restore cursor position
    const newPosition = new vscode.Position(
      Math.min(cursorPosition.line, editor.document.lineCount - 1),
      cursorPosition.character,
    )
    editor.selection = new vscode.Selection(newPosition, newPosition)
    editor.revealRange(
      new vscode.Range(newPosition, newPosition),
      vscode.TextEditorRevealType.InCenterIfOutsideViewport,
    )
  } catch (error) {
    vscode.window.showErrorMessage(
      `Tabula execution failed: ${error instanceof Error ? error.message : String(error)}`,
    )
  }
}

export const commandToggleAutoExecution = () => {
  const config = vscode.workspace.getConfiguration('tabula')
  const autoExecution = !config.get<boolean>('autoExecution', true)
  config.update(
    'autoExecution',
    autoExecution,
    vscode.ConfigurationTarget.Global,
  )
  vscode.window.showInformationMessage(
    `Tabula auto-execution ${autoExecution ? 'enabled' : 'disabled'}`,
  )
}

export const commandAutoExecute = async (document: vscode.TextDocument) => {
  // Check configuration
  const config = vscode.workspace.getConfiguration('tabula')
  const autoExecution = config.get<boolean>('autoExecution', true)

  if (!autoExecution) {
    return
  }

  // Check if it's a CSV file and if it's the active document
  const editor = vscode.window.activeTextEditor
  if (
    editor &&
    editor.document.uri.toString() === document.uri.toString() &&
    ['md', 'csv'].includes(document.languageId.toLowerCase())
  ) {
    // Execute the tabula.execute command
    await vscode.commands.executeCommand('tabula.execute')
  }
}
