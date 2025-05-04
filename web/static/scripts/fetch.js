// Получение структурных подразделений
async function FetchUnits(func) {
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
async function FetchGroups(func, unit) {
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

// Регистрация
async function FetchRegistration(user) {
  try {
    const response = await fetch("/auth/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    });

    if (!response.ok) {
      throw new Error(await response.text());
    }

    return true

  } catch (error) {
    serverMessage("show", error.message || "Возникла ошибка на сервере, попробуйте позже!");
    return false
  }
}

// Авторизация
async function FetchAuthorize(authData) {
  try {
    const response = await fetch("/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(authData),
      credentials: "include"
    });

    if (!response.ok) {
      const error = await response.text();
      throw new Error(
        response.status === 401
          ? "authError"
          : error || "Ошибка сервера"
      );
    }

    const html = await response.text();
    return html

  } catch (error) {
    if (error.message === "authError") {
      return "authError"
    } else {
      serverMessage("show", error.message || "Возникла ошибка на сервере, попробуйте позже!");
      return "error"
    }
  }
}

// Изменение иконки пользователя
async function FetchUpdateUserIcon(icon) {
  try {
    const response = await fetch("/user/me/icon", {
      method: "PATCH",
      body: icon,
      credentials: "include"
    });

    if (!response.ok) {
      throw new Error(await response.text());
    }

    const newIconUrl = await response.text();
    return newIconUrl
    
  } catch (error) {
    serverMessage("show", error.message);
    return ""
  }
}

// Изменение пароля пользователя
async function FetchUpdateUserPass(updatePass) {
  try {
    const response = await fetch("/user/me/pass", {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(updatePass),
      credentials: "include"
    });
    response.status 
    if (!response.ok) {
      if (response.status === 400) {
        return response.text();
      } else {
        throw new Error(error);
      }
    }
    return true

  } catch (error) {
    serverMessage("show", error.message);
    return false
  }
}

// Выход пользователя из аккаунта
async function FetchLogoutUser() {
  try {
    const response = await fetch("/auth/logout", {
      method: "DELETE",
      credentials: "include"
    });

    if (response.ok) {
      window.location.href = "/";
    } else {
      throw new Error(await response.text());
    }

  } catch (error) {
    serverMessage("show", error.message);
  }
}

async function FetchGetDocuments(settings) {
  try {
    const url = new URL('/documents', window.location.origin);
    
    Object.entries(settings).forEach(([key, value]) => {
      if (value !== undefined && value !== null) {
        url.searchParams.append(key, value);
      }
    });

    const response = await fetch(url.toString(), {
      method: 'GET',
      headers: {
        'Accept': 'application/json'
      }
    });

    if (!response.ok) {
      throw new Error(await response.text());
    }

    return await response.json();

  } catch (error) {
    serverMessage("show", error.message);
  }
}