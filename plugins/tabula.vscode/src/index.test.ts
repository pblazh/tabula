import * as assert from 'assert'
import * as vscode from 'vscode'
import * as path from 'path'
import * as fs from 'fs'

// Helper to wait for config change
const waitForConfigChange = (
  field: string,
  expectedValue: unknown,
  timeout = 5000,
): Promise<boolean> => {
  return new Promise((resolve) => {
    const startTime = Date.now()
    const check = () => {
      const config = vscode.workspace.getConfiguration('tabula')
      const currentValue = config.get<boolean>(field, true)
      if (currentValue === expectedValue) {
        resolve(true)
      } else if (Date.now() - startTime > timeout) {
        console.error(`Timeout waiting for ${field} to become ${expectedValue}`)
        resolve(false)
      } else {
        setTimeout(check, 50)
      }
    }
    check()
  })
}

suite('Tabula Extension Tests', () => {
  let extensionId: string

  suiteSetup(() => {
    // Find the extension
    const tabulaExtension = vscode.extensions.all.find(
      (ext) => ext.packageJSON.name === 'tabula',
    )

    if (tabulaExtension) {
      extensionId = tabulaExtension.id
    } else {
      extensionId = 'tabula.tabula'
    }
  })

  test('Should have correct package.json metadata', async () => {
    const extension = vscode.extensions.getExtension(extensionId)
    assert.ok(extension, 'Extension should exist for metadata test')

    const packageJSON = extension!.packageJSON

    assert.strictEqual(packageJSON.name, 'tabula')
    assert.strictEqual(packageJSON.displayName, 'Tabula')
    assert.ok(packageJSON.version)
    assert.ok(packageJSON.description)

    // Check commands are registered
    const commands = packageJSON.contributes.commands
    assert.ok(
      commands.find(
        (cmd: { command: string }) =>
          cmd.command === 'tabula.toggleAutoExecute',
      ),
      'toggleAutoExecute command should be in package.json',
    )

    assert.ok(
      commands.find(
        (cmd: { command: string }) => cmd.command === 'tabula.execute',
      ),
      'execute command should be in package.json',
    )

    // Check configuration is defined
    const properties = packageJSON.contributes.configuration.properties
    assert.ok(properties['tabula.autoExecute'])
    assert.ok(properties['tabula.executablePath'])
    assert.ok(properties['tabula.autoFormat'])
  })

  test('Extension should activate', async () => {
    const extension = vscode.extensions.getExtension(extensionId)

    await extension!.activate()
    assert.strictEqual(extension!.isActive, true, 'Extension should be active')
  })

  test('Should have correct configuration defaults', () => {
    const config = vscode.workspace.getConfiguration('tabula')

    const autoExecuteInspect = config.inspect('autoExecute')
    assert.ok(autoExecuteInspect, 'autoExecute configuration should exist')
    assert.ok(
      autoExecuteInspect.defaultValue,
      'autoExecute should default to true',
    )

    const executablePathInspect = config.inspect('executablePath')
    assert.ok(
      executablePathInspect,
      'executablePath configuration should exist',
    )
    assert.strictEqual(
      executablePathInspect.defaultValue,
      'tabula',
      "executablePath should default to 'tabula'",
    )

    const autoFormatInspect = config.inspect('autoFormat')
    assert.ok(autoFormatInspect, 'autoFormat configuration should exist')
    assert.strictEqual(
      autoFormatInspect.defaultValue,
      true,
      'autoFormat should default to true',
    )
  })

  test('Should toggle autoExecute setting', async () => {
    let config = vscode.workspace.getConfiguration('tabula')
    const initialValue = config.get<boolean>('autoExecute', true)

    // Execute toggle command
    await vscode.commands.executeCommand('tabula.toggleAutoExecute')

    // Wait for config to actually change
    await waitForConfigChange('autoExecute', !initialValue)

    // Verify the change
    config = vscode.workspace.getConfiguration('tabula')
    const newValue = config.get<boolean>('autoExecute', true)
    assert.strictEqual(
      newValue,
      !initialValue,
      `autoExecute should be ${!initialValue} after toggle`,
    )

    // Toggle back
    await vscode.commands.executeCommand('tabula.toggleAutoExecute')

    // Wait for config to change back
    await waitForConfigChange('autoExecute', initialValue)

    // Verify it's restored
    config = vscode.workspace.getConfiguration('tabula')
    const restoredValue = config.get<boolean>('autoExecute', true)
    assert.strictEqual(
      restoredValue,
      initialValue,
      `autoExecute should toggle back to ${initialValue}`,
    )
  })

  test('Should update executablePath configuration', async () => {
    let config = vscode.workspace.getConfiguration('tabula')
    const customPath = '/custom/path/to/tabula'

    await config.update(
      'executablePath',
      customPath,
      vscode.ConfigurationTarget.Global,
    )

    await waitForConfigChange('executablePath', customPath)
    // Re-get configuration to see changes
    config = vscode.workspace.getConfiguration('tabula')
    const updatedPath = config.get<string>('executablePath', 'tabula')
    assert.strictEqual(
      updatedPath,
      customPath,
      'executablePath should update correctly',
    )

    // Restore default
    await config.update(
      'executablePath',
      'tabula',
      vscode.ConfigurationTarget.Global,
    )
  })

  test('Should register tabula.execute command', async () => {
    const commands = await vscode.commands.getCommands(true)
    assert.ok(
      commands.includes('tabula.execute'),
      'tabula.execute command should be registered',
    )
  })

  test('Should register tabula.toggleAutoExecute command', async () => {
    const commands = await vscode.commands.getCommands(true)
    assert.ok(
      commands.includes('tabula.toggleAutoExecute'),
      'tabula.toggleAutoExecute command should be registered',
    )
  })

  test('Should update autoFormat configuration', async () => {
    let config = vscode.workspace.getConfiguration('tabula')

    // Test disabling auto format
    await config.update('autoFormat', false, vscode.ConfigurationTarget.Global)

    // Re-get configuration to see changes
    await waitForConfigChange('autoFormat', false)
    config = vscode.workspace.getConfiguration('tabula')
    const autoFormat = config.get<boolean>('autoFormat', true)
    assert.strictEqual(autoFormat, false, 'autoFormat should be disabled')

    // Re-enable
    await config.update('autoFormat', true, vscode.ConfigurationTarget.Global)

    // Verify it's enabled again
    await waitForConfigChange('autoFormat', true)
    config = vscode.workspace.getConfiguration('tabula')
    const autoFormatEnabled = config.get<boolean>('autoFormat', true)
    assert.strictEqual(autoFormatEnabled, true, 'autoFormat should be enabled')
  })

  test('Should handle CSV file detection', async () => {
    const testFilePath = path.join(__dirname, '..', 'example', 'input.csv')

    assert.ok(fs.existsSync(testFilePath), 'Test CSV file not found')

    const document = await vscode.workspace.openTextDocument(testFilePath)

    // Verify the file is recognized as a CSV by extension
    assert.equal(
      document.languageId,
      'csv',
      `Document languageId should be "csv" but got "${document.languageId}"`,
    )
  })

  test('When CSV file is saved with autoExecute enabled, tabula.execute command is called', async () => {
    // Ensure autoExecute is enabled
    const config = vscode.workspace.getConfiguration('tabula')
    await config.update('autoExecute', true, vscode.ConfigurationTarget.Global)

    // Verify autoExecute is enabled
    const autoExecute = config.get<boolean>('autoExecute')
    assert.strictEqual(
      autoExecute,
      true,
      'autoExecute should be enabled for this test',
    )

    // Verify that tabula.execute command is registered
    const commands = await vscode.commands.getCommands(true)
    assert.ok(
      commands.includes('tabula.execute'),
      'tabula.execute command should be registered',
    )

    // Create a test CSV file path
    const testFilePath = path.join(__dirname, '..', 'example', 'input.csv')

    // Track if tabula.execute was called
    let executeCommandCalled: boolean
    let commandCallCount = 0

    // Wrap the executeCommand to track calls
    const originalExecute = vscode.commands.executeCommand

    ;(
      vscode.commands as {
        executeCommand: (command: string, ...args: unknown[]) => void
      }
    ).executeCommand = function (command: string, ...args: unknown[]) {
      if (command === 'tabula.execute') {
        executeCommandCalled = true
        commandCallCount++
        // Don't actually execute - return a resolved promise
        return Promise.resolve()
      }
      // Pass through other commands
      return originalExecute.call(this, command, ...args)
    }

    try {
      // Open the test CSV file
      const document = await vscode.workspace.openTextDocument(testFilePath)
      const editor = await vscode.window.showTextDocument(document)

      // Verify the document is the active editor
      assert.ok(editor, 'There should be an active editor')
      assert.strictEqual(
        editor.document.uri.toString(),
        document.uri.toString(),
        'The CSV document should be the active editor',
      )

      // Make the document dirty by editing it
      const success = await editor.edit((editBuilder) => {
        editBuilder.insert(new vscode.Position(0, 0), ' ')
      })
      assert.ok(success, 'Edit should succeed')

      // Reset tracking
      executeCommandCalled = false
      commandCallCount = 0

      // Save the document - this should trigger tabula.execute
      const saved = await document.save()
      assert.ok(saved, 'Document should save successfully')

      // Wait for async save handlers to complete
      await new Promise((resolve) => setTimeout(resolve, 1000))

      // Verify that tabula.execute was called
      assert.ok(
        executeCommandCalled,
        'tabula.execute command should be called when CSV file is saved',
      )
      assert.ok(
        commandCallCount > 0,
        `tabula.execute should be called at least once (called ${commandCallCount} times)`,
      )

      // Restore the file by removing the inserted space
      const restoreSuccess = await editor.edit((editBuilder) => {
        editBuilder.delete(
          new vscode.Range(
            new vscode.Position(0, 0),
            new vscode.Position(0, 1),
          ),
        )
      })
      assert.ok(restoreSuccess, 'Restore edit should succeed')

      // Save the restored document
      await document.save()
    } finally {
      // Restore original executeCommand
      ;(
        vscode.commands as {
          executeCommand: (command: string, ...args: unknown[]) => void
        }
      ).executeCommand = originalExecute
    }
  })
})
