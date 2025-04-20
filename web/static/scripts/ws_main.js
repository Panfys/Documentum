let socket;
let reconnectAttempts = 0;
const MAX_RECONNECT_ATTEMPTS = 20;

function tryToConnect() {
  socket = new WebSocket("ws://localhost:8000");

  socket.onopen = onSocketOpen;
  socket.onmessage = onSocketMessage;
  socket.onclose = onSocketClose;         
  socket.onerror = onSocketError;
}

function onSocketOpen(e) {
  reconnectAttempts = 0; // сбрасываем число попыток при успешном подключении
  document.getElementById("otvet").innerHTML += "Успешное соединение<br>";
}

function onSocketClose(e) {
  document.getElementById("otvet").innerHTML +=
    "Соединение закрыто. Попытка повторного подключения<br>";

  if (reconnectAttempts < MAX_RECONNECT_ATTEMPTS) {
    setTimeout(() => tryToConnect(), 4000); // повторное подключение через 2 секунды
    reconnectAttempts++;
  } else {
    document.getElementById("otvet").innerHTML +=
      "Слишком много попыток соединения. Соединение невозможно, попробуйте позже";
  }
}

function onSocketError(e) {
  document.getElementById("otvet").innerHTML +=
    "Ошибка соединения. Попытка повторного подключения<br>";

  if (reconnectAttempts < MAX_RECONNECT_ATTEMPTS) {
    setTimeout(() => tryToConnect(), 4000); // повторное подключение через 2 секунды
    reconnectAttempts++;
  } else {
    document.getElementById("otvet").innerHTML +=
      "Слишком много попыток соединения. Соединение невозможно, попробуйте позже";
  }
}

function onSocketMessage(e) {

  const eventData = JSON.parse(e.data);
  // Assign JSON data to new Event Object
  const event = Object.assign(new Event(), eventData);
  // Let router manage message
  routeEvent(event);
  
}

document.addEventListener("DOMContentLoaded", () => {
  tryToConnect(); // начинаем соединение после загрузки DOM
});

