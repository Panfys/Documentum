date = new Date().toJSON().slice(0, 10).replace(/-/g, "/"); //дата

//-----Открытие боковых панелей------------

//------------------------------MAIN-------------------------------------

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

/*------Функция сортировки-------------
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
} */

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

*/
