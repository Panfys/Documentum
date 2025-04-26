//------Функция добавления фотографии-------------
icon = document.querySelector(".account__icon--intut");

icon.onchange = function () {
  user_icon = new FormData();
  user_icon.append("icon", icon.files[0]);

  $.ajax({
    method: "POST",
    url: "/protect/users/updateicon",
    data: user_icon,
    cache: false,
    contentType: false,
    processData: false,
    success: function (check) {
      if (document.querySelector(".account__icon")) {
        document.querySelector(".account__icon").className = "account__img";
      }
      img = document.querySelector(".account__img");

      img.style.backgroundImage = "url('')";
      img.style.backgroundImage = "url('" + check + "')";
    },
    error: function (jqXHR) {
      ServerMessage("show", jqXHR.responseText);
    },
  });
};

//-------------Открытие панели изменения пароля-----------
document.querySelector("#account-btn-open").onclick = () =>
  OpenAccountPassPanel("open");
document.querySelector("#account-btn-close").onclick = () =>
  OpenAccountPassPanel("close");

function OpenAccountPassPanel(act) {
  if (act == "open") {
    document.querySelector(".account__password--panel").style = "display: flex";
    document.getElementById("account-btn-open").style = "display: none";
  } else {
    document.querySelector(".account__password--panel").style = "display: none";
    document.getElementById("account-btn-open").style = "display: flex";
  }
}

pass_input = document.getElementById("account-pass-input");
newpass_input = document.getElementById("account-newpass-input");
repass_input = document.getElementById("account-repass-input");

//Проверка ввода нового пароля
pass_input.addEventListener("blur", WritePass);

function WritePass() {
  pass = pass_input.value.trim();
  if (pass.length <= 5) AlertMessages("account-pass", "Введите пароль!");
  else ReAlertMessages("account-pass");
}

//Проверка ввода нового пароля
newpass_input.addEventListener("blur", WriteNewPass);

function WriteNewPass() {
  newpass = newpass_input.value.trim();
  if (!IsValidPass(newpass))
    AlertMessages("account-newpass", "Пароль недостаточно надежный!");
  else ReAlertMessages("account-newpass");
}

//Проверка подтверждения пароля
repass_input.addEventListener("input", WriteRePass);

function WriteRePass() {
  repass = repass_input.value.trim();
  if (repass !== newpass)
    AlertMessages("account-repass", "Пароли не совпадают!");
  else ReAlertMessages("account-repass");
}

// Проверка пароля регулярным выражением
function IsValidPass(pass) {
  const pattern = /^.*(?=.{6,})(?=.*[a-zA-ZА-ЯЁа-яё]).*$/;
  return pattern.test(pass);
}

//Сообщение при ошибке
function AlertMessages(input, mess) {
  document.getElementById(input.concat("-input")).style.borderColor =
    "--error-color";
  document.getElementById(input.concat("-lable")).style.color = "--error-color";
  document.getElementById(input.concat("-message")).innerHTML = mess;
  document.getElementById(input.concat("-message")).classList.add("error");
  document.getElementById("account-btn-change").disabled = true;
}

//очистка сообщения
function ReAlertMessages(input) {
  document.getElementById(input.concat("-input")).style.borderColor =
    "--low-color";
  document.getElementById(input.concat("-lable")).style.color = "--low-color";
  document.getElementById(input.concat("-message")).innerHTML = "";
  document.getElementById(input.concat("-message")).classList.remove("error");
  document.getElementById("account-btn-change").disabled = false;
}

document.querySelector("#account-btn-change").onclick = function () {
  pass = pass_input.value.trim();
  newpass = newpass_input.value.trim();
  repass = repass_input.value.trim();

  WritePass();
  WriteNewPass();

  if (document.querySelector("#account-btn-change").disabled) return;

  $.ajax({
    method: "POST",
    url: "/protect/users/updatepass",
    data: {
      pass: pass,
      newpass: newpass,
      repass: repass,
    },
    success: function () {
      OpenAccountPassPanel("close");
      document.getElementById("account-chengepass-message").innerHTML =
        "Пароль успешно изменён!";
      document
        .getElementById("account-chengepass-message")
        .classList.add("message");
    },
    error: function (jqXHR) {
      if (jqXHR.status == 400) {
        AlertMessages("account-repass", jqXHR.responseText);
      } else ServerMessage("show", jqXHR.responseText);
    },
  });
};

function ReAlertAccountMessages() {
  document.getElementById("account-chengepass-message").innerHTML = "";
  document
    .getElementById("account-chengepass-message")
    .classList.remove("message");
}

// Выход пользователя из учетной записи
document.querySelector("#account-btn-exit").onclick = function () {
  if (confirm("Вы действительно хотите выйти?")) {
    $.ajax({
      method: "POST",
      url: "/protect/users/exit",
      success: function (check) {
        if ((document.querySelector(".container").innerHTML = check)) {
          script = document.createElement("script");
          script.src = "/static/scripts/entrance.js";
          document.head.appendChild(script);
        }
      },
      error: function (jqXHR) {
        ServerMessage("show", jqXHR.responseText);
      },
    });
  }
};
