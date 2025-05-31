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
    btn.addEventListener("click", submitNewDocumentForm);
  });

  // Кнопка сохранения изменений документа
  document.querySelectorAll("#btn-doc-save").forEach(btn => {
    btn.addEventListener("click", submitUpdateDocument);
  });

  // Кнопка поиска
  document.querySelectorAll("#btn-search").forEach(
    btn => {
      btn.addEventListener("click", () => { alert("Извините, поиск пока недоступен!") })
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

// Очистка панели добавления файла
function clearFilePanel() {
  const activeTab = document.querySelector(".main__tabs--active");
  const btnFile = activeTab.querySelector("#btn-newdoc-addfile");
  const filePanel = activeTab.querySelector("#newdoc-file-panel");
  const fileName = activeTab.querySelector("#newdoc-file-name");
  const fileSize = activeTab.querySelector("#newdoc-file-size");
  const fileImg = activeTab.querySelector("#newdoc-file-img");
  const fileInput = activeTab.querySelector("#input-newdoc-file");

  filePanel.style.display = "none";
  fileName.textContent = "";
  fileSize.textContent = "";
  fileImg.innerHTML = "";
  btnFile.textContent = "Файл";
  fileInput.value = "";
}

// Очистка панели резолюций
function clearResolutionPanel() {
  const activeTab = document.querySelector(".main__tabs--active");
  const resolutionPanel = activeTab.querySelector("#newdoc-resolution-panel");
  const inputIspolnitel = activeTab.querySelector("#input-newdoc-ispolnitel");

  resolutionPanel.innerHTML = "";
  inputIspolnitel.setAttribute("placeholder", "");

}

// Добавление файла новому документу
function handleFileUpload() {
  const activeTab = document.querySelector(".main__tabs--active");
  const fileInput = activeTab.querySelector("#input-newdoc-file");
  const filePanel = activeTab.querySelector("#newdoc-file-panel");
  const fileName = activeTab.querySelector("#newdoc-file-name");
  const fileSize = activeTab.querySelector("#newdoc-file-size");
  const fileImg = activeTab.querySelector("#newdoc-file-img");
  const btnFile = activeTab.querySelector("#btn-newdoc-addfile");

  fileInput.click();

  fileInput.onchange = function () {
    const file = fileInput.files[0];
    if (file && (!validDocFileType(file.type))) {
      const fileUrl = URL.createObjectURL(file);

      fileName.textContent = file.name;
      fileImg.innerHTML = getFilePreview(fileUrl, file.type);
      fileSize.textContent = formatFileSize(file.size);
      filePanel.style.display = "flex";
      btnFile.textContent = "Изменить файл";
    } else {
      clearFilePanel();
    }
  };
}

// показывает превью файла
function getFilePreview(url, type) {
  if (type.startsWith("image")) {
    return `<img src="${url}" alt="File preview">`;
  } else if (type === "application/pdf") {
    return `<embed src="${url}" scrolling="no" type="application/pdf">`;
  }
  return `<img src="/style/images/file error.png" alt="Invalid file">`;
}

// форматирует вывод размера файла
function formatFileSize(bytes) {
  const units = ["Б", "КБ", "МБ", "ГБ", "ТБ"];
  let i = 0;
  let n = parseInt(bytes, 10) || 0;

  while (n >= 1000 && ++i) {
    n /= 1000;
  }

  return n.toFixed(n < 10 && i > 0 ? 1 : 0) + " " + units[i];
}

async function submitNewDocumentForm() {
  const activeTab = document.querySelector(".main__tabs--active");
  const resolutionPanel = activeTab.querySelector("#newdoc-resolution-panel");
  const form = activeTab.querySelector("#form-newdoc");
  const fileInput = form.querySelector('input[type="file"]');
  const tabId = `#${activeTab.id}`;
  const docType = Object.values(DOCUMENT_TYPES).find(
    type => type.tabId === tabId
  );
  const docTable = Object.values(DOCUMENT_TYPES).find(
    table => table.tabId === tabId
  );
  const docData = {};
  const formData = new FormData(form);
  docData.type = docType.type;

  processResolutions(resolutionPanel, docData);

  // 1. Собираем данные формы в объект (без файла)
  for (const [key, value] of formData.entries()) {
    // Пропускаем файловые поля
    if (!(value instanceof File)) {
      docData[key] = value.trim();
    }
  }

  if (fileInput && fileInput.files[0]) {
    docData.fileType = fileInput.files[0].type;
  }
  else docData.fileType = ""

  if (docData.type === "directives") {
    if (validateDirectiveData(docData) > 0) {
      return;
    }
  } else if (docData.type === "inventory") {
    if (validateInventoryData(docData) > 0) {
      return;
    }
  } else {
    // 2. Валидация данных
    if (validateDocumentData(docData) > 0) {
      return;
    }
  }

  const uploadFormData = new FormData();

  uploadFormData.append('document', JSON.stringify(docData));

  if (fileInput && fileInput.files[0]) {
    uploadFormData.append('file', fileInput.files[0]);
  }

  if (await fetchAddDocument(docTable.table, uploadFormData)) {
    CloseActiveDoc();
    window.scrollTo({
      top: document.body.scrollHeight, // до конца страницы
      behavior: 'smooth' // плавная прокрутка
    });
  }


}

function processResolutions(resolutionPanel, docData) {
  if (!resolutionPanel) {
    return;
  }

  // Массив для хранения резолюций
  const resolutionsData = [];

  resolutionPanel.querySelectorAll(".table__resolution").forEach(resolution => {
    // Определяем тип резолюции
    const isTask = !!resolution.querySelector("#resolution-ispolnitel");

    // Базовый объект резолюции
    const resolutionData = {
      type: isTask ? "task" : "result",
      text: resolution.querySelector("#resolution-text")?.value || "",
      user: resolution.querySelector("#resolution-user")?.value || "",
      date: resolution.querySelector("#resolution-date")?.value || ""
    };

    // Добавляем специфичные поля
    if (isTask) {
      resolutionData.ispolnitel =
        resolution.querySelector("#resolution-ispolnitel")?.value || "";
      resolutionData.deadline =
        resolution.querySelector("#resolution-time")?.value || "";
    } else {
      resolutionData.result =
        resolution.querySelector("#resolution-result")?.value || "";
    }

    if (resolution.getAttribute("id") !== "ingoing-resolution") {
      resolutionsData.push(resolutionData);
    }
  });

  // Добавляем резолюции в основной объект формы
  docData.resolutions = resolutionsData;
}

async function submitUpdateDocument() {
  const activeTab = document.querySelector(".main__tabs--active");
  const activeTable = document.querySelector(".tubs__table--active-table");
  const documentID = activeTable.getAttribute("document-id");
  const resolutionPanel = document.querySelector(`#resolution-panel-${documentID}`);
  const tabId = `#${activeTab.id}`;
  const docTable = Object.values(DOCUMENT_TYPES).find(
    table => table.tabId === tabId
  );
  const docData = {};
  startResCount = resolutionPanel.getAttribute("res_count")
  processResolutions(resolutionPanel, docData);
  if (docData.resolutions.length > 0) {
    if (validResolutions(docData.resolutions, startResCount)) return;
  }
  const updateFormData = new FormData();

  updateFormData.append('document', JSON.stringify(docData));

  if (await fetchUpdateDocument(docTable.table, updateFormData, documentID)) {
    handleActiveDocumentResolution("removeAll");
  }
}

// Инициализация при загрузке страницы
document.addEventListener('DOMContentLoaded', initDocumentHandlers);