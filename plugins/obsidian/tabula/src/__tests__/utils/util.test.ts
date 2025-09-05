import {
  appendRow,
  prependRow,
  appendColumn,
  prependColumn,
  deleteRow,
  deleteColumn,
  moveRow,
  moveColumn,
  getColumnLabel,
  normalizeTableData
} from '../../utils/util';

describe('Table Utility Functions', () => {
  describe('Row operations', () => {
    describe('appendRow', () => {
      it('should append row after specified index', () => {
        const tableData = [
          ['A', 'B', 'C'],
          ['1', '2', '3'],
          ['4', '5', '6']
        ];
        
        appendRow(1, tableData);
        
        expect(tableData).toHaveLength(4);
        expect(tableData[2]).toEqual(['', '', '']);
        expect(tableData[3]).toEqual(['4', '5', '6']);
      });

      it('should append row at end when index is last', () => {
        const tableData = [
          ['A', 'B'],
          ['1', '2']
        ];
        
        appendRow(1, tableData);
        
        expect(tableData).toHaveLength(3);
        expect(tableData[2]).toEqual(['', '']);
      });

      it('should handle single row table', () => {
        const tableData = [['A', 'B']];
        
        appendRow(0, tableData);
        
        expect(tableData).toHaveLength(2);
        expect(tableData[1]).toEqual(['', '']);
      });

      it('should create empty cells matching column count', () => {
        const tableData = [
          ['A', 'B', 'C', 'D', 'E'],
          ['1', '2', '3', '4', '5']
        ];
        
        appendRow(0, tableData);
        
        expect(tableData[1]).toHaveLength(5);
        expect(tableData[1]).toEqual(['', '', '', '', '']);
      });
    });

    describe('prependRow', () => {
      it('should prepend row at specified index', () => {
        const tableData = [
          ['A', 'B', 'C'],
          ['1', '2', '3'],
          ['4', '5', '6']
        ];
        
        prependRow(1, tableData);
        
        expect(tableData).toHaveLength(4);
        expect(tableData[1]).toEqual(['', '', '']);
        expect(tableData[2]).toEqual(['1', '2', '3']);
        expect(tableData[3]).toEqual(['4', '5', '6']);
      });

      it('should prepend row at beginning when index is 0', () => {
        const tableData = [
          ['A', 'B'],
          ['1', '2']
        ];
        
        prependRow(0, tableData);
        
        expect(tableData).toHaveLength(3);
        expect(tableData[0]).toEqual(['', '']);
        expect(tableData[1]).toEqual(['A', 'B']);
      });
    });

    describe('deleteRow', () => {
      it('should delete row at specified index', () => {
        const tableData = [
          ['A', 'B', 'C'],
          ['1', '2', '3'],
          ['4', '5', '6']
        ];
        
        deleteRow(1, tableData);
        
        expect(tableData).toHaveLength(2);
        expect(tableData[0]).toEqual(['A', 'B', 'C']);
        expect(tableData[1]).toEqual(['4', '5', '6']);
      });

      it('should not delete when only one row remains', () => {
        const tableData = [['A', 'B', 'C']];
        
        deleteRow(0, tableData);
        
        expect(tableData).toHaveLength(1);
        expect(tableData[0]).toEqual(['A', 'B', 'C']);
      });

      it('should delete first row', () => {
        const tableData = [
          ['A', 'B'],
          ['1', '2'],
          ['3', '4']
        ];
        
        deleteRow(0, tableData);
        
        expect(tableData).toHaveLength(2);
        expect(tableData[0]).toEqual(['1', '2']);
      });

      it('should delete last row', () => {
        const tableData = [
          ['A', 'B'],
          ['1', '2'],
          ['3', '4']
        ];
        
        deleteRow(2, tableData);
        
        expect(tableData).toHaveLength(2);
        expect(tableData[1]).toEqual(['1', '2']);
      });
    });

    describe('moveRow', () => {
      it('should move row to different position', () => {
        const tableData = [
          ['A', 'B'],
          ['1', '2'],
          ['3', '4'],
          ['5', '6']
        ];
        
        moveRow(1, 3, tableData);
        
        expect(tableData).toEqual([
          ['A', 'B'],
          ['3', '4'],
          ['5', '6'],
          ['1', '2']
        ]);
      });

      it('should move row forward', () => {
        const tableData = [
          ['A'],
          ['B'],
          ['C'],
          ['D']
        ];
        
        moveRow(0, 2, tableData);
        
        expect(tableData).toEqual([
          ['B'],
          ['C'],
          ['A'],
          ['D']
        ]);
      });

      it('should move row backward', () => {
        const tableData = [
          ['A'],
          ['B'],
          ['C'],
          ['D']
        ];
        
        moveRow(3, 1, tableData);
        
        expect(tableData).toEqual([
          ['A'],
          ['D'],
          ['B'],
          ['C']
        ]);
      });

      it('should not move when indices are out of bounds', () => {
        const tableData = [
          ['A', 'B'],
          ['1', '2']
        ];
        const original = JSON.parse(JSON.stringify(tableData));
        
        moveRow(-1, 1, tableData);
        expect(tableData).toEqual(original);
        
        moveRow(0, 5, tableData);
        expect(tableData).toEqual(original);
        
        moveRow(5, 0, tableData);
        expect(tableData).toEqual(original);
      });
    });
  });

  describe('Column operations', () => {
    describe('appendColumn', () => {
      it('should append column after specified index', () => {
        const tableData = [
          ['A', 'B', 'C'],
          ['1', '2', '3'],
          ['4', '5', '6']
        ];
        
        appendColumn(1, tableData);
        
        expect(tableData[0]).toEqual(['A', 'B', '', 'C']);
        expect(tableData[1]).toEqual(['1', '2', '', '3']);
        expect(tableData[2]).toEqual(['4', '5', '', '6']);
      });

      it('should append column at end', () => {
        const tableData = [
          ['A', 'B'],
          ['1', '2']
        ];
        
        appendColumn(1, tableData);
        
        expect(tableData[0]).toEqual(['A', 'B', '']);
        expect(tableData[1]).toEqual(['1', '2', '']);
      });
    });

    describe('prependColumn', () => {
      it('should prepend column at specified index', () => {
        const tableData = [
          ['A', 'B', 'C'],
          ['1', '2', '3'],
          ['4', '5', '6']
        ];
        
        prependColumn(1, tableData);
        
        expect(tableData[0]).toEqual(['A', '', 'B', 'C']);
        expect(tableData[1]).toEqual(['1', '', '2', '3']);
        expect(tableData[2]).toEqual(['4', '', '5', '6']);
      });

      it('should prepend column at beginning', () => {
        const tableData = [
          ['A', 'B'],
          ['1', '2']
        ];
        
        prependColumn(0, tableData);
        
        expect(tableData[0]).toEqual(['', 'A', 'B']);
        expect(tableData[1]).toEqual(['', '1', '2']);
      });
    });

    describe('deleteColumn', () => {
      it('should delete column at specified index', () => {
        const tableData = [
          ['A', 'B', 'C'],
          ['1', '2', '3'],
          ['4', '5', '6']
        ];
        
        deleteColumn(1, tableData);
        
        expect(tableData[0]).toEqual(['A', 'C']);
        expect(tableData[1]).toEqual(['1', '3']);
        expect(tableData[2]).toEqual(['4', '6']);
      });

      it('should not delete when table is empty', () => {
        const tableData: string[][] = [];
        
        deleteColumn(0, tableData);
        
        expect(tableData).toEqual([]);
      });

      it('should delete first column', () => {
        const tableData = [
          ['A', 'B', 'C'],
          ['1', '2', '3']
        ];
        
        deleteColumn(0, tableData);
        
        expect(tableData[0]).toEqual(['B', 'C']);
        expect(tableData[1]).toEqual(['2', '3']);
      });

      it('should delete last column', () => {
        const tableData = [
          ['A', 'B', 'C'],
          ['1', '2', '3']
        ];
        
        deleteColumn(2, tableData);
        
        expect(tableData[0]).toEqual(['A', 'B']);
        expect(tableData[1]).toEqual(['1', '2']);
      });
    });

    describe('moveColumn', () => {
      it('should move column to different position', () => {
        const tableData = [
          ['A', 'B', 'C', 'D'],
          ['1', '2', '3', '4'],
          ['5', '6', '7', '8']
        ];
        
        moveColumn(1, 3, tableData);
        
        expect(tableData[0]).toEqual(['A', 'C', 'D', 'B']);
        expect(tableData[1]).toEqual(['1', '3', '4', '2']);
        expect(tableData[2]).toEqual(['5', '7', '8', '6']);
      });

      it('should move column forward', () => {
        const tableData = [
          ['A', 'B', 'C', 'D'],
          ['1', '2', '3', '4']
        ];
        
        moveColumn(0, 2, tableData);
        
        expect(tableData[0]).toEqual(['B', 'C', 'A', 'D']);
        expect(tableData[1]).toEqual(['2', '3', '1', '4']);
      });

      it('should move column backward', () => {
        const tableData = [
          ['A', 'B', 'C', 'D'],
          ['1', '2', '3', '4']
        ];
        
        moveColumn(3, 1, tableData);
        
        expect(tableData[0]).toEqual(['A', 'D', 'B', 'C']);
        expect(tableData[1]).toEqual(['1', '4', '2', '3']);
      });

      it('should not move when indices are out of bounds', () => {
        const tableData = [
          ['A', 'B'],
          ['1', '2']
        ];
        const original = JSON.parse(JSON.stringify(tableData));
        
        moveColumn(-1, 1, tableData);
        expect(tableData).toEqual(original);
        
        moveColumn(0, 5, tableData);
        expect(tableData).toEqual(original);
        
        moveColumn(5, 0, tableData);
        expect(tableData).toEqual(original);
      });
    });
  });

  describe('getColumnLabel', () => {
    it('should generate correct column labels for single letters', () => {
      expect(getColumnLabel(0)).toBe('A');
      expect(getColumnLabel(1)).toBe('B');
      expect(getColumnLabel(25)).toBe('Z');
    });

    it('should generate correct column labels for double letters', () => {
      expect(getColumnLabel(26)).toBe('AA');
      expect(getColumnLabel(27)).toBe('AB');
      expect(getColumnLabel(51)).toBe('AZ');
      expect(getColumnLabel(52)).toBe('BA');
    });

    it('should generate correct column labels for triple letters', () => {
      expect(getColumnLabel(702)).toBe('AAA');
      expect(getColumnLabel(703)).toBe('AAB');
    });

    it('should handle edge cases', () => {
      expect(getColumnLabel(675)).toBe('YZ');
      expect(getColumnLabel(676)).toBe('ZA');
      expect(getColumnLabel(677)).toBe('ZB');
      expect(getColumnLabel(701)).toBe('ZZ');
    });

    it('should generate sequence correctly', () => {
      const labels: string[] = [];
      for (let i = 0; i < 30; i++) {
        labels.push(getColumnLabel(i));
      }
      
      expect(labels.slice(0, 5)).toEqual(['A', 'B', 'C', 'D', 'E']);
      expect(labels.slice(25, 30)).toEqual(['Z', 'AA', 'AB', 'AC', 'AD']);
    });
  });

  describe('normalizeTableData', () => {
    it('should normalize table with uneven rows', () => {
      const tableData = [
        ['A', 'B', 'C'],
        ['1', '2'],
        ['3', '4', '5', '6']
      ];
      
      const result = normalizeTableData(tableData);
      
      expect(result).toEqual([
        ['A', 'B', 'C', ''],
        ['1', '2', '', ''],
        ['3', '4', '5', '6']
      ]);
    });

    it('should handle empty table data', () => {
      const result = normalizeTableData([]);
      
      expect(result).toEqual([['']]);
    });

    it('should handle null/undefined table data', () => {
      const result1 = normalizeTableData(null as any);
      expect(result1).toEqual([['']]);
      
      const result2 = normalizeTableData(undefined as any);
      expect(result2).toEqual([['']]);
    });

    it('should handle table with empty rows', () => {
      const tableData = [
        ['A', 'B'],
        [],
        ['1', '2', '3']
      ];
      
      const result = normalizeTableData(tableData);
      
      expect(result).toEqual([
        ['A', 'B', ''],
        ['', '', ''],
        ['1', '2', '3']
      ]);
    });

    it('should handle table with null rows', () => {
      const tableData = [
        ['A', 'B'],
        null,
        ['1', '2', '3']
      ] as any;
      
      const result = normalizeTableData(tableData);
      
      expect(result).toEqual([
        ['A', 'B', ''],
        ['', '', ''],
        ['1', '2', '3']
      ]);
    });

    it('should not modify already normalized table', () => {
      const tableData = [
        ['A', 'B', 'C'],
        ['1', '2', '3'],
        ['4', '5', '6']
      ];
      
      const result = normalizeTableData(tableData);
      
      expect(result).toEqual(tableData);
      expect(result).not.toBe(tableData); // Should be a copy
    });

    it('should handle single column table', () => {
      const tableData = [
        ['A'],
        ['1'],
        ['2']
      ];
      
      const result = normalizeTableData(tableData);
      
      expect(result).toEqual([
        ['A'],
        ['1'],
        ['2']
      ]);
    });

    it('should handle single row table', () => {
      const tableData = [
        ['A', 'B', 'C']
      ];
      
      const result = normalizeTableData(tableData);
      
      expect(result).toEqual([
        ['A', 'B', 'C']
      ]);
    });

    it('should handle table with varying row lengths', () => {
      const tableData = [
        ['A'],
        ['1', '2', '3', '4', '5'],
        ['X', 'Y'],
        []
      ];
      
      const result = normalizeTableData(tableData);
      
      expect(result).toEqual([
        ['A', '', '', '', ''],
        ['1', '2', '3', '4', '5'],
        ['X', 'Y', '', '', ''],
        ['', '', '', '', '']
      ]);
    });

    it('should preserve original data types', () => {
      const tableData = [
        ['string', 'another'],
        ['test']
      ];
      
      const result = normalizeTableData(tableData);
      
      expect(typeof result[0][0]).toBe('string');
      expect(typeof result[1][1]).toBe('string'); // empty string
      expect(result[1][1]).toBe('');
    });
  });
});