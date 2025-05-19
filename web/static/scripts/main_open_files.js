const SELECTORS = {
  PANEL_OPEN_DOC: '#panel-opendoc',
  IFRAME_OPEN_DOC: '#iframe-opendoc',
  RESOLUTIONS_OPEN_DOC: '#resolutions-opendoc',
  ACCOUNT_NAME: '#account-name',
  OPEN_DOC_BTN: '.panel__opendoc--btn',
  OPEN_NEW_DOC_BTN: '#btn-newdoc-open'
};

// Инициализация обработчиков событий
function initDocumentViewHandlers() {
  // Обработчик для кнопки просмотра документа
  document.querySelector(SELECTORS.OPEN_DOC_BTN).addEventListener('click', () => {
    openDocument('close', 'id');
  });

  // Обработчики для кнопок просмотра нового документа
  document.querySelectorAll(SELECTORS.OPEN_NEW_DOC_BTN).forEach(btn => {
    btn.addEventListener('click', openNewDocument);
  });
}

// Функция открытия/закрытия просмотра документа
async function openDocument(action, docId) {
  const panel = document.querySelector(SELECTORS.PANEL_OPEN_DOC);
  const iframe = document.querySelector(SELECTORS.IFRAME_OPEN_DOC);
  const resolutionsPanel = document.querySelector(SELECTORS.RESOLUTIONS_OPEN_DOC);

  if (action === 'close') {
    closeDocumentView(panel, iframe, resolutionsPanel);
    return;
  }

  const activeTab = document.querySelector('.main__tabs--active');
  tabId = "#"+activeTab.id;
  const docTypeConfig = Object.values(DOCUMENT_TYPES).find(
    type => type.tabId === tabId
  );

  // Открытие документа
  document.body.style.overflow = 'hidden';
  panel.style.display = 'flex';

  iframe.innerHTML = getFileViewUrl(action)

  // Обработка резолюций
  const docTable = document.getElementById(`document-table-${docId}`);
  if (docTable) {
    const familiarCell = docTable.querySelector('.table__column--familiar');
    const resolutionPanel = document.getElementById(`resolution-panel-${docId}`);

    if (resolutionPanel && resolutionPanel.innerHTML !== '') {
      resolutionsPanel.style.minWidth = '294px';
      resolutionsPanel.innerHTML = resolutionPanel.innerHTML;
    }

    // Запись просмотра документа
    const accountName = document.querySelector(SELECTORS.ACCOUNT_NAME)?.textContent.trim();
    const familiarText = familiarCell?.textContent || '';

    if (accountName && !familiarText.includes(accountName)) {
        
    }
    await fetchFamiliarDocument(docTypeConfig.type, docId);
  }
}

// Функция открытия нового документа
function openNewDocument() {
  const activeTab = document.querySelector('.main__tabs--active');
  const newDocResolutions = activeTab.querySelector('#newdoc-resolution-panel');
  const panel = document.querySelector(SELECTORS.PANEL_OPEN_DOC);
  const resolutionsPanel = document.querySelector(SELECTORS.RESOLUTIONS_OPEN_DOC);
  const iframe = document.querySelector(SELECTORS.IFRAME_OPEN_DOC);
  const fileInput = activeTab.querySelector('#input-newdoc-file');
  const file = fileInput.files[0];

  document.body.style.overflow = 'hidden';
  panel.style.display = 'flex';

  if (file && (file.type === "application/pdf" || file.type.startsWith("image"))) {
    const fileUrl = URL.createObjectURL(file);
    iframe.innerHTML = getFilePreview(fileUrl, file.type);
  }

  if (newDocResolutions && newDocResolutions.children.length > 0) {
    resolutionsPanel.style.minWidth = '280px';
    resolutionsPanel.innerHTML = ''; 
    
    // Копируем каждый элемент с его содержимым и значениями
    Array.from(newDocResolutions.children).forEach(child => {
      const clone = child.cloneNode(true); // Глубокое копирование
      
      // Восстанавливаем значения полей ввода
      const inputs = child.querySelectorAll('input, textarea');
      const cloneInputs = clone.querySelectorAll('input, textarea');
      
      inputs.forEach((input, index) => {
        if (input.type !== 'file') {
          cloneInputs[index].value = input.value;
        }
      });
      
      resolutionsPanel.appendChild(clone);
    });
  }
}

// Закрытие просмотра документа
function closeDocumentView(panel, iframe, resolutionsPanel) {
  panel.style.display = 'none';
  iframe.innerHTML = ""
  iframe.style.width = '70%';
  iframe.style.height = 'auto';
  resolutionsPanel.innerHTML = '';
  resolutionsPanel.style.minWidth = '80px';
  document.body.style.overflow = 'auto';
}

// Функция для получения URL превью файла
function getFilePreviewUrl(url, type) {
  if (type.startsWith('image')) {
    return url;
  } else if (type === 'application/pdf') {
    return url;
  }
  return '';
}
// Функция для получения URL файла только па названию
function getFileViewUrl(file) {
  const extension = file.split('.').pop().toLowerCase();
  const imageExtensions = ['jpg', 'jpeg', 'png', 'gif', 'webp', 'bmp'];
  const pdfExtensions = ['pdf'];

  if (imageExtensions.includes(extension)) {
    return `<img src="${file}" class="file-content">`
  }
  if (pdfExtensions.includes(extension)) {
    return `<embed src="${file}" type="application/pdf" class="file-content">`
  }
}

// Инициализация при загрузке страницы
document.addEventListener('DOMContentLoaded', initDocumentViewHandlers);