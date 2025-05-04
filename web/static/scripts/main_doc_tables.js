function WriteDocuments(documents, container) {

    container.innerHTML = '';
    let documentsString = '';

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

        // HTML для резолюций
        if (document.resolutions && document.resolutions.length) {
            documentsString += `<div class='table__resolution-panel' id='resolution-panel-${document.id}'>`;

            document.resolutions.forEach(resolution => {
                if (!resolution) return;

                documentsString += `
                <div class='table__resolution' id='ingoing-resolution'>
                    <div class='table__resolution--ispolnitel'>${resolution.ispolnitel}</div>
                    <div class='table__resolution--text'>&#171;${resolution.text}&#187;</div>
                    <div class='table__resolution--time'>${resolution.time}</div>
                    <div class='table__resolution--user'>${resolution.user}</div>
                    <div class='table__resolution--date'>${resolution.date}</div>
                </div>`;
            });

            documentsString += '</div>';
        }
    });

    container.innerHTML = documentsString;
}