//-------Переключение темы
table_theme = document.querySelectorAll(".settings__theme--table");

table_theme.forEach((table) => {
  // вешаем на каждую кнопку обработчик события клик
  table.addEventListener("click", () => ChengeTheme(table));

  function ChengeTheme(table) {
    theme = table.getAttribute("theme");
    color = table.getAttribute("color");
    localStorage.setItem("theme", theme);
    localStorage.setItem("color", color);

    if (theme == "light") {
      document.body.classList.add("light-theme");
      document.body.classList.remove("dark-theme");
    } else if (theme == "dark") {
      document.body.classList.remove("light-theme");
      document.body.classList.add("dark-theme");
    }

    switch (color) {
      case "blue":
        {
          document.body.style.setProperty("--main-rgb", "45, 104, 248");
        }
        break;
      case "orange":
        {
          document.body.style.setProperty("--main-rgb", "255, 104, 0");
        }
        break;
      case "purple":
        {
          document.body.style.setProperty("--main-rgb", "116, 66, 200");
        }
        break;
      case "green":
        {
          document.body.style.setProperty("--main-rgb", "3, 108, 86");
        }
        break;
    }
  }
});

//- переключение и открытие настроек в боковых панелях

btn_settings = document.querySelectorAll(".main__settings--btn");
// Проходимся по всем кнопкам
btn_settings.forEach((btn) => {
  // вешаем на каждую кнопку обработчик события клик
  btn.addEventListener("click", () => ChengeActivePanel(btn));

  function ChengeActivePanel(btn) {
    // Получаем предыдущую активную вкладку
    pre_active_btn = document.querySelector(".main__settings-active-btn");
    // Получаем предыдущую активную вкладку
    pre_active_panel = document.querySelector(".main__settings--active-panel");

    // Проверяем есть или нет предыдущая активная кнопка
    if (pre_active_btn) {
      //Удаляем класс _active у предыдущей кнопки если она есть
      pre_active_btn.classList.remove("main__settings-active-btn");
    }
    // Проверяем есть или нет предыдущая активная вкладка
    if (pre_active_panel) {
      // Удаляем класс _active у предыдущей вкладки если она есть
      pre_active_panel.classList.remove("main__settings--active-panel");
    }
    // получаем id новой активной вкладки, который мы перем из атрибута data-tab у кнопки
    const active_panel_id = "#" + btn.getAttribute("panel-id");
    const active_panel = document.querySelector(active_panel_id);

    if (active_panel !== pre_active_panel) {
      // добавляем класс _active кнопке на которую нажали
      btn.classList.add("main__settings-active-btn");
      // добавляем класс _active новой выбранной вкладке
      active_panel.classList.add("main__settings--active-panel");
    }
  }
});
