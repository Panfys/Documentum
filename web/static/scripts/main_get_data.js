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