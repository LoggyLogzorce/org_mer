// Функция для загрузки заявок с сервера и обновления списка
function fetchApplications() {
    fetch('/api/applications', {
        method: 'GET'
    })
        .then(response => response.json())
        .then(data => {
            // Очищаем секции
            document.getElementById('unacceptedApplications').innerHTML = '';
            document.getElementById('inProgressApplications').innerHTML = '';
            document.getElementById('completedApplications').innerHTML = '';

            // Проверяем, если данные пустые
            if (data.length === 0) {
                document.getElementById('unacceptedApplications').innerHTML = '<p>Нет заявок.</p>';
                return;
            }

            // Обрабатываем каждую заявку
            data.forEach(application => {
                const applicationElement = document.createElement('div');
                applicationElement.classList.add('col-12', 'col-md-6', 'col-lg-4'); // Сетка Bootstrap
                applicationElement.innerHTML = `
                    <div class="application card p-3 h-100">
                        <h5>Заявка №${application.id_zayavki}</h5>
                        <p><strong>Мероприятие:</strong> ${application.VidiPrazdnikov.naimenovanie_vida}</p>
                        <p><strong>Дата:</strong> ${new Date(application.data_provedeniya).toLocaleDateString()}</p>
                        <button class="show-modal btn btn-info mt-auto" data-id="${application.id_zayavki}">Подробнее</button>
                        ${application.status_zayavki.naimenovanie_statusa_zayavki === 'Выполнена' ? `<button class="btn btn-warning report-btn mt-2" data-id="${application.id_zayavki}">Отчёт</button>` : ''}
                    </div>
                `;

                // Добавляем заявку в соответствующую секцию
                if (application.status_zayavki.naimenovanie_statusa_zayavki === 'Не принята') {
                    document.getElementById('unacceptedApplications').appendChild(applicationElement);
                } else if (application.status_zayavki.naimenovanie_statusa_zayavki === 'В работе') {
                    document.getElementById('inProgressApplications').appendChild(applicationElement);
                } else if (application.status_zayavki.naimenovanie_statusa_zayavki === 'Выполнена') {
                    document.getElementById('completedApplications').appendChild(applicationElement);
                }

                // Обработчик кнопки "Подробнее"
                applicationElement.querySelector('.show-modal').addEventListener('click', function () {
                    showApplicationDetails(application);
                });

                // Обработчик кнопки "Отчёт"
                if (application.status_zayavki.naimenovanie_statusa_zayavki === 'Выполнена') {
                    applicationElement.querySelector('.report-btn').addEventListener('click', function () {
                        generateReport(application.id_zayavki);
                    });
                }
            });
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

    const modal = new bootstrap.Modal(document.getElementById('applicationModal'));
    modal.show();
}

// Функция для генерации отчёта
function generateReport(applicationId) {
    fetch(`/api/create/docx?app=${applicationId}`, {
        method: 'POST'
    })
    alert(`Генерация отчёта по заявке №${applicationId}...`)
    window.location.href = `/api/download-report?app=${applicationId}`
}

// Функция выхода
function logout() {
    document.cookie = 'token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
    window.location.href = '/login';
}

// Запуск начальной загрузки заявок
fetchApplications();
