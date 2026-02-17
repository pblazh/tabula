import * as vscode from 'vscode'

export const reloadDocument = async (uri: vscode.Uri): Promise<void> => {
  // Find all text editors showing this document
  const editors = vscode.window.visibleTextEditors.filter(
    (editor) => editor.document.uri.toString() === uri.toString(),
  )

  if (editors.length === 0) {
    return
  }

  // Close the document
  await vscode.commands.executeCommand('workbench.action.closeActiveEditor')

  // Small delay to ensure file system changes are visible
  await new Promise((resolve) => setTimeout(resolve, 100))

  // Reopen the document
  const document = await vscode.workspace.openTextDocument(uri)
  await vscode.window.showTextDocument(document, {
    preview: false,
    preserveFocus: false,
  })
}
