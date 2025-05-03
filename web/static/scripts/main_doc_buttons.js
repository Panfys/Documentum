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
            ispolnitel = "NULL"
            resolution_data.ispolnitel = ispolnitel;
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

   AddDocument(newdoc_data)
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
      ispolnitel = "NULL"
      resolution_data.ispolnitel = ispolnitel;
      result = resolution.querySelector("#resolution-result").value;
      resolution_data.result = result;
    }

    if (WalidResolutionData(resolution_data) === true) {
      return;
    } else doc_data.append("resolution", JSON.stringify(resolution_data));
  } else return;

  //запись данных в БД

  data_obj = Object.fromEntries(doc_data.entries());

  ChangeDocument(data_obj)
}