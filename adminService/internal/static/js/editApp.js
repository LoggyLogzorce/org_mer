// Функция для получения данных заявки и заполнения формы
async function fillFormFromJson(jsonData) {
    const zayavkaData = jsonData.Prazdnik.Zayavka;

    // Заполнение полей формы
    document.getElementById('idZayavki').value = zayavkaData.id_zayavki;
    document.getElementById('idVidaPrazdnika').value = zayavkaData.id_vida_prazdnika;
    document.getElementById('dataProvedeniya').value = zayavkaData.data_provedeniya.split('T')[0]; // Только дата без времени
    document.getElementById('kolichestvoChelovek').value = zayavkaData.kolichestvo_chelovek;
    document.getElementById('nachaloProvedeniya').value = zayavkaData.nachalo_provedeniya;
    document.getElementById('konecProvedeniya').value = zayavkaData.konec_provedeniya;
    document.getElementById('prodoljitelnost').value = zayavkaData.prodoljitelnost;

    // Заполнение статуса заявки
    const statusSelect = document.getElementById('idStatusaZayavki');
    const statusOptions = jsonData.Statusi;
    statusOptions.forEach(status => {
        const option = document.createElement('option');
        option.value = status.id_statusa_zayavki;
        option.textContent = status.naimenovanie_statusa_zayavki;
        if (status.id_statusa_zayavki === zayavkaData.id_statusa_zayavki) {
            option.selected = true;
        }
        statusSelect.appendChild(option);
    });

    // Заполнение данных заказчика
    const zakazchik = zayavkaData.Zakazchik;
    document.getElementById('idZakazchika').value = `${zakazchik.familiya} ${zakazchik.imya} ${zakazchik.otchestvo}`;
    document.getElementById('telephoneZakazchik').value = zakazchik.telephone;
    document.getElementById('emailZakazchik').value = zakazchik.email;

    // Заполнение данных вида праздника
    const vidiPrazdnikov = zayavkaData.VidiPrazdnikov;
    document.getElementById('idVidaPrazdnika').value = vidiPrazdnikov.naimenovanie_vida;

    // Заполнение данных площадки
    const ploshadkaSelect = document.getElementById('idPloshadki'); // Идентификатор выпадающего списка для площадок
    const option = document.createElement('option');
    option.innerHTML = '';
    option.value = "0";
    option.textContent = 'Площадка не выбрана';
    option.selected = true;
    ploshadkaSelect.appendChild(option);

    const ploshadki = jsonData.Ploshadki; // Получаем массив площадок
    if (ploshadki && ploshadki.length > 0) {
        ploshadki.forEach(ploshadka => {
            const option = document.createElement('option');
            option.value = ploshadka.id_ploshadki; // Значение элемента - ID площадки
            option.textContent = ploshadka.nazvanie; // Текст элемента - название площадки

            // Устанавливаем выбранную площадку, если ID совпадает
            if (ploshadka.id_ploshadki === jsonData.Prazdnik.id_ploshadki) {
                option.selected = true;
            }
            ploshadkaSelect.appendChild(option);
        });
    } else {
        // Если площадок нет, добавляем пустую опцию
        const option = document.createElement('option');
        option.textContent = 'Нет доступных площадок';
        option.disabled = true;
        ploshadkaSelect.appendChild(option);
    }

    // Заполнение данных ведущего
    const vedushiySelect = document.getElementById('vedushiy'); // Идентификатор выпадающего списка для площадок
    const optionVedushiy = document.createElement('option');
    optionVedushiy.innerHTML = '';
    optionVedushiy.value = "0";
    optionVedushiy.textContent = 'Ведущий не выбран';
    optionVedushiy.selected = true;
    vedushiySelect.appendChild(optionVedushiy);

    const vedushie = jsonData.Vedushie; // Получаем массив площадок
    if (vedushie && vedushie.length > 0) {
        vedushie.forEach(vedushiy => {
            const option = document.createElement('option');
            option.value = vedushiy.id_vedushego;
            option.textContent = `${vedushiy.familiya} ${vedushiy.imya} ${vedushiy.otchestvo}`;

            if (vedushiy.id_vedushego === jsonData.Prazdnik.id_vedushevo) {
                option.selected = true;
            }
            vedushiySelect.appendChild(option);
        });
    } else {
        const option = document.createElement('option');
        option.textContent = 'Нет доступных ведущих';
        option.disabled = true;
        vedushiySelect.appendChild(option);
    }

    // Дополнительные услуги
    // Дополнительные услуги
    const dopUslugi = jsonData.Uslugi;
    const dopUslugiContainer = document.getElementById('dopUslugiContainer'); // Контейнер для услуг
    dopUslugiContainer.innerHTML = ''; // Очистить контейнер

    dopUslugi.forEach(service => {
        const serviceWrapper = document.createElement('div');
        serviceWrapper.style.marginBottom = '5px'; // Между строками услуг

        const checkbox = document.createElement('input');
        checkbox.type = 'checkbox';
        checkbox.name = 'dop_uslugi'; // Имя для сбора значений
        checkbox.value = service.id_uslugi_v_zayavke; // ID услуги

        // Если услуга была выбрана (поле complete == true), то чекбокс будет отмечен
        checkbox.checked = service.complete;

        const label = document.createElement('label');
        label.textContent = ` ${service.Usluga.naimenovanie}`; // Добавляем пробел перед названием услуги
        label.prepend(checkbox);

        serviceWrapper.appendChild(label);
        dopUslugiContainer.appendChild(serviceWrapper);
    });

}

// Функция для получения параметра из URL
function getQueryParam(param) {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get(param);
}

// Основной процесс получения данных и их заполнения
async function init() {
    const idZayavki = getQueryParam('id');

    if (!idZayavki) {
        console.error('ID заявки не найден в URL');
        return;
    }

    // Имитация загрузки данных JSON с сервером
    const jsonData = await fetch(`/api/get/holiday?id=${idZayavki}`).then(res => res.json());

    // Заполнение формы на основе полученных данных
    fillFormFromJson(jsonData);
}

async function saveChanges() {
    const id_zayavki = document.getElementById('idZayavki').value;
    const id_statusa_zayavki = document.getElementById('idStatusaZayavki').value;
    const id_ploshadki = document.getElementById('idPloshadki').value;
    const id_vedushego = document.getElementById('vedushiy').value;

    // Собираем все выбранные услуги
    const dop_uslugi = Array.from(document.querySelectorAll('input[name="dop_uslugi"]:checked'))
        .map(checkbox => checkbox.value);

    // Формируем объект с данными для отправки
    const data = {
        id_zayavki,
        id_statusa_zayavki,
        id_ploshadki,
        id_vedushego,
        dop_uslugi // Сюда добавляем массив с услугами
    };

    console.log("Data to save:", data);

    // Отправляем данные на сервер
    fetch('/api/save/holiday', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data) // Отправляем данные как JSON
    })
        .then(response => response.json())
        .then(data => {
            console.log('Row updated:', data);
            alert("Данные сохранены!")
            window.location.href = '/tasks'
        })
        .catch(error => {
            console.error('Error updating row:', error);
        });
}


// Обработчик для кнопки "Сохранить изменения"
document.getElementById("saveChanges").addEventListener("click", function (event) {
    event.preventDefault(); // Предотвращение стандартного поведения формы (отправка)

    saveChanges();
});

// Запуск инициализации после загрузки страницы
document.addEventListener('DOMContentLoaded', init);

function logout() {
    document.cookie = 'token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
    window.location.href = '/login';
}