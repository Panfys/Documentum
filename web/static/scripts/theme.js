// Управление темой и цветом
function applyThemeSettings() {
  // Применение темы
  const savedTheme = localStorage.getItem("theme") || "dark";
  
  document.body.classList.toggle("light-theme", savedTheme === "light");
  document.body.classList.toggle("dark-theme", savedTheme === "dark");

  // Применение цвета
  const savedColor = localStorage.getItem("color") || "blue";
  const colorValues = {
    blue: "45, 104, 248",
    orange: "255, 104, 0",
    purple: "116, 66, 200",
    green: "3, 108, 86"
  };

  document.body.style.setProperty("--main-rgb", colorValues[savedColor]);
}

// Инициализация темы при загрузке
function initializeTheme() {
  if (!localStorage.getItem("theme")) {
    localStorage.setItem("theme", "dark");
  }
  
  if (!localStorage.getItem("color")) {
    localStorage.setItem("color", "blue");
  }

  applyThemeSettings();
}

// Вызываем инициализацию при загрузке
initializeTheme();
