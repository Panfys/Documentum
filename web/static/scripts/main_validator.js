//---Проверка введенных данных при создании или изменении документа
function WalidDocumentData(data) {
  active_tub = document.querySelector(".main__tabs--active");
  error = false;
  fnum = data.fnum.trim();
  fdate = data.fdate;
  ldate = data.ldate;
  name = data.name.trim();
  sender = data.sender.trim();
  ispolnitel = data.ispolnitel.trim();
  count = data.count.trim();
  copy = data.copy.trim();
  width = data.width.trim();
  file = data.file;
  type = data.type;

  //Для всех типов документов

  //---Проверка добавления файла------
  if (!(file.type.slice(0, 5) == "image" || file.type == "application/pdf")) {
    active_tub.querySelector("#btn-newdoc-addfile").style.backgroundColor =
      "var(--error-color)";
    setTimeout(() => {
      active_tub.querySelector("#btn-newdoc-addfile").style.backgroundColor =
        "var(--low-color)";
    }, 3000);
    error = true;
  } else {
    active_tub.querySelector("#btn-newdoc-addfile").style.backgroundColor =
      "var(--low-color)";
  }

  //---Проверка номера документа------
  if (
    fnum == "" ||
    fnum == "№" ||
    fnum == "330/" ||
    fnum == "Приказ" ||
    fnum == "Приказ №" ||
    fnum == "Вх." ||
    fnum == "Вх" ||
    fnum == "Вх. №" ||
    fnum == "Исх." ||
    fnum == "Исх. № 330/" ||
    fnum == "Исх. №" ||
    fnum == "Инв. №" ||
    fnum == "Инв." ||
    fnum == "Инв"
  ) {
    {
      WalidDocumentError("fnum");
      error = true;
    }
  } else WalidDocumentReError("fnum");

  //---Проверка даты учета------
  if (fdate == "") {
    {
      WalidDocumentError("fdate");
      error = true;
    }
  } else WalidDocumentReError("fdate");

  //---Проверка имени документа------
  if (name == "") {
    {
      WalidDocumentError("name");
      error = true;
    }
  } else WalidDocumentReError("name");

  //---Проверка отправителя (получателя) документа------
  if (sender == "") {
    {
      WalidDocumentError("sender");
      error = true;
    }
  } else if (count > 1 && count < 6 && type == "Исходящий") {
    for (i = 1; i < count; i++) {
      if (data[`sender${i}`] == "") {
        {
          WalidDocumentError("sender" + i);
          error = true;
        }
      } else {
        WalidDocumentReError("sender" + i);
        //sender += " <br> " + data[`sender${i}`];
      }
    }
  } else WalidDocumentReError("sender");

  //---Проверка количества экземпляров документа------
  if (count == "") {
    {
      WalidDocumentError("count");
      error = true;
    }
  } else WalidDocumentReError("count");

  //---Проверка номеров экземпляров документа------
  if ((copy == "" || copy <= 0) && type !== "Приказ") {
    {
      WalidDocumentError("copy");
      error = true;
    }
  } else if (count > 1 && count < 6 && type == "Исходящий") {
    for (i = 1; i < count; i++) {
      if (data[`copy${i}`] == "") {
        {
          WalidDocumentError("copy" + i);
          error = true;
        }
      } else {
        WalidDocumentReError("copy" + i);
      }
    }
  } else WalidDocumentReError("copy");

  // Для входящих и исходящих

  if (type == "Исходящий" || type == "Входящий") {
    //---Проверка исполнителя документа------
    if (!data.resolutions || data.resolutions.length == 2) {
      if (ispolnitel == "") {
        WalidDocumentError("ispolnitel");
        error = true;
      } else WalidDocumentReError("ispolnitel");
    } else {
      resolutions = data.resolutions;
      data_resolutions = JSON.parse(resolutions);
      data_resolutions.forEach((resolution) => {
        if (WalidResolutionData(resolution) == true) {
          error = true;
        }
      });
      WalidDocumentReError("ispolnitel");
    }

    //---Проверка количества листов документа------
    if (width == "") {
      {
        WalidDocumentError("width");
        error = true;
      }
    } else WalidDocumentReError("width");
  }

  // Для приказов и директив
  if (type == "Приказ") {
    //---Проверка даты учета------
    if (ldate == "") {
      {
        WalidDocumentError("ldate");
        error = true;
      }
    } else WalidDocumentReError("ldate");
  }

  return error;
}

//---Ошибка введенных данных документа----
function WalidDocumentError(input) {
  active_tub = document.querySelector(".main__tabs--active");
  active_tub.querySelector("#input-newdoc-" + input).style.borderColor =
    "var(--error-color)";
  active_tub.querySelector("#input-newdoc-" + input).style.color =
    "var(--error-color)";
}

//---Удаление ошибки введенных данных----
function WalidDocumentReError(input) {
  active_tub = document.querySelector(".main__tabs--active");
  input = active_tub.querySelector("#input-newdoc-" + input);

  if (input) {
    input.style.borderColor = "var(--low-color)";
    input.style.color = "var(--mid-color)";
  }
}

//---Проверка введенных данных резолюции----
function WalidResolutionData(resolution) {
  text = resolution.text.trim();
  user = resolution.user.trim();
  date = resolution.date.trim();

  //---Проверка исполнителя в резолюции------
  if ("ispolnitel" in resolution) {
    ispolnitel = resolution.ispolnitel.trim();
    if (ispolnitel == "") {
      WalidResolutionError(resolution.id, "ispolnitel");
      error = true;
    } else {
      WalidResolutionReError(resolution.id, "ispolnitel");
    }
  }

  //---Проверка текста в резолюции------
  if (text == "") {
    WalidResolutionError(resolution.id, "text");
    error = true;
  } else {
    WalidResolutionReError(resolution.id, "text");
  }

  if (user == "") {
    WalidResolutionError(resolution.id, "user");
    error = true;
  } else {
    WalidResolutionReError(resolution.id, "user");
  }

  if (date == "") {
    WalidResolutionError(resolution.id, "date");
    error = true;
  } else {
    WalidResolutionReError(resolution.id, "date");
  }
  return error;
}

//---Ошибка введенных данных резолюции----
function WalidResolutionError(resolution_id, input) {
  resolution = document.getElementById(resolution_id);
  input = "#resolution-" + input;

  resolution.querySelector(input).style.borderColor = "var(--error-color)";
  resolution.querySelector(input).style.color = "var(--error-color)";
}

//---Удаление ошибки введенных данных резолюции----
function WalidResolutionReError(resolution_id, input) {
  resolution = document.getElementById(resolution_id);
  input = "#resolution-" + input;

  if (resolution.querySelector(input)) {
    resolution.querySelector(input).style.borderColor = "var(--low-color)";
    resolution.querySelector(input).style.color = "var(--mid-color)";
  }
}