//---Проверка введенных данных при создании или изменении документа

// Функция валидации приказов и директив
function validateInventoryData(data) {
  validate = 0

  if (validDirNumber(data.number)) validate++;
  if (validDirDate(data.date)) validate++;
  if (validDirName(data.name)) validate++;
  if (validDirAutor(data.autor)) validate++;
  if (validDirCountCopy(data.countCopy)) validate++;

  if (validate === 0) serverMessage("close");

  if (validDocFileType(data.fileType)) validate++

  return validate;
}

// Проверка порядкового номера
function validDirNumber(number) {
  if (number === '' || number === "Приказ №" || number === '№' || number === "Приказание №" || number === "Директива №") {
    AlertValidDocError("number")
    serverMessage("show", 'порядковый номер документа не указан');
    return true;
  }
  ReAlertValidDocError('number')

  return false;
}

// Проверка даты подписания
function validDirDate(date) {
  if (date === '') {
    AlertValidDocError("date")
    serverMessage("show", 'дата подписи не указана');
    return true;
  }

  ReAlertValidDocError('date')
  return false;
}

// Проверка названия документа
function validDirName(name) {
  if (name === '') {
    AlertValidDocError("name")
    serverMessage("show", 'краткое содержание не указано');
    return true;
  }

  if (!/^[А-Я]/.test(name)) {
    AlertValidDocError("name");
    serverMessage("show", 'краткое содержание  должно начинаться с заглавной буквы');
    return true;
  }

  ReAlertValidDocError('name')
  return false;
}

// Проверка автора приказа/директивы
function validDirAutor(autor) {
  if (autor === '') {
    AlertValidDocError("autor")
    serverMessage("show", 'лицо, подписавшее документ не указано');
    return true;
  }

  ReAlertValidDocError('autor')
  return false;
}

// Проверка количества экз. документа
function validDirCountCopy(count) {
  if (count < 1) {
    AlertValidDocError("countCopy")
    serverMessage("show", 'количестов экземпляров должно быть больше нуля');
    return true;
  }

  ReAlertValidDocError('countCopy');
  return false;
}