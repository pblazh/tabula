import { enUS } from "./en";

export const LOCALE = {
  en: enUS,
};

export type Locale = keyof typeof LOCALE;

export class I18n {
  private locale: Locale = "en";

  constructor(locale?: string) {
    this.setLocale(locale || "en");
  }

  setLocale(locale: string) {
    console.log(`I18n: Attempting to set locale to '${locale}'`);
    const lowerLocale = locale.toLowerCase();

    if (lowerLocale.startsWith("en")) {
      this.locale = "en";
    } else {
      this.locale = "en";
    }
  }

  t(key: string, params?: Record<string, string | number>): string {
    // Try to get translation from current language
    let translatedText = this.getTranslation(key, this.locale);

    // If not found in current language and current language is not English, try fallback to English
    if (translatedText === null && this.locale !== "en") {
      console.warn(
        `I18n: Key '${key}' not found in '${this.locale}', falling back to 'en'.`,
      );
      translatedText = this.getTranslation(key, "en");
    }

    // If not found use the key itself
    let result = translatedText ?? key;

    // If parameters provided, perform string interpolation
    if (params) {
      Object.keys(params).forEach((paramKey) => {
        const placeholder = `{${paramKey}}`;
        result = result.replace(
          new RegExp(placeholder, "g"),
          String(params[paramKey]),
        );
      });
    }

    return result;
  }

  // Helper method for finding translations
  private getTranslation(key: string, locale: Locale): string | null {
    const translation = LOCALE[locale];
    const keys = key.split(".");
    let result: any = translation;

    for (const k of keys) {
      if (result && typeof result === "object" && k in result) {
        result = result[k];
      } else {
        return null; // Return null if not found
      }
    }

    return typeof result === "string" ? result : null;
  }

  getCurrentLocale(): Locale {
    return this.locale;
  }
}

export const i18n = new I18n();
