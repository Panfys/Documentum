//сообщения от сервера
server_message_btn = document.querySelector(".server__messages--btn");
server_message_btn.onclick = () => ServerMessage("clouse", "");

function ServerMessage(action, message) {
  div = document.querySelector(".server__messages");
  text = document.querySelector(".server__messages--text");
  switch (action) {
    case "show":
      {
        new_date = new Date();
        now_date = new_date.toLocaleString();
        div.style.display = "flex";
        text.innerHTML +=
          "(" + now_date + ") " + "Сообщение от сервера: " + message + "</br>";
        setTimeout(ServerMessage, 10000, "clouse", "");
      }
      break;
    case "clouse":
      {
        div.style.display = "none";
        setTimeout(ServerMessage, 600000, "clear", "");
      }
      break;
    case "clear":
      {
        text.innerHTML = "";
      }
      break;
  }
}

//Определение темы
// Вывод страницы входа
if (localStorage.getItem("theme")) {
  if (localStorage.getItem("theme") == "light") {
    document.body.classList.add("light-theme");
    document.body.classList.remove("dark-theme");
  } else if (localStorage.getItem("theme") == "dark") {
    document.body.classList.remove("light-theme");
    document.body.classList.add("dark-theme");
  }
} else {
  localStorage.setItem("theme", "dark");
  document.body.classList.remove("light-theme");
  document.body.classList.add("dark-theme");
}
if (localStorage.getItem("color")) {
  switch (localStorage.getItem("color")) {
    case "blue":
      {
        document.body.style.setProperty("--main-rgb", "45, 104, 248");
      }
      break;
    case "orange":
      {
        document.body.style.setProperty("--main-rgb", "255, 104, 0");
      }
      break;
    case "purple":
      {
        document.body.style.setProperty("--main-rgb", "116, 66, 200");
      }
      break;
    case "green":
      {
        document.body.style.setProperty("--main-rgb", "3, 108, 86");
      }
      break;
  }
} else {
  localStorage.setItem("color", "blue");
  document.body.style.setProperty("--main-rgb", "45, 104, 248");
}

/*   ПОЛЕЗНОЕ

События для всех див
var divs = document.querySelectorAll("div");

for (var i = 0; i < divs.length; i++) {
  divs[i].onclick = function (e) {
    e.target.style.backgroundColor = bgChange();
  };
}

Валидация 
function isValidLogin(login) {
  // Проверка имени регулярным выражением
  const pattern = /^[a-zA-Z0-9]+$/;
  return pattern.test(login);
}

function isValidPassword(password) {
  // Проверка пароля регулярным выражением
  const pattern = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[a-zA-Z0-9]{8,20}$/;
  return pattern.test(password);
}
*/
