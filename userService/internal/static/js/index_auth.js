document.addEventListener("DOMContentLoaded", function () {
    function fetchHolidayData() {
        fetch('/api/get/vidi-prazdnikov', {
            method: "GET"
        })
            .then(response => response.json())  // Обрабатываем ответ как JSON
            .then(data => {
                // Обработка полученных данных и динамическое создание блоков
                const holidayBlocksContainer = document.getElementById('events-container');
                holidayBlocksContainer.innerHTML = '<br><h2 id="vidi-prazdnikov" class="text-center mb-4">Виды праздников</h2>';  // Очистка контейнера перед добавлением новых блоков

                // Перебор всех праздников из ответа
                data.forEach(holiday => {
                    const holidayBlock = document.createElement('div');
                    holidayBlock.classList.add('col-lg-4', 'col-md-6', 'mb-4');

                    // Структура для блока праздника
                    holidayBlock.innerHTML = `
              <div class="card">
                <img src="${holiday.photo}" class="card-img-top" alt="${holiday.name}">
                <div class="card-body">
                  <h5 class="card-title">${holiday.naimenovanie_vida}</h5>
                  <p class="card-text">${holiday.opisanie}</p>
                  <p class="card-text">От ${holiday.summa} рублей.</p>
                  <button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#applicationModal">
                        Оставить заявку
                  </button>
                </div>
              </div>
            `;

                    holidayBlocksContainer.appendChild(holidayBlock);
                });
            })
            .catch(error => {
                console.error('Ошибка при загрузке данных:', error);
            });
    }

    function fetchUslugiData() {
        fetch('/api/get/dop-uslugi', {
            method: "GET"
        })
            .then(response => response.json())  // Обрабатываем ответ как JSON
            .then(data => {
                // Обработка полученных данных и динамическое создание блоков
                const uslugiBlocksContainer = document.getElementById('dop-uslugi');
                uslugiBlocksContainer.innerHTML = '<br><h2 id="dop-uslugi" class="text-center mb-4">Дополнительные услуги</h2>';  // Очистка контейнера перед добавлением новых блоков

                // Перебор всех праздников из ответа
                data.forEach(uslugi => {
                    const uslugiBlock = document.createElement('div');
                    uslugiBlock.classList.add('col-lg-4', 'col-md-6', 'mb-4');

                    // Структура для блока праздника
                    uslugiBlock.innerHTML = `
              <div class="card">
                <div class="card-body">
                  <h5 class="card-title">${uslugi.naimenovanie}</h5> 
                  <p class="card-text">${uslugi.opisanie}</p>
                  <p class="card-text">${uslugi.stoimost} рублей.</p>
                </div>
              </div>
            `;

                    uslugiBlocksContainer.appendChild(uslugiBlock);
                });

                const servicesContainer = document.getElementById('additional-services');

                // Очищаем контейнер перед добавлением новых чекбоксов
                servicesContainer.innerHTML = '';

                // Перебор всех дополнительных услуг
                data.forEach(service => {
                    const checkboxWrapper = document.createElement('div');
                    checkboxWrapper.classList.add('form-check');

                    const checkbox = document.createElement('input');
                    checkbox.type = 'checkbox';
                    checkbox.classList.add('form-check-input');
                    checkbox.id = service.id_uslugi;
                    checkbox.name = service.naimenovanie;
                    checkbox.value = service.id_uslugi;

                    const label = document.createElement('label');
                    label.classList.add('form-check-label');
                    label.setAttribute('for', checkbox.id);
                    label.innerText = service.naimenovanie;

                    checkboxWrapper.appendChild(checkbox);
                    checkboxWrapper.appendChild(label);
                    servicesContainer.appendChild(checkboxWrapper);
                });
            })
            .catch(error => {
                console.error('Ошибка при загрузке данных:', error);
            });
    }

    // Загружаем данные при загрузке страницы
    fetchHolidayData();
    fetchUslugiData();
});

function logout() {
    document.cookie = 'token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
    window.location.href = '/';
}