//-------Переключение табов документов
const btn_menu = document.querySelectorAll(".header__menu--btn");

//-----Открытие боковых панелей------------
document.querySelector(".main__header--settings").onclick = () =>
    OpenMainSide("settings-panel");
document.querySelector(".main__header--account").onclick = () =>
    OpenMainSide("account-panel");

function OpenMainSide(panel_id) {
    const panel = document.getElementById(panel_id);
    if (panel.style.display == "flex") {
        panel.style = "display: none";
    } else {
        panel.style = "display: flex";
    }
}

// Проходимся по всем кнопкам
btn_menu.forEach((btn) => {
    btn.addEventListener("click", () => ChengeActiveTab(btn));
});

// Set initial active tab
let button;
if (sessionStorage.getItem("active_btn_id")) {
    button = document.getElementById(sessionStorage.getItem("active_btn_id"));
} else {
    button = document.getElementById("menu-btn-ingoing");
}
ChengeActiveTab(button);

function ChengeActiveTab(btn) {
    // Получаем предыдущую активную кнопку
    const pre_active_tub = document.querySelector(".main__tabs.main__tabs--active");
    const pre_active_btn = document.querySelector(".header__menu--btn.menu__btn--active");

    // Удаляем активные классы
    if (pre_active_btn) pre_active_btn.classList.remove("menu__btn--active");
    if (pre_active_tub) pre_active_tub.classList.remove("main__tabs--active");

    // Получаем новую активную вкладку
    const active_tub_id = "#" + btn.getAttribute("data-tab");
    sessionStorage.setItem("active_btn_id", btn.getAttribute("id"));
    const active_tub = document.querySelector(active_tub_id);

    // Добавляем активные классы
    btn.classList.add("menu__btn--active");
    active_tub.classList.add("main__tabs--active");

    // Обновляем вкладку с документами
    switch (active_tub_id) {
        case "#main-tab-ingoing":
            GetDocuments("ASC", "Входящий", "id", "2000-01-01", "3000-01-01");
            break;
        case "#main-tab-outgoing":
            GetDocuments("ASC", "Исходящий", "id", "2000-01-01", "3000-01-01");
            break;
        case "#main-tab-directive":
            GetDocuments("ASC", "Приказ", "id", "2000-01-01", "3000-01-01");
            break;
        case "#main-tab-inventory":
            GetDocuments("ASC", "Издание", "id", "2000-01-01", "3000-01-01");
            break;
    }
}

//-------Открытие панели документа
function CheckDocumentTable() {
    // получаем все таблицы документов
    doc_table = document.querySelectorAll(".tubs__table--document");
    // Проходимся по всем таблицам
    doc_table.forEach((doc) => {
        // вешаем на каждую таблицу обработчик события клик
        doc.addEventListener("click", () => ViewDocumentTable(doc, event));
    });
}
