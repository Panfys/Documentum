class WSClient {
  constructor() {
    this.socket = null;
    this.reconnectAttempts = 0;
    this.maxFastReconnectAttempts = 3;
    this.pingInterval = 30000; // 30 секунд
    this.fastReconnectDelay = 3000; // 3 секунды
    this.longReconnectDelay = 60000; // 1 минута
    this.pingTimeout = null;
    this.reconnectTimeout = null;
    this.messageHandlers = new Map(); // Для обработки разных типов сообщений
    this.connect();
  }

  connect() {
    if (this.socket) {
      this.socket.close();
    }

    this.socket = new WebSocket(`ws://localhost:8000/ws`);

    this.socket.onopen = () => {
      console.log('WebSocket connected');
      this.reconnectAttempts = 0;
      this.startPing();
    };

    this.socket.onclose = (e) => {
      console.log(`Disconnected (code: ${e.code}, reason: ${e.reason || 'none'})`);
      this.clearPing();
      this.scheduleReconnect();
    };

    this.socket.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    this.socket.onmessage = (e) => {
      try {
        const message = JSON.parse(e.data);

        switch (message.action) {
          case "PONG": {
            return;
          }
          case "disconnect": {
            alert("Обнаружено новое подключение с другого устройства");
            FetchLogoutUser()
          }
          default: {
            // Вызов обработчика для конкретного action
            const handler = this.messageHandlers.get(message.action);
            if (handler) {
              handler(message.content);
            } else {
              console.log('Unhandled message:', message);
              this.onMessage?.(message); // Общий обработчик
            }
          }
        }
      } catch (err) {
        console.error('Message parse error:', err, 'Raw data:', e.data);
      }
    };
  }

  // Регистрация обработчиков для разных типов сообщений
  on(action, handler) {
    this.messageHandlers.set(action, handler);
    return this; // Для чейнинга
  }

  startPing() {
    this.clearPing();
    this.pingTimeout = setInterval(() => {
      if (this.socket?.readyState === WebSocket.OPEN) {
        this.send({ action: 'PING' });
      }
    }, this.pingInterval);
  }

  clearPing() {
    if (this.pingTimeout) {
      clearInterval(this.pingTimeout);
      this.pingTimeout = null;
    }
  }

  scheduleReconnect() {
    if (this.reconnectTimeout) {
      clearTimeout(this.reconnectTimeout);
    }

    const delay = this.reconnectAttempts < this.maxFastReconnectAttempts
      ? this.fastReconnectDelay
      : this.longReconnectDelay;

    console.log(`Reconnect in ${delay / 1000} sec (attempt ${this.reconnectAttempts + 1})`);

    this.reconnectTimeout = setTimeout(() => {
      this.reconnectAttempts++;
      this.connect();
    }, delay);
  }

  // Отправка структуры Message
  send(message) {
    if (this.socket?.readyState === WebSocket.OPEN) {
      this.socket.send(JSON.stringify(message));
    } else {
      console.error('Cannot send - connection not ready');
      // Можно добавить очередь сообщений при необходимости
    }
  }

  close() {
    this.clearPing();
    if (this.reconnectTimeout) {
      clearTimeout(this.reconnectTimeout);
    }
    if (this.socket) {
      this.socket.close();
    }
    this.messageHandlers.clear();
  }
}

// Использование
const client = new WSClient()
  .on('updDocFam', (content) => {
    const activeTab = document.querySelector(".main__tabs--active");
    const activeTabId = `#${activeTab.id}`;
    const docTabIdData = Object.values(DOCUMENT_TYPES).find(
      tabId => tabId.type === content.type
    );
    if (docTabIdData.tabId === activeTabId) {
      const list = document.querySelector(`#familiar-list-${content.docID}`);
      if (list) {
        const li = document.createElement('li');
        li.textContent = content.familiar;
        list.appendChild(li);
      }
    }
  })
  .on('updDocRes', (content) => {

    const activeTab = document.querySelector(".main__tabs--active");
    const activeTabId = `#${activeTab.id}`;
    const tabId = "#main-tab-ingoing";
    if (tabId === activeTabId) {
      const doc = activeTab.querySelector(`[document-id="${content.id}"]`);
      const docResult = doc.querySelector(".table__column--result");
      const docIspolnitel = doc.querySelector(".table__column--ispolnitel");
      const resolutionPanel = activeTab.querySelector("#resolution-panel-" + content.id);
      if (content.result !== "") {
        if (docResult.innerHTML == "") {
          docResult.innerHTML += content.result;
        } else {
          docResult.innerHTML += "<br>" + content.result;
        }
      }
      if (content.ispolnitel !== "") {
        docIspolnitel.innerHTML = content.ispolnitel;
      }
      if (content.resolutions && content.resolutions.length) {
        content.resolutions.forEach(resolution => {
          if (!resolution) return;

          resolutionPanel.innerHTML += `
                <div class='table__resolution' id='ingoing-resolution'>
                    <div class='table__resolution--ispolnitel'>${resolution.ispolnitel}</div>
                    <div class='table__resolution--text'>&#171;${resolution.text}&#187;</div>
                    <div class='table__resolution--time'>${resolution.deadline || resolution.result}</div>
                    <div class='table__resolution--user'>${resolution.user}</div>
                    <div class='table__resolution--date'>${resolution.date}</div>
                </div>`;
        });
      }
    }
  })
  .on('addDoc', (content) => {
    const activeTab = document.querySelector(".main__tabs--active");
    const activeTabId = `#${activeTab.id}`;
    const docTabIdData = Object.values(DOCUMENT_TYPES).find(
      tabId => tabId.type === content.type
    );
    if (docTabIdData.tabId === activeTabId) {
      const docTableIdData = Object.values(DOCUMENT_TYPES).find(
        documentTableId => documentTableId.type === content.type
      );
      const container = activeTab.querySelector(`#${docTableIdData.documentTableId}`);
      content.familiars = [content.familiar];
      container.innerHTML += WriteDocumentsInTable([content],content.type);
      setupDocumentTables()
    }
  });

// Для тестирования в консоли
window.wsClient = {
  disconnect: () => client.socket.close(),
  reconnect: () => client.connect(),
};