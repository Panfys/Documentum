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