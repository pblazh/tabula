/**
 * Load CSS content and inject it into the document head
 */
export async function loadStylesheet(cssContent: string, id?: string): Promise<() => void> {
  const style = document.createElement("style");
  
  if (id) {
    style.id = id;
    // Remove existing stylesheet with same ID
    const existing = document.getElementById(id);
    if (existing) {
      existing.remove();
    }
  }
  
  style.textContent = cssContent;
  document.head.appendChild(style);
  
  // Return cleanup function
  return () => {
    if (style.parentNode) {
      style.parentNode.removeChild(style);
    }
  };
}

/**
 * Load CSS file from the styles directory
 */
export async function loadCSSFile(filename: string, id?: string): Promise<() => void> {
  try {
    // In production, CSS files should be bundled or loaded differently
    // For now, we'll use a simple approach that works with the build system
    const response = await fetch(`./styles/${filename}`);
    if (!response.ok) {
      throw new Error(`Failed to load CSS file: ${filename}`);
    }
    const cssContent = await response.text();
    return loadStylesheet(cssContent, id);
  } catch (error) {
    console.warn(`Could not load CSS file ${filename}:`, error);
    // Return no-op cleanup function
    return () => {};
  }
}