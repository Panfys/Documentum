//------------------------------ENTRANCE-------------------------------------

/*/-------Прелоадер
window.onload = function() {
  document.querySelector('.main__preloader--panel').style.display = 'none';
}; */

// Variables
const btnShowPassword = document.querySelector(".password__checkbox");
const loginInput = document.querySelector("#regist-login-input");
const nameInput = document.querySelector("#regist-name-input");
const funcInput = document.querySelector("#regist-func-input");
const unitInput = document.querySelector("#regist-unit-input");
const groupInput = document.querySelector("#regist-group-input");
const passInput = document.querySelector("#regist-pass-input");
const repassInput = document.querySelector("#regist-repass-input");
const registButton = document.querySelector("#regist-button");
const authorButton = document.querySelector("#author-button");
let errorRegistration = false;
let currentPass = "";

//-------Показать пароль
btnShowPassword.addEventListener("click", showPassword);

function showPassword() {
  const passwordStatus = document.querySelector("#auth-password-input");
  passwordStatus.type = btnShowPassword.checked ? "text" : "password";
}

// Проверка ввода логина
loginInput.addEventListener("blur", writeLogin);

function writeLogin() {
  const login = loginInput.value.trim();
  if (login === "") {
    alertMessages("regist-login", "Введите логин!");
  } else if (login.length < 3) {
    alertMessages("regist-login", "Минимальная длина логина - 3 символа!");
  } else if (login.length > 12) {
    alertMessages("regist-login", "Максимальная длина логина - 12 символов!");
  } else if (!isValidLogin(login)) {
    alertMessages("regist-login", "Используйте только латинские буквы и цифры!");
  } else {
    reAlertMessages("regist-login");
  }
}

// Проверка ввода имени
nameInput.addEventListener("blur", writeName);

function writeName() {
  const name = nameInput.value.trim();
  if (name === "") {
    alertMessages("regist-name", "Введите фамилию и инициалы!");
  } else if (!isValidName(name)) {
    alertMessages("regist-name", "Введенные данные некорректны!");
  } else {
    reAlertMessages("regist-name");
  }
}

funcInput.addEventListener("blur", writeFunc);
funcInput.addEventListener("input", writeFunc);
funcInput.addEventListener("input", writeGroups);

// Проверка ввода должности
function writeFunc() {
  const func = funcInput.value;
  if (func === "0") {
    alertMessages("regist-func", "Укажите должность!");
  } else {
    reAlertMessages("regist-func");
    fetchUnits(func);
  }
}

// Получение подразделений
async function fetchUnits(func) {
  try {
    const response = await fetch(`/funcs/${encodeURIComponent(func)}/units`, {
      method: "GET",
      headers: {
        "Accept": "text/html",
      },
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(errorText || "Ошибка сервера");
    }

    const units = await response.text();
    document.getElementById("regist-unit-input").innerHTML = units;
  } catch (error) {
    serverMessage("show", error.message || "Возникла ошибка на сервере, попробуйте позже!");
  }
}

// Получение групп
unitInput.addEventListener("input", writeGroups);

async function writeGroups() {
  const func = funcInput.value;
  const unit = unitInput.value;

  try {
    // Формируем URL с параметрами
    const url = `/funcs/${encodeURIComponent(func)}/${encodeURIComponent(unit)}/groups`;

    const response = await fetch(url, {
      method: "GET",
      headers: {
        "Accept": "text/html", // Указываем ожидаемый тип ответа
      },
    });

    if (!response.ok) {
      // Пытаемся получить текст ошибки от сервера
      const errorText = await response.text();
      throw new Error(errorText || "Ошибка сервера");
    }

    const groups = await response.text();
    const groupBox = document.getElementById("regist-group-box");

    if (groups.length > 0) {
      groupBox.style.display = "block";
      document.getElementById("regist-group-input").innerHTML = groups;
    } else {
      groupBox.style.display = "none";
    }
  } catch (error) {
    // Выводим конкретное сообщение об ошибке
    serverMessage("show", error.message || "Возникла ошибка на сервере, попробуйте позже!");

    // Скрываем блок групп при ошибке
    document.getElementById("regist-group-box").style.display = "none";
  }
}

// Проверка ввода пароля
passInput.addEventListener("blur", writePass);

function writePass() {
  currentPass = passInput.value.trim();
  if (!isValidPass(currentPass)) {
    alertMessages("regist-pass", "Пароль недостаточно надежный!");
  } else {
    reAlertMessages("regist-pass");
  }
}

// Проверка подтверждения пароля
repassInput.addEventListener("input", writeRePass);

function writeRePass() {
  const repass = repassInput.value.trim();
  if (repass !== currentPass) {
    alertMessages("regist-repass", "Пароли не совпадают!");
  } else {
    reAlertMessages("regist-repass");
  }
}

// Проверка регистрации
registButton.addEventListener("click", validateRegistration);

async function validateRegistration() {
  const user = {
    login: loginInput.value.trim(),
    name: nameInput.value.trim(),
    func: funcInput.value.trim(),
    unit: unitInput.value.trim(),
    group: groupInput.value.trim(),
    pass: passInput.value.trim(),
    repass: repassInput.value.trim(),
  };

  errorRegistration = false;
  writeLogin();
  writeName();

  if (user.func === "0") {
    alertMessages("regist-func", "Укажите должность!");
  }

  writePass();
  writeRePass();

  if (errorRegistration) return;

  try {
    const response = await fetch("/auth/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    });

    if (!response.ok) {
      throw new Error(await response.text());
    }

    document.getElementById("auth-link").click();
    const authMessage = document.getElementById("auth-message");
    authMessage.innerHTML = "Аккаунт зарегистрирован, входите!";
    authMessage.classList.add("message");

  } catch (error) {
    serverMessage("show", error.message || "Возникла ошибка на сервере, попробуйте позже!");
  }
}

// Авторизация
authorButton.addEventListener("click", authorize);

async function authorize() {
  const authMessage = document.querySelector("#auth-message");
  authMessage.textContent = "";
  authMessage.classList.remove("error", "message");

  const login = document.getElementById("auth-login-input").value;
  const pass = document.getElementById("auth-password-input").value;
  const remember = document.getElementById("remember").checked;

  try {
    const response = await fetch("/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
        "X-Requested-With": "XMLHttpRequest"
      },
      body: new URLSearchParams({
        login: login,
        pass: pass,
        remember: remember
      }),
      credentials: "include"
    });

    if (!response.ok) {
      const error = await response.text();
      throw new Error(
        response.status === 401
          ? "authError"
          : error || "Ошибка сервера"
      );
    }

    const html = await response.text();
    document.querySelector(".container").innerHTML = html;

    // Динамическая загрузка скриптов
    await loadScripts([
      "/static/scripts/main.js",
      "/static/scripts/main_account.js",
      "/static/scripts/main_settings.js"
    ]);

  } catch (error) {
    if (error.message === "authError") {
      authMessage.textContent = "Неверный логин или пароль!";
      authMessage.classList.add("error");
    } else {
      serverMessage("show", error.message || "Возникла ошибка на сервере, попробуйте позже!");
    }
  }
}

// Проверка логина регулярным выражением
function isValidLogin(login) {
  const pattern = /^[a-zA-Z0-9-_]+$/;
  return pattern.test(login);
}

// Проверка имени регулярным выражением
function isValidName(name) {
  const pattern = /^[А-ЯЁ][а-яё]+[ ][А-ЯЁ][.][А-ЯЁ][.]$/;
  return pattern.test(name);
}

// Проверка пароля регулярным выражением
function isValidPass(pass) {
  const pattern = /^.*(?=.{6,})(?=.*[a-zA-ZА-ЯЁа-яё]).*$/;
  return pattern.test(pass);
}

// Сообщение при ошибке
function alertMessages(input, mess) {
  document.getElementById(`${input}-input`).style.borderColor = "var(--error-color)";
  document.getElementById(`${input}-lable`).style.color = "var(--error-color)";
  const messageElement = document.getElementById(`${input}-message`);
  messageElement.innerHTML = mess;
  messageElement.classList.add("error");
  errorRegistration = true;
}

// Очистка сообщения
function reAlertMessages(input) {
  document.getElementById(`${input}-input`).style.borderColor = "var(--low-color)";
  document.getElementById(`${input}-lable`).style.color = "var(--low-color)";
  const messageElement = document.getElementById(`${input}-message`);
  messageElement.innerHTML = "";
  messageElement.classList.remove("error");
}

// Показать сообщение сервера
function serverMessage(action, message) {
  // Реализация этой функции зависит от вашего UI
  console.log(message);
  // Пример реализации:
  const serverMessageElement = document.getElementById("server-message");
  if (action === "show") {
    serverMessageElement.textContent = message;
    serverMessageElement.style.display = "block";
  } else {
    serverMessageElement.style.display = "none";
  }
}

// Функция для загрузки скриптов
function loadScripts(urls) {
  return Promise.all(
    urls.map(url => {
      return new Promise((resolve, reject) => {
        const script = document.createElement("script");
        script.src = url;
        script.async = true;
        script.onload = resolve;
        script.onerror = reject;
        document.head.appendChild(script);
      });
    })
  );
}