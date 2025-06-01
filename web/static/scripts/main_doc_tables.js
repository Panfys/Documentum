function WriteDocumentsInTable(documents, type) {

    let documentsString = '';

    if (type == "directives") {
        if (documents) {

            documents.forEach(document => {
                if (!document) return "";

                documentsString += `
                <table class="tubs__table tubs__table--document" id="document-table-${document.id}" document-id='${document.id}'>
                    <tr>
                        <td>${document.number}</td>
                        <td class="table__column--number">${document.date}</td>
                        <td colspan="2" id="table__column--name">${document.name}</td>
                        <td>${document.autor}</td>
                        <td>${document.numCoverLetter}<br>${document.dateCoverLetter}</td>
                        <td>${document.countCopy}</td>
                        <td>${document.sender}</td>
                        <td>${document.numSendLetter}<br>${document.dateSendLetter}</td>
                        <td>${document.countSendCopy}</td>
                        <td class="table__column--familiar">${prepareFamiliars(document.id, document.familiars)}</td>
                        <td class="table__column--location">${document.location}</td>
                        <td class="table__column--button">
                            <button class="table__btn--opendoc" file="${document.fileURL}"></button>
                        </td>
                    </tr>
                </table>`;
            });

            return documentsString;
        }
        return
    } else if (type == "inventory") {
        if (documents) {

            documents.forEach(document => {
                if (!document) return "";

                documentsString += `
                <table class="tubs__table tubs__table--document" id="document-table-${document.id}" document-id='${document.id}'>
                   <tr>
                        <td class="table__column--number">${document.number}</td>
                        <td colspan="2">${document.numCoverLetter}<br>${document.dateCoverLetter}</td>
                        <td colspan="3" id="table__column--name">${document.name}</td>
                        <td colspan="2">${document.sender}</td>
                        <td>${document.countCopy}</td>
                        <td>${document.copy}</td>
                        <td colspan="2">${document.addressee}</td>
                        <td colspan="2">${document.numSendLetter}<br>${document.dateSendLetter}</td>
                        <td>${document.sendCopy}</td>
                        <td class="table__column--familiar">${prepareFamiliars(document.id, document.familiars)}</td>
                        <td class="table__column--location">${document.location}</td>
                        <td class="table__column--button">
                            <button class="table__btn--opendoc" file="${document.fileURL}"></button>
                        </td>
                    </tr>
                </table>`;
            });

            return documentsString;
        }
        return
    } else if (type == "ingoing") {
        if (documents) {

            documents.forEach(document => {
                if (!document) return "";

                documentsString += `
        <table class='tubs__table tubs__table--document' id='document-table-${document.id}' document-id='${document.id}'>
            <tr>
                <td class='table__column--number'>${document.fnum}<br>${document.fdate}</td>
                <td class='table__column--number'>${document.lnum}<br>${document.ldate}</td>
                <td class='table__column--name'>${document.name}</td>
                <td class='table__column--sender'>${document.sender}</td>
                <td class='table__column--ispolnitel'>${document.ispolnitel}</td>
                <td class='table__column--result'>${document.result}</td>
                <td class="table__column--familiar">${prepareFamiliars(document.id, document.familiars)}</td>
                <td class='table__column--count'>${document.count}</td>
                <td class='table__column--copy'>${document.copy}</td>
                <td class='table__column--width'>${document.width}</td>
                <td class='table__column--location'>${document.location}</td>
                <td class='table__column--button'>
                  <button class='table__btn--opendoc' file="${document.file}"></button>
                </td>
            </tr>
        </table>`;

                documentsString += `<div class='table__resolution-panel' id='resolution-panel-${document.id}'>`;

                if (document.resolutions && document.resolutions.length) {
                    document.resolutions.forEach(resolution => {
                        if (!resolution) return;

                        documentsString += `
                <div class='table__resolution' id='ingoing-resolution'>
                    <div class='table__resolution--ispolnitel'>${resolution.ispolnitel}</div>
                    <div class='table__resolution--text'>&#171;${resolution.text}&#187;</div>
                    <div class='table__resolution--time'>${resolution.deadline || resolution.result}</div>
                    <div class='table__resolution--user'>${resolution.user}</div>
                    <div class='table__resolution--date'>${resolution.date}</div>
                </div>`;
                    });
                }
                documentsString += '</div>';
            });

            return documentsString;
        }
    } else if (type == "outgoing"){
        if (documents) {

            documents.forEach(document => {
                if (!document) return "";

                documentsString += `
        <table class='tubs__table tubs__table--document' id='document-table-${document.id}' document-id='${document.id}'>
            <tr>
                <td class='table__column--number'>${document.fnum}<br>${document.fdate}</td>
                <td class='table__column--number'>${document.lnum}<br>${document.ldate}</td>
                <td class='table__column--name'>${document.name}</td>
                <td class='table__column--sender'>${document.sender}</td>
                <td class='table__column--ispolnitel'>${document.ispolnitel}</td>
                <td class='table__column--result'>${document.result}</td>
                <td class="table__column--familiar">${prepareFamiliars(document.id, document.familiars)}</td>
                <td class='table__column--count'>${document.count}</td>
                <td class='table__column--copy'>${document.copy}</td>
                <td class='table__column--width'>${document.width}</td>
                <td class='table__column--location'>${document.location}</td>
                <td class='table__column--button'>
                  <button class='table__btn--opendoc' file="${document.file}"></button>
                </td>
            </tr>
        </table>`;
            });

            return documentsString;
        }
    } else return ""
}

function prepareFamiliars(id, familiars) {
    const familiarsList = familiars?.length
                    ? `<ul class="familiar-list" id="familiar-list-${id}">
                ${familiars.map(f => `<li>${f}</li>`).join('')}
                    </ul>`
                    : `<ul class="familiar-list" id="familiar-list-${id}"></ul>`;
    return familiarsList
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

// При нажатии на документ выделяет и открывает панель документа

// Объявляем глобальные переменные (если они действительно нужны)
let resolution_id = 0;
let newresolution_id = 0;

function ViewDocumentTable(doc, event) {
    // Получаем элементы DOM с проверкой на null
    const active_tub = document.querySelector(".main__tabs--active");
    if (!active_tub) return;

    const pre_active_doc = document.querySelector(".tubs__table--active-table");
    const docpanel = active_tub.querySelector("#title-docpanel");
    const btn_search = active_tub.querySelector("#btn-search");
    const btn_newdoc = active_tub.querySelector("#btn-newdoc");
    const tubs_folder = active_tub.querySelector(".tubs__folder");
    const tubs_title_span = active_tub.querySelector(".tubs__title--span");

    // Обработка клика на кнопку открытия документа
    if (event.target.classList.contains("table__btn--opendoc")) {
        openDocument(
            event.target.getAttribute("file"),
            doc.getAttribute("document-id")
        );
        if (pre_active_doc && pre_active_doc !== doc) {
            CloseActiveDoc();
        }
        return;
    } else {
        CloseActiveDoc();
    }

    // Проверка повторного клика на тот же документ
    if (pre_active_doc === doc) return;

    // Активация нового документа
    doc.classList.add("tubs__table--active-table");
    const doc_id = doc.getAttribute("document-id");
    const resolution_panel = document.getElementById("resolution-panel-" + doc_id);

    // Установка названия документа
    const nameElement = doc.querySelector(".table__column--name, #table__column--name");
    if (nameElement && tubs_folder) {
        tubs_folder.innerHTML = nameElement.innerHTML;
    }

    // Обновление UI
    if (tubs_title_span) tubs_title_span.style.display = "flex";
    if (btn_search) btn_search.style.display = "none";
    if (btn_newdoc) btn_newdoc.style.display = "none";
    if (docpanel) docpanel.style.display = "flex";

    // Открытие панели резолюций
    if (resolution_panel) {
        resolution_panel.style.display = "flex";
        resolution_id = resolution_panel.childElementCount;
        newresolution_id = resolution_panel.childElementCount;
    }
}

// Функция закрытия активного документа
function CloseActiveDoc() {
    const active_tub = document.querySelector(".main__tabs--active");
    if (!active_tub) return;

    const pre_active_doc = document.querySelector(".tubs__table--active-table");
    if (!pre_active_doc) return;

    const doc_id = pre_active_doc.getAttribute("document-id");
    const resolution_panel = document.getElementById("resolution-panel-" + doc_id);

    // Получаем UI элементы, если active_tub доступен
    let tubs_folder, tubs_title_span, btn_search, btn_newdoc, docpanel;
    if (active_tub) {
        tubs_folder = active_tub.querySelector(".tubs__folder");
        tubs_title_span = active_tub.querySelector(".tubs__title--span");
        btn_search = active_tub.querySelector("#btn-search");
        btn_newdoc = active_tub.querySelector("#btn-newdoc");
        docpanel = active_tub.querySelector("#title-docpanel");
    }

    // Обработка новой таблицы документа
    if (pre_active_doc.id === "table-newdoc") {
        toggleNewDocumentForm();
    }
    // Обработка обычной таблицы документа
    else {
        // Закрытие панели резолюций
        if (resolution_panel) {
            handleActiveDocumentResolution("removeAll");
            AddDocResolution("back");
            resolution_panel.style.display = "none";
            if (resolution_id !== newresolution_id) {
                resolution_panel.removeChild(resolution_panel.lastChild);
                newresolution_id -= 1;
            }
        }

        pre_active_doc.classList.remove("tubs__table--active-table");

        // Сброс UI элементов
        if (tubs_folder) tubs_folder.innerHTML = "";
        if (tubs_title_span) tubs_title_span.style.display = "none";
        if (btn_search) btn_search.style.display = "flex";
        if (btn_newdoc) btn_newdoc.style.display = "flex";
        if (docpanel) docpanel.style.display = "none";
    }
}