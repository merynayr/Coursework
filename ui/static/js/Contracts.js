function getContracts() {
    fetch('/contract/select')
        .then(response => response.json())
        .then(data => {
            var resultDiv = document.getElementById('Contracts');
            resultDiv.innerHTML = '<h2>Список договоров</h2>';
            var tableHTML = '<table><tr><th>№\nДоговора</th><th>Клиент</th><th>Боксы</th><th>Дата заключения</th><th>Начало аренды</th><th>Конец аренды</th></tr>';

            data?.forEach(contract => {
                tableHTML += `<tr><td>${contract.ContractID ?? ''}</td><td>${contract.ClientName ?? ''}</td><td>${contract.BoxID ?? ''}</td><td>${formatDate(contract.DateSigned) ?? ''}</td><td>${formatDate(contract.StartDate) ?? ""}</td><td>${formatDate(contract.EndDate) ?? ''}</td></tr>`;
            });

            tableHTML += '</table>';
            resultDiv.innerHTML += tableHTML;
        })
        .catch(error => {
            console.error('Произошла ошибка при получении списка договоров:', error);
        });
}


function addContracts() {
    var ContractID = parseInt(document.getElementById('ContractID').value);
    var ClientID = parseInt(document.getElementById('ClientID').value);
    var BoxID = parseInt(document.getElementById('BoxID').value);
    var DateSigned = document.getElementById('DateSigned').value;
    var StartDate = document.getElementById('StartDate').value;
    var EndDate = document.getElementById('EndDate').value;

    var data = {
        ContractID: ContractID,
        ClientID: ClientID,
        BoxID: BoxID,
        DateSigned: DateSigned,
        StartDate: StartDate,
        EndDate: EndDate
    };
    console.log(data);
    fetch('/contract/add', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => response.text())
        .then(result => {
            document.getElementById('Contracts').innerText = result;
        })
        .then(data => {
            getContracts();
        })
        .catch(error => {
            console.error('Произошла ошибка при отправке данных:', error);
        });
}
