const API_URL = "http://localhost:8080";

// Storage helper
const Store = {
  get: (k) => localStorage.getItem(k),
  set: (k, v) => localStorage.setItem(k, v),
};

const USER = Store.get("user");
const ROOM = Store.get("room");

/**
 * Generates a random username if value is empty
 */
function handleUsername() {
  if (!USER) {
    var username = "guest" + Math.floor(Math.random() * 100000);
    Store.set("user", username);
  }
}

// ── populate header if present ──
const headerUsername = document.getElementById("header-username");
const headerRoomCode = document.getElementById("header-room-code");
if (headerUsername) headerUsername.textContent = USER;
if (headerRoomCode) headerRoomCode.textContent = ROOM;

// ── html escape ──
function escapeHtml(str) {
  return str.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;");
}

// ── chat ──
const chatBox = document.getElementById('chat-box');
 
/**
 * Appends Message to chat box.
 * 
 * @param {string} user 
 * @param {string} message 
 * @param {boolean} isSelf 
 * @param {boolean} isSystem 
 * @returns 
 */
function appendMessage(user, message, isSelf = false, isSystem = false) {
  if (!chatBox) return;
 
  const wrapper = document.createElement('div');
  wrapper.classList.add('msg');
  if (isSelf)    wrapper.classList.add('self');
  if (isSystem)  wrapper.classList.add('msg-system');
 
  if (!isSystem) {
    const meta = document.createElement('div');
    meta.classList.add('msg-meta');
    meta.innerHTML = `<span class="name">${escapeHtml(user)}</span>`;
    wrapper.appendChild(meta);
  }
 
  const bubble = document.createElement('div');
  bubble.classList.add('msg-bubble');
  bubble.textContent = message;
  wrapper.appendChild(bubble);
 
  chatBox.appendChild(wrapper);
  chatBox.scrollTop = chatBox.scrollHeight;
}

let ws=null;

/**
 * Handle Websocket Connection
 *
 */
function handleWebsocket() {
  ws = new WebSocket(`${API_URL}/join/${ROOM}`);

  ws.onopen = function () {
    appendMessage('', `Connected to WebSocket server room ${ROOM}`, false, true);
  };

  ws.onmessage = function (event) {
    var chat = JSON.parse(event.data);
    if (chat.user != USER) appendMessage(chat.user, chat.message);
  };

  ws.onclose = function () {
    appendMessage('', 'disconnected', false, true);
    setTimeout(() => {
      appendMessage('', 'Retrying', false, true);
      handleWebsocket();
    }, 1000);
  };

  ws.onerror = function (error) {
    console.error("WebSocket error:", error);
  };
}

/**
 * Send message to server
 *
 * @returns
 */
function sendMessage() {
  const input = document.getElementById("messageInput");
  if (!input) return;
  const text = input.value.trim();
  if (!text || !ws || ws.readyState != WebSocket.OPEN) return;
  let message = { user: USER, message: text };
  ws.send(JSON.stringify(message));
  appendMessage(USER, text, true);
  input.value = "";
}


// ── enter key for chat ──
document.addEventListener("DOMContentLoaded", () => {
  const msgInput = document.getElementById("messageInput");
  if (msgInput) {
    msgInput.addEventListener("keydown", (e) => {
      if (e.key === "Enter") sendMessage();
    });
  }
});

