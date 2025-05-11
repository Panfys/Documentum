//---Проверка введенных данных при создании или изменении документа

// Функция валидации входящего документа
function validateDocumentData(data) {
  validate = 0
/*
  if (data.type === "Входящий") {
    if (validDocFNum(data.fnum, "Вх. № ")) validate++;
    if (validDocSender(data.sender)) validate++;
    if (validDocCopy(data.copy)) validate++;
  } else {
    if (validDocFNum(data.fnum, "Исх. № 330/")) validate++;
    if (1 < data.count && data.count < 6) {
      if (validDocSender(data.sender)) validate++;
      if (validDocCopy(data.copy)) validate++;
      for (let i = 1; i < data.count; i++) {
        const sender = data[`sender${i}`];
        const copy = data[`copy${i}`];
        if (validDocSender(sender, i)) validate++;
        if (validDocCopy(copy, i)) validate++;
      }
    } else {
      if (validDocSender(data.sender)) validate++;
      if (validDocCopy(data.copy)) validate++;
    }
  }

  if (validDocFDate(data.fdate)) validate++;
  if ((data.lnum != "" || data.ldate != "") && validDocLNum(data.lnum, data.ldate)) { validate++; } else {
    ReAlertValidDocError('lnum')
  }

  if (validDocName(data.name)) validate++;
  if (validDocIspolnitel(data.ispolnitel, data.resolutions)) validate++;
  if (validDocCount(data.count)) validate++;

  if (validDocWidth(data.width)) validate++;

  if (data.resolutions.length > 0) {
    if (validResolutions(data.resolutions)) validate++;
  }

  if (validate === 0) serverMessage("close");

  if (validDocFileType(data.fileType)) validate++
*/
  return validate;
}

// Проверка типа файла
function validDocFileType(fileType) {
  // Список разрешенных MIME-типов
  const allowedTypes = [
    'image/jpeg',
    'image/png',
    'image/jpg',
    'application/pdf',
    'application/x-pdf',
    'application/acrobat',
    'application/vnd.pdf'
  ];

  // Проверяем, есть ли тип файла в списке разрешенных
  const isValid = allowedTypes.includes(fileType.toLowerCase());

  // Если тип недопустим - мигаем кнопкой
  if (!isValid) {
    const activeTab = document.querySelector(".main__tabs--active");
    if (activeTab) {
      const addFileBtn = activeTab.querySelector("#btn-newdoc-addfile");
      if (addFileBtn) flashButton(addFileBtn);
    }
  }

  return !isValid;
}

// Проверка порядкового номера
function validDocFNum(fnum, type) {
  if (fnum === '' || fnum === type || fnum === '№') {
    AlertValidDocError("fnum")
    serverMessage("show", 'порядковый номер документа не указан');
    return true;
  }

  if (fnum.startsWith(type)) {
    const numberPart = fnum.slice(type.length);
    if (!validDocNum(numberPart)) {
      AlertValidDocError("fnum")
      serverMessage("show", `порядковый номер документа указан неверно, примеры верного номера: "${type}123", "${type}123дсп", "${type}123/124", "${type}123/126дсп"`);
      return true;
    }
  } else {
    AlertValidDocError("fnum")
    serverMessage("show", `порядковый номер документа указан неверно, примеры верного номера: "${type}123", "${type}123дсп", "${type}123/124", "${type}123/126дсп"`);
    return true;
  }

  ReAlertValidDocError('fnum')
  return false;
}

// Проверка даты учета
function validDocFDate(fdate) {
  if (fdate === '') {
    AlertValidDocError("fdate")
    serverMessage("show", 'дата учета документа не указана');
    return true;
  }

  ReAlertValidDocError('fdate')
  return false;
}

// Проверка номера документа
function validDocLNum(lnum, ldate) {
  if (lnum === '') {
    AlertValidDocError("lnum")
    serverMessage("show", 'номер документа не указан');
    return true;
  }

  // Проверка формата с "Исх. №"
  if (lnum.startsWith('Исх. № ')) {
    const numberPart = lnum.slice('Исх. № '.length);
    if (numberPart.length === 0) {
      AlertValidDocError("lnum")
      serverMessage("show", 'после "Исх. № " должен следовать номер');
      return true
    }
    if (!validDocNum(numberPart)) {
      AlertValidDocError("lnum")
      serverMessage("show", 'номер документа указан неверно, примеры верного номера: "Исх. № 123", "Исх. № 123дсп", "123/124", "123дсп"');
      return true;
    }
  }
  // Проверка формата без "Исх. №"
  else {
    if (!validDocNum(lnum)) {
      AlertValidDocError("lnum")
      serverMessage("show", 'номер документа указан неверно, примеры верного номера: "Исх. № 123", "Исх. № 123дсп", "123/124", "123дсп"');
      return true;
    }
  }

  if (ldate == "") {
    AlertValidDocError("ldate")
    serverMessage("show", 'номер документа указан, а дата документа не указана');
    return true;
  } else {
    ReAlertValidDocError('ldate')
  }

  ReAlertValidDocError('lnum')
  return false;
}

// Проверка названия документа
function validDocName(name) {
  if (name === '') {
    AlertValidDocError("name")
    serverMessage("show", 'наименование документа не указано');
    return true;
  }

  if (!/^[А-Я]/.test(name)) {
    AlertValidDocError("name");
    serverMessage("show", 'наименование документа должно начинаться с заглавной буквы');
    return true;
  }

  ReAlertValidDocError('name')
  return false;
}

// Проверка отправителя документа
function validDocSender(sender, id) {
  if (!id) {
    id = ""
  }
  
  if (sender === '') {
    AlertValidDocError(`sender${id}`)
    serverMessage("show", 'отправитель документа не указан');
    return true;
  }

  ReAlertValidDocError(`sender${id}`)
  return false;
}

// Проверка исполнителя документа
function validDocIspolnitel(ispolnitel, resolutions) {
  if (resolutions.length == 0) {
    if (ispolnitel === '') {
      AlertValidDocError("ispolnitel")
      serverMessage("show", 'исполнитель документа не указан');
      return true;
    } else if (!IspPattern.test(ispolnitel)) {
      AlertValidDocError("ispolnitel")
      serverMessage("show", `исполнитель указан неверно, пример: "Панфилов А.П."`);
      return true;
    }
  }
  ReAlertValidDocError('ispolnitel')
  return false;
}

// Проверка количества экз. документа
function validDocCount(count) {
  if (count < 1) {
    AlertValidDocError("count")
    serverMessage("show", 'количестов экземпляров должно быть больше нуля');
    return true;
  }

  ReAlertValidDocError('count');
  return false;
}

// Проверка количества экз. документа

function validDocCopy(copy, id) {

  if (!id) {
    id = ""
  }

  if (copy === '') {
    AlertValidDocError(`copy${id}`);
    serverMessage("show", 'номер экземпляра не указан');
    return true;
  }

  // Проверяем что первый символ - цифра
  if (!/^\d/.test(copy)) {
    AlertValidDocError(`copy${id}`);
    serverMessage("show", 'номер экземпляра должен начинаться с цифры');
    return true;
  }

  // Проверяем что первая цифра > 0
  if (copy[0] === '0') {
    AlertValidDocError(`copy${id}`);
    serverMessage("show", 'первая цифра номера экземпляра должна быть больше нуля');
    return true;
  }

  ReAlertValidDocError(`copy${id}`);
  return false;
}

// Проверка количества экз. документа
function validDocWidth(width) {
  if (width == "") {
    AlertValidDocError("width")
    serverMessage("show", 'количестов листов не указано');
    return true;
  }

  const parts = width.split("/");

  // Проверяем что частей не больше 2
  if (parts.length > 2) {
    AlertValidDocError("width")
    serverMessage("show", 'количество листов указано неверно, пример: "1", "1/25"');
    return true;
  }

  // Проверяем каждую часть
  for (const part of parts) {
    const num = parseInt(part, 10);

    // Проверка на NaN (не число)
    if (isNaN(num)) {
      AlertValidDocError("width")
      serverMessage("show", 'количество листов указано неверно, пример: "1", "1/25"');
      return true;
    }

    // Проверка что число целое и положительное
    if (num < 1 || !Number.isInteger(num)) {
      AlertValidDocError("width")
      serverMessage("show", "количество листов должно быть больше нуля");
      return true;
    }
  }

  ReAlertValidDocError('width');
  return false;
}

// Вспомогательные функции проверки
function validDocNum(num) {
  if (!/^[0-9]/.test(num)) {
    return false
  }
  if (!/^[0-9а-яА-Я\/\-\.]+$/.test(num)) {
    return false
  }

  return true
}

// Проверка резолюций
function validResolutions(resolutions) {
  validate = 0
  resolutions.forEach((resolution, id) => {
    if (validResolution(id, resolution)) validate++
  })
  if (validate > 0) return true; else return false
}

function validResolution(id, resolution) {
  validate = 0
  res_id = `resolution-${id}`;
  // Валидация текста резолюции
  if (!resolution.text || resolution.text === '') {
    ValidResolutionError(res_id, 'text');
    serverMessage("show", `в резолюции ${id + 1} не заполнен текст`);
    validate++
  } else {
    ValidResolutionReError(res_id, 'text');
  }

  // Валидация даты
  if (!resolution.date) {
    ValidResolutionError(res_id, 'date');
    serverMessage("show", `в резолюции ${id + 1} не заполнена дата`);
    validate++
  } else {
    ValidResolutionReError(res_id, 'date');
  }

  // Дополнительная валидация по типу резолюции
  if (resolution.type === 'task') {
    if (!resolution.ispolnitel || resolution.ispolnitel === '') {
      ValidResolutionError(res_id, 'ispolnitel');
      serverMessage("show", `в резолюции ${id + 1} не указан исполнитель`);
      validate++
    } else if (!multiIspPattern.test(resolution.ispolnitel)) {
      ValidResolutionError(res_id, 'ispolnitel');
      serverMessage("show", `в резолюции ${id + 1} исполнитель указан неверно, пример: "Панфилов А.П.", "Панфилов А.П., Якель Е.В."`);
      validate++
    } else {
      ValidResolutionReError(res_id, 'ispolnitel');
    }
  }

  if (!resolution.user || resolution.user === '') {
    ValidResolutionError(res_id, 'user');
    serverMessage("show", `в резолюции ${id + 1} не указан исполнитель`);
    validate++
  } else if (!UserPattern.test(resolution.user)) {
    ValidResolutionError(res_id, 'user');
    serverMessage("show", `автор резолюции ${id + 1} указан неверно, пример: "Е.Лыков"`);
    validate++
  } else {
    ValidResolutionReError(res_id, 'user');
  }

  if (validate > 0) return true; else return false
}

const UserPattern = /^[А-ЯЁ]\.[А-ЯЁ][а-яё]+$/;
const IspPattern = /^[А-ЯЁ][а-яё]+ [А-ЯЁ]\.[А-ЯЁ]\.$/;
const singleIspPattern = '[А-ЯЁ][а-яё]+ [А-ЯЁ]\\.[А-ЯЁ]\\.';
const multiIspPattern = new RegExp(
  `^(${singleIspPattern})(, ${singleIspPattern})*$`
);