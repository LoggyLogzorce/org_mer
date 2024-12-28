const loginForm = document.getElementById('loginForm');
const errorMessage = document.getElementById('error-message');

loginForm.addEventListener('submit', (event) => {
    event.preventDefault();

    const login = document.getElementById('login').value;
    const password = document.getElementById('password').value;

    if (!login || !password) {
        alert("Заполните все поля.")
        return
    }

    const data = {
        login,
        password
    };

    const jsonData = JSON.stringify(data);

    fetch('/api/auth/ruk', {
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
                errorMessage.textContent = 'Неверный логин или пароль';
            }
        })
        .catch(error => {
            console.error('Ошибка:', error);
            errorMessage.textContent = 'Произошла ошибка при отправке запроса';
        });
});