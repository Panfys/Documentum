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
