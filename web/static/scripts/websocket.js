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
    this.connect();
  }

  connect() {
    if (this.socket) {
      this.socket.close();
    }

    this.socket = new WebSocket(`ws://localhost/ws`);

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
      if (e.data === 'PONG') {
        //console.debug('PONG received');
      } else {
        console.log('Message:', e.data);
      }
    };
  }

  startPing() {
    this.clearPing();
    this.pingTimeout = setInterval(() => {
      if (this.socket.readyState === WebSocket.OPEN) {
        //console.debug('Sending PING');
        this.socket.send('PING');
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

    let delay;
    if (this.reconnectAttempts < this.maxFastReconnectAttempts) {
      delay = this.fastReconnectDelay;
      console.log(`Fast reconnect in ${delay/1000} sec (attempt ${this.reconnectAttempts + 1}/${this.maxFastReconnectAttempts})`);
    } else {
      delay = this.longReconnectDelay;
      console.log(`Slow reconnect in ${delay/1000/60} minutes`);
    }

    this.reconnectTimeout = setTimeout(() => {
      this.reconnectAttempts++;
      this.connect();
    }, delay);
  }

  send(message) {
    if (this.socket?.readyState === WebSocket.OPEN) {
      this.socket.send(message);
    } else {
      console.error('Cannot send - connection not ready');
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
  }
}

// Использование
const client = new WSClient();

// Для тестирования в консоли:
window.testWS = {
  disconnect: () => client.socket.close(),
  reconnect: () => client.connect()
};