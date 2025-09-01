import { i18n } from "../../i18n";
import { Action } from "./types";

export const render = (items: Action[], x: number, y: number): VoidFunction => {
  const menu = document.createElement("div");
  menu.className = "context-menu";

  const outsideHandler = (ev: MouseEvent) => {
    if (!menu.contains(ev.target as Node)) closeMenu();
  };

  const keyHandler = (ev: KeyboardEvent) => {
    if (ev.key === "Escape") closeMenu();
  };

  const closeMenu = () => {
    menu.remove();
    document.removeEventListener("mousedown", outsideHandler!);
    document.removeEventListener("keydown", keyHandler!);
  };

  document.addEventListener("mousedown", outsideHandler!);
  document.addEventListener("keydown", keyHandler!);

  Object.assign(menu.style, {
    left: `${x}px`,
    top: `${y}px`,
  });

  items.forEach((item) => {
    const menuItem = document.createElement("div");
    menuItem.className = "context-menu-item";
    menuItem.textContent = i18n.t(item.label);
    menuItem.onclick = (ev) => {
      ev.stopPropagation();
      ev.preventDefault();
      item.action();
      closeMenu();
    };
    menu.appendChild(menuItem);
  });
  document.body.appendChild(menu);

  return closeMenu;
};
