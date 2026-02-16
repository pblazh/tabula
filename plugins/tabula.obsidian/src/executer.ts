import { TabulaSettings } from "./types";
import * as path from "node:path";
import { spawn } from "node:child_process";
import * as os from "node:os";
import * as fs from "node:fs/promises";
import * as crypto from "node:crypto";

export class Executer {
  constructor(
    private settings: TabulaSettings,
    private sourcePath: string,
  ) {}

  async execute(data: string, code: string): Promise<string> {
    const filePath = path.join(
      os.tmpdir(),
      `tabula_${crypto.randomBytes(6).toString("hex")}.tbl`,
    );

    try {
      await fs.writeFile(filePath, fixIncludes(this.sourcePath, code), "utf8");
      const args = [
        this.settings.autoFormat ? "-a" : "",
        "-s",
        filePath,
      ].filter(Boolean);

      return await run(this.settings.executablePath, args, data);
    } finally {
      await fs.unlink(filePath).catch((err) => {
        return {
          result: "",
          error: String(err),
        };
      });
    }
  }
}

function fixIncludes(base: string, code: string): string {
  return code.replace(/#include "/g, `#include "${base}/`);
}

function run(
  cmd: string,
  args: string[] = [],
  input: string = "",
): Promise<string> {
  return new Promise((resolve, reject) => {
    const child = spawn(cmd, args);

    let stdout = "";
    let stderr = "";

    child.stdout.on("data", (data: string) => {
      stdout += data;
    });

    child.stderr.on("data", (data: string) => {
      stderr += data;
    });

    child.on("error", reject);

    child.on("close", (code: null) => {
      if (code === 0) {
        resolve(stdout);
      } else {
        reject(new Error(stderr));
      }
    });

    if (input) {
      child.stdin.write(input);
    }
    child.stdin.end();
  });
}
