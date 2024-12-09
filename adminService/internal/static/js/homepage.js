// Функция для загрузки заявок с сервера и обновления списка
function fetchApplications() {
    fetch('/api/applications', {
        method: 'GET'
    })
        .then(response => response.json())
        .then(data => {
            const applicationsSection = document.getElementById('applications');
            applicationsSection.innerHTML = '<h3>Непринятые заявки</h3>'; // Очищаем существующий список

            if (data.length === 0) {
                applicationsSection.innerHTML += '<p>На данный момент нет новых заявок.</p>';
            } else {
                data.forEach(application => {
                    const applicationElement = document.createElement('div');
                    applicationElement.classList.add('application', 'card', 'mb-3', 'p-3');
                    applicationElement.innerHTML = `
                        <h3>Заявка №${application.id_zayavki}</h3>
                        <p>Мероприятие: ${application.vid_prazdnika}</p>
                        <p>Дата: ${new Date(application.data_provedeniya).toLocaleDateString()}</p>
                        <button class="show-modal btn btn-info" data-id="${application.id_zayavki}">Подробнее</button>
                    `;
                    applicationsSection.appendChild(applicationElement);
                });

                // Добавляем обработчик для кнопок "Подробнее"
                const showModalButtons = document.querySelectorAll('.show-modal');
                showModalButtons.forEach(button => {
                    button.addEventListener('click', function () {
                        const applicationId = button.getAttribute('data-id');
                        const application = data.find(app => app.id_zayavki == applicationId);
                        showApplicationDetails(application);
                    });
                });
            }
        })
        .catch(error => {
            console.error('Ошибка при загрузке заявок:', error);
        });
}

// Функция для отображения деталей заявки в модальном окне
function showApplicationDetails(application) {
    // Заполняем данные заявки в модальном окне
    document.getElementById('familiya_zakazchika').textContent = application.familiya_zakazchika;
    document.getElementById('imya_zakazchika').textContent = application.imya_zakazchika;
    document.getElementById('otchestvo_zakazchika').textContent = application.otchestvo_zakazchika;
    document.getElementById('stasus_zakazchika').textContent = application.stasus_zakazchika;
    document.getElementById('telephone_zakazchika').textContent = application.telephone_zakazchika;
    document.getElementById('email_zakazchika').textContent = application.email_zakazchika;
    document.getElementById('vid_prazdnika').textContent = application.vid_prazdnika;
    document.getElementById('data_provedeniya').textContent = new Date(application.data_provedeniya).toLocaleDateString();
    document.getElementById('kolichestvo_chelovek').textContent = application.kolichestvo_chelovek;
    document.getElementById('nachalo_provedeniya').textContent = application.nachalo_provedeniya;
    document.getElementById('konec_provedeniya').textContent = application.konec_provedeniya;
    document.getElementById('prodoljitelnost').textContent = `${application.prodoljitelnost} часов`;

    // Устанавливаем ID заявки для кнопки
    document.getElementById('takeApplicationBtn').setAttribute('data-id', application.id_zayavki);

    // Открываем модальное окно
    const modal = new bootstrap.Modal(document.getElementById('applicationModal'));
    modal.show();
}

// Функция для отправки заявки в работу
function takeApplication(applicationId) {
    fetch(`/api/application/accept?app=${applicationId}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => {
            if (response.ok) {
                alert('Заявка успешно взята в работу!');
                fetchApplications(); // Обновляем список заявок
                const modal = bootstrap.Modal.getInstance(document.getElementById('applicationModal'));
                modal.hide();
            } else {
                throw new Error('Не удалось взять заявку в работу');
            }
        })
        .catch(error => {
            console.error('Ошибка:', error);
            alert('Произошла ошибка при взятии заявки в работу.');
        });
}

// Добавление обработчика кнопки "Взять заявку в работу"
document.getElementById('takeApplicationBtn').addEventListener('click', function () {
    const applicationId = this.getAttribute('data-id'); // ID заявки
    if (applicationId) {
        takeApplication(applicationId);
    } else {
        console.error('ID заявки не найден.');
    }
});

// Функция выхода
function logout() {
    document.cookie = 'token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
    window.location.href = '/login';
}

// Запуск начальной загрузки заявок
fetchApplications();
