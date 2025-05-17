// Переменные
let ErrorAuthMessages = false;
let timeoutId = null;

// Серверные сообщения
const serverMessageBtn = document.querySelector(".server__messages--btn");
serverMessageBtn.addEventListener("click", () => serverMessage("close", ""));

function serverMessage(action, message) {
  const messageDiv = document.querySelector(".server__messages");
  const messageText = document.querySelector(".server__messages--text");

  switch (action) {
    case "show":
      {
        // Если уже есть активный таймер, отменяем его
        if (timeoutId) {
          clearTimeout(timeoutId);
          timeoutId = null;
        }

        const now = new Date().toLocaleString();
        messageText.innerHTML = `(${now}) Сообщение от сервера: ${message}<br>`;
        // Сохраняем ID нового таймера
        timeoutId = setTimeout(() => serverMessage("close", ""), 10000);
        messageDiv.style.display = "flex";
      }
      break;
    case "close":
      {
        // Если есть активный таймер, отменяем его
        if (timeoutId) {
          clearTimeout(timeoutId);
          timeoutId = null;
        }

        messageDiv.style.display = "none";
        messageText.innerHTML = "";
      }
      break;
    case "clear":
      {
        messageText.innerHTML = "";
      }
      break;
  }
}

// Сообщение при ошибке в форму регистрации и авторизации
function AlertAuthMessages(input, mess) {
  document.getElementById(`${input}-input`).style.borderColor = "var(--error-color)";
  document.getElementById(`${input}-lable`).style.color = "var(--error-color)";
  const messageElement = document.getElementById(`${input}-message`);
  messageElement.innerHTML = mess;
  messageElement.classList.add("error");
  ErrorAuthMessages = true;
}

// Очистка сообщения в форме регистрации и авторизации
function ReAlertAuthMessages(input) {
  document.getElementById(`${input}-input`).style.borderColor = "var(--low-color)";
  document.getElementById(`${input}-lable`).style.color = "var(--low-color)";
  const messageElement = document.getElementById(`${input}-message`);
  messageElement.innerHTML = "";
  messageElement.classList.remove("error");
}

//---Ошибка введенных данных нового документа----
function AlertValidDocError(input) {
  const activeTab = document.querySelector(".main__tabs--active");
  activeTab.querySelector("#input-newdoc-" + input).style.borderColor =
    "var(--error-color)";
  activeTab.querySelector("#input-newdoc-" + input).style.color =
    "var(--error-color)";
}

//---Удаление ошибки введенных нового данных----
function ReAlertValidDocError(input) {
  const activeTab = document.querySelector(".main__tabs--active");
  input = activeTab.querySelector("#input-newdoc-" + input);

  if (input) {
    input.style.borderColor = "var(--low-color)";
    input.style.color = "var(--mid-color)";
  }
}

//---Ошибка введенных данных резолюции----
function ValidResolutionError(resolution_id, input) {
  resolution = document.getElementById(resolution_id);
  input = "#resolution-" + input;

  resolution.querySelector(input).style.borderColor = "var(--error-color)";
  resolution.querySelector(input).style.color = "var(--error-color)";
}

//---Удаление ошибки введенных данных резолюции----
function ValidResolutionReError(resolution_id, input) {
  resolution = document.getElementById(resolution_id);
  input = "#resolution-" + input;

  if (resolution.querySelector(input)) {
    resolution.querySelector(input).style.borderColor = "var(--low-color)";
    resolution.querySelector(input).style.color = "var(--mid-color)";
  }
}