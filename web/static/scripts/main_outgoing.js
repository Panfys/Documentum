//------Функция изменения количества экземпляров Исходящего-----

const outgoing_tub = document.querySelector("#main-tab-outgoing");
const count_input = outgoing_tub.querySelector("#input-newdoc-count");
const sender_column = outgoing_tub.querySelector("#column-newdoc-sender");
const copy_column = outgoing_tub.querySelector("#column-newdoc-copy");

// Инициализация начального состояния
document.addEventListener('DOMContentLoaded', updateInputs);

count_input.addEventListener('change', updateInputs);

function updateInputs() {
  const count = parseInt(count_input.value) || 1;
  const current_sender = outgoing_tub.querySelector("#input-newdoc-sender")?.value || '';
  
  // Очищаем колонки
  sender_column.innerHTML = '';
  copy_column.innerHTML = '';
  
  // Обработка разных случаев количества экземпляров
  if (count < 2) {
    // Один экземпляр - обычные поля
    sender_column.innerHTML = createSenderInput(current_sender);
    copy_column.innerHTML = createCopyInput();
  } 
  else if (count > 5) {
    // Более 5 экземпляров - специальное значение
    sender_column.innerHTML = createSenderInput("По расчету-рассылки");
    copy_column.innerHTML = createCopyInput();
  } 
  else {
    // От 2 до 5 экземпляров - динамические поля
    for (let i = 0; i < count; i++) {
      const sender_value = i === 0 ? current_sender : '';
      sender_column.appendChild(createSenderInputElement(sender_value, i));
      copy_column.appendChild(createCopyInputElement(i));
    }
  }
}

// Функции для создания элементов
function createSenderInput(value = '', index = 0) {
  return `<input class="table__text--input" id="input-newdoc-sender${index || ''}" 
          name="sender${index || ''}" type="text" value="${value}">`;
}

function createCopyInput(index = 0) {
  return `<input class="table__text--input" id="input-newdoc-copy${index || ''}" 
          name="copy${index || ''}" type="text">`;
}

// Альтернативные версии для appendChild
function createSenderInputElement(value = '', index = 0) {
  const input = document.createElement('input');
  input.className = 'table__text--input';
  input.id = `input-newdoc-sender${index || ''}`;
  input.name = `sender${index || ''}`;
  input.type = 'text';
  input.value = value;
  return input;
}

function createCopyInputElement(index = 0) {
  const input = document.createElement('input');
  input.className = 'table__text--input';
  input.id = `input-newdoc-copy${index || ''}`;
  input.name = `copy${index || ''}`;
  input.type = 'text';
  return input;
}

// Функция для подготовки данных к отправке
function prepareDataForSubmission() {
  const count = parseInt(count_input.value) || 1;
  let senderValue = '';
  let copyValue = '';
  
  if (count === 1) {
    // Один экземпляр
    senderValue = outgoing_tub.querySelector("#input-newdoc-sender").value;
    copyValue = outgoing_tub.querySelector("#input-newdoc-copy").value;
  } 
  else if (count > 5) {
    // Более 5 экземпляров
    senderValue = "По расчету-рассылки";
    copyValue = outgoing_tub.querySelector("#input-newdoc-copy").value;
  } 
  else {
    // От 2 до 5 экземпляров
    const senderInputs = outgoing_tub.querySelectorAll('[id^="input-newdoc-sender"]');
    const copyInputs = outgoing_tub.querySelectorAll('[id^="input-newdoc-copy"]');
    
    senderValue = Array.from(senderInputs).map(input => input.value).join(' <br> ');
    copyValue = Array.from(copyInputs).map(input => input.value).join(' <br> ');
  }
  
  return {
    count: count,
    sender: senderValue,
    copy: copyValue
  };
}

