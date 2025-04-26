date = new Date().toJSON().slice(0, 10).replace(/-/g, "/"); //дата

//-----Открытие боковых панелей------------

//------------------------------MAIN-------------------------------------

document.querySelector(".main__header--settings").onclick = () =>
  OpenMainSide("settings-panel");
document.querySelector(".main__header--account").onclick = () =>
  OpenMainSide("account-panel");

function OpenMainSide(panel_id) {
  panel = document.getElementById(panel_id);
  if (panel.style.display == "flex") {
    panel.style = "display: none";
  } else {
    panel.style = "display: flex";
  }
}

//-------Переключение темы
table_theme = document.querySelectorAll(".settings__theme--table");

table_theme.forEach((table) => {
  // вешаем на каждую кнопку обработчик события клик
  table.addEventListener("click", () => ChengeTheme(table));

  function ChengeTheme(table) {
    theme = table.getAttribute("theme");
    color = table.getAttribute("color");
    localStorage.setItem("theme", theme);
    localStorage.setItem("color", color);

    if (theme == "light") {
      document.body.classList.add("light-theme");
      document.body.classList.remove("dark-theme");
    } else if (theme == "dark") {
      document.body.classList.remove("light-theme");
      document.body.classList.add("dark-theme");
    }

    switch (color) {
      case "blue":
        {
          document.body.style.setProperty("--main-rgb", "45, 104, 248");
        }
        break;
      case "orange":
        {
          document.body.style.setProperty("--main-rgb", "255, 104, 0");
        }
        break;
      case "purple":
        {
          document.body.style.setProperty("--main-rgb", "116, 66, 200");
        }
        break;
      case "green":
        {
          document.body.style.setProperty("--main-rgb", "3, 108, 86");
        }
        break;
    }
  }
});

//- переключение и открытие настроек в боковых панелях

btn_settings = document.querySelectorAll(".main__settings--btn");
// Проходимся по всем кнопкам
btn_settings.forEach((btn) => {
  // вешаем на каждую кнопку обработчик события клик
  btn.addEventListener("click", () => ChengeActivePanel(btn));

  function ChengeActivePanel(btn) {
    // Получаем предыдущую активную вкладку
    pre_active_btn = document.querySelector(".main__settings-active-btn");
    // Получаем предыдущую активную вкладку
    pre_active_panel = document.querySelector(".main__settings--active-panel");

    // Проверяем есть или нет предыдущая активная кнопка
    if (pre_active_btn) {
      //Удаляем класс _active у предыдущей кнопки если она есть
      pre_active_btn.classList.remove("main__settings-active-btn");
    }
    // Проверяем есть или нет предыдущая активная вкладка
    if (pre_active_panel) {
      // Удаляем класс _active у предыдущей вкладки если она есть
      pre_active_panel.classList.remove("main__settings--active-panel");
    }
    // получаем id новой активной вкладки, который мы перем из атрибута data-tab у кнопки
    const active_panel_id = "#" + btn.getAttribute("panel-id");
    const active_panel = document.querySelector(active_panel_id);

    if (active_panel !== pre_active_panel) {
      // добавляем класс _active кнопке на которую нажали
      btn.classList.add("main__settings-active-btn");
      // добавляем класс _active новой выбранной вкладке
      active_panel.classList.add("main__settings--active-panel");
    }
  }
});

//-------Переключение табов документов

// получаем все кнопки навигации
btn_menu = document.querySelectorAll(".header__menu--btn");
// Проходимся по всем кнопкам
btn_menu.forEach((btn) => {
  // вешаем на каждую кнопку обработчик события клик
  btn.addEventListener("click", () => ChengeActiveTab(btn));

  if (sessionStorage.getItem("active_btn_id")) {
    button = document.getElementById(sessionStorage.getItem("active_btn_id"));
    ChengeActiveTab(button);
  } else {
    button = document.getElementById("menu-btn-ingoing");
    ChengeActiveTab(button);
  }
  function ChengeActiveTab(btn) {
    // Получаем предыдущую активную кнопку
    pre_active_tub = document.querySelector(".main__tabs.main__tabs--active");
    // Получаем предыдущую активную вкладку
    pre_active_btn = document.querySelector(
      ".header__menu--btn.menu__btn--active"
    );

    // Проверяем есть или нет предыдущая активная кнопка
    if (pre_active_btn) {
      //Удаляем класс _active у предыдущей кнопки если она есть
      pre_active_btn.classList.remove("menu__btn--active");
    }
    // Проверяем есть или нет предыдущая активная вкладка
    if (pre_active_tub) {
      // Удаляем класс _active у предыдущей вкладки если она есть
      pre_active_tub.classList.remove("main__tabs--active");
    }
    // получаем id новой активной вкладки, который мы перем из атрибута data-tab у кнопки
    const active_tub_id = "#" + btn.getAttribute("data-tab");
    sessionStorage.setItem("active_btn_id", btn.getAttribute("id"));
    // получаем новую активную вкладку по id
    const active_tub = document.querySelector(active_tub_id);

    // добавляем класс _active кнопке на которую нажали
    btn.classList.add("menu__btn--active");
    // добавляем класс _active новой выбранной вкладке
    active_tub.classList.add("main__tabs--active");

    // обновляем вкладку с документами
    switch (active_tub_id) {
      case "#main-tab-ingoing":
        {
          ViewDocuments("ASC", "Входящий", "id", "2000-01-01", "3000-01-01");
        }
        break;

      case "#main-tab-outgoing":
        {
          ViewDocuments("ASC", "Исходящий", "id", "2000-01-01", "3000-01-01");
        }
        break;

      case "#main-tab-directive":
        {
          ViewDocuments("ASC", "Приказ", "id", "2000-01-01", "3000-01-01");
        }
        break;

      case "#main-tab-inventory":
        {
          ViewDocuments("ASC", "Издание", "id", "2000-01-01", "3000-01-01");
        }
        break;
    }
  }
});

//------Функция добавления фотографии-------------
icon = document.querySelector(".account__icon--intut");

icon.onchange = function () {
  user_icon = new FormData();
  user_icon.append("icon", icon.files[0]);

  $.ajax({
    method: "POST",
    url: "/users/addicon",
    data: user_icon,
    cache: false,
    contentType: false,
    processData: false,
    success: function (check) {
      if (document.querySelector(".account__icon")) {
        document.querySelector(".account__icon").className = "account__img";
      }
      img = document.querySelector(".account__img");

      img.style.backgroundImage = "url('')";
      img.style.backgroundImage = "url('" + check + "')";
    },
    error: function (jqXHR) {
      ServerMessage("show", jqXHR.responseText);
    },
  });
};

//-------------Открытие панели изменения пароля-----------
document.querySelector("#account-btn-open").onclick = () =>
  OpenAccountPassPanel("open");
document.querySelector("#account-btn-close").onclick = () =>
  OpenAccountPassPanel("close");

function OpenAccountPassPanel(act) {
  if (act == "open") {
    document.querySelector(".account__password--panel").style = "display: flex";
    document.getElementById("account-btn-open").style = "display: none";
  } else {
    document.querySelector(".account__password--panel").style = "display: none";
    document.getElementById("account-btn-open").style = "display: flex";
  }
}

pass_input = document.getElementById("account-pass-input");
newpass_input = document.getElementById("account-newpass-input");
repass_input = document.getElementById("account-repass-input");

//Проверка ввода нового пароля
pass_input.addEventListener("blur", WritePass);

function WritePass() {
  pass = pass_input.value.trim();
  if (pass.length <= 5) AlertMessages("account-pass", "Введите пароль!");
  else ReAlertMessages("account-pass");
}

//Проверка ввода нового пароля
newpass_input.addEventListener("blur", WriteNewPass);

function WriteNewPass() {
  newpass = newpass_input.value.trim();
  if (!IsValidPass(newpass))
    AlertMessages("account-newpass", "Пароль недостаточно надежный!");
  else ReAlertMessages("account-newpass");
}

//Проверка подтверждения пароля
repass_input.addEventListener("input", WriteRePass);

function WriteRePass() {
  repass = repass_input.value.trim();
  if (repass !== newpass)
    AlertMessages("account-repass", "Пароли не совпадают!");
  else ReAlertMessages("account-repass");
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
  document.getElementById("account-btn-change").disabled = true;
}

//очистка сообщения
function ReAlertMessages(input) {
  document.getElementById(input.concat("-input")).style.borderColor =
    "--low-color";
  document.getElementById(input.concat("-lable")).style.color = "--low-color";
  document.getElementById(input.concat("-message")).innerHTML = "";
  document.getElementById(input.concat("-message")).classList.remove("error");
  document.getElementById("account-btn-change").disabled = false;
}

document.querySelector("#account-btn-change").onclick = function () {
  pass = pass_input.value.trim();
  newpass = newpass_input.value.trim();
  repass = repass_input.value.trim();

  WritePass();
  WriteNewPass();

  if (document.querySelector("#account-btn-change").disabled) return;

  $.ajax({
    method: "POST",
    url: "/protect/users/updatepass",
    data: {
      pass: pass,
      newpass: newpass,
      repass: repass,
    },
    success: function () {
      OpenAccountPassPanel("close");
      document.getElementById("account-chengepass-message").innerHTML =
        "Пароль успешно изменён!";
      document
        .getElementById("account-chengepass-message")
        .classList.add("message");
    },
    error: function (jqXHR) {
      if (jqXHR.status == 400) {
        AlertMessages("account-repass", jqXHR.responseText);
      } else ServerMessage("show", jqXHR.responseText);
    },
  });
};

function ReAlertAccountMessages() {
  document.getElementById("account-chengepass-message").innerHTML = "";
  document
    .getElementById("account-chengepass-message")
    .classList.remove("message");
}

// Выход пользователя из учетной записи
document.querySelector("#account-btn-exit").onclick = function () {
  if (confirm("Вы действительно хотите выйти?")) {
    $.ajax({
      method: "POST",
      url: "/protect/users/exit",
      success: function (check) {
        if ((document.querySelector(".container").innerHTML = check)) {
          script = document.createElement("script");
          script.src = "/static/scripts/script_entrance.js";
          document.head.appendChild(script);
        }
      },
      error: function (jqXHR) {
        ServerMessage("show", jqXHR.responseText);
      },
    });
  }
};

//------Функция прокрутки и закрепления шапки таблицы-------
window.addEventListener("scroll", function () {
  active_tub = document.querySelector(".main__tabs--active");
  head_containers = active_tub.querySelector(".tubs__container--head");
  hidden_containers = active_tub.querySelector(".tubs__container--hidden");

  if (window.scrollY > 30) {
    head_containers.classList.add("tubs__container--tabscroll");
    hidden_containers.style.display = "block";
  } else {
    head_containers.classList.remove("tubs__container--tabscroll");
    hidden_containers.style.display = "none";
  }
});

//------Функция кнопки "новый документ" или "Отмена"-------
btn_newdocument = document.querySelectorAll("#btn-newdoc");
// Проходимся по всем кнопкам
btn_newdocument.forEach((btn) => {
  // вешаем на каждую кнопку обработчик события клик
  btn.addEventListener("click", () => OpenNewDocForm());
});

function OpenNewDocForm() {
  active_tub = document.querySelector(".main__tabs--active");
  tubs_form = active_tub.querySelector("#form-newdoc");
  tubs_table = tubs_form.querySelector("#table-newdoc");
  tubs_btn = active_tub.querySelector("#btn-newdoc");
  tubs_folder = active_tub.querySelector("#head-folder");
  tubs_panel = active_tub.querySelector("#title-newdocpanel");
  tubs_span = active_tub.querySelector("#title-span");
  btn_search = active_tub.querySelector("#btn-search");
  resolution_btn_panel = active_tub.querySelector("#btn-resolution-panel");
  if (tubs_form.style.display == "flex") {
    if (tubs_btn.style.display == "none") {
      AddNewdocResolution("back");
    }
    tubs_form.style = "display: none";
    tubs_btn.innerHTML = "Новый документ";
    tubs_span.style = "display: none";
    tubs_folder.innerHTML = "";
    tubs_panel.style = "display: none";
    btn_search.style = "display: flex";
    tubs_table.classList.remove("tubs__table--active-table");
  } else {
    tubs_form.style = "display: flex";
    tubs_btn.innerHTML = "Отмена";
    tubs_span.style = "display: flex";
    tubs_folder.innerHTML = "Запись нового документа";
    tubs_panel.style = "display: flex";
    btn_search.style = "display: none";
    tubs_table.classList.add("tubs__table--active-table");
    window.scrollTo({
      top: 0,
      left: 0,
      behavior: "smooth",
    });
  }
}

//------Функция кнопки "Очистить"-------

btn_clearnewdoc = document.querySelectorAll("#btn-newdoc-clearnewdoc");
// Проходимся по всем кнопкам
btn_clearnewdoc.forEach((btn) => {
  // вешаем на каждую кнопку обработчик события клик
  btn.addEventListener("click", function () {
    active_tub = document.querySelector(".main__tabs--active");
    active_tub.querySelector("#form-newdoc").reset();
    ClearNewdocFilePanel();
    ClearNewdocResolutionPanel();
  });
});

//------Функция кнопки "Прикрепить файл"-------
btn_addfile = document.querySelectorAll("#btn-newdoc-addfile");
// Проходимся по всем кнопкам
btn_addfile.forEach((btn) => {
  // вешаем на каждую кнопку обработчик события клик
  btn.addEventListener("click", function AddNewdocFile() {
    active_tub = document.querySelector(".main__tabs--active");
    btn_file = active_tub.querySelector("#btn-newdoc-addfile");
    file_input = active_tub.querySelector("#input-newdoc-file");
    file_panel = active_tub.querySelector("#newdoc-file-panel");
    file_name = active_tub.querySelector("#newdoc-file-name");
    file_size = active_tub.querySelector("#newdoc-file-size");
    file_img = active_tub.querySelector("#newdoc-file-img");

    file_input.click();
    file_input.onchange = function () {
      file = file_input.files[0];
      if (file.type == "application/pdf" || file.type.slice(0, 5) == "image") {
        file_name.innerHTML = file.name;
        file_url = URL.createObjectURL(file);
        file_img.innerHTML = DefineFile(file_url, file.type);
        file_size.innerHTML = DefineFileSize(file_input.files[0].size);
        file_panel.style.display = "flex";
        btn_file.innerHTML = "Изменить файл";
      }
    };
  });
});

//------Функция Очистки панели файлов------
function ClearNewdocFilePanel() {
  active_tub = document.querySelector(".main__tabs--active");
  btn_file = active_tub.querySelector("#btn-newdoc-addfile");
  file_panel = active_tub.querySelector("#newdoc-file-panel");
  file_name = active_tub.querySelector("#newdoc-file-name");
  file_size = active_tub.querySelector("#newdoc-file-size");
  file_img = active_tub.querySelector("#newdoc-file-img");

  active_tub.querySelector("#input-newdoc-file").files[0] = "";
  file_name.innerHTML = "";
  file_url = "";
  file_img.innerHTML = "";
  file_size.innerHTML = "";
  file_panel.style.display = "none";
  btn_file.innerHTML = "Файл";
}

//------Функция Корректности введенного файла"------
function DefineFile(file_url, file_type) {
  if (file_type.slice(0, 5) == "image") {
    return "<img src=" + file_url + ">";
  } else if (file_type == "application/pdf") {
    return '<embed src="' + file_url + '" scrolling="no"></embed>';
  } else {
    return "<img src='/style/images/file error.png'>";
  }
}

//------Функция Вывода размера файла"------
function DefineFileSize(file_size) {
  const units = ["Б", "КБ", "МБ", "ГБ", "ТБ"];

  i = 0;
  n = parseInt(file_size, 10) || 0;

  while (n >= 1000 && ++i) {
    n = n / 1000;
  }

  return n.toFixed(n < 10 && i > 0 ? 1 : 0) + " " + units[i];
}

//------Функция кнопки "Добавить резолюции"-------
document.querySelector("#btn-newdoc-resolution").onclick = () =>
  AddNewdocResolution("open");

document.querySelector("#btn-resolution-add").onclick = function () {
  active_doc = document.querySelector(".tubs__table--active-table");
  if (active_doc.getAttribute("id") == "table-newdoc") {
    AddNewdocResolution("add");
  } else AddDocResolution("add");
};

document.querySelector("#btn-resolution-result").onclick = function () {
  active_doc = document.querySelector(".tubs__table--active-table");
  if (active_doc.getAttribute("id") == "table-newdoc") {
    AddNewdocResolution("result");
  } else AddDocResolution("result");
};

document.querySelector("#btn-resolution-remove").onclick = function () {
  active_doc = document.querySelector(".tubs__table--active-table");
  if (active_doc.getAttribute("id") == "table-newdoc") {
    AddNewdocResolution("remove");
  } else AddDocResolution("remove");
};

document.querySelector("#btn-resolution-back").onclick = function () {
  active_doc = document.querySelector(".tubs__table--active-table");
  if (active_doc.getAttribute("id") == "table-newdoc") {
    AddNewdocResolution("back");
  } else AddDocResolution("back");
};

function AddNewdocResolution(action) {
  active_tub = document.querySelector(".main__tabs--active");
  resolution_panel = active_tub.querySelector("#newdoc-resolution-panel");
  resolution_id = resolution_panel.childElementCount + 1;
  resolution_btn_add = active_tub.querySelector("#btn-resolution-add");
  resolution_btn_panel = active_tub.querySelector("#btn-resolution-panel");
  newdoc_panel = active_tub.querySelector("#title-newdocpanel");
  newdoc_btn = active_tub.querySelector("#btn-newdoc");
  input_ispolnitel = active_tub.querySelector("#input-newdoc-ispolnitel");
  new_resolution = document.createElement("div");
  new_resolution.setAttribute("class", "table__resolution");
  new_resolution.setAttribute("id", "newdoc-resolution-1");
  new_resolution.innerHTML = `
    <input id="resolution-ispolnitel" type="text" placeholder="Исполнитель">
    <textarea id="resolution-text"></textarea>
    <div class="resolution__time--block">Срок исполнения 
      <input id="resolution-time" type="date">
    </div>
    <input class="resolution__user--input" type="text" id="resolution-user" placeholder="Руководитель">
    <input class="resolution__date--input" type="date" id="resolution-date">
  `;

  switch (action) {
    case "open":
      {
        newdoc_panel.style.display = "none";
        resolution_btn_panel.style.display = "flex";
        newdoc_btn.style.display = "none";
      }
      break;
    case "add":
      {
        new_resolution.setAttribute("id", "newdoc-resolution-" + resolution_id);
        new_resolution.setAttribute("res_id", resolution_id);
        resolution_panel.appendChild(new_resolution);
        resolution_panel.scrollTo({
          top: 0,
          left: 5000,
          behavior: "smooth",
        });
        resolution_id += 1;
        input_ispolnitel.setAttribute(
          "placeholder",
          "Заполняется автоматически"
        );
        resolution_btn_add.style.backgroundColor = "var(--low-color)";
      }
      break;
    case "result":
      {
        if (resolution_id == 1) {
          resolution_btn_add.style.backgroundColor = "var(--error-color)";
          setTimeout(() => {
            resolution_btn_add.style.backgroundColor = "var(--low-color)";
          }, 3000);
        } else {
          new_resolution.innerHTML = `
            <textarea id="resolution-text"></textarea>
            <input id="resolution-result" type="text" placeholder="Исполненный документ">
            <input class="resolution__user--input" type="text" id="resolution-user" placeholder="Исполнитель">
            <input class="resolution__date--input" type="date" id="resolution-date">
          `;
          new_resolution.setAttribute(
            "id",
            "newdoc-resolution-" + resolution_id
          );
          new_resolution.setAttribute("res_id", resolution_id);
          resolution_panel.appendChild(new_resolution);
          resolution_panel.scrollTo({
            top: 0,
            left: 5000,
            behavior: "smooth",
          });
          resolution_id += 1;
        }
      }
      break;
    case "back":
      {
        resolution_btn_panel.style = "display: none";
        newdoc_panel.style = "display: flex";
        newdoc_btn.style.display = "flex";
      }
      break;
    case "remove":
      {
        if (resolution_id == 2) {
          resolution_id = 1;
          input_ispolnitel.setAttribute("placeholder", "");
        } else if (resolution_id == 1) return;
        else resolution_id -= 1;

        resolution = document.getElementById(
          "newdoc-resolution-" + resolution_id
        );
        resolution_panel.removeChild(resolution);
      }
      break;
  }
}

//------Функция Очистки панели резолюций-------
function ClearNewdocResolutionPanel() {
  active_tub = document.querySelector(".main__tabs--active");
  resolution_panel = active_tub.querySelector("#newdoc-resolution-panel");
  input_ispolnitel = active_tub.querySelector("#input-newdoc-ispolnitel");

  resolution_panel.innerHTML = "";
  input_ispolnitel.setAttribute("placeholder", "");
}

//------Функция кнопки "Записать"-------
btn_addnewdoc = document.querySelectorAll("#btn-newdoc-addnewdoc");
// Проходимся по всем кнопкам
btn_addnewdoc.forEach((btn) => {
  // вешаем на каждую кнопку обработчик события клик
  btn.addEventListener("click", function AddNewDocument() {
    active_tub = document.querySelector(".main__tabs--active");
    newdoc_form = active_tub.querySelector("#form-newdoc");
    newdoc_data = new FormData(newdoc_form);

    switch (active_tub.getAttribute("id")) {
      case "main-tab-ingoing":
        newdoc_data.append("type", "Входящий");

        //обработка резолюций
        resolution_panel = active_tub.querySelector("#newdoc-resolution-panel");
        resolutions = resolution_panel.querySelectorAll(".table__resolution");

        resolutions_data = [];

        resolutions.forEach((resolution) => {
          resolution_data = new Object();
          resolution_id = resolution.getAttribute("id");
          text = resolution.querySelector("#resolution-text").value;
          user = resolution.querySelector("#resolution-user").value;
          date = resolution.querySelector("#resolution-date").value;
          resolution_data.text = text;
          resolution_data.user = user;
          resolution_data.date = date;
          resolution_data.id = resolution_id;

          if (resolution.querySelector("#resolution-ispolnitel")) {
            ispolnitel = resolution.querySelector(
              "#resolution-ispolnitel"
            ).value;
            time = resolution.querySelector("#resolution-time").value;
            resolution_data.ispolnitel = ispolnitel;
            resolution_data.time = time;
            resolutions_data.push(resolution_data);
          } else {
            result = resolution.querySelector("#resolution-result").value;
            resolution_data.result = result;
            resolutions_data.push(resolution_data);
          }
        });
        newdoc_data.append("resolutions", JSON.stringify(resolutions_data));
        break;

      case "main-tab-outgoing":
        newdoc_data.append("type", "Исходящий");
        break;

      case "main-tab-directive":
        newdoc_data.append("type", "Приказ");
        break;

      case "main-tab-inventory":
        newdoc_data.append("type", "Издание");
        break;
    }

    data_obj = Object.fromEntries(newdoc_data.entries());
    if (WalidDocumentData(data_obj) === true) {
      return;
    }

    $.ajax({
      method: "POST",
      url: "../documents/adddoc",
      data: newdoc_data,
      cache: false,
      contentType: false,
      processData: false,
      success: function () {
        alert("Документ успешно загружен, перезагрузите страницу!");
      },
      error: function (jqXHR) {
        ServerMessage("show", jqXHR.responseText);
      },
    });
  });
});

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
    fnum == "Исх" ||
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
    if (data.resolutions.length == 2) {
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

//------Вывод документов---------
function ViewDocuments(set, type, col, datain, datato) {
  $.ajax({
    method: "POST",
    url: "/documents/getdoc",
    data: {
      type: type,
      col: col,
      set: set,
      datain: datain,
      datato: datato,
    },
    success: function (documents) {
      switch (type) {
        case "Входящий":
          document.getElementById("ingoing-documents-container").innerHTML =
            documents;
          break;
        case "Исходящий":
          document.getElementById("outgoing-documents-container").innerHTML =
            documents;
          break;
        case "Приказ":
          document.getElementById("directive-documents-container").innerHTML =
            documents;
          break;
        case "Издание":
          document.getElementById("inventory-documents-container").innerHTML =
            documents;
          break;
      }
      CheckDocumentTable();
    },
    error: () =>
      ServerMessage("show", "Возникла ошибка на сервере, попробуйте позже!"),
  });
}

//------Функция сортировки-------------
function SortDocuments() {
  val = document.getElementById("SetSort-input").value;
  datain = document.getElementById("DataInSet-input").value;
  datato = document.getElementById("DataToSet-input").value;
  if (datain == "") datain = "2000-01-01";
  if (datato == "") datato = "3000-01-01";

  switch (val) {
    case "0":
      ViewDocuments("ASC", "Входящий", "id", datain, datato);
      break;
    case "1":
      ViewDocuments("DESC", "Входящий", "id", datain, datato);
      break;
    case "2":
      ViewDocuments("ASC", "Входящий", "name", datain, datato);
      break;
    case "3":
      ViewDocuments("ASC", "Входящий", "sender", datain, datato);
      break;
  }
}

//-------Открытие панели документа
function CheckDocumentTable() {
  // получаем все таблицы документов
  doc_table = document.querySelectorAll(".tubs__table--document");
  // Проходимся по всем таблицам
  doc_table.forEach((doc) => {
    // вешаем на каждую таблицу обработчик события клик
    doc.addEventListener("click", () => ViewDocumentTable(doc, event));
  });
}

function ViewDocumentTable(doc, event) {
  // Получаем предыдущую активную таблицу
  active_tub = document.querySelector(".main__tabs--active");
  pre_active_doc = document.querySelector(".tubs__table--active-table");
  docpanel = active_tub.querySelector("#title-docpanel");
  btn_search = active_tub.querySelector("#btn-search");
  btn_newdoc = active_tub.querySelector("#btn-newdoc");
  tubs_folder = active_tub.querySelector(".tubs__folder");
  tubs_title_span = active_tub.querySelector(".tubs__title--span");
  // Проверяем есть или нет предыдущая активная таблица

  function CloseActiveDoc() {
    if (pre_active_doc) {
      //Удаляем класс _active у таблицы нового документа
      if (pre_active_doc.getAttribute("id") == "table-newdoc") {
        OpenNewDocForm();
      }
      //Удаляем класс _active у предыдущей активной таблицы
      else {
        if (docpanel.style.display == "none") {
          AddDocResolution("back");
        }
        pre_active_doc.classList.remove("tubs__table--active-table");
        pre_active_doc_id = pre_active_doc.getAttribute("document-id");
        pre_active_res = document.getElementById(
          "resolution-panel-" + pre_active_doc_id
        );
        tubs_folder.innerHTML = "";
        tubs_title_span.style.display = "none";
        btn_search.style.display = "flex";
        btn_newdoc.style.display = "flex";
        docpanel.style.display = "none";
        // закрываем панель резолюций у предыдущего документа
        if (resolution_panel) {
          resolution_panel.style.display = "none";
          if (resolution_id !== newresolution_id) {
            resolution_panel.removeChild(resolution_panel.lastChild);
            newresolution_id -= 1;
          }
        }
      }
    }
  }

  if (event.target.classList == "table__btn--opendoc") {
    OpenDocument(
      event.target.getAttribute("file"),
      doc.getAttribute("document-id")
    );
    if (!(pre_active_doc == doc)) CloseActiveDoc();
    return;
  } else CloseActiveDoc();
  // проверяем нажание на одну таблицу
  if (pre_active_doc == doc) return;
  // добавляем класс _active новой выбранной вкладке
  doc.classList.add("tubs__table--active-table");
  doc_id = doc.getAttribute("document-id");
  resolution_panel = document.getElementById("resolution-panel-" + doc_id);
  if (doc.querySelector(".table__column--name")) {
    tubs_folder.innerHTML = doc.querySelector(".table__column--name").innerHTML;
  } else if (doc.querySelector("#table__column--name")) {
    tubs_folder.innerHTML = doc.querySelector("#table__column--name").innerHTML;
  }
  tubs_title_span.style.display = "flex";
  btn_search.style.display = "none";
  btn_newdoc.style.display = "none";
  docpanel.style.display = "flex";
  // открываем панель резолюций
  if (resolution_panel) {
    resolution_panel.style.display = "flex";
    resolution_id = resolution_panel.childElementCount;
    newresolution_id = resolution_panel.childElementCount;
  }
}

document.querySelector(".panel__opendoc--btn").onclick = () =>
  OpenDocument("close", "id");

//------Функция отрытия вкладки просмотра документа
function OpenDocument(ask, id) {
  panel_opendoc = document.querySelector("#panel-opendoc");
  iframe_opendoc = document.querySelector("#iframe-opendoc");
  resolutions_opendoc = document.querySelector("#resolutions-opendoc");

  if (ask == "close") {
    panel_opendoc.style = "display: none";
    iframe_opendoc.setAttribute("src", "");
    iframe_opendoc.style.width = "70%";
    iframe_opendoc.style.height = "auto";
    resolutions_opendoc.innerHTML = "";
    resolutions_opendoc.style = "min-width: 80px;";
    document.body.style.overflow = "auto";
  } else {
    document.body.style.overflow = "hidden";
    panel_opendoc.style = "display: flex";
    iframe_opendoc.setAttribute("src", "/documents/wievdoc?file=" + ask);
    familiars_opendoc = document
      .getElementById("document-table-" + id)
      .querySelector(".table__column--familiar");

    if (document.getElementById("resolution-panel-" + id)) {
      if (document.getElementById("resolution-panel-" + id).innerHTML !== "") {
        resolutions_opendoc.style = "min-width: 294px;";
      }
      resolutions_opendoc.innerHTML = document.getElementById(
        "resolution-panel-" + id
      ).innerHTML;
    }

    //------ Запись просмотра документа------------
    account_name = document.getElementById("account-name").innerHTML;
    familiar_opendoc = familiars_opendoc.innerHTML;

    if (familiar_opendoc.match(account_name.trim()) == null) {
      $.ajax({
        method: "POST",
        url: "/documents/look",
        data: { id: id },
        error: function (jqXHR, exception) {
          ServerMessage("show", jqXHR.responseText);
        },
      });
    }
  }
}

//------Функция отрытия вкладки просмотра нового документа
btn_opennewdoc = document.querySelectorAll("#btn-newdoc-open");
// Проходимся по всем кнопкам
btn_opennewdoc.forEach((btn) => {
  // вешаем на каждую кнопку обработчик события клик
  btn.addEventListener("click", function () {
    active_tub = document.querySelector(".main__tabs--active");
    newdoc_resolutions = active_tub.querySelector("#newdoc-resolution-panel");
    panel_opendoc = document.querySelector("#panel-opendoc");
    resolutions_opendoc = document.querySelector("#resolutions-opendoc");
    iframe_opendoc = document.querySelector("#iframe-opendoc");
    file_input = active_tub.querySelector("#input-newdoc-file");
    file = file_input.files[0];
    document.body.style.overflow = "hidden";

    if (file) {
      file_url = URL.createObjectURL(file);
      iframe_opendoc.setAttribute(
        "src",
        "/documents/wievnewdoc?file=" + DefineFile(file_url, file.type)
      );
    }

    panel_opendoc.style.display = "flex";
    if (newdoc_resolutions.innerHTML !== "") {
      resolutions_opendoc.style = "min-width: 280px;";
      resolutions_opendoc.innerHTML = newdoc_resolutions.innerHTML;
    }
  });
});

//------Функция изменения данных документа-----

//------Функция кнопки "Резолюция"-------
document.querySelector("#btn-doc-resolution").onclick = () =>
  AddDocResolution("open");

function AddDocResolution(action) {
  active_tub = document.querySelector(".main__tabs--active");
  active_doc = document.querySelector(".tubs__table--active-table");
  active_doc_id = active_doc.getAttribute("document-id");
  resolution_panel = active_tub.querySelector(
    "#resolution-panel-" + active_doc_id
  );
  resolution_btn_add = active_tub.querySelector("#btn-resolution-add");
  resolution_btn_panel = active_tub.querySelector("#btn-resolution-panel");
  doc_panel = active_tub.querySelector("#title-docpanel");
  new_resolution = document.createElement("div");
  new_resolution.setAttribute("class", "table__resolution");
  new_resolution.setAttribute("id", "doc-newresolution");

  switch (action) {
    case "open":
      {
        doc_panel.style.display = "none";
        resolution_btn_panel.style.display = "flex";
      }
      break;
    case "add":
      {
        if (resolution_id == newresolution_id) {
          new_resolution.innerHTML = `
          <input id="resolution-ispolnitel" name="resolution-ispolnitel" type="text" placeholder="Исполнитель">
          <textarea id="resolution-text" name="resolution-text"></textarea>
          <div class="resolution__time--block">Срок исполнения 
            <input id="resolution-time" name="resolution-time" type="date">
          </div>
          <input class="resolution__user--input" type="text" id="resolution-user" name="resolution-user" placeholder="Руководитель">
          <input class="resolution__date--input" type="date" id="resolution-date" name="resolution-date" value="">
        `;
          resolution_panel.appendChild(new_resolution);
          resolution_panel.scrollTo({
            top: 0,
            left: 5000,
            behavior: "smooth",
          });
          newresolution_id += 1;
        }
      }
      break;
    case "result":
      {
        if (resolution_id == 0) {
          resolution_btn_add.style.backgroundColor = "var(--error-color)";
          setTimeout(() => {
            resolution_btn_add.style.backgroundColor = "var(--low-color)";
          }, 3000);
        } else if (resolution_id == newresolution_id) {
          new_resolution.innerHTML = `
            <textarea id="resolution-text" name="resolution-text"></textarea>
            <input id="resolution-result" name="resolution-result" type="text" placeholder="Исполненный документ">
            <input class="resolution__user--input" type="text" id="resolution-user" name="resolution-user" placeholder="Исполнитель">
            <input class="resolution__date--input" type="date" id="resolution-date" name="resolution-date">
          `;
          resolution_panel.appendChild(new_resolution);
          resolution_panel.scrollTo({
            top: 0,
            left: 5000,
            behavior: "smooth",
          });
          newresolution_id += 1;
        }
      }
      break;
    case "back":
      {
        resolution_btn_panel.style = "display: none";
        doc_panel.style = "display: flex";
      }
      break;
    case "remove":
      {
        if (resolution_id !== newresolution_id) {
          resolution_panel.removeChild(resolution_panel.lastChild);
          newresolution_id -= 1;
        }
      }
      break;
  }
}

/*/------Функция кнопки "Изменить"-------НЕ ДОДЕЛАЛ
document.querySelector("#btn-doc-change").onclick = () => DocСhange();

function DocСhange() {
  active_tub = document.querySelector(".main__tabs--active");
  active_doc = document.querySelector(".tubs__table--active-table");
  active_doc_id = active_doc.getAttribute("document-id");
} */

//------Функция кнопки "Сохранить"-------
document.querySelector("#btn-doc-save").onclick = () => DocSave();

function DocSave() {
  active_tub = document.querySelector(".main__tabs--active");
  doc = document.querySelector(".tubs__table--active-table");
  doc_id = doc.getAttribute("document-id");
  doc_data = new FormData();
  doc_data.append("action", "ChangeDocument");
  doc_data.append("id", doc_id);
  error = false;

  //обработка резолюции
  resolution_panel = active_tub.querySelector("#resolution-panel-" + doc_id);
  resolution = resolution_panel.querySelector("#doc-newresolution");

  if (resolution) {
    resolution_data = new Object();
    text = resolution.querySelector("#resolution-text").value;
    user = resolution.querySelector("#resolution-user").value;
    date = resolution.querySelector("#resolution-date").value;
    resolution_data.text = text;
    resolution_data.user = user;
    resolution_data.date = date;
    resolution_data.id = resolution.getAttribute("id");

    if (resolution.querySelector("#resolution-ispolnitel")) {
      ispolnitel = resolution.querySelector("#resolution-ispolnitel").value;
      time = resolution.querySelector("#resolution-time").value;
      resolution_data.ispolnitel = ispolnitel;
      resolution_data.time = time;
    } else {
      result = resolution.querySelector("#resolution-result").value;
      resolution_data.result = result;
    }

    if (WalidResolutionData(resolution_data) === true) {
      return;
    } else doc_data.append("resolution", JSON.stringify(resolution_data));
  } else return;

  //запись данных в БД

  data_obj = Object.fromEntries(doc_data.entries());

  $.ajax({
    method: "POST",
    url: "../router.php",
    data: doc_data,
    cache: false,
    contentType: false,
    processData: false,
    success: function (check) {
      if (check == "OK") {
        alert(check);
      } else ServerMessage("show", check);
    },
    error: () =>
      ServerMessage("show", "Возникла ошибка на сервере, попробуйте позже!"),
  });
}

//------Функция изменения количества экземпляров Исходящего-----

outgoing_tub = document.querySelector("#main-tab-outgoing");
count_input = outgoing_tub.querySelector("#input-newdoc-count");

count_input.onchange = function () {
  sender_input = outgoing_tub.querySelector("#input-newdoc-sender");
  outgoing_sender =
    '<input class="table__text--input" id="input-newdoc-sender" name="sender" type="text" value="' +
    sender_input.value +
    '">';

  senders =
    '<input class="table__text--input" id="input-newdoc-sender" name="sender" type="text" value="По расчету-рассылки">';
  sender =
    '<input class="table__text--input" id="input-newdoc-sender" name="sender" type="text">';
  copy =
    '<input class="table__text--input" id="input-newdoc-copy" name="copy" type="text">';

  if (count_input.value < 2) {
    outgoing_tub.querySelector("#input-newdoc-count").value = 1;
    outgoing_tub.querySelector("#column-newdoc-sender").innerHTML =
      outgoing_sender;
    outgoing_tub.querySelector("#column-newdoc-copy").innerHTML = copy;
  }

  if (count_input.value > 5) {
    outgoing_tub.querySelector("#column-newdoc-sender").innerHTML = senders;
    outgoing_tub.querySelector("#column-newdoc-copy").innerHTML = copy;
  }

  if (count_input.value > 1 && count_input.value < 6) {
    sender = outgoing_sender;
    for (i = 1; i < count_input.value; i++) {
      sender +=
        '<input class="table__text--input" id="input-newdoc-sender' +
        i +
        '" name="sender' +
        i +
        '" type="text">';
      copy +=
        '<input class="table__text--input" id="input-newdoc-copy' +
        i +
        '" name="copy' +
        i +
        '" type="text">';
    }
    outgoing_tub.querySelector("#column-newdoc-sender").innerHTML = sender;
    outgoing_tub.querySelector("#column-newdoc-copy").innerHTML = copy;
  }
};

/*/------Функция изменения количества экземпляров Приказов-----

directive_tub = document.querySelector("#main-tab-directive");
count_dir_input = directive_tub.querySelector("#input-newdoc-count");

count_dir_input.onchange = function () {
  sender_input = directive_tub.querySelector("#column-newdoc-sender");
  letter_input = directive_tub.querySelector("#column-newdoc-letter");
  copy_input = directive_tub.querySelector("#column-newdoc-copy");

  sender =
    '<input class="table__text--input" id="input-newdoc-result" name="result" type="text">';

  letter =
    '<input class="table__text--input" id="input-newdoc-copy" name="copy" type="text">';

  copy =
    '<input class="table__text--input" id="input-newdoc-width" name="width" type="number">';

  if (count_dir_input.value < 2 || count_dir_input.value > 10) {
    count_dir_input.value = 1;
    sender_input.innerHTML = sender;
    letter_input.innerHTML = letter;
    copy_input.innerHTML = copy;
  }

  if (count_dir_input.value > 1 && count_dir_input.value < 11) {
    for (i = 1; i < count_dir_input.value; i++) {
      sender +=
        '<input class="table__text--input" id="input-newdoc-result' +
        i +
        '" name="result' +
        i +
        '" type="text">';

      letter +=
        '<input class="table__text--input" id="input-newdoc-copy ' +
        i +
        '" name="copy' +
        i +
        '" type="text">';

      copy +=
        '<input class="table__text--input" id="input-newdoc-width' +
        i +
        '" name="width' +
        i +
        '" type="text">';
    }
    sender_input.innerHTML = sender;
    letter_input.innerHTML = letter;
    copy_input.innerHTML = copy;
  }
};

//---Проверка введенных данных при создании или изменении документа
function WalidDocumentData(data) {
  active_tub = document.querySelector(".main__tabs--active");
  error = false;
  fnum = data.fnum.trim();
  fdate = data.fdate.trim();
  name = data.name.trim();
  sender = data.sender.trim();
  ispolnitel = data.ispolnitel.trim();
  count = data.count.trim();
  copy = data.copy.trim();
  width = data.width.trim();
  file = data.file;

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
    fnum == "Вх. №" ||
    fnum == "№" ||
    fnum == "Вх." ||
    fnum == "Вх"
  ) {
    {
      WalidDocumentError("fnum");
      error = true;
    }
  } else WalidDocumentReError("fnum");

  //---Проверка даты документа------
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

  //---Проверка отправителя документа------
  if (sender == "") {
    {
      WalidDocumentError("sender");
      error = true;
    }
  } else WalidDocumentReError("sender");

  //---Проверка исполнителя документа------
  if (!data.resolutions) {
    if (ispolnitel == "") {
      WalidDocumentError("ispolnitel");
      error = true;
    } else WalidDocumentReError("ispolnitel");
  } else {
    if (WalidResolutionData(data.resolutions) == true) {
      error = true;
    }
    WalidDocumentReError("ispolnitel");
  }

  //---Проверка количества экземпляров документа------
  if (count == "") {
    {
      WalidDocumentError("count");
      error = true;
    }
  } else WalidDocumentReError("count");

  //---Проверка номеров экземпляров документа------
  if (copy == "" || copy <= 0) {
    {
      WalidDocumentError("copy");
      error = true;
    }
  } else WalidDocumentReError("copy");

  //---Проверка количества листов документа------
  if (width == "") {
    {
      WalidDocumentError("width");
      error = true;
    }
  } else WalidDocumentReError("width");
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
function WalidResolutionData(resolutions) {
  data = JSON.parse(resolutions);
  data.forEach((resolution) => {
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
  });
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
/*
//------Функция запоминания документа-----
function LikeDoc (id)
{
  alert ("Запомнить " + id);
}


*/
