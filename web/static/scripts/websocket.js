        
login = document.getElementById('user-login'); 
ustatus = document.getElementById('user-status').innerHTML.trim();
i = 1;


function ErrorsMessage (ErrorMessage)
{
  document.getElementById('server-error-div').style.display = "flex";
  document.getElementById('server-error-p').innerHTML += ErrorMessage + "</br>";
  setTimeout(ErrorsClouse, 60000);
}

function ErrorsClouse ()
{
  document.getElementById('server-error-div').style.display = "none";
  document.getElementById('server-error-p').innerHTML = "";
}

function OnlineUp (user)
{
  if (ustatus == "Администратор")
  {
  userid = "online-"+ user;
  document.getElementById(userid).innerHTML = "Онлайн";
  document.getElementById(userid).style.color = "#2d68f8";
  document.getElementById(userid).style.fontWeight="bold";
  }
}

function OnlineDoun (user)
{
  if (ustatus == "Администратор")
  {
  userid = "online-"+ user;
  document.getElementById(userid).innerHTML = "Офлайн";
  document.getElementById(userid).style.color = "#f2f2f2";
  document.getElementById(userid).style.fontWeight="normal"; 
  } 
}

let socket = new WebSocket('ws://localhost:8080');
          // Connection opened
  socket.onopen = function(e) 
{
  let ConnectData = 
    {
        action : "Connect",
        user : login.innerHTML.trim().toLowerCase()
    }
    ErrorsMessage ("Сообщение от сервера : Соединение c серевером установлено!");
  i++;
  socket.send(JSON.stringify(ConnectData));
};

socket.onmessage = function(event) 
{
  data = JSON.parse(event.data);
  users = data.users;
  
  if (data.action == "Online")
  {
    for (var i in users) 
    {   
      if (users[i].userLogin != null)
      {
        ErrorsMessage("Сообщение от сервера ("+ data.action +") : " + users[i].userLogin.trim());
        OnlineUp (users[i].userLogin.trim());
      }
    }   
  }   
  else if (data.action == "Authorized")
  {
    ErrorsMessage("Сообщение от сервера ("+ data.action +") : " + data.userLogin.trim());
    OnlineUp (data.userLogin.trim());
  }
  else if (data.action == "Disconnected")
  {
    ErrorsMessage("Сообщение от сервера ("+ data.action +") : " + data.userLogin.trim());
    OnlineDoun (data.userLogin.trim());
  }
  else if (data.action == "Message")
  {
    ErrorsMessage("Сообщение от сервера ("+ data.action +") : " + data.message.trim());
    OnlineDoun (data.userLogin.trim());
  }
  else if (data.action == "NewDocument")
  {
    ErrorsMessage("Сообщение от сервера ("+ data.action +") : Добавлен новый документ! ");
    switch (data.type)
    {
      case 'Входящий': document.getElementById('documents-vhod').innerHTML += data.document; break;
      case 'Исходящий': document.getElementById('documents-isx').innerHTML += data.document; break;
      case 'Директива': document.getElementById('documents-dir').innerHTML += data.document; break;
    }
  }
  else if (data.action == "LookDocument")
  {
    if (document.getElementById(data.id +'-doc').querySelector(".col7"))
    {
      document.getElementById(data.id +'-doc').querySelector(".col7").innerHTML = data.familiar;
    }
    else {document.getElementById(data.id +'-doc').querySelector(".col13").innerHTML = data.familiar;}
    ErrorsMessage("Сообщение от сервера ("+ data.action +") : Кто-то просмотрел документ! ");
  }
  else if (data.action == "AddResolution")
  {
    ErrorsMessage("Сообщение от сервера ("+ data.action +") : Кто-то поставил резолюцию на документ № " + data.id);
    document.getElementById(data.id +'-res').innerHTML += data.resolution;
    isp = data.ispolnitel;
    setTimeout(function(){document.getElementById(data.id +'-doc').querySelector(".col5").innerHTML = isp;},500);
  }
  else if (data.action == "NewLocation")
  {
    ErrorsMessage("Сообщение от сервера ("+ data.action +") : Кто-то подшил документ № " + data.id);
    document.getElementById(data.id +'-doc').querySelector(".col11").innerHTML = data.location;
  }
  else if (data.action == "NewResult")
  {
    ErrorsMessage("Сообщение от сервера ("+ data.action +") : Кто-то исполнил документ № " + data.id);
    document.getElementById(data.id +'-doc').querySelector(".col6").innerHTML = data.result;
  }
  else
  {
    alert (event.data);  
  }
};

socket.onclose = function(event) 
{
  if (event.wasClean) {
    ErrorsMessage(`Сообщение от сервера : Соединение закрыто, код=${event.code} причина=${event.reason}`);
  } else {
    // например, сервер убил процесс или сеть недоступна
    // обычно в этом случае event.code 1006
    ErrorsMessage('Сообщение от сервера : Сервер недоступен');
  }
};

socket.onerror = function(error) {
  ErrorsMessage('Сообщение от сервера: Ошибка соединения');
};