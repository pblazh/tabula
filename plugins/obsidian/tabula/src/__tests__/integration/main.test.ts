import Tabula from "../../main";
import { validateSettings } from "../../services/validateSettings";
import { showWarningNotice } from "../../utils/error-utils";
import { I18n } from "../../i18n";

// Mock dependencies
jest.mock("../../services/csv-service");
jest.mock("../../utils/error-utils");
jest.mock("../../i18n", () => ({
  I18n: {
    setLocale: jest.fn(),
  },
}));

const mockValidateCSVSettings = validateSettings as jest.MockedFunction<
  typeof validateSettings
>;
const mockShowWarningNotice = showWarningNotice as jest.MockedFunction<
  typeof showWarningNotice
>;

// Mock Obsidian APIs
const mockApp = (global as any).createMockApp();

describe("Main Plugin Integration", () => {
  let plugin: Tabula;

  beforeEach(() => {
    jest.clearAllMocks();
    plugin = new Tabula(mockApp, {
      id: "tabula",
      name: "Tabula",
      version: "1.0.0",
      dir: "/test/path",
    } as any);
  });

  describe("Plugin lifecycle", () => {
    it("should load successfully with valid settings", async () => {
      // Mock successful settings validation
      mockValidateCSVSettings.mockReturnValue([]);

      // Mock successful data loading
      plugin.loadData = jest.fn().mockResolvedValue({
        delimiter: ",",
        comment: "#",
      });

      plugin.registerView = jest.fn();
      plugin.registerExtensions = jest.fn();
      plugin.addSettingTab = jest.fn();

      await plugin.onload();

      expect(plugin.registerView).toHaveBeenCalledTimes(2); // CSV and source views
      expect(plugin.registerExtensions).toHaveBeenCalledWith(
        ["csv"],
        "csv-view",
      );
      expect(plugin.addSettingTab).toHaveBeenCalled();
      expect(mockValidateCSVSettings).toHaveBeenCalledWith(plugin.settings);
      expect(mockShowWarningNotice).not.toHaveBeenCalled();
    });

    it("should show warnings for invalid settings on load", async () => {
      const validationErrors = ["Delimiter cannot be empty"];
      mockValidateCSVSettings.mockReturnValue(validationErrors);

      plugin.loadData = jest.fn().mockResolvedValue({});
      plugin.registerView = jest.fn();
      plugin.registerExtensions = jest.fn();
      plugin.addSettingTab = jest.fn();

      await plugin.onload();

      expect(mockShowWarningNotice).toHaveBeenCalledWith(
        "Settings validation warnings: Delimiter cannot be empty",
      );
    });

    it("should handle load errors gracefully", async () => {
      plugin.loadData = jest.fn().mockRejectedValue(new Error("Load failed"));
      plugin.registerView = jest.fn();
      plugin.registerExtensions = jest.fn();
      plugin.addSettingTab = jest.fn();

      await plugin.onload();

      expect(mockShowWarningNotice).toHaveBeenCalledWith(
        expect.stringContaining("Failed to load Tabula plugin"),
      );
    });

    it("should use default settings if loading fails", async () => {
      plugin.loadData = jest.fn().mockRejectedValue(new Error("Load failed"));

      await plugin.loadSettings();

      expect(plugin.settings).toEqual({
        tabula: "tabula",
        delimiter: ",",
        comment: "#",
        quote: '"',
      });
    });

    it("should register both view types", async () => {
      mockValidateCSVSettings.mockReturnValue([]);
      plugin.loadData = jest.fn().mockResolvedValue({});
      plugin.registerView = jest.fn();
      plugin.registerExtensions = jest.fn();
      plugin.addSettingTab = jest.fn();

      await plugin.onload();

      expect(plugin.registerView).toHaveBeenCalledWith(
        "csv-view",
        expect.any(Function),
      );
      expect(plugin.registerView).toHaveBeenCalledWith(
        "csv-source-view",
        expect.any(Function),
      );
    });
  });

  describe("Settings management", () => {
    it("should load settings successfully", async () => {
      const settingsData = {
        tabula: "/usr/bin/tabula",
        delimiter: ";",
        comment: "//",
        quote: "'",
      };

      plugin.loadData = jest.fn().mockResolvedValue(settingsData);

      await plugin.loadSettings();

      expect(plugin.settings).toEqual(settingsData);
    });

    it("should merge with default settings", async () => {
      const partialSettings = {
        delimiter: ";",
      };

      plugin.loadData = jest.fn().mockResolvedValue(partialSettings);

      await plugin.loadSettings();

      expect(plugin.settings).toEqual({
        tabula: "tabula", // default
        delimiter: ";", // overridden
        comment: "#", // default
        quote: '"', // default
      });
    });

    it("should handle empty loaded data", async () => {
      plugin.loadData = jest.fn().mockResolvedValue(null);

      await plugin.loadSettings();

      expect(plugin.settings).toEqual({
        tabula: "tabula",
        delimiter: ",",
        comment: "#",
        quote: '"',
      });
    });

    it("should save settings successfully", async () => {
      plugin.saveData = jest.fn().mockResolvedValue(undefined);
      plugin.settings = {
        tabula: "custom-tabula",
        delimiter: ";",
        comment: "//",
        quote: "'",
      };

      await plugin.saveSettings();

      expect(plugin.saveData).toHaveBeenCalledWith(plugin.settings);
    });

    it("should handle save errors", async () => {
      plugin.saveData = jest.fn().mockRejectedValue(new Error("Save failed"));

      await expect(plugin.saveSettings()).rejects.toThrow("Save failed");
    });
  });

  describe("View factories", () => {
    beforeEach(async () => {
      mockValidateCSVSettings.mockReturnValue([]);
      plugin.loadData = jest.fn().mockResolvedValue({});
      plugin.registerView = jest.fn();
      plugin.registerExtensions = jest.fn();
      plugin.addSettingTab = jest.fn();

      await plugin.onload();
    });

    it("should create table view with settings", () => {
      const mockLeaf = { view: null } as any;
      const registerViewCall = (
        plugin.registerView as jest.Mock
      ).mock.calls.find((call) => call[0] === "csv-view");

      expect(registerViewCall).toBeDefined();

      const viewFactory = registerViewCall[1];
      const view = viewFactory(mockLeaf);

      expect(view).toBeDefined();
    });

    it("should create source view", () => {
      const mockLeaf = { view: null } as any;
      const registerViewCall = (
        plugin.registerView as jest.Mock
      ).mock.calls.find((call) => call[0] === "csv-source-view");

      expect(registerViewCall).toBeDefined();

      const viewFactory = registerViewCall[1];
      const view = viewFactory(mockLeaf);

      expect(view).toBeDefined();
    });
  });

  describe("File extension registration", () => {
    it("should register CSV extension", async () => {
      mockValidateCSVSettings.mockReturnValue([]);
      plugin.loadData = jest.fn().mockResolvedValue({});
      plugin.registerView = jest.fn();
      plugin.registerExtensions = jest.fn();
      plugin.addSettingTab = jest.fn();

      await plugin.onload();

      expect(plugin.registerExtensions).toHaveBeenCalledWith(
        ["csv"],
        "csv-view",
      );
    });
  });

  describe("Internationalization", () => {
    it("should set locale from Obsidian", async () => {
      const mockMoment = { locale: jest.fn().mockReturnValue("fr") };
      (global as any).moment = mockMoment;

      mockValidateCSVSettings.mockReturnValue([]);
      plugin.loadData = jest.fn().mockResolvedValue({});
      plugin.registerView = jest.fn();
      plugin.registerExtensions = jest.fn();
      plugin.addSettingTab = jest.fn();

      await plugin.onload();

      expect(I18n.setLocale).toHaveBeenCalledWith("fr");
    });
  });

  describe("Error handling", () => {
    it("should not crash on view factory errors", async () => {
      mockValidateCSVSettings.mockReturnValue([]);
      plugin.loadData = jest.fn().mockResolvedValue({});
      plugin.registerView = jest.fn();
      plugin.registerExtensions = jest.fn();
      plugin.addSettingTab = jest.fn();

      await plugin.onload();

      // Get the view factory
      const tableViewFactory = (
        plugin.registerView as jest.Mock
      ).mock.calls.find((call) => call[0] === "csv-view")[1];

      // This should not crash even with invalid leaf
      expect(() => tableViewFactory(null)).not.toThrow();
    });

    it("should handle registration errors gracefully", async () => {
      mockValidateCSVSettings.mockReturnValue([]);
      plugin.loadData = jest.fn().mockResolvedValue({});
      plugin.registerView = jest.fn().mockImplementation(() => {
        throw new Error("Registration failed");
      });
      plugin.registerExtensions = jest.fn();
      plugin.addSettingTab = jest.fn();

      await plugin.onload();

      expect(mockShowWarningNotice).toHaveBeenCalledWith(
        expect.stringContaining("Failed to load Tabula plugin"),
      );
    });

    it("should handle multiple validation errors", async () => {
      const validationErrors = [
        "Delimiter cannot be empty",
        "Delimiter and quote character cannot be the same",
      ];
      mockValidateCSVSettings.mockReturnValue(validationErrors);

      plugin.loadData = jest.fn().mockResolvedValue({});
      plugin.registerView = jest.fn();
      plugin.registerExtensions = jest.fn();
      plugin.addSettingTab = jest.fn();

      await plugin.onload();

      expect(mockShowWarningNotice).toHaveBeenCalledWith(
        "Settings validation warnings: Delimiter cannot be empty, Delimiter and quote character cannot be the same",
      );
    });
  });

  describe("Plugin metadata", () => {
    it("should have correct default settings", () => {
      expect(plugin.settings).toBeUndefined(); // Not loaded yet

      // Test default settings structure
      const expectedDefaults = {
        tabula: "tabula",
        delimiter: ",",
        comment: "#",
        quote: '"',
      };

      // We can't access DEFAULT_SETTINGS directly, but we can verify through loading
      plugin.loadData = jest.fn().mockResolvedValue({});
      return plugin.loadSettings().then(() => {
        expect(plugin.settings).toEqual(expectedDefaults);
      });
    });
  });
});

