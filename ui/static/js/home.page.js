function getClients() {
    // Отправляем GET запрос на сервер для получения списка клиентов
    fetch('/clients')
        .then(response => response.json())
        .then(data => {
            // Создаем таблицу
            var resultDiv = document.getElementById('result');
            resultDiv.innerHTML = '<h2>Список клиентов</h2>';
            var tableHTML = '<table><tr><th>ID</th><th>Имя</th><th>Тип</th><th>Телефон</th></tr>';

            // Заполняем таблицу данными о клиентах
            data.forEach(client => {
                tableHTML += `<tr><td>${client.ClientID}</td><td>${client.Name}</td><td>${client.Type}</td><td>${client.Phone}</td></tr>`;
            });

            tableHTML += '</table>';
            resultDiv.innerHTML += tableHTML;
        })
        .catch(error => {
            console.error('Произошла ошибка при получении списка клиентов:', error);
        });
}

function sendData() {
    // Получаем данные из формы
    var clientID = parseInt(document.getElementById('clientID').value);
    var name = document.getElementById('name').value;
    var type = document.getElementById('type').value;
    var phone = document.getElementById('phone').value;

    // Создаем объект с данными для отправки на сервер
    var data = {
        clientID: clientID,
        name: name,
        type: type,
        phone: phone
    };

    // Отправляем данные на сервер
    fetch('/client', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.text())
    .then(result => {
        // Отображаем результат отправки данных
        document.getElementById('result').innerText = result;
    })
    .catch(error => {
        console.error('Произошла ошибка при отправке данных:', error);
    });
}