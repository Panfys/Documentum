// Инициализация обработчиков 
function initDocumentHandlers() {
  // Кнопка нового документа/отмены
  document.querySelectorAll("#btn-newdoc").forEach(btn => {
    btn.addEventListener("click", toggleNewDocumentForm);
  });
  
  // Кнопка очистки формы
  document.querySelectorAll("#btn-newdoc-clearnewdoc").forEach(btn => {
    btn.addEventListener("click", clearNewDocumentForm);
  });
  
  // Кнопка прикрепления файла
  document.querySelectorAll("#btn-newdoc-addfile").forEach(btn => {
    btn.addEventListener("click", handleFileUpload);
  });
  
  // Кнопка записи документа
  document.querySelectorAll("#btn-newdoc-addnewdoc").forEach(btn => {
    btn.addEventListener("click", saveNewDocument);
  });
  
  // Кнопки работы с резолюциями
  initResolutionHandlers();
}

// Открытие/Закрытие формы нового документа
function toggleNewDocumentForm() {
  const activeTab = document.querySelector(".main__tabs--active");
  const form = activeTab.querySelector("#form-newdoc");

  if (form.style.display === "flex") {
    closeNewDocumentForm(activeTab);
  } else {
    openNewDocumentForm(activeTab);
  }
}

// Открытие формы нового документа
function openNewDocumentForm(activeTab) {
  const form = activeTab.querySelector("#form-newdoc");
  const table = form.querySelector("#table-newdoc");
  const btn = activeTab.querySelector("#btn-newdoc");
  const folder = activeTab.querySelector("#head-folder");
  const panel = activeTab.querySelector("#title-newdocpanel");
  const span = activeTab.querySelector("#title-span");
  const searchBtn = activeTab.querySelector("#btn-search");
  
  form.style.display = "flex";
  btn.textContent = "Отмена";
  span.style.display = "flex";
  folder.textContent = "Запись нового документа";
  panel.style.display = "flex";
  searchBtn.style.display = "none";
  table.classList.add("tubs__table--active-table");
  
  window.scrollTo({
    top: 0,
    left: 0,
    behavior: "smooth"
  });
}

// Закрытие формы нового документа
function closeNewDocumentForm(activeTab) {
  const form = activeTab.querySelector("#form-newdoc");
  const table = form.querySelector("#table-newdoc");
  const btn = activeTab.querySelector("#btn-newdoc");
  const folder = activeTab.querySelector("#head-folder");
  const panel = activeTab.querySelector("#title-newdocpanel");
  const span = activeTab.querySelector("#title-span");
  const searchBtn = activeTab.querySelector("#btn-search");
  
  if (btn.style.display === "none") {
    AddNewdocResolution("back"); // Вызываем закрытие панели резолюций
  }
  
  form.style.display = "none";
  btn.textContent = "Новый документ";
  span.style.display = "none";
  folder.textContent = "";
  panel.style.display = "none";
  searchBtn.style.display = "flex";
  table.classList.remove("tubs__table--active-table");
}

// Очистка нового докумнета
function clearNewDocumentForm() {
  const activeTab = document.querySelector(".main__tabs--active");
  activeTab.querySelector("#form-newdoc").reset();
  clearFilePanel();
  clearResolutionPanel();
}

function clearFilePanel() {
  const activeTab = document.querySelector(".main__tabs--active");
  const btnFile = activeTab.querySelector("#btn-newdoc-addfile");
  const filePanel = activeTab.querySelector("#newdoc-file-panel");
  const fileName = activeTab.querySelector("#newdoc-file-name");
  const fileSize = activeTab.querySelector("#newdoc-file-size");
  const fileImg = activeTab.querySelector("#newdoc-file-img");

  filePanel.style.display = "none";
  fileName.textContent = "";
  fileSize.textContent = "";
  fileImg.innerHTML = "";
  btnFile.textContent = "Файл";
}

function clearResolutionPanel() {
  const activeTab = document.querySelector(".main__tabs--active");
  const resolutionPanel = activeTab.querySelector("#newdoc-resolution-panel");
  const inputIspolnitel = activeTab.querySelector("#input-newdoc-ispolnitel");

  resolutionPanel.innerHTML = "";
  inputIspolnitel.setAttribute("placeholder", "");
}

function handleFileUpload() {
  const activeTab = document.querySelector(".main__tabs--active");
  const fileInput = activeTab.querySelector("#input-newdoc-file");
  const filePanel = activeTab.querySelector("#newdoc-file-panel");
  const fileName = activeTab.querySelector("#newdoc-file-name");
  const fileSize = activeTab.querySelector("#newdoc-file-size");
  const fileImg = activeTab.querySelector("#newdoc-file-img");
  const btnFile = activeTab.querySelector("#btn-newdoc-addfile");

  fileInput.click();
  
  fileInput.onchange = function() {
    const file = fileInput.files[0];
    if (file && (file.type === "application/pdf" || file.type.startsWith("image"))) {
      const fileUrl = URL.createObjectURL(file);
      
      fileName.textContent = file.name;
      fileImg.innerHTML = getFilePreview(fileUrl, file.type);
      fileSize.textContent = formatFileSize(file.size);
      filePanel.style.display = "flex";
      btnFile.textContent = "Изменить файл";
    }
  };
}

function getFilePreview(url, type) {
  if (type.startsWith("image")) {
    return `<img src="${url}" alt="File preview">`;
  } else if (type === "application/pdf") {
    return `<embed src="${url}" scrolling="no" type="application/pdf">`;
  }
  return `<img src="/style/images/file error.png" alt="Invalid file">`;
}

function formatFileSize(bytes) {
  const units = ["Б", "КБ", "МБ", "ГБ", "ТБ"];
  let i = 0;
  let n = parseInt(bytes, 10) || 0;

  while (n >= 1000 && ++i) {
    n /= 1000;
  }

  return n.toFixed(n < 10 && i > 0 ? 1 : 0) + " " + units[i];
}

function saveNewDocument() {
  const activeTab = document.querySelector(".main__tabs--active");
  const form = activeTab.querySelector("#form-newdoc");
  const formData = new FormData(form);
  
  // Определение типа документа
  const tabId = activeTab.id;
  let docType;
  
  switch (tabId) {
    case "main-tab-ingoing":
      docType = DOCUMENT_TYPES.INGOING.type;
      processResolutions(activeTab, formData);
      break;
    case "main-tab-outgoing":
      docType = DOCUMENT_TYPES.OUTGOING.type;
      break;
    case "main-tab-directive":
      docType = DOCUMENT_TYPES.DIRECTIVE.type;
      break;
    case "main-tab-inventory":
      docType = DOCUMENT_TYPES.INVENTORY.type;
      break;
  }
  
  formData.append("type", docType);
  
  // Вызов сохраненной функции без изменений
  const dataObj = Object.fromEntries(formData.entries());
  if (WalidDocumentData(dataObj) === true) {
    return;
  }
  
  // Вызов сохраненной функции без изменений
  AddDocument(formData);
}

function processResolutions(activeTab, formData) {
  const resolutionPanel = activeTab.querySelector("#newdoc-resolution-panel");
  const resolutions = resolutionPanel.querySelectorAll(".table__resolution");
  const resolutionsData = [];
  
  resolutions.forEach(resolution => {
    const resolutionData = {
      id: resolution.getAttribute("id"),
      text: resolution.querySelector("#resolution-text")?.value || "",
      user: resolution.querySelector("#resolution-user")?.value || "",
      date: resolution.querySelector("#resolution-date")?.value || ""
    };
    
    const ispolnitelInput = resolution.querySelector("#resolution-ispolnitel");
    if (ispolnnitelInput) {
      resolutionData.ispolnitel = ispolnitelInput.value || "";
      resolutionData.time = resolution.querySelector("#resolution-time")?.value || "";
    } else {
      resolutionData.ispolnitel = "NULL";
      resolutionData.result = resolution.querySelector("#resolution-result")?.value || "";
    }
    
    resolutionsData.push(resolutionData);
  });
  
  formData.append("resolutions", JSON.stringify(resolutionsData));
}

// Инициализация при загрузке страницы
document.addEventListener('DOMContentLoaded', initDocumentHandlers);