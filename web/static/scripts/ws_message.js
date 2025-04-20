class Event {
  constructor(type, payload) {
    this.type = type;
    this.payload = payload;
  }
}

class Payload {
  constructor(message, from) {
    this.message = message;
    this.from = from;
  }
}

document.getElementById("go").addEventListener("click", () => sendMessage("new_user", "mess", "from"));

function sendMessage(action, message, from) {
    let payload = new Payload(message, from);
    const event = new Event(action, payload);
    socket.send(JSON.stringify(event));
    document.getElementById("otvet").innerHTML +=
      "Сообщение от клиента: " + JSON.stringify(event) + "<br><br>";
  }

function routeEvent(event) {
  if (event.type === undefined) {
    document.getElementById("otvet").innerHTML =
      "Сообщение от сервера: no 'type' field in event<br>";
  }
  switch (event.type) {
    case "new_document":
      // format message
      document.getElementById("otvet").innerHTML +=
        "Сообщение от сервера:\n" + event.payload.message + "<br>";
      console.log(event);
      break;

    case "new_user":
      // format message
      document.getElementById("otvet").innerHTML +=
        "Сообщение от сервера:\n" + event.payload.message + "<br>";
      console.log(event);
      break;

    default:
      document.getElementById("otvet").innerHTML +=
        "unsupported message type<br>";
      break;
  }
}