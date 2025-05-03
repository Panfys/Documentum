// Серверные сообщения
const serverMessageBtn = document.querySelector(".server__messages--btn");
serverMessageBtn.addEventListener("click", () => serverMessage("close", ""));

function serverMessage(action, message) {
  const messageDiv = document.querySelector(".server__messages");
  const messageText = document.querySelector(".server__messages--text");

  switch (action) {
    case "show":
      {
        const now = new Date().toLocaleString();
        messageDiv.style.display = "flex";
        messageText.innerHTML = `(${now}) Сообщение от сервера: ${message}<br>`;
        setTimeout(() => serverMessage("close", ""), 10000);
      }
      break;
    case "close":
      {
        messageDiv.style.display = "none";
        setTimeout(() => serverMessage("clear", ""), 600000);
      }
      break;
    case "clear":
      {
        messageText.innerHTML = "";
      }
      break;
  }
}

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