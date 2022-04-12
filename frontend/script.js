const connectBtn = document.getElementById('connect-btn');
const createSessionBtn = document.getElementById('create-session-btn');
const disconnectBtn = document.getElementById('disconnect-btn');
const sessionIdField = document.getElementById('session-id');

let ws = null;
let sessionId = null;

sessionIdField.addEventListener('input', e => {
  const text = e.target.value;
  const isEmpty = text.length === 0;

  setBoolAttribute(connectBtn, 'disabled', isEmpty);
  setBoolAttribute(createSessionBtn, 'disabled', !isEmpty);

  sessionId = text;
});

createSessionBtn.addEventListener('click', async () => {
  sessionId = await getSessionId();

  setBoolAttribute(createSessionBtn, 'disabled', true);
  sessionIdField.value = sessionId;
});

connectBtn.addEventListener('click', () => {
  if (!sessionId) {
    return;
  }

  ws = connectWs();
  ws.onopen = () => {
    console.log('we are live!');
    setBoolAttribute(disconnectBtn, 'disabled', false);
  };

  ws.onclose = () => {
    console.log('ws disconnected');
    setBoolAttribute(disconnectBtn, 'disabled', true);
  };

  ws.onerror = () => {
    console.log('connection error');
  };
});

disconnectBtn.addEventListener('click', () => {
  if (!ws) {
    return;
  }

  ws.close();
  ws = null;
});

function setBoolAttribute(node, attribute, enabled) {
  if (enabled) {
    node.setAttribute(attribute, true);
  } else {
    node.removeAttribute(attribute);
  }
}

async function getSessionId() {
  const inputText = sessionIdField.value;

  return inputText.length > 0 ? inputText : await newSession();
}

async function newSession() {
  const res = await fetch('http://localhost:8000/api/session', {
    method: 'POST'
  });
  const json = await res.json();

  return json.id;
}

function connectWs() {
  return ws ?? new WebSocket('ws://localhost:8000/ws');
}
