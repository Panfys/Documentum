//-----Открытие боковых панелей------------

//------------------------------MAIN-------------------------------------

//------Функция прокрутки и закрепления шапки таблицы-------
window.addEventListener("scroll", function () {
  active_tub = document.querySelector(".main__tabs--active");
  head_containers = active_tub.querySelector(".tubs__container--head");
  hidden_containers = active_tub.querySelector(".tubs__container--hidden");

  if (window.scrollY > 30) {
    head_containers.classList.add("tubs__container--tabscroll");
    hidden_containers.style.display = "block";
  } else {
    head_containers.classList.remove("tubs__container--tabscroll");
    hidden_containers.style.display = "none";
  }
});

//------Функция изменения количества экземпляров Исходящего-----

outgoing_tub = document.querySelector("#main-tab-outgoing");
count_input = outgoing_tub.querySelector("#input-newdoc-count");

count_input.onchange = function () {
  sender_input = outgoing_tub.querySelector("#input-newdoc-sender");
  outgoing_sender =
    '<input class="table__text--input" id="input-newdoc-sender" name="sender" type="text" value="' +
    sender_input.value +
    '">';

  senders =
    '<input class="table__text--input" id="input-newdoc-sender" name="sender" type="text" value="По расчету-рассылки">';
  sender =
    '<input class="table__text--input" id="input-newdoc-sender" name="sender" type="text">';
  copy =
    '<input class="table__text--input" id="input-newdoc-copy" name="copy" type="text">';

  if (count_input.value < 2) {
    outgoing_tub.querySelector("#input-newdoc-count").value = 1;
    outgoing_tub.querySelector("#column-newdoc-sender").innerHTML =
      outgoing_sender;
    outgoing_tub.querySelector("#column-newdoc-copy").innerHTML = copy;
  }

  if (count_input.value > 5) {
    outgoing_tub.querySelector("#column-newdoc-sender").innerHTML = senders;
    outgoing_tub.querySelector("#column-newdoc-copy").innerHTML = copy;
  }

  if (count_input.value > 1 && count_input.value < 6) {
    sender = outgoing_sender;
    for (i = 1; i < count_input.value; i++) {
      sender +=
        '<input class="table__text--input" id="input-newdoc-sender' +
        i +
        '" name="sender' +
        i +
        '" type="text">';
      copy +=
        '<input class="table__text--input" id="input-newdoc-copy' +
        i +
        '" name="copy' +
        i +
        '" type="text">';
    }
    outgoing_tub.querySelector("#column-newdoc-sender").innerHTML = sender;
    outgoing_tub.querySelector("#column-newdoc-copy").innerHTML = copy;
  }
};

