//---Проверка введенных данных при создании или изменении документа

// Функция валидации приказов и директив
function validateInventoryData(data) {
  validate = 0

  if (validInvNumber(data.number)) validate++;
  if (validInvName(data.name)) validate++;
  if (validInvSender(data.sender)) validate++;
  if (validInvCountCopy(data.countCopy)) validate++;
  if (validDocCopy(data.copy)) validate++;

  if (validate === 0) serverMessage("close");

  if (validDocFileType(data.fileType)) validate++

  return validate;
}

// Проверка порядкового номера
function validInvNumber(number) {
  if (number === '' || number === "Инв. №" || number === '№') {
    AlertValidDocError("number")
    serverMessage("show", 'порядковый (инвентарный) номер документа не указан');
    return true;
  }
  ReAlertValidDocError('number')

  if (number.startsWith("Инв. № ")) {
    const numberPart = number.slice("Инв. № ".length);
    if (!validDocNum(numberPart)) {
      AlertValidDocError("number")
      serverMessage("show", `порядковый (инвентарный) номер издания указан неверно, примеры верного номера: "Инв. № 123", "Инв. № 123дсп"`);
      return true;
    }
  } else {
    AlertValidDocError("number")
    serverMessage("show", `порядковый (инвентарный) номер издания указан неверно, примеры верного номера: "Инв. № 123", "Инв. № 123дсп"`);
    return true;
  }

  return false;
}

// Проверка названия документа
function validInvName(name) {
  if (name === '') {
    AlertValidDocError("name")
    serverMessage("show", 'название издания не указано');
    return true;
  }

  if (!/^[А-Я]/.test(name)) {
    AlertValidDocError("name");
    serverMessage("show", 'название издания должно начинаться с заглавной буквы');
    return true;
  }

  ReAlertValidDocError('name')
  return false;
}

// Проверка отправителя документа
function validInvSender(sender) {
  if (sender === '') {
    AlertValidDocError("sender")
    serverMessage("show", 'отправитель, издатель и год издания не указаны');
    return true;
  }

  ReAlertValidDocError('sender')
  return false;
}

// Проверка количества экз. документа
function validInvCountCopy(count) {
  if (count < 1) {
    AlertValidDocError("countCopy")
    serverMessage("show", 'количестов экземпляров должно быть больше нуля');
    return true;
  }

  ReAlertValidDocError('countCopy');
  return false;
}