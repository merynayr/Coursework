function getBoxes() {
    fetch('/box/select')
        .then(response => response.json())
        .then(data => {
            var resultDiv = document.getElementById('Boxes');
            resultDiv.innerHTML = '<h2>Список боксов</h2>';
            var tableHTML = '<table><tr><th>Статус</th><th>Номер</th><th>Этаж</th><th>Площадь, м<sup>2</sup></th><th>Последний договор</th><th>Действия</th></tr>';

            data?.forEach(box => {
                tableHTML += `<tr><td>${box.Status ?? ''}</td><td>${box.BoxID ?? ''}</td><td>${box.Floor ?? ''}</td><td>${box.Area ?? ''}</td><td>${box.Contract_id ?? ""}</td><td>${box.Contract_start ?? ''}</td><td>${box.Contract_end ?? ''}</td></tr>`;
            });

            tableHTML += '</table>';
            resultDiv.innerHTML += tableHTML;
        })
        .catch(error => {
            console.error('Произошла ошибка при получении списка боксов:', error);
        });
}

function addBoxes() {
    var BoxID = parseInt(document.getElementById('boxID').value);
    var status = document.getElementById('status').value;
    var floor = parseInt(document.getElementById('floor').value);
    var area = parseFloat(document.getElementById('area').value);

    var data = {
        box_id: BoxID,
        status: status,
        floor: floor,
        area: area
    };

    fetch('/box/add', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => response.text())
        .then(result => {
            document.getElementById('Boxes').innerText = result;
        })
        .then(data => {
            getBoxes();
        })
        .catch(error => {
            console.error('Произошла ошибка при отправке данных:', error);
        });
}
