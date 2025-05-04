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

function ChangeDocument(data) {
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