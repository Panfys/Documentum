async function AddDocument(data) {
    $.ajax({
        method: "POST",
        url: "/document",
        data: data,
        cache: false,
        contentType: false,
        processData: false,
        success: function () {
            alert("Документ успешно загружен, перезагрузите страницу!");
        },
        error: function (jqXHR) {
            serverMessage("show", jqXHR.responseText);
        },
    });
}

//------Вывод документов---------
function GetDocuments(set, type, col, datain, datato) {
    $.ajax({
        method: "GET",
        url: "/documents",
        data: {
            type: type,
            col: col,
            set: set,
            datain: datain,
            datato: datato,
        },
        success: function (documents) {
            switch (type) {
                case "Входящий":
                    document.getElementById("ingoing-documents-container").innerHTML =
                        documents;
                    break;
                case "Исходящий":
                    document.getElementById("outgoing-documents-container").innerHTML =
                        documents;
                    break;
                case "Приказ":
                    document.getElementById("directive-documents-container").innerHTML =
                        documents;
                    break;
                case "Издание":
                    document.getElementById("inventory-documents-container").innerHTML =
                        documents;
                    break;
            }
            CheckDocumentTable();
        },
        error: () =>
            serverMessage("show", "Возникла ошибка на сервере, попробуйте позже!"),
    });
}

function AddViewDocument(id) {
    $.ajax({
        method: "PATCH",
        url: `/document/${encodeURIComponent(id)}/view`,
        error: function (jqXHR, exception) {
            serverMessage("show", jqXHR.responseText);
        },
    });
}

function ChangeDocument (data) {
    $.ajax({
        method: "POST",
        url: "../router.php",
        data: data,
        cache: false,
        contentType: false,
        processData: false,
        success: function (check) {
          if (check == "OK") {
            alert(check);
          } else serverMessage("show", check);
        },
        error: () =>
          serverMessage("show", "Возникла ошибка на сервере, попробуйте позже!"),
      });
}