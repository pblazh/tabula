import { parseCSVContent, unparseCSVContent } from "../../services/csv-service";
import { validateCSVSettings } from "../../services/validateSettings";
import { Settings } from "../../types";

describe("CSV Service", () => {
  const defaultSettings: Settings = {
    tabula: "tabula",
    delimiter: ",",
    quote: '"',
    comment: "#",
  };

  describe("parseCSVContent", () => {
    it("should parse simple CSV data", () => {
      const csvString = "name,age,city\nJohn,25,NYC\nJane,30,LA";
      const result = parseCSVContent(defaultSettings, csvString);

      expect(result.data).toEqual([
        ["name", "age", "city"],
        ["John", "25", "NYC"],
        ["Jane", "30", "LA"],
      ]);
      expect(result.comments).toEqual({});
      expect(result.errors).toBeUndefined();
    });

    it("should handle CSV with comments", () => {
      const csvString =
        "# This is a comment\nname,age\n# Another comment\nJohn,25";
      const result = parseCSVContent(defaultSettings, csvString);

      expect(result.data).toEqual([
        ["name", "age"],
        ["John", "25"],
      ]);
      expect(result.comments).toEqual({
        0: "# This is a comment",
        2: "# Another comment",
      });
    });

    it("should handle empty CSV string", () => {
      const result = parseCSVContent(defaultSettings, "");

      expect(result.data).toEqual([[""]]);
      expect(result.comments).toEqual({});
      expect(result.errors).toEqual(["Invalid CSV content"]);
    });

    it("should handle null/undefined input", () => {
      const result = parseCSVContent(defaultSettings, null as any);

      expect(result.data).toEqual([[""]]);
      expect(result.comments).toEqual({});
      expect(result.errors).toEqual(["Invalid CSV content"]);
    });

    it("should use custom delimiter", () => {
      const settings = { ...defaultSettings, delimiter: ";" };
      const csvString = "name;age;city\nJohn;25;NYC";
      const result = parseCSVContent(settings, csvString);

      expect(result.data).toEqual([
        ["name", "age", "city"],
        ["John", "25", "NYC"],
      ]);
    });

    it("should use custom comment prefix", () => {
      const settings = { ...defaultSettings, comment: "//" };
      const csvString = "// Comment\nname,age\nJohn,25";
      const result = parseCSVContent(settings, csvString);

      expect(result.comments).toEqual({
        0: "// Comment",
      });
      expect(result.data).toEqual([
        ["name", "age"],
        ["John", "25"],
      ]);
    });

    it("should trim whitespace from cells", () => {
      const csvString = " name , age , city \n John , 25 , NYC ";
      const result = parseCSVContent(defaultSettings, csvString);

      expect(result.data).toEqual([
        ["name", "age", "city"],
        ["John", "25", "NYC"],
      ]);
    });

    it("should handle quoted fields with commas", () => {
      const csvString = 'name,description\nJohn,"Lives in NYC, works in tech"';
      const result = parseCSVContent(defaultSettings, csvString);

      expect(result.data).toEqual([
        ["name", "description"],
        ["John", "Lives in NYC, works in tech"],
      ]);
    });

    it("should handle malformed CSV gracefully", () => {
      const csvString = "name,age\nJohn,25,extra,data\nJane";
      const result = parseCSVContent(defaultSettings, csvString);

      expect(result.data).toHaveLength(3);
      expect(result.errors).toBeUndefined(); // PapaParse handles this gracefully
    });
  });

  describe("unparseCSVContent", () => {
    it("should convert table data back to CSV string", () => {
      const tableData = [
        ["name", "age", "city"],
        ["John", "25", "NYC"],
        ["Jane", "30", "LA"],
      ];
      const comments = {};

      const result = unparseCSVContent(defaultSettings, tableData, comments);

      expect(result).toContain("name,age,city");
      expect(result).toContain("John,25,NYC");
      expect(result).toContain("Jane,30,LA");
    });

    it("should include comments in output", () => {
      const tableData = [
        ["name", "age"],
        ["John", "25"],
      ];
      const comments = {
        0: "# Header comment",
        1: "# Data comment",
      };

      const result = unparseCSVContent(defaultSettings, tableData, comments);

      expect(result).toContain("# Header comment");
      expect(result).toContain("# Data comment");
    });

    it("should handle empty table data", () => {
      const result = unparseCSVContent(defaultSettings, [], {});
      expect(result).toBe("");
    });

    it("should handle null/undefined table data", () => {
      const result = unparseCSVContent(defaultSettings, null as any, {});
      expect(result).toBe("");
    });

    it("should use custom delimiter in output", () => {
      const settings = { ...defaultSettings, delimiter: ";" };
      const tableData = [
        ["name", "age"],
        ["John", "25"],
      ];

      const result = unparseCSVContent(settings, tableData, {});

      expect(result).toContain("name;age");
      expect(result).toContain("John;25");
    });

    it("should handle fields that need quoting", () => {
      const tableData = [
        ["name", "description"],
        ["John", "Lives in NYC, works in tech"],
      ];

      const result = unparseCSVContent(defaultSettings, tableData, {});

      expect(result).toContain('"Lives in NYC, works in tech"');
    });

    it("should add remaining comments at end", () => {
      const tableData = [["name"], ["John"]];
      const comments = {
        5: "# Comment after data",
        10: "# Another comment",
      };

      const result = unparseCSVContent(defaultSettings, tableData, comments);

      expect(result).toContain("# Comment after data");
      expect(result).toContain("# Another comment");
    });

    it("should throw error for malformed data", () => {
      const malformedData = [null, undefined, "not an array"] as any;

      expect(() => {
        unparseCSVContent(defaultSettings, malformedData, {});
      }).toThrow("Failed to unparse CSV");
    });
  });

  describe("validateCSVSettings", () => {
    it("should return no errors for valid settings", () => {
      const errors = validateCSVSettings(defaultSettings);
      expect(errors).toEqual([]);
    });

    it("should detect empty delimiter", () => {
      const settings = { ...defaultSettings, delimiter: "" };
      const errors = validateCSVSettings(settings);

      expect(errors).toContain("Delimiter cannot be empty");
    });

    it("should detect delimiter same as quote", () => {
      const settings = { ...defaultSettings, delimiter: '"', quote: '"' };
      const errors = validateCSVSettings(settings);

      expect(errors).toContain(
        "Delimiter and quote character cannot be the same",
      );
    });

    it("should detect delimiter same as comment prefix", () => {
      const settings = { ...defaultSettings, delimiter: "#", comment: "#" };
      const errors = validateCSVSettings(settings);

      expect(errors).toContain(
        "Delimiter and comment prefix cannot be the same",
      );
    });

    it("should return multiple errors for multiple issues", () => {
      const settings = {
        ...defaultSettings,
        delimiter: "",
        quote: "#",
        comment: "#",
      };
      const errors = validateCSVSettings(settings);

      expect(errors).toHaveLength(2); // empty delimiter + delimiter/comment same
      expect(errors).toContain("Delimiter cannot be empty");
    });

    it("should handle edge case delimiters", () => {
      const settings1 = { ...defaultSettings, delimiter: " " }; // space
      const errors1 = validateCSVSettings(settings1);
      expect(errors1).toEqual([]);

      const settings2 = { ...defaultSettings, delimiter: "\t" }; // tab
      const errors2 = validateCSVSettings(settings2);
      expect(errors2).toEqual([]);
    });
  });

  describe("Integration tests", () => {
    it("should round-trip CSV data correctly", () => {
      const originalCSV =
        '# Comment\nname,age,city\nJohn,25,"New York, NY"\nJane,30,LA';

      // Parse the CSV
      const parsed = parseCSVContent(defaultSettings, originalCSV);

      // Unparse it back
      const unparsed = unparseCSVContent(
        defaultSettings,
        parsed.data,
        parsed.comments,
      );

      // Parse again to verify consistency
      const reparsed = parseCSVContent(defaultSettings, unparsed);

      expect(reparsed.data).toEqual(parsed.data);
      expect(reparsed.comments).toEqual(parsed.comments);
    });

    it("should handle complex CSV with mixed content", () => {
      const complexCSV = `# File header
name,age,"address info",notes
# User data starts here
John,25,"123 Main St, NYC","Has ""quotes"" in notes"
Jane,30,"456 Oak Ave
Second line",Multi-line address
# End of data`;

      const result = parseCSVContent(defaultSettings, complexCSV);

      expect(result.data).toHaveLength(3); // header + 2 data rows
      expect(Object.keys(result.comments)).toHaveLength(3); // 3 comment lines
      expect(result.errors).toBeUndefined();
    });
  });
});

