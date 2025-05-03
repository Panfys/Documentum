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
