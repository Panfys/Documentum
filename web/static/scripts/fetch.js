// Получение подразделений
async function fetchUnits(func) {
  try {
    const response = await fetch(`/funcs/${encodeURIComponent(func)}/units`);

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(errorText || "Ошибка сервера");
    }

    selectUnits = "<option value=0></option>"
    const units = await response.json();
    units.forEach(unit => {
     selectUnits += `<option value=${unit.ID}>${unit.Name}</option>`
    });
    return selectUnits;

  } catch (error) {
    serverMessage("show", error.message || "Возникла ошибка на сервере, попробуйте позже!");
    return []
  }
}

// Получение подразделений
async function fetchGroups(func, unit) {
  try {
    const response = await fetch(`/funcs/${encodeURIComponent(func)}/${encodeURIComponent(unit)}/groups`);

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(errorText || "Ошибка сервера");
    }

    selectGroups = "<option value=0></option>"
    const groups = await response.json();
    if (groups != null) {
      groups.forEach(group => {
        selectGroups += `<option value=${group.ID}>${group.Name}</option>`
       });
    }
    return selectGroups;

  } catch (error) {
    serverMessage("show", error.message || "Возникла ошибка на сервере, попробуйте позже!");
    return ""
  }
}