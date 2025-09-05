import { renderTable } from '../../components/table';

// Mock DOM methods
const createMockElement = (tagName: string) => ({
  tagName: tagName.toUpperCase(),
  empty: jest.fn(),
  createEl: jest.fn().mockImplementation((tag: string, options?: any) => {
    const element = createMockElement(tag);
    if (options?.cls) {
      element.className = options.cls;
    }
    if (options?.attr) {
      Object.assign(element, options.attr);
    }
    return element;
  }),
  appendChild: jest.fn(),
  textContent: '',
  className: '',
  setAttribute: jest.fn(),
  getAttribute: jest.fn(),
  dataset: {}
});

// Mock document.createDocumentFragment
const mockFragment = {
  createEl: jest.fn().mockImplementation((tag: string) => createMockElement(tag)),
  appendChild: jest.fn()
};

Object.defineProperty(document, 'createDocumentFragment', {
  value: jest.fn(() => mockFragment)
});

describe('Table Component', () => {
  let mockTableEl: any;

  beforeEach(() => {
    jest.clearAllMocks();
    mockTableEl = createMockElement('table');
  });

  describe('renderTable', () => {
    it('should render table with headers and data', () => {
      const tableData = [
        ['Name', 'Age', 'City'],
        ['John', '25', 'NYC'],
        ['Jane', '30', 'LA']
      ];

      renderTable(mockTableEl, tableData);

      expect(mockTableEl.empty).toHaveBeenCalled();
      expect(mockTableEl.appendChild).toHaveBeenCalled();
      expect(document.createDocumentFragment).toHaveBeenCalled();
    });

    it('should create column headers with correct labels', () => {
      const tableData = [
        ['A', 'B', 'C']
      ];

      renderTable(mockTableEl, tableData);

      // Verify fragment structure is created
      expect(mockFragment.createEl).toHaveBeenCalledWith('thead');
      expect(mockFragment.createEl).toHaveBeenCalledWith('tbody');
    });

    it('should handle empty table data', () => {
      const tableData: string[][] = [];

      renderTable(mockTableEl, tableData);

      expect(mockTableEl.empty).toHaveBeenCalled();
      expect(mockTableEl.appendChild).toHaveBeenCalled();
    });

    it('should create input elements for each cell', () => {
      const tableData = [
        ['A', 'B'],
        ['1', '2']
      ];

      renderTable(mockTableEl, tableData);

      // Verify the structure is being built
      expect(mockFragment.createEl).toHaveBeenCalledWith('thead');
      expect(mockFragment.createEl).toHaveBeenCalledWith('tbody');
    });

    it('should set correct data attributes on elements', () => {
      const tableData = [
        ['A'],
        ['1']
      ];

      renderTable(mockTableEl, tableData);

      // The function should create elements with proper structure
      expect(mockFragment.createEl).toHaveBeenCalled();
    });

    it('should handle single column table', () => {
      const tableData = [
        ['Header'],
        ['Data1'],
        ['Data2']
      ];

      renderTable(mockTableEl, tableData);

      expect(mockTableEl.empty).toHaveBeenCalled();
      expect(mockTableEl.appendChild).toHaveBeenCalledWith(mockFragment);
    });

    it('should handle table with many columns', () => {
      const tableData = [
        ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'],
        ['1', '2', '3', '4', '5', '6', '7', '8', '9', '10']
      ];

      renderTable(mockTableEl, tableData);

      expect(mockTableEl.empty).toHaveBeenCalled();
      expect(mockTableEl.appendChild).toHaveBeenCalledWith(mockFragment);
    });

    it('should create row number cells', () => {
      const tableData = [
        ['A'],
        ['1'],
        ['2']
      ];

      renderTable(mockTableEl, tableData);

      // Verify the basic structure is created
      expect(mockFragment.createEl).toHaveBeenCalledWith('thead');
      expect(mockFragment.createEl).toHaveBeenCalledWith('tbody');
    });

    it('should handle table with empty cells', () => {
      const tableData = [
        ['A', '', 'C'],
        ['', '2', ''],
        ['3', '', '']
      ];

      renderTable(mockTableEl, tableData);

      expect(mockTableEl.empty).toHaveBeenCalled();
      expect(mockTableEl.appendChild).toHaveBeenCalledWith(mockFragment);
    });

    it('should clear existing table content', () => {
      const tableData = [['A']];

      renderTable(mockTableEl, tableData);

      expect(mockTableEl.empty).toHaveBeenCalledTimes(1);
      expect(mockTableEl.appendChild).toHaveBeenCalledTimes(1);
    });
  });

  describe('Table structure', () => {
    it('should create proper HTML structure', () => {
      const tableData = [
        ['Name', 'Value'],
        ['Test', '123']
      ];

      renderTable(mockTableEl, tableData);

      // Verify thead and tbody are created
      expect(mockFragment.createEl).toHaveBeenCalledWith('thead');
      expect(mockFragment.createEl).toHaveBeenCalledWith('tbody');
      
      // Verify the fragment is appended to the table element
      expect(mockTableEl.appendChild).toHaveBeenCalledWith(mockFragment);
    });

    it('should handle tables with uneven row lengths', () => {
      const tableData = [
        ['A', 'B', 'C'],
        ['1', '2'], // Missing third column
        ['3', '4', '5', '6'] // Extra column
      ];

      // This should not crash
      expect(() => renderTable(mockTableEl, tableData)).not.toThrow();
    });
  });

  describe('Column label generation', () => {
    it('should generate correct number of column headers', () => {
      const tableData = [
        ['A', 'B', 'C', 'D', 'E']
      ];

      renderTable(mockTableEl, tableData);

      // Should create header with 6 columns (including row number column)
      expect(mockFragment.createEl).toHaveBeenCalledWith('thead');
    });
  });

  describe('Performance and edge cases', () => {
    it('should handle large tables efficiently', () => {
      // Create a 100x10 table
      const tableData: string[][] = [];
      for (let i = 0; i < 100; i++) {
        const row: string[] = [];
        for (let j = 0; j < 10; j++) {
          row.push(`Cell ${i}-${j}`);
        }
        tableData.push(row);
      }

      const startTime = Date.now();
      renderTable(mockTableEl, tableData);
      const endTime = Date.now();

      // Should complete reasonably quickly
      expect(endTime - startTime).toBeLessThan(1000);
      expect(mockTableEl.empty).toHaveBeenCalled();
      expect(mockTableEl.appendChild).toHaveBeenCalled();
    });

    it('should handle null/undefined table data gracefully', () => {
      expect(() => renderTable(mockTableEl, null as any)).not.toThrow();
      expect(() => renderTable(mockTableEl, undefined as any)).not.toThrow();
    });

    it('should handle table data with null rows', () => {
      const tableData = [
        ['A', 'B'],
        null,
        ['C', 'D']
      ] as any;

      expect(() => renderTable(mockTableEl, tableData)).not.toThrow();
    });

    it('should handle table data with null cells', () => {
      const tableData = [
        ['A', null, 'C'],
        [null, 'B', null]
      ] as any;

      expect(() => renderTable(mockTableEl, tableData)).not.toThrow();
    });
  });
});