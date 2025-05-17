// Константы для элементов
const RESOLUTION_TYPES = {
  NEW_DOC: 'newdoc',
  EXISTING_DOC: 'doc'
};
let resolutionStartCount = null;

// Инициализация обработчиков резолюций
function initResolutionHandlers() {
  document.querySelector("#btn-newdoc-resolution").onclick = () =>
    AddNewdocResolution("open");

  document.querySelector("#btn-doc-resolution").onclick = () =>
    AddDocResolution("open");

  document.querySelector("#btn-resolution-add").onclick = () =>
    handleActiveDocumentResolution("add");

  document.querySelector("#btn-resolution-result").onclick = () =>
    handleActiveDocumentResolution("result");

  document.querySelector("#btn-resolution-remove").onclick = () =>
    handleActiveDocumentResolution("remove");

  document.querySelector("#btn-resolution-back").onclick = () =>
    handleActiveDocumentResolution("back");
}

// Обработка действий для активного документа
function handleActiveDocumentResolution(action) {
  const activeDoc = document.querySelector(".tubs__table--active-table");
  const docType = activeDoc.id === "table-newdoc"
    ? RESOLUTION_TYPES.NEW_DOC
    : RESOLUTION_TYPES.EXISTING_DOC;

  if (docType === RESOLUTION_TYPES.NEW_DOC) {
    AddNewdocResolution(action);
  } else {
    AddDocResolution(action);
  }
}

// Основная функция для работы с резолюциями нового документа
function AddNewdocResolution(action) {
  const activeTab = document.querySelector(".main__tabs--active");
  const resolutionPanel = activeTab.querySelector("#newdoc-resolution-panel");
  const resolutionCount = resolutionPanel.childElementCount;
  const btnAdd = activeTab.querySelector("#btn-resolution-add");
  const btnPanel = activeTab.querySelector("#btn-resolution-panel");
  const docPanel = activeTab.querySelector("#title-newdocpanel");
  const docBtn = activeTab.querySelector("#btn-newdoc");
  const inputIspolnitel = activeTab.querySelector("#input-newdoc-ispolnitel");

  switch (action) {
    case "open":
      docPanel.style.display = "none";
      btnPanel.style.display = "flex";
      docBtn.style.display = "none";
      break;

    case "add":
      addNewResolution(resolutionPanel, false, resolutionPanel.childElementCount);
      inputIspolnitel.setAttribute("placeholder", "Заполняется автоматически");
      btnAdd.style.backgroundColor = "var(--low-color)";
      break;

    case "result":
      if (resolutionCount === 0) {
        flashButton(btnAdd);
      } else {
        addNewResolution(resolutionPanel, true, resolutionPanel.childElementCount);
      }
      break;

    case "back":
      btnPanel.style.display = "none";
      docPanel.style.display = "flex";
      docBtn.style.display = "flex";
      break;

    case "remove":
      if (resolutionCount > 0) {
        removeResolution(resolutionPanel);
        if (resolutionCount === 1) {
          inputIspolnitel.setAttribute("placeholder", "");
        }
      }
      break;
  }
}

//------Функция кнопки "Резолюция"------


function AddDocResolution(action) {
  const activeTab = document.querySelector(".main__tabs--active");
  const activeDoc = document.querySelector(".tubs__table--active-table");
  const activeDocId = activeDoc.getAttribute("document-id");
  const resolutionPanel = activeTab.querySelector("#resolution-panel-" + activeDocId);
  const resolutionCount = resolutionPanel.childElementCount;
  const btnAdd = activeTab.querySelector("#btn-resolution-add");
  const btnPanel = activeTab.querySelector("#btn-resolution-panel");
  const docPanel = activeTab.querySelector("#title-docpanel");
  const newResolution = document.createElement("div");

  newResolution.setAttribute("class", "table__resolution");
  newResolution.setAttribute("id", "doc-newresolution");

  switch (action) {
    case "open":
      if (resolutionStartCount === null) {
        resolutionStartCount = resolutionPanel.childElementCount;
        resolutionPanel.setAttribute("res_count", resolutionPanel.childElementCount);
      }
      docPanel.style.display = "none";
      btnPanel.style.display = "flex";
      break;

    case "add":
      addNewResolution(resolutionPanel, false, resolutionPanel.childElementCount);
      break;

    case "result":
      if (resolutionCount === 0) {
        flashButton(btnAdd);
      } else {
        addNewResolution(resolutionPanel, true, resolutionPanel.childElementCount);
      }
      break;

    case "back":
      btnPanel.style.display = "none";
      docPanel.style.display = "flex";
      checkSaveBtn (resolutionCount - resolutionStartCount)
      break;

    case "remove":
      // Используем сохраненное начальное значение
      console.log(resolutionStartCount)
      if (resolutionStartCount !== null && resolutionCount > resolutionStartCount) {
        resolutionPanel.removeChild(resolutionPanel.lastChild);
      }
      break;
  }
}

// Вспомогательные функции
function addNewResolution(panel, isResult, id) {
  const resolution = document.createElement("div");
  resolution.className = "table__resolution";
  resolution.id = `resolution-${id}`;

  resolution.innerHTML = isResult
    ? `
      <textarea id="resolution-text"></textarea>
      <input id="resolution-result" type="text" placeholder="Исполненный документ">
      <input class="resolution__user--input" type="text" id="resolution-user" placeholder="Исполнитель">
      <input class="resolution__date--input" type="date" id="resolution-date">
    `
    : `
      <input id="resolution-ispolnitel" type="text" placeholder="Исполнитель">
      <textarea id="resolution-text"></textarea>
      <div class="resolution__time--block">Срок исполнения 
        <input id="resolution-time" type="date">
      </div>
      <input class="resolution__user--input" type="text" id="resolution-user" placeholder="Руководитель">
      <input class="resolution__date--input" type="date" id="resolution-date">
    `;

  panel.appendChild(resolution);
  scrollPanel(panel);
}

function removeResolution(panel) {
  const resolutions = panel.querySelectorAll(".table__resolution");
  if (resolutions.length > 0) {
    const lastResolution = resolutions[resolutions.length - 1];
    panel.removeChild(lastResolution);
  }
}

function scrollPanel(panel) {
  panel.scrollTo({
    top: 0,
    left: 5000,
    behavior: "smooth"
  });
}

// добавление/удаление кнопки сохранить
function checkSaveBtn(countRes) {
  const activeTab = document.querySelector(".main__tabs--active");
  const btnSave = activeTab.querySelector("#btn-doc-save");
  if (countRes > 0) {
    btnSave.style.display = "flex";
  } else {
    btnSave.style.display = "none";
  }
}

// Подствечиваем красным кнопку, если валидация не пройдена
function flashButton(button) {
  button.style.backgroundColor = "var(--error-color)";
  setTimeout(() => {
    button.style.backgroundColor = "var(--low-color)";
  }, 3000);
}

// Функция очистки панели резолюций
function ClearNewdocResolutionPanel() {
  const activeTab = document.querySelector(".main__tabs--active");
  const resolutionPanel = activeTab.querySelector("#newdoc-resolution-panel");
  const inputIspolnitel = activeTab.querySelector("#input-newdoc-ispolnitel");

  resolutionPanel.innerHTML = "";
  inputIspolnitel.setAttribute("placeholder", "");
}

// Инициализация при загрузке страницы
document.addEventListener('DOMContentLoaded', initResolutionHandlers);