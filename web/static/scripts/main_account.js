
const iconInput = document.querySelector(".account__icon--input");

iconInput.onchange = async function () {
  const userIcon = new FormData();
  userIcon.append("icon", iconInput.files[0]);

  try {
    const response = await fetch("/user/me/icon", {
      method: "PATCH",
      body: userIcon,
      credentials: "include"
    });

    if (!response.ok) {
      throw new Error(await response.text());
    }

    const newIconUrl = await response.text();
    const imgElement = document.querySelector(".account__icon") ||
      document.querySelector(".account__img");

    if (imgElement) {
      imgElement.className = "account__img";
      imgElement.style.backgroundImage = `url('${newIconUrl}')`;
    }
  } catch (error) {
    serverMessage("show", error.message);
  }
};

//-------------Открытие панели изменения пароля-----------
document.querySelector("#account-btn-open").addEventListener("click", () =>
  openAccountPassPanel("open"));
document.querySelector("#account-btn-close").addEventListener("click", () =>
  openAccountPassPanel("close"));

function openAccountPassPanel(action) {
  const panel = document.querySelector(".account__password--panel");
  const openBtn = document.getElementById("account-btn-open");

  if (action === "open") {
    panel.style.display = "flex";
    openBtn.style.display = "none";
    reAlertAccountMessages();
  } else {
    panel.style.display = "none";
    openBtn.style.display = "flex";
  }
}

const accpassInput = document.getElementById("account-pass-input");
const newPassInput = document.getElementById("account-newpass-input");
const rePassInput = document.getElementById("account-repass-input");
const changePassBtn = document.getElementById("account-btn-change");

// Проверка ввода текущего пароля
accpassInput.addEventListener("blur", validateCurrentPass);

function validateCurrentPass() {
  const pass = accpassInput.value.trim();
  if (pass.length <= 5) {
    showAlert("account-pass", "Введите пароль!");
  } else {
    clearAlert("account-pass");
  }
}

// Проверка ввода нового пароля
newPassInput.addEventListener("blur", validateNewPass);

function validateNewPass() {
  const newPass = newPassInput.value.trim();
  if (!isValidPass(newPass)) {
    showAlert("account-newpass", "Пароль недостаточно надежный!");
  } else {
    clearAlert("account-newpass");
  }
}

// Проверка подтверждения пароля
rePassInput.addEventListener("input", validateRePass);

function validateRePass() {
  const rePass = rePassInput.value.trim();
  const newPass = newPassInput.value.trim();

  if (rePass !== newPass) {
    showAlert("account-repass", "Пароли не совпадают!");
  } else {
    clearAlert("account-repass");
  }
}

// Проверка пароля регулярным выражением
function isValidPass(pass) {
  const pattern = /^.*(?=.{6,})(?=.*[a-zA-ZА-ЯЁа-яё]).*$/;
  return pattern.test(pass);
}

// Показать сообщение об ошибке
function showAlert(input, message) {
  const inputElement = document.getElementById(`${input}-input`);
  const labelElement = document.getElementById(`${input}-lable`);
  const messageElement = document.getElementById(`${input}-message`);

  inputElement.style.borderColor = "var(--error-color)";
  labelElement.style.color = "var(--error-color)";
  messageElement.textContent = message;
  messageElement.classList.add("error");
  changePassBtn.disabled = true;
}

// Очистить сообщение об ошибке
function clearAlert(input) {
  const inputElement = document.getElementById(`${input}-input`);
  const labelElement = document.getElementById(`${input}-lable`);
  const messageElement = document.getElementById(`${input}-message`);

  inputElement.style.borderColor = "var(--low-color)";
  labelElement.style.color = "var(--low-color)";
  messageElement.textContent = "";
  messageElement.classList.remove("error");

  // Проверяем все поля перед разблокировкой кнопки
  if (!document.querySelectorAll(".error").length) {
    changePassBtn.disabled = false;
  }
}

// Изменение пароля
changePassBtn.addEventListener("click", async function () {
  validateCurrentPass();
  validateNewPass();
  validateRePass();

  if (changePassBtn.disabled) return;

  try {
    const response = await fetch("/user/me/pass", {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        pass: accpassInput.value.trim(),
        newpass: newPassInput.value.trim()
      }),
      credentials: "include"
    });

    if (!response.ok) {
      const error = await response.text();
      if (response.status === 400) {
        console.log(error)
        if (error.includes("текущий пароль")) {
          showAlert("account-pass", "Текущий пароль неверный!");
        } else showAlert("account-newpass", "Неверный формат нового пароля!");
      } else {
        throw new Error(error);
      }
      return;
    }

    openAccountPassPanel("close");
    const messageElement = document.getElementById("account-chengepass-message");
    messageElement.textContent = "Пароль успешно изменён!";
    messageElement.classList.add("message");

    // Очищаем поля
    accpassInput.value = "";
    newPassInput.value = "";
    rePassInput.value = "";

  } catch (error) {
    serverMessage("show", error.message);
  }
});

function reAlertAccountMessages() {
  const messageElement = document.getElementById("account-chengepass-message");
  messageElement.textContent = "";
  messageElement.classList.remove("message");
}

// Выход пользователя из учетной записи
document.querySelector("#account-btn-exit").addEventListener("click", async function () {
  if (confirm("Вы действительно хотите выйти?")) {
    try {
      const response = await fetch("/auth/logout", {
        method: "DELETE",
        credentials: "include"
      });

      if (response.ok) {
        window.location.href = "/";
      } else {
        throw new Error(await response.text());
      }

    } catch (error) {
      serverMessage("show", error.message);
    }
  }
});

function loadScript(src) {
  return new Promise((resolve, reject) => {
    const script = document.createElement("script");
    script.src = src;
    script.onload = resolve;
    script.onerror = reject;
    document.head.appendChild(script);
  });
}