import { Plugin, WorkspaceLeaf } from 'obsidian'

export function highlightSyntax(self: Plugin) {
  // The defineSimpleMode function is not immediately available during
  // onload, so continue to try and define the language until it is.
  const setupInterval = self.registerInterval(
    window.setInterval(() => {
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
                /\b(SUM|AVERAGE|MIN|MAX|ABS|ROUND|SQRT|MOD|POWER|FLOOR|CEIL|TRUNC|INT|SIGN|RAND|RANDBETWEEN|GCD|LCM|CONCATENATE|LEFT|RIGHT|MID|UPPER|LOWER|TRIM|LEN|SUBSTITUTE|FIND|REPLACE|TEXT|VALUE|SPLIT|JOIN|DATE|DATEVALUE|YEAR|MONTH|DAY|HOUR|MINUTE|SECOND|NOW|TODAY|DATEDIF|WEEKDAY|WORKDAY|NETWORKDAYS|TIME|TIMEVALUE|IF|AND|OR|NOT|TRUE|FALSE|ISNUMBER|ISTEXT|ISLOGICAL|ISBLANK|ISERROR|COLUMN|ROW|COLUMNS|ROWS|ADDRESS|REF|INDEX|MATCH|VLOOKUP|HLOOKUP|EXEC|COUNT|COUNTA|COUNTBLANK|SUMIF|AVERAGEIF|MAXIFS|MINIFS)\b/i,
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
    }, 100),
  )
}
