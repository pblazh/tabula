import { FileUtils } from '../../utils/file-utils';

// Mock Notice from setup
const mockNotice = (global as any).Notice;

describe('FileUtils', () => {
  beforeEach(() => {
    jest.clearAllMocks();
    jest.useFakeTimers();
  });

  afterEach(() => {
    jest.useRealTimers();
  });

  describe('withRetry', () => {
    it('should succeed on first attempt', async () => {
      const operation = jest.fn().mockResolvedValue('success');
      
      const result = await FileUtils.withRetry(operation, 3, 100);
      
      expect(result).toBe('success');
      expect(operation).toHaveBeenCalledTimes(1);
      expect(mockNotice).not.toHaveBeenCalled();
    });

    it('should retry on file busy error and eventually succeed', async () => {
      const busyError = new Error('EBUSY: resource busy or locked');
      const operation = jest.fn()
        .mockRejectedValueOnce(busyError)
        .mockRejectedValueOnce(busyError)
        .mockResolvedValueOnce('success');
      
      const retryPromise = FileUtils.withRetry(operation, 3, 500);
      
      // Fast-forward through the delays
      setTimeout(() => {
        jest.advanceTimersByTime(500); // First retry delay
        setTimeout(() => {
          jest.advanceTimersByTime(500); // Second retry delay
        }, 0);
      }, 0);
      
      const result = await retryPromise;
      
      expect(result).toBe('success');
      expect(operation).toHaveBeenCalledTimes(3);
      expect(mockNotice).toHaveBeenCalledWith('File is busy. Retrying... (1/3)');
    });

    it('should retry on different busy error messages', async () => {
      const busyErrors = [
        new Error('File is busy'),
        new Error('Resource locked'),
        new Error('EBUSY: resource busy')
      ];
      
      for (const error of busyErrors) {
        const operation = jest.fn()
          .mockRejectedValueOnce(error)
          .mockResolvedValueOnce('success');
        
        const retryPromise = FileUtils.withRetry(operation, 2, 100);
        
        setTimeout(() => {
          jest.advanceTimersByTime(100);
        }, 0);
        
        const result = await retryPromise;
        
        expect(result).toBe('success');
        expect(operation).toHaveBeenCalledTimes(2);
      }
    });

    it('should not retry on non-busy errors', async () => {
      const nonBusyError = new Error('File not found');
      const operation = jest.fn().mockRejectedValue(nonBusyError);
      
      await expect(FileUtils.withRetry(operation, 3, 100)).rejects.toThrow('File not found');
      expect(operation).toHaveBeenCalledTimes(1);
      expect(mockNotice).not.toHaveBeenCalled();
    });

    it('should throw last error after max retries', async () => {
      const busyError = new Error('EBUSY: resource busy');
      const operation = jest.fn().mockRejectedValue(busyError);
      
      const retryPromise = FileUtils.withRetry(operation, 2, 100);
      
      // Fast-forward through delays
      setTimeout(() => {
        jest.advanceTimersByTime(100);
      }, 0);
      
      await expect(retryPromise).rejects.toThrow('EBUSY: resource busy');
      expect(operation).toHaveBeenCalledTimes(2); // Initial attempt + 1 retry
      expect(mockNotice).toHaveBeenCalledWith('File is busy. Retrying... (1/2)');
    });

    it('should use default parameters', async () => {
      const busyError = new Error('busy');
      const operation = jest.fn()
        .mockRejectedValueOnce(busyError)
        .mockResolvedValueOnce('success');
      
      const retryPromise = FileUtils.withRetry(operation); // No parameters
      
      setTimeout(() => {
        jest.advanceTimersByTime(500); // Default delay
      }, 0);
      
      const result = await retryPromise;
      
      expect(result).toBe('success');
      expect(operation).toHaveBeenCalledTimes(2);
    });

    it('should respect custom max retries', async () => {
      const busyError = new Error('EBUSY');
      const operation = jest.fn().mockRejectedValue(busyError);
      
      const retryPromise = FileUtils.withRetry(operation, 1, 100);
      
      await expect(retryPromise).rejects.toThrow('EBUSY');
      expect(operation).toHaveBeenCalledTimes(1); // No retries with maxRetries=1
    });

    it('should respect custom delay', async () => {
      const busyError = new Error('locked');
      const operation = jest.fn()
        .mockRejectedValueOnce(busyError)
        .mockResolvedValueOnce('success');
      
      const retryPromise = FileUtils.withRetry(operation, 3, 200);
      
      // Verify specific delay is used
      setTimeout(() => {
        jest.advanceTimersByTime(199); // Just before delay
        expect(operation).toHaveBeenCalledTimes(1); // Should not have retried yet
        
        jest.advanceTimersByTime(1); // Complete the delay
      }, 0);
      
      const result = await retryPromise;
      expect(result).toBe('success');
    });

    it('should show notice only on first retry', async () => {
      const busyError = new Error('EBUSY');
      const operation = jest.fn()
        .mockRejectedValueOnce(busyError)
        .mockRejectedValueOnce(busyError)
        .mockRejectedValueOnce(busyError);
      
      const retryPromise = FileUtils.withRetry(operation, 3, 100);
      
      setTimeout(() => {
        jest.advanceTimersByTime(100);
        setTimeout(() => {
          jest.advanceTimersByTime(100);
        }, 0);
      }, 0);
      
      await expect(retryPromise).rejects.toThrow('EBUSY');
      
      // Notice should only be shown once (on first retry)
      expect(mockNotice).toHaveBeenCalledTimes(1);
      expect(mockNotice).toHaveBeenCalledWith('File is busy. Retrying... (1/3)');
    });

    it('should handle async operations that throw synchronously', async () => {
      const operation = jest.fn().mockImplementation(() => {
        throw new Error('EBUSY: sync error');
      });
      
      const retryPromise = FileUtils.withRetry(operation, 2, 100);
      
      setTimeout(() => {
        jest.advanceTimersByTime(100);
      }, 0);
      
      await expect(retryPromise).rejects.toThrow('EBUSY: sync error');
      expect(operation).toHaveBeenCalledTimes(2);
    });

    it('should handle operations returning rejected promises', async () => {
      const busyError = new Error('busy');
      const operation = jest.fn().mockReturnValue(Promise.reject(busyError));
      
      const retryPromise = FileUtils.withRetry(operation, 2, 100);
      
      setTimeout(() => {
        jest.advanceTimersByTime(100);
      }, 0);
      
      await expect(retryPromise).rejects.toThrow('busy');
      expect(operation).toHaveBeenCalledTimes(2);
    });

    it('should detect busy errors case-insensitively', async () => {
      const busyErrors = [
        new Error('File is BUSY'),
        new Error('resource LOCKED'),
        new Error('ebusy: error')
      ];
      
      for (const error of busyErrors) {
        const operation = jest.fn()
          .mockRejectedValueOnce(error)
          .mockResolvedValueOnce('success');
        
        const retryPromise = FileUtils.withRetry(operation, 2, 100);
        
        setTimeout(() => {
          jest.advanceTimersByTime(100);
        }, 0);
        
        const result = await retryPromise;
        expect(result).toBe('success');
      }
    });

    it('should handle non-Error objects', async () => {
      const stringError = 'EBUSY: file busy';
      const operation = jest.fn().mockRejectedValue(stringError);
      
      await expect(FileUtils.withRetry(operation, 1, 100)).rejects.toBe(stringError);
      expect(operation).toHaveBeenCalledTimes(1);
    });

    it('should handle operations that return non-promise values', async () => {
      const operation = jest.fn().mockReturnValue('immediate success');
      
      const result = await FileUtils.withRetry(operation);
      
      expect(result).toBe('immediate success');
      expect(operation).toHaveBeenCalledTimes(1);
    });

    it('should handle zero retries', async () => {
      const busyError = new Error('EBUSY');
      const operation = jest.fn().mockRejectedValue(busyError);
      
      await expect(FileUtils.withRetry(operation, 0, 100)).rejects.toThrow('EBUSY');
      expect(operation).toHaveBeenCalledTimes(1);
      expect(mockNotice).not.toHaveBeenCalled();
    });

    it('should handle negative retries', async () => {
      const busyError = new Error('EBUSY');
      const operation = jest.fn().mockRejectedValue(busyError);
      
      await expect(FileUtils.withRetry(operation, -1, 100)).rejects.toThrow('EBUSY');
      expect(operation).toHaveBeenCalledTimes(1);
    });
  });
});