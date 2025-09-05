export const InvalidSettingsError = (errors: string[]) =>
  `Settings validation warnings: ${errors.join(", ")}`;

export const PluginFaildedToStartError = (error: Error) =>
  `Failed to load Tabula plugin: ${error.message}`;
