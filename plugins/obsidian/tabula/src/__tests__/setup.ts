// Jest setup file for global test configuration

// Mock Obsidian API
export {};
const mockNotice = jest.fn();
const mockApp = {
  workspace: {
    getLeavesOfType: jest.fn(() => []),
    setActiveLeaf: jest.fn(),
    getLeaf: jest.fn(() => ({
      openFile: jest.fn(),
      setViewState: jest.fn()
    }))
  },
  vault: {
    adapter: {
      basePath: '/test/path'
    }
  }
};

// Global mocks for Obsidian
(global as any).Notice = mockNotice;
(global as any).moment = {
  locale: () => 'en'
};

// Mock DOM methods
Object.defineProperty(document, 'createElement', {
  value: jest.fn(() => ({
    id: '',
    textContent: '',
    appendChild: jest.fn(),
    removeChild: jest.fn(),
    parentNode: {
      removeChild: jest.fn()
    },
    querySelector: jest.fn(),
    querySelectorAll: jest.fn(() => []),
    addEventListener: jest.fn(),
    removeEventListener: jest.fn(),
    classList: {
      contains: jest.fn(),
      add: jest.fn(),
      remove: jest.fn()
    },
    dataset: {},
    style: {}
  }))
});

Object.defineProperty(document, 'head', {
  value: {
    appendChild: jest.fn(),
    removeChild: jest.fn()
  }
});

Object.defineProperty(document, 'getElementById', {
  value: jest.fn(() => null)
});

// Mock process.env
process.env.PATH = '/usr/bin:/bin';
process.env.HOME = '/home/test';

// Global test utilities
(global as any).createMockApp = () => mockApp;
(global as any).createMockNotice = () => mockNotice;

// Reset mocks before each test
beforeEach(() => {
  jest.clearAllMocks();
});