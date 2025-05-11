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
const authButton = document.querySelector("#author-button");
const authMessage = document.querySelector("#auth-message");
const groupBox = document.getElementById("regist-group-box");
const groupMess = document.getElementById("regist-group-message");

// Проверка логина
loginInput.addEventListener("blur", () => ValidLogin(loginInput.value.trim(), "regist-login"));

// Проверка имени
nameInput.addEventListener("blur", () => ValidName(nameInput.value.trim(), "regist-name"));

// Получение структурных подразделений
funcInput.addEventListener("input", () => {
  GetUnits(funcInput.value)
  GetGroups(funcInput.value, unitInput.value)
});

async function GetUnits(func) {
  if (ValidFunc(func, "regist-func")) {
    unitInput.innerHTML = await FetchUnits(func);
  }
}

// Получение подразделений
unitInput.addEventListener("input", () => GetGroups(funcInput.value, unitInput.value));

async function GetGroups(func, unit) {
  if (ValidUnit(unit, "regist-unit")) {
    groupInput.innerHTML = await FetchGroups(func, unit);
    ValidGroups(groupInput.innerHTML, groupBox, groupMess)
  }
}

// Проверка пароля
passInput.addEventListener("blur", () => ValidPass(passInput.value, "regist-pass"));

// Проверка подтверждения пароля
repassInput.addEventListener("input", () => ValidRepass(passInput.value, repassInput.value, "regist-repass"));

// Регистрации
registButton.addEventListener("click", Registration);

async function Registration() {
  const user = {
    login: loginInput.value.trim(),
    name: nameInput.value.trim(),
    func: funcInput.value.trim(),
    unit: unitInput.value.trim(),
    group: groupInput.value.trim(),
    pass: passInput.value,
    repass: repassInput.value,
  };

  if (ValidRegistration(user, groupBox, groupMess)) return;

  const successReg = await FetchRegistration(user);
  if (successReg) {
    document.getElementById("auth-link").click();
    const authMessage = document.getElementById("auth-message");
    authMessage.innerHTML = "Аккаунт зарегистрирован, входите!";
    authMessage.classList.add("message");
  }
}

// Авторизация
authButton.addEventListener("click", authorize);

async function authorize() {
  authMessage.textContent = "";
  authMessage.classList.remove("error", "message");

  const authData = {
    login: document.getElementById("auth-login-input").value.trim(),
    pass: document.getElementById("auth-password-input").value,
    remember: document.getElementById("remember").checked,
  };

  const successAuth = await FetchAuthorize(authData);
  if (successAuth === "authError") {
    authMessage.textContent = "Неверный логин или пароль!";
    authMessage.classList.add("error");
  } else if (successAuth != "error") {
    document.querySelector(".container").innerHTML = successAuth;

    // Динамическая загрузка скриптов
    await loadScripts([
      "/static/scripts/main.js",
      "/static/scripts/main_account.js",
      "/static/scripts/main_settings.js",
      "/static/scripts/main_header_panel.js",
      "/static/scripts/main_open_files.js",
      "/static/scripts/main_panel_buttons.js",
      "/static/scripts/main_res_buttons.js",
      "/static/scripts/main_outgoing.js"
    ]);

     initDocumentViewHandlers();
     initDocumentHandlers();
     initResolutionHandlers();
     updateInputs();
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

//-------Показать пароль
btnShowPassword.addEventListener("click", showPassword);

function showPassword() {
  const passwordStatus = document.querySelector("#auth-password-input");
  passwordStatus.type = btnShowPassword.checked ? "text" : "password";
}