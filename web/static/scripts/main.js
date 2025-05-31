//-----Открытие боковых панелей------------

//------------------------------MAIN-------------------------------------

//------Функция прокрутки и закрепления шапки таблицы-------
window.addEventListener("scroll", function () {
  const active_tub = document.querySelector(".main__tabs--active");
  const head_containers = active_tub.querySelector(".tubs__container--head");
  const hidden_containers = active_tub.querySelector(".tubs__container--hidden");

  if (window.scrollY > 30) {
    hidden_containers.style.display = "block";
    head_containers.classList.add("tubs__container--tabscroll");
  } else {
    head_containers.classList.remove("tubs__container--tabscroll");
    hidden_containers.style.display = "none";
  }
});
