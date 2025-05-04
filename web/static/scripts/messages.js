// Переменные
let ErrorAuthMessages = false;

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
