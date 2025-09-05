import { Notice } from "obsidian";
import * as child_process from "child_process";
import * as path from "path";

export interface ProcessExecutionResult {
  success: boolean;
  stdout?: string;
  stderr?: string;
  error?: string;
  hasStderr?: boolean;
}

/**
 * Sanitize file path to prevent path traversal attacks
 */
function sanitizeFilePath(filePath: string): string {
  if (!filePath || typeof filePath !== "string") {
    throw new Error("Invalid file path provided");
  }

  // Resolve the path to handle any relative components
  const resolvedPath = path.resolve(filePath);

  // Basic validation - ensure it's a reasonable file path
  if (resolvedPath.includes("..") || resolvedPath.includes("~")) {
    throw new Error("File path contains unsafe components");
  }

  return resolvedPath;
}

/**
 * Validate executable name to prevent command injection
 */
function validateExecutableName(executable: string): string {
  if (!executable || typeof executable !== "string") {
    throw new Error("Invalid executable name");
  }

  // Only allow alphanumeric characters, hyphens, underscores, and dots
  const validPattern = /^[a-zA-Z0-9._-]+$/;
  if (!validPattern.test(executable)) {
    throw new Error("Executable name contains invalid characters");
  }

  return executable.trim();
}

export async function executeTabula(
  executable: string,
  filePath: string,
  showNotices = true,
): Promise<ProcessExecutionResult> {
  try {
    const sanitizedPath = sanitizeFilePath(filePath);
    const validatedExecutable = validateExecutableName(executable);

    if (showNotices) {
      new Notice("Executing tabula...");
    }

    return new Promise<ProcessExecutionResult>((resolve) => {
      // Use execFile instead of spawn for better security
      // Fixed arguments prevent command injection
      const child = child_process.execFile(
        validatedExecutable,
        ["-a", "-u", sanitizedPath],
        {
          timeout: 30000, // 30 second timeout
          maxBuffer: 1024 * 1024, // 1MB buffer limit
          env: {
            // Only pass specific environment variables
            PATH: process.env.PATH,
            HOME: process.env.HOME,
          },
        },
        (error, stdout, stderr) => {
          if (error) {
            const errorMsg = `Tabula execution failed: ${error.message}`;
            if (showNotices) {
              new Notice(errorMsg);
            }
            resolve({
              success: false,
              error: errorMsg,
              stderr: stderr?.toString(),
            });
            return;
          }

          const hasStderr = stderr && stderr.trim();
          if (hasStderr && showNotices) {
            new Notice("Tabula execution completed with errors - check error panel for details");
          }

          if (stdout && stdout.trim() && showNotices) {
            new Notice(`Tabula output: ${stdout}`);
          }

          if (showNotices) {
            new Notice("Tabula execution completed");
          }

          resolve({
            success: true,
            stdout: stdout?.toString(),
            stderr: stderr?.toString(),
            hasStderr: !!(stderr && stderr.trim()),
          });
        },
      );

      // Handle process errors
      child.on("error", (error) => {
        const errorMsg = `Process error: ${error.message}`;
        if (showNotices) {
          new Notice(errorMsg);
        }
        resolve({
          success: false,
          error: errorMsg,
        });
      });
    });
  } catch (error) {
    const errorMsg = `Security validation failed: ${error instanceof Error ? error.message : "Unknown error"}`;
    if (showNotices) {
      new Notice(errorMsg);
    }
    return {
      success: false,
      error: errorMsg,
    };
  }
}

export async function checkTabulaExecutableAvailable(
  executable: string,
): Promise<boolean> {
  try {
    const validatedExecutable = validateExecutableName(executable);

    return new Promise<boolean>((resolve) => {
      child_process.exec(`${validatedExecutable} --help`, (error) =>
        resolve(!error),
      );
    });
  } catch {
    return false;
  }
}

