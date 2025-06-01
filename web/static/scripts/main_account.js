
const iconInput = document.querySelector(".account__icon--input");

// Обновление иконки пользователя
iconInput.onchange = async function () {
  const userIcon = new FormData();
  userIcon.append("icon", iconInput.files[0]);

  const iconUrl = await FetchUpdateUserIcon(userIcon)
  const imgElement = document.querySelector(".account__icon") ||
    document.querySelector(".account__img");

  if (imgElement) {
    if (iconUrl !== "") {
      imgElement.className = "account__img";
      imgElement.style.backgroundImage = `url('${iconUrl}')`;
    }
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

function reAlertAccountMessages() {
  const messageElement = document.getElementById("account-chengepass-message");
  messageElement.textContent = "";
  messageElement.classList.remove("message");
}

const accpassInput = document.getElementById("account-pass-input");
const newPassInput = document.getElementById("account-newpass-input");
const rePassInput = document.getElementById("account-repass-input");
const changePassBtn = document.getElementById("account-btn-change");

// Проверка ввода текущего пароля
accpassInput.addEventListener("blur", () => ValidCurrentPass(accpassInput.value));

// Проверка ввода нового пароля
newPassInput.addEventListener("blur", () => ValidPass(newPassInput.value, "account-newpass"));

// Проверка подтверждения пароля
rePassInput.addEventListener("input", () => ValidRepass(rePassInput.value, newPassInput.value, "account-repass"));

// Изменение пароля
changePassBtn.addEventListener("click", async function () {
  const updatePass = {
    pass: accpassInput.value,
    newpass: newPassInput.value,
    repass: rePassInput.value
  }

  if (ValidUpdateUserPass(updatePass)) return

  const error = await FetchUpdateUserPass(updatePass)
  if (typeof error === 'string') {
    if (error.includes("текущий пароль")) {
      AlertAuthMessages("account-pass", "Текущий пароль неверный!");
    } else AlertAuthMessages("account-newpass", "Неверный формат нового пароля!");
    return
  } else if (!error) return

  openAccountPassPanel("close");
  const messageElement = document.getElementById("account-chengepass-message");
  messageElement.textContent = "Пароль успешно изменён!";
  messageElement.classList.add("message");

  // Очищаем поля
  accpassInput.value = "";
  newPassInput.value = "";
  rePassInput.value = "";
});

// Выход пользователя из учетной записи
document.querySelector("#account-btn-exit").addEventListener("click", async function () {
  if (confirm("Вы действительно хотите выйти?")) {
    FetchLogoutUser()
  }
});