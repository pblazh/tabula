import { parseCSV, unparseCSV } from '../../utils/csv-utils';
import { parseCSVContent, unparseCSVContent } from '../../services/csv-service';
import { Settings } from '../../types';

// Mock the service functions
jest.mock('../../services/csv-service');
const mockParseCSVContent = parseCSVContent as jest.MockedFunction<typeof parseCSVContent>;
const mockUnparseCSVContent = unparseCSVContent as jest.MockedFunction<typeof unparseCSVContent>;

describe('CSV Utils (Deprecated Wrapper)', () => {
  const settings: Settings = {
    tabula: 'tabula',
    delimiter: ',',
    quote: '"',
    comment: '#'
  };

  beforeEach(() => {
    jest.clearAllMocks();
  });

  describe('parseCSV', () => {
    it('should call parseCSVContent service and return compatible format', () => {
      const csvString = 'name,age\nJohn,25';
      const mockResult = {
        data: [['name', 'age'], ['John', '25']],
        comments: { 0: '# comment' },
        errors: undefined
      };
      
      mockParseCSVContent.mockReturnValue(mockResult);
      
      const result = parseCSV(settings, csvString);
      
      expect(mockParseCSVContent).toHaveBeenCalledWith(settings, csvString);
      expect(result).toEqual({
        data: mockResult.data,
        comments: mockResult.comments
      });
    });

    it('should handle service errors gracefully', () => {
      const csvString = 'invalid,csv';
      const mockResult = {
        data: [['']],
        comments: {},
        errors: ['Parse error']
      };
      
      mockParseCSVContent.mockReturnValue(mockResult);
      
      const result = parseCSV(settings, csvString);
      
      expect(result).toEqual({
        data: mockResult.data,
        comments: mockResult.comments
      });
    });

    it('should pass through settings correctly', () => {
      const customSettings = {
        ...settings,
        delimiter: ';',
        comment: '//'
      };
      const csvString = 'a;b\n1;2';
      
      mockParseCSVContent.mockReturnValue({
        data: [['a', 'b'], ['1', '2']],
        comments: {},
      });
      
      parseCSV(customSettings, csvString);
      
      expect(mockParseCSVContent).toHaveBeenCalledWith(customSettings, csvString);
    });

    it('should handle empty CSV input', () => {
      mockParseCSVContent.mockReturnValue({
        data: [['']],
        comments: {},
        errors: ['Invalid CSV content']
      });
      
      const result = parseCSV(settings, '');
      
      expect(result.data).toEqual([['']]);
      expect(result.comments).toEqual({});
    });

    it('should handle null/undefined input', () => {
      mockParseCSVContent.mockReturnValue({
        data: [['']],
        comments: {},
        errors: ['Invalid CSV content']
      });
      
      const result = parseCSV(settings, null as any);
      
      expect(mockParseCSVContent).toHaveBeenCalledWith(settings, null);
      expect(result.data).toEqual([['']]);
    });
  });

  describe('unparseCSV', () => {
    it('should call unparseCSVContent service', () => {
      const tableData = [['name', 'age'], ['John', '25']];
      const comments = { 0: '# comment' };
      const expectedResult = '# comment\nname,age\nJohn,25';
      
      mockUnparseCSVContent.mockReturnValue(expectedResult);
      
      const result = unparseCSV(settings, tableData, comments);
      
      expect(mockUnparseCSVContent).toHaveBeenCalledWith(settings, tableData, comments);
      expect(result).toBe(expectedResult);
    });

    it('should handle empty table data', () => {
      mockUnparseCSVContent.mockReturnValue('');
      
      const result = unparseCSV(settings, [], {});
      
      expect(mockUnparseCSVContent).toHaveBeenCalledWith(settings, [], {});
      expect(result).toBe('');
    });

    it('should handle comments properly', () => {
      const tableData = [['A'], ['1']];
      const comments = { 
        0: '# Header comment',
        2: '# Footer comment'
      };
      const expectedResult = '# Header comment\nA\n1\n# Footer comment';
      
      mockUnparseCSVContent.mockReturnValue(expectedResult);
      
      const result = unparseCSV(settings, tableData, comments);
      
      expect(result).toBe(expectedResult);
    });

    it('should pass through custom settings', () => {
      const customSettings = {
        ...settings,
        delimiter: '\t',
        quote: "'"
      };
      const tableData = [['a', 'b']];
      
      mockUnparseCSVContent.mockReturnValue('a\tb');
      
      unparseCSV(customSettings, tableData, {});
      
      expect(mockUnparseCSVContent).toHaveBeenCalledWith(customSettings, tableData, {});
    });

    it('should handle null/undefined inputs', () => {
      mockUnparseCSVContent.mockReturnValue('');
      
      const result = unparseCSV(settings, null as any, {});
      
      expect(mockUnparseCSVContent).toHaveBeenCalledWith(settings, null, {});
      expect(result).toBe('');
    });
  });

  describe('Integration with service layer', () => {
    it('should maintain backward compatibility', () => {
      const csvString = 'name,age,city\nJohn,25,NYC';
      const mockParseResult = {
        data: [['name', 'age', 'city'], ['John', '25', 'NYC']],
        comments: {},
      };
      
      mockParseCSVContent.mockReturnValue(mockParseResult);
      mockUnparseCSVContent.mockReturnValue(csvString);
      
      // Parse then unparse should work seamlessly
      const parsed = parseCSV(settings, csvString);
      const unparsed = unparseCSV(settings, parsed.data, parsed.comments);
      
      expect(mockParseCSVContent).toHaveBeenCalledWith(settings, csvString);
      expect(mockUnparseCSVContent).toHaveBeenCalledWith(settings, mockParseResult.data, mockParseResult.comments);
    });

    it('should handle service exceptions', () => {
      mockParseCSVContent.mockImplementation(() => {
        throw new Error('Service error');
      });
      
      expect(() => parseCSV(settings, 'test')).toThrow('Service error');
    });

    it('should handle unparse service exceptions', () => {
      mockUnparseCSVContent.mockImplementation(() => {
        throw new Error('Unparse error');
      });
      
      expect(() => unparseCSV(settings, [['A']], {})).toThrow('Unparse error');
    });
  });

  describe('Deprecation warnings', () => {
    it('should indicate functions are deprecated', () => {
      // These functions should have JSDoc comments indicating deprecation
      // We can't test the comments directly, but we can verify the functions exist
      // and delegate to the service layer
      expect(typeof parseCSV).toBe('function');
      expect(typeof unparseCSV).toBe('function');
    });

    it('should maintain original function signatures', () => {
      mockParseCSVContent.mockReturnValue({
        data: [['test']],
        comments: {}
      });
      
      // Original signature: (settings, csvString) => { data, comments }
      const result = parseCSV(settings, 'test');
      
      expect(result).toHaveProperty('data');
      expect(result).toHaveProperty('comments');
      expect(result).not.toHaveProperty('errors'); // Errors are filtered out for compatibility
    });
  });

  describe('Error handling compatibility', () => {
    it('should not expose errors in parseCSV for backward compatibility', () => {
      mockParseCSVContent.mockReturnValue({
        data: [['name'], ['John']],
        comments: {},
        errors: ['Some warning']
      });
      
      const result = parseCSV(settings, 'name\nJohn');
      
      expect(result).toEqual({
        data: [['name'], ['John']],
        comments: {}
      });
      expect(result).not.toHaveProperty('errors');
    });

    it('should handle service returning minimal data', () => {
      mockParseCSVContent.mockReturnValue({
        data: [['']],
        comments: {},
        errors: ['Invalid input']
      });
      
      const result = parseCSV(settings, 'invalid');
      
      expect(result.data).toEqual([['']]);
      expect(result.comments).toEqual({});
    });
  });
});