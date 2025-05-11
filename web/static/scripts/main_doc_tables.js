function WriteDocumentsInTable(documents, container) {

    container.innerHTML = '';
    let documentsString = '';

    if (documents) {

        documents.forEach(document => {
            if (!document) return;

            documentsString += `
        <table class='tubs__table tubs__table--document' id='document-table-${document.id}' document-id='${document.id}'>
            <tr>
                <td class='table__column--number'>${document.fnum}<br>${document.fdate}</td>
                <td class='table__column--number'>${document.lnum}<br>${document.ldate}</td>
                <td class='table__column--name'>${document.name}</td>
                <td class='table__column--sender'>${document.sender}</td>
                <td class='table__column--ispolnitel'>${document.ispolnitel}</td>
                <td class='table__column--result'>${document.result}</td>
                <td class='table__column--familiar'>${document.familiar}</td>
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
                    <div class='table__resolution--time'>${resolution.deadline}</div>
                    <div class='table__resolution--user'>${resolution.user}</div>
                    <div class='table__resolution--date'>${resolution.date}</div>
                </div>`;
                });
            }
            documentsString += '</div>';
        });

        container.innerHTML = documentsString;
        setupDocumentTables()
    }
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

function ViewDocumentTable(doc, event) {
    // Получаем элементы DOM
    const active_tub = document.querySelector(".main__tabs--active");
    const pre_active_doc = document.querySelector(".tubs__table--active-table");
    const docpanel = active_tub.querySelector("#title-docpanel");
    const btn_search = active_tub.querySelector("#btn-search");
    const btn_newdoc = active_tub.querySelector("#btn-newdoc");
    const tubs_folder = active_tub.querySelector(".tubs__folder");
    const tubs_title_span = active_tub.querySelector(".tubs__title--span");

    // Функция закрытия активного документа
    function CloseActiveDoc() {
        if (!pre_active_doc) return;

        // Обработка новой таблицы документа
        if (pre_active_doc.getAttribute("id") === "table-newdoc") {
            toggleNewDocumentForm();
        }
        // Обработка обычной таблицы документа
        else {
            if (docpanel.style.display === "none") {
                AddDocResolution("back");
            }

            resolutionStartCount = null;
            pre_active_doc.classList.remove("tubs__table--active-table");
            const pre_active_doc_id = pre_active_doc.getAttribute("document-id");
            const pre_active_res = document.getElementById("resolution-panel-" + pre_active_doc_id);

            // Сброс UI элементов
            tubs_folder.innerHTML = "";
            tubs_title_span.style.display = "none";
            btn_search.style.display = "flex";
            btn_newdoc.style.display = "flex";
            docpanel.style.display = "none";

            // Закрытие панели резолюций
            if (resolution_panel) {
                resolution_panel.style.display = "none";
                if (resolution_id !== newresolution_id) {
                    resolution_panel.removeChild(resolution_panel.lastChild);
                    newresolution_id -= 1;
                }
            }
        }
    }

    // Обработка клика на кнопку открытия документа
    if (event.target.classList.contains("table__btn--opendoc")) {
        openDocument(
            event.target.getAttribute("file"),
            doc.getAttribute("document-id")
        );
        if (pre_active_doc !== doc) CloseActiveDoc();
        return;
    } else {
        CloseActiveDoc();
    }

    // Проверка повторного клика на тот же документ
    if (pre_active_doc === doc) return;

    // Активация нового документа
    doc.classList.add("tubs__table--active-table");
    const doc_id = doc.getAttribute("document-id");
    resolution_panel = document.getElementById("resolution-panel-" + doc_id);

    // Установка названия документа
    const nameElement = doc.querySelector(".table__column--name") || doc.querySelector("#table__column--name");
    if (nameElement) {
        tubs_folder.innerHTML = nameElement.innerHTML;
    }

    // Обновление UI
    tubs_title_span.style.display = "flex";
    btn_search.style.display = "none";
    btn_newdoc.style.display = "none";
    docpanel.style.display = "flex";

    // Открытие панели резолюций
    if (resolution_panel) {
        resolution_panel.style.display = "flex";
        resolution_id = resolution_panel.childElementCount;
        newresolution_id = resolution_panel.childElementCount;
    }
}