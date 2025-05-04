// Константы для элементов интерфейса
const menuButtons = document.querySelectorAll(".header__menu--btn");
const settingsButton = document.querySelector(".main__header--settings");
const accountButton = document.querySelector(".main__header--account");

// Типы документов для табов
const DOCUMENT_TYPES = {
  INGOING: {
    tabId: "#main-tab-ingoing",
    type: "Входящий",
    defaultBtnId: "menu-btn-ingoing",
    documentTableId: "ingoing-documents-container"
  },
  OUTGOING: {
    tabId: "#main-tab-outgoing",
    type: "Исходящий",
    documentTableId: "outgoing-documents-container"
  },
  DIRECTIVE: {
    tabId: "#main-tab-directive",
    type: "Приказ",
    documentTableId: "directive-documents-container"
  },
  INVENTORY: {
    tabId: "#main-tab-inventory",
    type: "Издание",
    documentTableId: "inventory-documents-container"
  }
};

// Открытие боковых панелей
settingsButton.addEventListener("click", () => togglePanel("settings-panel"));
accountButton.addEventListener("click", () => togglePanel("account-panel"));

function togglePanel(panelId) {
  const panel = document.getElementById(panelId);
  panel.style.display = panel.style.display === "flex" ? "none" : "flex";
}

// Инициализация табов
menuButtons.forEach(btn => {
  btn.addEventListener("click", () => changeActiveTab(btn));
});

// Установка начального активного таба
function initializeActiveTab() {
  const savedButtonId = sessionStorage.getItem("active_btn_id");
  const defaultButtonId = DOCUMENT_TYPES.INGOING.defaultBtnId;
  const initialButton = document.getElementById(savedButtonId || defaultButtonId);
  changeActiveTab(initialButton);
}

initializeActiveTab();

function changeActiveTab(btn) {
  // Получаем предыдущие активные элементы
  const prevActiveTab = document.querySelector(".main__tabs.main__tabs--active");
  const prevActiveBtn = document.querySelector(".header__menu--btn.menu__btn--active");

  // Удаляем активные классы
  prevActiveBtn?.classList.remove("menu__btn--active");
  prevActiveTab?.classList.remove("main__tabs--active");

  // Устанавливаем новые активные элементы
  const activeTabId = `#${btn.dataset.tab}`;
  sessionStorage.setItem("active_btn_id", btn.id);
  const activeTab = document.querySelector(activeTabId);

  btn.classList.add("menu__btn--active");
  activeTab.classList.add("main__tabs--active");

  // Обновляем документы в зависимости от активного таба
  updateDocumentsForTab(activeTabId);
}

async function updateDocumentsForTab(tabId) {
  // Находим соответствующий тип документа
  const docTypeConfig = Object.values(DOCUMENT_TYPES).find(
    type => type.tabId === tabId
  );

  if (!docTypeConfig) {
    console.error(`Не найден тип документа для tabId: ${tabId}`);
    return;
  }

  // Получаем документы
  const documents = await FetchGetDocuments({ type: docTypeConfig.type });

  // Обновляем соответствующий контейнер
  const container = document.getElementById(docTypeConfig.documentTableId);
  WriteDocuments(documents, container)
  setupDocumentTables()
}

// Обработка кликов по таблицам документов
function setupDocumentTables() {
  const documentTables = document.querySelectorAll(".tubs__table--document");

  documentTables.forEach(table => {
    table.addEventListener("click", (event) => {
      ViewDocumentTable(table, event);
    });
  });
}