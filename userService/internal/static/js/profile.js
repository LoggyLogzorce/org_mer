document.addEventListener("DOMContentLoaded", async function() {
    // Функция для загрузки данных
    async function loadData() {
        try {
            // Асинхронный запрос для получения данных с сервера
            const response = await fetch('/api/user/applications'); // Замените на ваш URL для запроса
            const data = await response.json();

            // Загружаем информацию о заказчике
            const customer = data.applications[0].Zakazchik;
            document.getElementById('customer-name').textContent = `${customer.familiya} ${customer.imya} ${customer.otchestvo}`;
            document.getElementById('customer-email').textContent = customer.email;
            document.getElementById('customer-phone').textContent = customer.telephone;
            document.getElementById('customer-status').textContent = customer.StatusZakazchika.naimenovanie_statusa;

            // Загружаем заявки со статусом "Не принята"
            const applicationsPending = data.applications;
            const applicationsPendingContainer = document.getElementById('applicationsListPending');
            applicationsPendingContainer.innerHTML = ''; // Очищаем контейнер перед добавлением новых элементов
            applicationsPending.forEach(application => {
                const applicationElement = document.createElement('div');
                applicationElement.classList.add('col-md-6', 'col-lg-4');
                applicationElement.innerHTML = `
          <div class="card">
            <div class="card-body">
              <h5 class="card-title">${application.VidiPrazdnikov.naimenovanie_vida}</h5>
              <p><strong>Дата:</strong> ${new Date(application.data_provedeniya).toLocaleDateString()}</p>
              <p><strong>Начала проведения:</strong> ${application.nachalo_provedeniya}</p>
              <p><strong>Конец проведения:</strong> ${application.konec_provedeniya}</p>
              <p><strong>Кол-во человек:</strong> ${application.kolichestvo_chelovek}</p>
              <p><strong>Статус:</strong> ${application.status_zayavki.naimenovanie_statusa_zayavki}</p>
            </div>
          </div>
        `;
                applicationsPendingContainer.appendChild(applicationElement);
            });

            // Загружаем заявки со статусом "В работе" или "Выполнена"
            const prazdnikiInProgress = data.prazdniki;
            const applicationsInProgressContainer = document.getElementById('applicationsListInProgress');
            applicationsInProgressContainer.innerHTML = ''; // Очищаем контейнер перед добавлением новых элементов
            prazdnikiInProgress.forEach(prazdnik => {
                const applicationElement = document.createElement('div');
                applicationElement.classList.add('col-md-6', 'col-lg-4');
                applicationElement.innerHTML = `
          <div class="card">
            <div class="card-body">
              <h5 class="card-title">${prazdnik.Zayavka.VidiPrazdnikov.naimenovanie_vida}</h5>
              <p><strong>Дата:</strong> ${new Date(prazdnik.Zayavka.data_provedeniya).toLocaleDateString()}</p>
              <p><strong>Начала проведения:</strong> ${prazdnik.Zayavka.nachalo_provedeniya}</p>
              <p><strong>Конец проведения:</strong> ${prazdnik.Zayavka.konec_provedeniya}</p>
              <p><strong>Кол-во человек:</strong> ${prazdnik.Zayavka.kolichestvo_chelovek}</p>
              <p><strong>Статус:</strong> ${prazdnik.Zayavka.status_zayavki.naimenovanie_statusa_zayavki}</p>
              <p><strong>Полная стоимость:</strong> ${prazdnik.polnaya_stoimost}</p>
            </div>
          </div>
        `;
                applicationsInProgressContainer.appendChild(applicationElement);
            });

            const showModalButtons = document.querySelectorAll('.show-modal');
            showModalButtons.forEach(button => {
                button.addEventListener('click', function () {
                    const applicationId = button.getAttribute('data-id');
                    const application = data.prazdniki.find(app => app.id_zayavki == applicationId);
                    showApplication(application);
                });
            });

        } catch (error) {
            console.error('Ошибка при загрузке данных:', error);
        }
    }

    // Загрузить данные после загрузки страницы
    await loadData();
});

function showApplication(application) {
    console.log(application)
    document.getElementById('familiya_zakazchika').textContent = application.Zayavka.Zakazchik.familiya;
    document.getElementById('imya_zakazchika').textContent = application.Zayavka.Zakazchik.imya;
    document.getElementById('otchestvo_zakazchika').textContent = application.Zayavka.Zakazchik.otchestvo;
    document.getElementById('stasus_zakazchika').textContent = application.Zayavka.Zakazchik.StatusZakazchika.naimenovanie_statusa;
    document.getElementById('telephone_zakazchika').textContent = application.Zayavka.Zakazchik.telephone;
    document.getElementById('email_zakazchika').textContent = application.Zayavka.Zakazchik.email;
    document.getElementById('vid_prazdnika').textContent = application.VidiPrazdnikov.naimenovanie_vida;
    document.getElementById('data_provedeniya').textContent = new Date(application.Zayavka.data_provedeniya).toLocaleDateString();
    document.getElementById('kolichestvo_chelovek').textContent = application.Zayavka.kolichestvo_chelovek;
    document.getElementById('nachalo_provedeniya').textContent = application.Zayavka.nachalo_provedeniya;
    document.getElementById('konec_provedeniya').textContent = application.Zayavka.konec_provedeniya;
    document.getElementById('prodoljitelnost').textContent = `${application.Zayavka.prodoljitelnost} часа`;

    document.getElementById('editApplicationBtn').setAttribute('data-id', application.id_zayavki);

    const modal = new bootstrap.Modal(document.getElementById('applicationModal'));
    modal.show();
}

function logout() {
    document.cookie = 'token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
    window.location.href = '/login';
}