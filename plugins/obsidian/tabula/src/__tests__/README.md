# Tabula Plugin Tests

This directory contains comprehensive tests for the Tabula Obsidian plugin, covering all major components and functionality.

## Test Structure

```
src/__tests__/
├── setup.ts                 # Global test configuration and mocks
├── services/                # Tests for service layer
│   ├── csv-service.test.ts  # CSV parsing and validation
│   └── process-service.test.ts # Secure process execution
├── utils/                   # Tests for utility functions
│   ├── error-utils.test.ts  # Error handling utilities
│   ├── util.test.ts         # Table manipulation utilities
│   ├── file-utils.test.ts   # File retry utilities
│   └── csv-utils.test.ts    # CSV wrapper utilities
├── components/              # Tests for UI components
│   └── table.test.ts        # Table rendering component
├── integration/             # Integration tests
│   └── main.test.ts         # Main plugin lifecycle
└── test-runner.ts           # Custom test runner utility
```

## Test Categories

### Unit Tests (`npm run test:unit`)
- **Services**: Core business logic and data processing
- **Utils**: Helper functions and utilities
- **Components**: UI rendering and interaction logic

### Integration Tests (`npm run test:integration`)
- **Main Plugin**: Plugin lifecycle and initialization
- **View Integration**: Complete view workflows
- **Settings Integration**: Settings validation and persistence

## Running Tests

```bash
# Run all tests
npm test

# Run with coverage
npm run test:coverage

# Run in watch mode
npm run test:watch

# Run specific test suites
npm run test:unit
npm run test:integration

# Run for CI/CD
npm run test:ci
```

## Test Configuration

### Jest Configuration (`jest.config.js`)
- **Environment**: jsdom (for DOM testing)
- **TypeScript**: Full TypeScript support with ts-jest
- **Coverage**: Comprehensive coverage reporting
- **Mocks**: Automatic Obsidian API mocking

### Global Setup (`setup.ts`)
- **Obsidian API Mocks**: Notice, App, Workspace APIs
- **DOM Mocks**: Document manipulation methods
- **Test Utilities**: Helper functions for test creation

## Key Features Tested

### Security
✅ **Command Injection Prevention**: Process execution security  
✅ **Path Traversal Protection**: File path sanitization  
✅ **Input Validation**: User input sanitization  

### Data Processing
✅ **CSV Parsing**: Complex CSV handling with comments  
✅ **Error Recovery**: Graceful error handling and fallbacks  
✅ **Data Validation**: Settings and input validation  

### UI Components
✅ **Table Rendering**: Dynamic table generation  
✅ **Event Handling**: User interaction processing  
✅ **State Management**: View state consistency  

### File Operations
✅ **Retry Logic**: File busy error handling  
✅ **Atomic Operations**: Safe file operations  
✅ **Error Recovery**: Graceful failure handling  

## Coverage Goals

- **Statements**: >90%
- **Branches**: >85%
- **Functions**: >90%
- **Lines**: >90%

## Test Patterns

### Service Testing
```typescript
describe('ServiceName', () => {
  it('should handle normal operation', () => {
    const result = service.method(input);
    expect(result).toEqual(expected);
  });

  it('should handle error cases', () => {
    expect(() => service.method(badInput)).toThrow('Expected error');
  });
});
```

### Mock Usage
```typescript
// Mock external dependencies
jest.mock('../../services/external-service');
const mockService = externalService as jest.MockedFunction<typeof externalService>;

beforeEach(() => {
  jest.clearAllMocks();
  mockService.mockReturnValue(testData);
});
```

### Async Testing
```typescript
it('should handle async operations', async () => {
  const promise = asyncFunction();
  
  // Fast-forward timers if needed
  jest.advanceTimersByTime(1000);
  
  const result = await promise;
  expect(result).toBe(expected);
});
```

## Writing New Tests

### 1. Test File Naming
- Use `.test.ts` suffix
- Match source file structure: `src/utils/helper.ts` → `src/__tests__/utils/helper.test.ts`

### 2. Test Organization
```typescript
describe('ComponentName', () => {
  describe('methodName', () => {
    it('should handle normal case', () => {});
    it('should handle edge cases', () => {});
    it('should handle error cases', () => {});
  });
});
```

### 3. Mock Guidelines
- Mock external dependencies at module level
- Use `jest.clearAllMocks()` in `beforeEach`
- Mock Obsidian APIs through global setup
- Use specific mocks for individual test needs

### 4. Assertion Patterns
```typescript
// Exact matching
expect(result).toEqual(expected);

// Partial matching
expect(result).toMatchObject({ key: value });

// Function calls
expect(mockFn).toHaveBeenCalledWith(expectedArgs);
expect(mockFn).toHaveBeenCalledTimes(1);

// Errors
expect(() => fn()).toThrow('Expected message');
await expect(asyncFn()).rejects.toThrow('Error');
```

## Debugging Tests

### Running Single Test
```bash
npm test -- --testNamePattern="specific test name"
```

### Debugging with VSCode
Add to `.vscode/launch.json`:
```json
{
  "name": "Debug Jest Tests",
  "type": "node",
  "request": "launch",
  "program": "${workspaceFolder}/node_modules/.bin/jest",
  "args": ["--runInBand"],
  "console": "integratedTerminal",
  "internalConsoleOptions": "neverOpen"
}
```

### Test Output
- Use `console.log()` for debugging (will show in test output)
- Use `--verbose` flag for detailed test results
- Check coverage reports in `coverage/` directory

## Continuous Integration

The test suite is designed to run in CI/CD environments:

```bash
npm run test:ci
```

This command:
- Runs all tests once (no watch mode)
- Generates coverage reports
- Exits with proper error codes
- Works in headless environments

## Maintenance

### Regular Tasks
- [ ] Update tests when adding new features
- [ ] Maintain >90% coverage
- [ ] Review and update mocks for API changes
- [ ] Clean up deprecated test patterns

### Performance
- Tests should complete in <30 seconds
- Use `jest.useFakeTimers()` for time-dependent tests
- Mock expensive operations
- Keep test data minimal but comprehensive

## Troubleshooting

### Common Issues

**Tests timeout**
- Check for unresolved promises
- Use `jest.useFakeTimers()` and `jest.advanceTimersByTime()`

**Mock not working**
- Verify mock is at module level
- Check mock is cleared between tests
- Ensure proper TypeScript types

**Coverage not accurate**
- Check file paths in coverage configuration
- Verify test files are not included in coverage
- Update ignore patterns for generated files