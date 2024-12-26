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
                  <button id="buttonForm" type="button" class="btn btn-primary" onclick=ClickFormZayavki()>
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
            })
            .catch(error => {
                console.error('Ошибка при загрузке данных:', error);
            });
    }

    // Загружаем данные при загрузке страницы
    fetchHolidayData();
    fetchUslugiData();
});


function ClickFormZayavki() {
    alert('Сначала нужно создать вам профиль.')
    window.location.href = '/registration'
}