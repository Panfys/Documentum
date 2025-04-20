//------------------------------ENTRANCE-------------------------------------

/*/-------Прелоадер
window.onload = function() {
  document.querySelector('.main__preloader--panel').style.display = 'none';
}; */

//var
btn_show_password = document.querySelector(".password__checkbox");
login_input = document.querySelector("#regist-login-input");
name_input = document.querySelector("#regist-name-input");
func_input = document.querySelector("#regist-func-input");
unit_input = document.querySelector("#regist-unit-input");
group_input = document.querySelector("#regist-group-input");
pass_input = document.querySelector("#regist-pass-input");
repass_input = document.querySelector("#regist-repass-input");
regist_button = document.querySelector("#regist-button");
author_button = document.querySelector("#author-button");
error_registration = false;

// Проверка работоспособности сервисов
$.ajax({
  method: "POST",
  url: "/check",
  success: (check) => {
    if (check != "ok") ServerMessage("show", check);
  },
});

//-------Показать пароль
btn_show_password.addEventListener("click", ShowPassword);

function ShowPassword() {
  password_status = document.querySelector("#auth-password-input");
  if (btn_show_password.checked) password_status.type = "text";
  else password_status.type = "password";
}

//Проверка ввода логина
login_input.addEventListener("blur", WriteLogin);

function WriteLogin() {
  login = login_input.value.trim();
  if (login === "") AlertMessages("regist-login", "Введите логин!");
  else if (login.length < 3)
    AlertMessages("regist-login", "Минимальная длина логина - 3 символа!");
  else if (login.length > 12)
    AlertMessages("regist-login", "Максимальная длина логина - 12 символов!");
  else if (!IsValidLogin(login))
    AlertMessages(
      "regist-login",
      "Используйте только латинские буквы и цифры!"
    );
  else ReAlertMessages("regist-login");
}

//Проверка ввода имени
name_input.addEventListener("blur", WriteName);

function WriteName() {
  name = name_input.value.trim();
  if (name === "") AlertMessages("regist-name", "Введите фамилию и инициалы!");
  else if (!IsValidName(name))
    AlertMessages("regist-name", "Введенные данные некорректны!");
  else ReAlertMessages("regist-name");
}

func_input.addEventListener("blur", WriteFunc);
func_input.addEventListener("input", WriteFunc);
func_input.addEventListener("input", WriteGroups);

//Проверка ввода должности
function WriteFunc() {
  func = func_input.value;
  if (func === "0") AlertMessages("regist-func", "Укажите должность!");
  else {
    ReAlertMessages("regist-func");

    $.ajax({
      method: "POST",
      url: "/users/units",
      data: { func: func },
      success: function (units) {
        document.getElementById("regist-unit-input").innerHTML = units;
      },
      error: () =>
        ServerMessage("show", "Возникла ошибка на сервере, попробуйте позже!"),
    });
  }
}

//Изменение групп
unit_input.addEventListener("input", WriteGroups);
function WriteGroups() {
  func = func_input.value;
  unit = unit_input.value;

  $.ajax({
    method: "POST",
    url: "/users/groups",
    data: { unit: unit, func: func },
    success: function (groups) {
      if (groups.length > 0) {
        document.getElementById("regist-group-box").style.display = "block";
        document.getElementById("regist-group-input").innerHTML = groups;
      } else {
        document.getElementById("regist-group-box").style.display = "none";
      }
    },
    error: () =>
      ServerMessage("show", "Возникла ошибка на сервере, попробуйте позже!"),
  });
}

//Проверка ввода пароля
pass_input.addEventListener("blur", WritePass);

function WritePass() {
  pass = pass_input.value.trim();
  if (!IsValidPass(pass))
    AlertMessages("regist-pass", "Пароль недостаточно надежный!");
  else ReAlertMessages("regist-pass");
}

//Проверка подтверждения пароля
repass_input.addEventListener("input", WriteRePass);

function WriteRePass() {
  repass = repass_input.value.trim();
  if (repass !== pass) AlertMessages("regist-repass", "Пароли не совпадают!");
  else ReAlertMessages("regist-repass");
}

// Проверка регистрации
regist_button.addEventListener("click", Walid);

function Walid() {
  user = {
    login: login_input.value.trim(),
    name: name_input.value.trim(),
    func: func_input.value.trim(),
    unit: unit_input.value.trim(),
    group: group_input.value.trim(),
    pass: pass_input.value.trim(),
    repass: repass_input.value.trim(), 
  };
  error_registration = false;
  //WriteLogin();
  //WriteName();
  //if (user["func"] === "0") AlertMessages("regist-func", "Укажите должность!");
  //WritePass();
  //WriteRePass();
  if (error_registration == true) return;

  $.ajax({
    method: "POST",
    url: "users/add",
    contentType: "application/json",
    data: JSON.stringify(user),
    success: function (check) {
      if (check == "ok") {
        document.getElementById("auth-link").click();
        document.getElementById("auth-message").innerHTML =
          "Аккаунт зарегистрирован, входите!";
        document.getElementById("auth-message").classList.add("message");
      } else ServerMessage("show", check);
    },
    error: function (jqXHR, exception) {
      ServerMessage("show", jqXHR.responseText); 
    },
  });
}
// Авторизация
author_button.addEventListener("click", Authorization);

function Authorization() {
  auth_message = document.querySelector("#auth-message");
  auth_message.innerHTML = "";
  auth_message.classList.remove("error");
  auth_message.classList.remove("message");
  login = document.getElementById("auth-login-input").value;
  pass = document.getElementById("auth-password-input").value;
  remember = document.getElementById("remember").checked;

  $.ajax({
    method: "POST",
    url: "/users/auth",
    data: {
      login: login,
      pass: pass,
      remember: remember,
    },
    success: function (check) {
      if ((document.querySelector(".container").innerHTML = check)) {
        script = document.createElement("script");
        script.src = "/static/scripts/script_main.js";
        document.head.appendChild(script);
      }
    },
    error: function (jqXHR) {
      if (jqXHR.status == 401) {
        document.getElementById("auth-message").innerHTML =
          "Неверный логин или пароль!";
        document.getElementById("auth-message").classList.add("error");
      } else ServerMessage("show", jqXHR.responseText);
    },
  });
}

// Проверка логина регулярным выражением
function IsValidLogin(login) {
  const pattern = /^[a-zA-Z0-9-_]+$/;
  return pattern.test(login);
}

// Проверка имени регулярным выражением
function IsValidName(name) {
  const pattern = /^[А-ЯЁ][а-яё]+[ ][А-ЯЁ][.][А-ЯЁ][.]$/;
  return pattern.test(name);
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
  error_registration = true;
}

//очистка сообщения
function ReAlertMessages(input) {
  document.getElementById(input.concat("-input")).style.borderColor =
    "--low-color";
  document.getElementById(input.concat("-lable")).style.color = "--low-color";
  document.getElementById(input.concat("-message")).innerHTML = "";
  document.getElementById(input.concat("-message")).classList.remove("error");
}
