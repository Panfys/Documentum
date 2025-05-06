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

  filePanel.style.display = "none";
  fileName.textContent = "";
  fileSize.textContent = "";
  fileImg.innerHTML = "";
  btnFile.textContent = "Файл";
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
  const form = activeTab.querySelector("#form-newdoc");
  const tabId = `#${activeTab.id}`;
  const docType = Object.values(DOCUMENT_TYPES).find(
    type => type.tabId === tabId
  );
  const docData = {};
  const formData = new FormData(form);
  docData.createdAt = new Date().toISOString();
  docData.type = docType.type;

  processResolutions(activeTab, docData);

  // 1. Собираем данные формы в объект (без файла)
  for (const [key, value] of formData.entries()) {
    // Пропускаем файловые поля
    if (!(value instanceof File)) {
      docData[key] = value;
    }
  }

  console.log(docData)
  return
  // 2. Валидация данных
  if (!validateFormData(formDataObj)) {
    return false;
  }

  // 3. Создаем новый FormData для отправки
  const uploadFormData = new FormData();

  // Добавляем текстовые данные как JSON
  uploadFormData.append('data', JSON.stringify(formDataObj));

  // Добавляем файл отдельно
  const fileInput = form.querySelector('input[type="file"]');
  if (fileInput && fileInput.files[0]) {
    uploadFormData.append('file', fileInput.files[0]);
  }

  // 4. Отправка на сервер
  try {
    const response = await fetch('/api/upload', {
      method: 'POST',
      body: uploadFormData
    });

    if (!response.ok) throw new Error('Ошибка сервера');

    const result = await response.json();
    console.log('Успешно отправлено:', result);
    return true;
  } catch (error) {
    console.error('Ошибка отправки:', error);
    return false;
  }
}

// Функция валидации
function validateDocumentData(data) {
  // Пример валидации - проверка обязательных полей
  if (!data.name || data.name.trim() === '') {
    alert('Поле "Имя" обязательно для заполнения');
    return false;
  }

  // Добавьте другие проверки по необходимости
  return true;
}

function processResolutions(activeTab, docData) {
  const resolutionPanel = activeTab.querySelector("#newdoc-resolution-panel");

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
      createdAt: new Date().toISOString(),
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

    resolutionsData.push(resolutionData);
  });

  // Добавляем резолюции в основной объект формы
  docData.resolutions = resolutionsData;
}

// Инициализация при загрузке страницы
document.addEventListener('DOMContentLoaded', initDocumentHandlers);