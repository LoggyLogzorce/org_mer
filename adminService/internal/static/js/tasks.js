// Функция для загрузки заявок с сервера и обновления списка
function fetchApplications() {
    fetch('/api/applications?search=my_tasks', {
        method: 'GET'
    })
        .then(response => response.json())
        .then(data => {
            const applicationsList = document.getElementById('applicationsList');
            applicationsList.innerHTML = ''; // Очищаем существующий список

            if (data.length === 0) {
                applicationsList.innerHTML = '<p>На данный момент вы не работаете над заявками.</p>';
            } else {
                data.forEach(application => {
                    const applicationElement = document.createElement('div');
                    applicationElement.classList.add('col-12', 'col-md-6', 'col-lg-4'); // Сетка Bootstrap
                    applicationElement.innerHTML = `
                        <div class="application card p-3 h-100">
                            <h5>Заявка №${application.id_zayavki}</h5>
                            <p><strong>Мероприятие:</strong> ${application.VidiPrazdnikov.naimenovanie_vida}</p>
                            <p><strong>Дата:</strong> ${new Date(application.data_provedeniya).toLocaleDateString()}</p>
                            <button class="show-modal btn btn-info mt-auto" data-id="${application.id_zayavki}">Подробнее</button>
                        </div>
                    `;
                    applicationsList.appendChild(applicationElement);
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
    document.getElementById('familiya_zakazchika').textContent = application.Zakazchik.familiya;
    document.getElementById('imya_zakazchika').textContent = application.Zakazchik.imya;
    document.getElementById('otchestvo_zakazchika').textContent = application.Zakazchik.otchestvo;
    document.getElementById('stasus_zakazchika').textContent = application.Zakazchik.StatusZakazchika.naimenovanie_statusa;
    document.getElementById('telephone_zakazchika').textContent = application.Zakazchik.telephone;
    document.getElementById('email_zakazchika').textContent = application.Zakazchik.email;
    document.getElementById('vid_prazdnika').textContent = application.VidiPrazdnikov.naimenovanie_vida;
    document.getElementById('data_provedeniya').textContent = new Date(application.data_provedeniya).toLocaleDateString();
    document.getElementById('kolichestvo_chelovek').textContent = application.kolichestvo_chelovek;
    document.getElementById('nachalo_provedeniya').textContent = application.nachalo_provedeniya;
    document.getElementById('konec_provedeniya').textContent = application.konec_provedeniya;
    document.getElementById('prodoljitelnost').textContent = `${application.prodoljitelnost} часа`;

    document.getElementById('editApplicationBtn').setAttribute('data-id', application.id_zayavki);
    document.getElementById('cancelApplicationBtn').setAttribute('data-id', application.id_zayavki);

    const modal = new bootstrap.Modal(document.getElementById('applicationModal'));
    modal.show();
}

// Функция для отправки заявки в работу
function cancelApplication(applicationId) {
    fetch(`/api/application/cancel?app=${applicationId}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => {
            if (response.ok) {
                alert('Вы больше не работаете над этой заявкой.');
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

// Добавление обработчика кнопки "Отказаться от заявки"
document.getElementById('cancelApplicationBtn').addEventListener('click', function () {
    const applicationId = this.getAttribute('data-id');
    if (applicationId) {
        cancelApplication(applicationId);
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
