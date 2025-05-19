// Валидация логина пользователя
function ValidLogin(login, input) {
  const pattern = /^[a-zA-Z0-9-_]+$/;
  if (login === "") {
    AlertAuthMessages(input, "Введите логин!");
  } else if (login.length < 3) {
    AlertAuthMessages(input, "Минимальная длина логина - 3 символа!");
  } else if (login.length > 12) {
    AlertAuthMessages(input, "Максимальная длина логина - 12 символов!");
  } else if (!pattern.test(login)) {
    AlertAuthMessages(input, "Используйте только латинские буквы и цифры!");
  } else {
    ReAlertAuthMessages(input);
  }
}

// Валидация имени пользователя
function ValidName(name, input) {
  const pattern = /^[А-ЯЁ][а-яё]+[ ][А-ЯЁ][.][А-ЯЁ][.]$/;
  if (name === "") {
    AlertAuthMessages(input, "Введите фамилию и инициалы!");
  } else if (!pattern.test(name)) {
    AlertAuthMessages(input, "Введенные данные некорректны!");
  } else {
    ReAlertAuthMessages(input);
  }
}

// Валидация ввода должности
function ValidFunc(func, input) {
  if (func === "0") {
    AlertAuthMessages(input, "Укажите должность!");
    return false
  } else {
    ReAlertAuthMessages(input);
    return true
  }
}

// Валидация ввода структурного подразделения
function ValidUnit(unit, input) {
  if (unit === "0") {
    AlertAuthMessages(input, "Укажите структурное подразделение!");
    return false
  } else {
    ReAlertAuthMessages(input);
    return true
  }
}

// Валидация ввода подразделения
function ValidGroups(groups, groupBox, groupMess) {
  if (groups != `<option value="1"></option>`) {
    groupBox.style.display = "block";
    groupMess.style.display = "block";
  } else {
    groupBox.style.display = "none";
    groupMess.style.display = "none";
  }
}

// Валидация пароля 
function ValidPass(pass, input) {
  const pattern = /^.*(?=.{6,})(?=.*[a-zA-ZА-ЯЁа-яё]).*$/;
  if (!pattern.test(pass)) {
    AlertAuthMessages(input, "Пароль недостаточно надежный!");
  } else {
    ReAlertAuthMessages(input);
  }
}

// Валидация повторного ввода пароля 
function ValidRepass(pass, repass, input) {
  if (repass !== pass) {
    AlertAuthMessages(input, "Пароли не совпадают!");
  } else {
    ReAlertAuthMessages(input);
  }
}

// Валидация при регистрации
function ValidRegistration(user, groupBox, groupMess) {
  ErrorAuthMessages = false
  ValidLogin(user.login, "regist-login")
  ValidName(user.name, "regist-name")
  ValidFunc(user.func, "regist-func")
  ValidUnit(user.unit, "regist-unit")
  ValidGroups(user.groups, groupBox, groupMess)
  ValidPass(user.pass, "regist-pass")
  ValidRepass(user.pass, user.repass, "regist-repass")
  return ErrorAuthMessages
}

// Валидация ввода действующего пароля при его изменении
function ValidCurrentPass(pass) {
  if (pass.length <= 5) {
    AlertAuthMessages("account-pass", "Введите пароль!");
  } else {
    ReAlertAuthMessages("account-pass");
  }
}

// Валидация при изменении пароля
function ValidUpdateUserPass(updatePass) {
  ErrorAuthMessages = false
  ValidCurrentPass(updatePass.pass)
  ValidPass(updatePass.newpass, "account-newpass")
  ValidRepass(updatePass.newpass, updatePass.repass, "account-repass")
  return ErrorAuthMessages
}
