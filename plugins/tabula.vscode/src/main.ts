import * as vscode from 'vscode'
import {
  commandAutoExecute,
  commandExecute,
  commandToggleAutoExecution,
} from './commands'

export function activate(context: vscode.ExtensionContext) {
  // Auto-execute Tabula on CSV file save
  const autoExecuteDisposable =
    vscode.workspace.onDidSaveTextDocument(commandAutoExecute)

  context.subscriptions.push(autoExecuteDisposable)

  // Manual command to toggle auto-execution
  const toggleCommand = vscode.commands.registerCommand(
    'tabula.toggleAutoExecution',
    commandToggleAutoExecution,
  )

  // Manual command to execture tabula on a current file
  const executeCommand = vscode.commands.registerCommand(
    'tabula.execute',
    commandExecute,
  )

  context.subscriptions.push(toggleCommand)
  context.subscriptions.push(executeCommand)
}

export function deactivate() {}
