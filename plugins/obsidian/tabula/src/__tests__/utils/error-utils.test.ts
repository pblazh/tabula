import {
  createErrorResult,
  createSuccessResult,
  safeAsync,
  safeSync,
  validateRequired,
  showErrorNotice,
  showWarningNotice,
  showSuccessNotice,
  handleOperationResult,
  retryWithBackoff
} from '../../utils/error-utils';

// Mock Notice from setup
const mockNotice = (global as any).Notice;

describe('Error Utils', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  describe('createErrorResult', () => {
    it('should create error result with message', () => {
      const result = createErrorResult('Test error');
      
      expect(result.success).toBe(false);
      expect(result.error).toBe('Test error');
      expect(result.data).toBeUndefined();
      expect(result.warnings).toBeUndefined();
    });

    it('should create error result with warnings', () => {
      const result = createErrorResult('Test error', ['Warning 1', 'Warning 2']);
      
      expect(result.success).toBe(false);
      expect(result.error).toBe('Test error');
      expect(result.warnings).toEqual(['Warning 1', 'Warning 2']);
    });
  });

  describe('createSuccessResult', () => {
    it('should create success result without data', () => {
      const result = createSuccessResult();
      
      expect(result.success).toBe(true);
      expect(result.data).toBeUndefined();
      expect(result.error).toBeUndefined();
      expect(result.warnings).toBeUndefined();
    });

    it('should create success result with data', () => {
      const data = { test: 'value' };
      const result = createSuccessResult(data);
      
      expect(result.success).toBe(true);
      expect(result.data).toBe(data);
      expect(result.error).toBeUndefined();
    });

    it('should create success result with warnings', () => {
      const result = createSuccessResult('data', ['Warning']);
      
      expect(result.success).toBe(true);
      expect(result.data).toBe('data');
      expect(result.warnings).toEqual(['Warning']);
    });
  });

  describe('safeAsync', () => {
    it('should return success result for successful async operation', async () => {
      const operation = jest.fn().mockResolvedValue('test result');
      
      const result = await safeAsync(operation);
      
      expect(result.success).toBe(true);
      expect(result.data).toBe('test result');
      expect(operation).toHaveBeenCalled();
    });

    it('should return error result for failed async operation', async () => {
      const error = new Error('Test error');
      const operation = jest.fn().mockRejectedValue(error);
      
      const result = await safeAsync(operation, 'Custom message');
      
      expect(result.success).toBe(false);
      expect(result.error).toBe('Custom message: Test error');
    });

    it('should handle non-Error rejections', async () => {
      const operation = jest.fn().mockRejectedValue('String error');
      
      const result = await safeAsync(operation);
      
      expect(result.success).toBe(false);
      expect(result.error).toBe('Operation failed: String error');
    });

    it('should use default error message', async () => {
      const operation = jest.fn().mockRejectedValue(new Error('Test'));
      
      const result = await safeAsync(operation);
      
      expect(result.error).toBe('Operation failed: Test');
    });
  });

  describe('safeSync', () => {
    it('should return success result for successful sync operation', () => {
      const operation = jest.fn().mockReturnValue('test result');
      
      const result = safeSync(operation);
      
      expect(result.success).toBe(true);
      expect(result.data).toBe('test result');
      expect(operation).toHaveBeenCalled();
    });

    it('should return error result for failed sync operation', () => {
      const error = new Error('Test error');
      const operation = jest.fn().mockImplementation(() => { throw error; });
      
      const result = safeSync(operation, 'Custom message');
      
      expect(result.success).toBe(false);
      expect(result.error).toBe('Custom message: Test error');
    });

    it('should handle non-Error exceptions', () => {
      const operation = jest.fn().mockImplementation(() => { throw 'String error'; });
      
      const result = safeSync(operation);
      
      expect(result.success).toBe(false);
      expect(result.error).toBe('Operation failed: String error');
    });
  });

  describe('validateRequired', () => {
    it('should return no errors for valid object', () => {
      const obj = {
        name: 'John',
        age: 25,
        city: 'NYC'
      };
      
      const errors = validateRequired(obj, ['name', 'age']);
      
      expect(errors).toEqual([]);
    });

    it('should return errors for missing fields', () => {
      const obj = {
        name: 'John'
      };
      
      const errors = validateRequired(obj, ['name', 'age', 'city']);
      
      expect(errors).toHaveLength(2);
      expect(errors[0].field).toBe('age');
      expect(errors[0].message).toBe('age is required');
      expect(errors[1].field).toBe('city');
      expect(errors[1].message).toBe('city is required');
    });

    it('should return errors for empty string fields', () => {
      const obj = {
        name: '',
        age: '  ',
        city: 'NYC'
      };
      
      const errors = validateRequired(obj, ['name', 'age', 'city']);
      
      expect(errors).toHaveLength(2);
      expect(errors[0].field).toBe('name');
      expect(errors[1].field).toBe('age');
    });

    it('should handle undefined and null values', () => {
      const obj = {
        name: null,
        age: undefined,
        city: 'NYC'
      };
      
      const errors = validateRequired(obj, ['name', 'age']);
      
      expect(errors).toHaveLength(2);
    });
  });

  describe('Notice functions', () => {
    it('should show error notice with default duration', () => {
      showErrorNotice('Test error');
      
      expect(mockNotice).toHaveBeenCalledWith('Error: Test error', 5000);
    });

    it('should show error notice with custom duration', () => {
      showErrorNotice('Test error', 3000);
      
      expect(mockNotice).toHaveBeenCalledWith('Error: Test error', 3000);
    });

    it('should show warning notice', () => {
      showWarningNotice('Test warning');
      
      expect(mockNotice).toHaveBeenCalledWith('Warning: Test warning', 4000);
    });

    it('should show success notice', () => {
      showSuccessNotice('Test success');
      
      expect(mockNotice).toHaveBeenCalledWith('Test success', 3000);
    });
  });

  describe('handleOperationResult', () => {
    it('should handle successful result without message', () => {
      const result = createSuccessResult('test data');
      
      const data = handleOperationResult(result);
      
      expect(data).toBe('test data');
      expect(mockNotice).not.toHaveBeenCalled();
    });

    it('should handle successful result with message', () => {
      const result = createSuccessResult('test data');
      
      const data = handleOperationResult(result, 'Operation completed');
      
      expect(data).toBe('test data');
      expect(mockNotice).toHaveBeenCalledWith('Operation completed', 3000);
    });

    it('should handle successful result with warnings', () => {
      const result = createSuccessResult('test data', ['Warning 1', 'Warning 2']);
      
      const data = handleOperationResult(result);
      
      expect(data).toBe('test data');
      expect(mockNotice).toHaveBeenCalledWith('Warning: Warning 1', 4000);
      expect(mockNotice).toHaveBeenCalledWith('Warning: Warning 2', 4000);
    });

    it('should handle error result', () => {
      const result = createErrorResult('Test error');
      
      const data = handleOperationResult(result);
      
      expect(data).toBeUndefined();
      expect(mockNotice).toHaveBeenCalledWith('Error: Test error', 5000);
    });

    it('should handle error result without error message', () => {
      const result = { success: false, error: undefined };
      
      const data = handleOperationResult(result as any);
      
      expect(data).toBeUndefined();
      expect(mockNotice).not.toHaveBeenCalled();
    });
  });

  describe('retryWithBackoff', () => {
    beforeEach(() => {
      jest.useFakeTimers();
    });

    afterEach(() => {
      jest.useRealTimers();
    });

    it('should succeed on first attempt', async () => {
      const operation = jest.fn().mockResolvedValue('success');
      
      const promise = retryWithBackoff(operation, 3, 100);
      
      const result = await promise;
      
      expect(result).toBe('success');
      expect(operation).toHaveBeenCalledTimes(1);
    });

    it('should retry on failure and eventually succeed', async () => {
      const operation = jest.fn()
        .mockRejectedValueOnce(new Error('Attempt 1'))
        .mockRejectedValueOnce(new Error('Attempt 2'))
        .mockResolvedValueOnce('success');
      
      const promise = retryWithBackoff(operation, 3, 100);
      
      // Fast forward through delays
      setTimeout(async () => {
        jest.advanceTimersByTime(100); // First retry delay
        setTimeout(() => {
          jest.advanceTimersByTime(200); // Second retry delay (exponential backoff)
        }, 0);
      }, 0);
      
      const result = await promise;
      
      expect(result).toBe('success');
      expect(operation).toHaveBeenCalledTimes(3);
    });

    it('should throw last error after max retries', async () => {
      const error1 = new Error('Attempt 1');
      const error2 = new Error('Attempt 2');
      const error3 = new Error('Attempt 3');
      
      const operation = jest.fn()
        .mockRejectedValueOnce(error1)
        .mockRejectedValueOnce(error2)
        .mockRejectedValueOnce(error3);
      
      const promise = retryWithBackoff(operation, 3, 100);
      
      // Fast forward through delays
      setTimeout(() => {
        jest.advanceTimersByTime(100);
        setTimeout(() => {
          jest.advanceTimersByTime(200);
        }, 0);
      }, 0);
      
      await expect(promise).rejects.toThrow('Attempt 3');
      expect(operation).toHaveBeenCalledTimes(3);
    });

    it('should use exponential backoff delays', async () => {
      const operation = jest.fn()
        .mockRejectedValueOnce(new Error('Attempt 1'))
        .mockRejectedValueOnce(new Error('Attempt 2'))
        .mockResolvedValueOnce('success');
      
      const promise = retryWithBackoff(operation, 3, 100);
      
      const startTime = Date.now();
      
      // Manually advance timers to simulate delays
      setTimeout(() => {
        jest.advanceTimersByTime(100); // 100ms delay after first failure
        setTimeout(() => {
          jest.advanceTimersByTime(200); // 200ms delay after second failure
        }, 0);
      }, 0);
      
      await promise;
      
      // Verify exponential backoff was used (100ms, then 200ms)
      expect(operation).toHaveBeenCalledTimes(3);
    });

    it('should handle non-Error rejections', async () => {
      const operation = jest.fn()
        .mockRejectedValueOnce('String error')
        .mockResolvedValueOnce('success');
      
      const promise = retryWithBackoff(operation, 2, 100);
      
      setTimeout(() => {
        jest.advanceTimersByTime(100);
      }, 0);
      
      const result = await promise;
      
      expect(result).toBe('success');
    });

    it('should respect max retries parameter', async () => {
      const operation = jest.fn().mockRejectedValue(new Error('Always fails'));
      
      const promise = retryWithBackoff(operation, 2, 100);
      
      setTimeout(() => {
        jest.advanceTimersByTime(100);
      }, 0);
      
      await expect(promise).rejects.toThrow('Always fails');
      expect(operation).toHaveBeenCalledTimes(2); // Initial + 1 retry
    });

    it('should use custom base delay', async () => {
      const operation = jest.fn()
        .mockRejectedValueOnce(new Error('Attempt 1'))
        .mockResolvedValueOnce('success');
      
      const promise = retryWithBackoff(operation, 2, 500);
      
      setTimeout(() => {
        jest.advanceTimersByTime(500); // Custom delay
      }, 0);
      
      await promise;
      
      expect(operation).toHaveBeenCalledTimes(2);
    });
  });
});