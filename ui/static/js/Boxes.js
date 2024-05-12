function getBoxes() {
    fetch('/box/select')
    .then(response => response.json())
    .then(data => {
        var resultDiv = document.getElementById('result');
        resultDiv.innerHTML = '<h2>Список боксов</h2>';
        var tableHTML = '<table><tr><th>ID</th><th>статус</th><th>Номер</th><th>Этаж</th><th>Площадь, м<sup>2</sup></th><th>Последний договор</th><th>Действия</th></tr>';

        data.forEach(box => {
            tableHTML += `<tr><td>${Box.Status}</td><td>${Box.BoxID}</td><td>${Box.Floor}</td><td>${Box.Area}</td><td>${Box.Contract_id}</td><td>${Box.Contract_start}</td><td>${Box.Contract_end}</td></tr>`;
        });

        tableHTML += '</table>';
        resultDiv.innerHTML += tableHTML;
    })
    .catch(error => {
        console.error('Произошла ошибка при получении списка боксов:', error);
    });
}

function sendBox() {
    var status = document.getElementById('status').value;
    var floor = document.getElementById('floor').value;
    var area = document.getElementById('area').value;

    var data = {
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
        document.getElementById('result').innerText = result;
    })
    .then(data => {
        getBoxes();
    })
    .catch(error => {
        console.error('Произошла ошибка при отправке данных:', error);
    });
}
