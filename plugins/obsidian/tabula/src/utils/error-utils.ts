import { Notice } from "obsidian";
import { OperationResult, ValidationError } from "../types";

/**
 * Create a standardized error result
 */
export function createErrorResult<T = void>(
  error: string, 
  warnings?: string[]
): OperationResult<T> {
  return {
    success: false,
    error,
    warnings,
  };
}

/**
 * Create a standardized success result
 */
export function createSuccessResult<T = void>(
  data?: T, 
  warnings?: string[]
): OperationResult<T> {
  return {
    success: true,
    data,
    warnings,
  };
}

/**
 * Safe async operation wrapper with error handling
 */
export async function safeAsync<T>(
  operation: () => Promise<T>,
  errorMessage: string = "Operation failed"
): Promise<OperationResult<T>> {
  try {
    const result = await operation();
    return createSuccessResult(result);
  } catch (error) {
    const message = error instanceof Error ? error.message : String(error);
    return createErrorResult(`${errorMessage}: ${message}`);
  }
}

/**
 * Safe sync operation wrapper with error handling
 */
export function safeSync<T>(
  operation: () => T,
  errorMessage: string = "Operation failed"
): OperationResult<T> {
  try {
    const result = operation();
    return createSuccessResult(result);
  } catch (error) {
    const message = error instanceof Error ? error.message : String(error);
    return createErrorResult(`${errorMessage}: ${message}`);
  }
}

/**
 * Validate required fields
 */
export function validateRequired(
  obj: Record<string, unknown>,
  requiredFields: string[]
): ValidationError[] {
  const errors: ValidationError[] = [];
  
  for (const field of requiredFields) {
    if (!obj[field] || (typeof obj[field] === 'string' && !(obj[field] as string).trim())) {
      errors.push({
        field,
        message: `${field} is required`,
      });
    }
  }
  
  return errors;
}

/**
 * Show error notice to user
 */
export function showErrorNotice(error: string, duration: number = 5000): void {
  new Notice(`Error: ${error}`, duration);
}

/**
 * Show warning notice to user
 */
export function showWarningNotice(warning: string, duration: number = 4000): void {
  new Notice(`Warning: ${warning}`, duration);
}

/**
 * Show success notice to user
 */
export function showSuccessNotice(message: string, duration: number = 3000): void {
  new Notice(message, duration);
}

/**
 * Handle operation result and show appropriate notices
 */
export function handleOperationResult<T>(
  result: OperationResult<T>,
  successMessage?: string
): T | undefined {
  if (result.success) {
    if (successMessage) {
      showSuccessNotice(successMessage);
    }
    if (result.warnings && result.warnings.length > 0) {
      result.warnings.forEach(warning => showWarningNotice(warning));
    }
    return result.data;
  } else {
    if (result.error) {
      showErrorNotice(result.error);
    }
    return undefined;
  }
}

/**
 * Retry operation with exponential backoff
 */
export async function retryWithBackoff<T>(
  operation: () => Promise<T>,
  maxRetries: number = 3,
  baseDelay: number = 1000
): Promise<T> {
  let lastError: Error;
  
  for (let attempt = 0; attempt < maxRetries; attempt++) {
    try {
      return await operation();
    } catch (error) {
      lastError = error instanceof Error ? error : new Error(String(error));
      
      if (attempt < maxRetries - 1) {
        const delay = baseDelay * Math.pow(2, attempt);
        await new Promise(resolve => setTimeout(resolve, delay));
      }
    }
  }
  
  throw lastError!;
}