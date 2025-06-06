const menuButtons = document.querySelectorAll(".header__menu--btn");
const settingsButton = document.querySelector(".main__header--settings");
const accountButton = document.querySelector(".main__header--account");

// Типы документов для табов
const DOCUMENT_TYPES = {
  INGOING: {
    tabId: "#main-tab-ingoing",
    type: "ingoing",
    defaultBtnId: "menu-btn-ingoing",
    documentTableId: "ingoing-documents-container",
    table: "inouts"
  },
  OUTGOING: {
    tabId: "#main-tab-outgoing",
    type: "outgoing",
    documentTableId: "outgoing-documents-container",
    table: "inouts"
  },
  DIRECTIVE: {
    tabId: "#main-tab-directive",
    type: "directives",
    documentTableId: "directive-documents-container",
    table: "directives"
  },
  INVENTORY: {
    tabId: "#main-tab-inventory",
    type: "inventory",
    documentTableId: "inventory-documents-container",
    table: "inventory"
  }
};

// Открытие боковых панелей
settingsButton.addEventListener("click", () => togglePanel("settings-panel"));
accountButton.addEventListener("click", () => togglePanel("account-panel"));

function togglePanel(panelId) {
  const panel = document.getElementById(panelId);
  panel.style.display = panel.style.display === "flex" ? "none" : "flex";
}

// Установка начального активного таба
function initializeActiveTab() {
  const savedButtonId = sessionStorage.getItem("active_btn_id");
  const defaultButtonId = DOCUMENT_TYPES.INGOING.defaultBtnId;
  const initialButton = document.getElementById(savedButtonId || defaultButtonId);
  changeActiveTab(initialButton);
}

// Инициализация табов
menuButtons.forEach(btn => {
  btn.addEventListener("click", () => changeActiveTab(btn));
});

initializeActiveTab();

function changeActiveTab(btn) {
  // Получаем предыдущие активные элементы
  CloseActiveDoc();
  const prevActiveTab = document.querySelector(".main__tabs.main__tabs--active");
  const prevActiveBtn = document.querySelector(".header__menu--btn.menu__btn--active");
  const activeTabId = `#${btn.dataset.tab}`;

  if (prevActiveTab !== null) {
    const prevActiveTabId = `#${prevActiveTab.id}`;
    if (prevActiveTabId == activeTabId) return;
  }

  // Удаляем активные классы
  prevActiveBtn?.classList.remove("menu__btn--active");
  prevActiveTab?.classList.remove("main__tabs--active");

  // Устанавливаем новые активные элементы
  sessionStorage.setItem("active_btn_id", btn.id);
  const activeTab = document.querySelector(activeTabId);

  btn.classList.add("menu__btn--active");
  activeTab.classList.add("main__tabs--active");

  // Обновляем документы в зависимости от активного таба
  updateDocumentsForTab(activeTabId);
}

// Получает документы и отображает их
async function updateDocumentsForTab(tabId) {
  // Находим соответствующий тип документа
  const docTypeConfig = Object.values(DOCUMENT_TYPES).find(
    type => type.tabId === tabId
  );

  const docTableConfig = Object.values(DOCUMENT_TYPES).find(
    table => table.tabId === tabId
  );

  // Получаем документы
  const documents = await FetchGetDocuments(docTableConfig.table, { type: docTypeConfig.type });
  if (documents != null) {
    // Обновляем соответствующий контейнер
    const container = document.getElementById(docTypeConfig.documentTableId);
    container.innerHTML = WriteDocumentsInTable(documents, docTypeConfig.type)
    setupDocumentTables()
  }
}

