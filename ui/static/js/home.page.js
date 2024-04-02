function getClients() {
    fetch('/client/select')
        .then(response => response.json())
        .then(data => {
            var resultDiv = document.getElementById('result');
            resultDiv.innerHTML = '<h2>Список клиентов</h2>';
            var tableHTML = '<table><tr><th>ID</th><th>Имя</th><th>Тип</th><th>Телефон</th><th>Действия</th></tr>';

            data.forEach(client => {
                tableHTML += `<tr><td>${client.ClientID}</td><td>${client.Name}</td><td>${client.Type}</td><td>${client.Phone}</td><td><button onclick="editClient(${client.ClientID})">Редактировать</button></td><td><button onclick="deleteClient(${client.ClientID})">Удалить</button></td></tr>`;
            });

            tableHTML += '</table>';
            resultDiv.innerHTML += tableHTML;
        })
        .catch(error => {
            console.error('Произошла ошибка при получении списка клиентов:', error);
        });
}

function sendData() {
    var clientID = parseInt(document.getElementById('clientID').value);
    var name = document.getElementById('name').value;
    var type = document.getElementById('type').value;
    var phone = document.getElementById('phone').value;

    var data = {
        clientID: clientID,
        name: name,
        type: type,
        phone: phone
    };

    fetch('/сlient/add', {
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
        getClients();
    })
    .catch(error => {
        console.error('Произошла ошибка при отправке данных:', error);
    });
}

function deleteClient(id) {
    fetch('/client/del', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ clientID: id })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to delete client');
        }
        return response.json();
    })
    .then(data => {
        getClients();
    })
    .catch(error => {
        console.error('Произошла ошибка при удалении клиента:', error);
    });
}



// function getClients() {
//     fetch('/client/select')
//         .then(response => response.json())
//         .then(data => {
//             var table = document.getElementById('clientTable');
//             table.innerHTML = ''; // Очищаем таблицу перед добавлением новых клиентов

//             data.forEach(client => {
//                 var row = table.insertRow();
//                 row.id = 'row_' + client.ClientID;
//                 row.insertCell(0).innerText = client.ClientID;
//                 row.insertCell(1).innerText = client.Name;
//                 row.insertCell(2).innerText = client.Type;
//                 row.insertCell(3).innerText = client.Phone;
//                 var editButton = document.createElement('button');
//                 editButton.innerText = 'Редактировать';
//                 editButton.onclick = function() {
//                     editClient(client.ClientID);
//                 };
//                 row.insertCell(4).appendChild(editButton);
//                 var deleteButton = document.createElement('button');
//                 deleteButton.innerText = 'Удалить';
//                 deleteButton.onclick = function() {
//                     deleteClient(client.ClientID);
//                 };
//                 row.insertCell(5).appendChild(deleteButton);
//             });
//         })
//         .catch(error => {
//             console.error('Произошла ошибка при получении списка клиентов:', error);
//         });
// }

function editClient(id) {
    var row = document.getElementById('row_' + id);
    var cells = row.getElementsByTagName('td');
    cells[1].contentEditable = true; // Разрешаем редактирование ячейки с именем клиента
    cells[2].contentEditable = true; // Разрешаем редактирование ячейки с типом клиента
    cells[3].contentEditable = true; // Разрешаем редактирование ячейки с телефоном клиента
    var editButton = row.getElementsByTagName('button')[0];
    editButton.innerText = 'Сохранить';
    editButton.onclick = function() {
        updateClient(id);
    };
}

function updateClient(id) {
    var row = document.getElementById('row_' + id);
    var cells = row.getElementsByTagName('td');
    var newName = cells[1].innerText;
    var newType = cells[2].innerText;
    var newPhone = cells[3].innerText;
    var data = {
        clientID: id,
        name: newName,
        type: newType,
        phone: newPhone
    };
    fetch('/client/update', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to update client');
        }
        return response.json();
    })
    .then(data => {
        // Обновляем список клиентов после успешного обновления
        getClients();
        // Возвращаем кнопку редактирования в исходное состояние
        var editButton = row.getElementsByTagName('button')[0];
        editButton.innerText = 'Редактировать';
        editButton.onclick = function() {
            editClient(id);
        };
    })
    .catch(error => {
        console.error('Произошла ошибка при обновлении клиента:', error);
    });
}
