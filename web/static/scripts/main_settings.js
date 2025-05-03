// Константы для цветовых тем
const COLOR_VALUES = {
  blue: "45, 104, 248",
  orange: "255, 104, 0",
  purple: "116, 66, 200",
  green: "3, 108, 86"
};

// Переключение темы
function setupThemeSwitcher() {
  const themeButtons = document.querySelectorAll(".settings__theme--table");
  
  themeButtons.forEach(button => {
    button.addEventListener("click", () => {
      const theme = button.getAttribute("theme");
      const color = button.getAttribute("color");
      
      // Сохраняем настройки
      localStorage.setItem("theme", theme);
      localStorage.setItem("color", color);
      
      // Применяем изменения
      applyThemeSettings();
    });
  });
}

// Переключение панелей настроек
function setupSettingsPanels() {
  const settingButtons = document.querySelectorAll(".main__settings--btn");
  
  settingButtons.forEach(button => {
    button.addEventListener("click", () => {
      const panelId = button.getAttribute("panel-id");
      console.log("TAB")
      toggleActivePanel(button, panelId);
    });
  });
}

// Управление активной панелью
function toggleActivePanel(activeButton, panelId) {
  // Проверяем, является ли кнопка уже активной
  const isAlreadyActive = activeButton.classList.contains("main__settings-active-btn");
  
  // Удаляем активные классы у всех элементов
  document.querySelector(".main__settings-active-btn")?.classList.remove("main__settings-active-btn");
  document.querySelector(".main__settings--active-panel")?.classList.remove("main__settings--active-panel");
  
  // Если кнопка не была активной и panelId существует - активируем
  if (!isAlreadyActive && panelId) {
    const targetPanel = document.getElementById(panelId);
    if (targetPanel) {
      activeButton.classList.add("main__settings-active-btn");
      targetPanel.classList.add("main__settings--active-panel");
    }
  }
}

setupThemeSwitcher();
setupSettingsPanels();
