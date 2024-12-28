const loginForm = document.getElementById('registerForm');
const errorMessage = document.getElementById('error-message');

loginForm.addEventListener('submit', (event) => {
    event.preventDefault();

    const last_name = document.getElementById('last-name').value;
    const first_name = document.getElementById('first-name').value;
    const middle_name = document.getElementById('middle-name').value;
    const telephone = document.getElementById('telephone').value;
    const login = document.getElementById('login').value;
    const password = document.getElementById('password').value;
    const status = document.getElementById('statusiZak').value

    const data = {
        login,
        password,
        last_name,
        first_name,
        middle_name,
        telephone,
        status
    };

    const jsonData = JSON.stringify(data);

    fetch('api/register/user', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: jsonData
    })
        .then(response => response.json())
        .then(data => {
            if (data.ok) {
                window.location.href = '/';
            } else {
                errorMessage.textContent = 'Данный email уже используется';
            }
        })
        .catch(error => {
            console.error('Ошибка:', error);
            errorMessage.textContent = 'Произошла ошибка при отправке запроса';
        });
});

function GetStatusi() {
    fetch('/api/get/statusi-zakazchikov', {
        method: 'GET'
    })
        .then(responce => responce.json())
        .then(data => {
            const statusiZakBlock = document.getElementById('statusiZak')
            statusiZakBlock.innerHTML = ''

            const option = document.createElement('option');
            option.innerHTML = '';
            option.value = "0";
            option.textContent = 'Выбрать';
            option.selected = true;
            statusiZakBlock.appendChild(option);

            data.forEach(statusZak => {
                const option = document.createElement('option')
                option.value = statusZak.id_statusa;
                option.textContent = statusZak.naimenovanie_statusa

                statusiZakBlock.appendChild(option)
            })
        })
}

GetStatusi();