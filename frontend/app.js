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

// renders the leaderboard from a players array
function renderLeaderboard(players) {
  const list = document.getElementById('player-list');
  const sorted = [...players].sort((a, b) => b.points - a.points);
  list.innerHTML = sorted.map((p, i) => {
    const rank = i + 1;
    const isYou = p.username === USER;
    const initials = p.username.substring(0, 2);
    return `
      <div class="player-row rank-${rank}">
        <div class="player-rank">${rank <= 3 ? ['&#129351;','&#129352;','&#129353;'][rank-1] : rank}</div>
        <div class="player-avatar">${initials}</div>
        <div class="player-info">
          <div class="player-name ${isYou ? 'is-you' : ''}">${p.username}</div>
        </div>
        <div class="player-score">${p.points}</div>
      </div>`;
  }).join('');
}
 

let ws=null;

/**
 * Handle Websocket Connection
 *
 */
function handleWebsocket() {
  ws = new WebSocket(`${API_URL}/join/${ROOM}?user=${USER}`);

  ws.onopen = function () {
    appendMessage('', `Connected to WebSocket server room ${ROOM}`, false, true);
  };

  ws.onmessage = function (event) {
    const message = JSON.parse(event.data);
    switch (message.type) {
      case 'chat':
        console.log(message)
        if (message.user != USER) appendMessage(message.user, message.message);
        break;
      case 'leaderboard':
        renderLeaderboard(message.player);
        break;
      // handle other message types (e.g., game updates) here
    }
  };

  ws.onclose = function () {
    appendMessage('', 'disconnected', false, true);
    // setTimeout(() => {
    //   appendMessage('', 'Retrying', false, true);
    //   handleWebsocket();
    // }, 1000);
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
  let message = { user: USER, message: text, type: "chat"};
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


/*
============== Canvas ==================
*/
const canvas = document.querySelector('canvas');
const canvasRect = canvas.getBoundingClientRect();

const ctx = canvas.getContext("2d");
ctx.fillStyle = "white";

//states
let running = true;
let isDrawing = false;
let xPos = null;
let yPos = null;
let currX = null;
let currY = null;

canvas.addEventListener('mousedown', (e) => {
  isDrawing=true;
  xPos = e.clientX - canvasRect.left;
  yPos = Math.floor(e.clientY - canvasRect.top);
  ctx.moveTo(xPos, yPos);
  console.log(xPos, yPos)
})

canvas.addEventListener('mouseup', (e) => {
  isDrawing=false;
})

canvas.addEventListener('mouseleave', (e) => {
  isDrawing = false;
})

canvas.addEventListener('mousemove', (e) => {
  if (isDrawing) {
    const x = e.clientX - canvasRect.left;
    const y = e.clientY - canvasRect.top;

    ctx.beginPath();
    ctx.moveTo(xPos, yPos);
    ctx.lineTo(x, y);
    ctx.strokeStyle = '#000';
    ctx.lineWidth = 3;
    ctx.lineCap = 'round';    // smooth joins
    ctx.stroke();

    xPos = x;
    yPos = y;
  }
})
