
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
    iframe_opendoc.setAttribute("src", "/document/file?file=" + ask);
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
      AddViewDocument(id)
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
        "/document/new/file?file=" + DefineFile(file_url, file.type)
      );
    }

    panel_opendoc.style.display = "flex";
    if (newdoc_resolutions.innerHTML !== "") {
      resolutions_opendoc.style = "min-width: 280px;";
      resolutions_opendoc.innerHTML = newdoc_resolutions.innerHTML;
    }
  });
});

