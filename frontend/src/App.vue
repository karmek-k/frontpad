<template>
  <header>
    <h1>Frontpad</h1>
    <h3>a basic PoC frontend!</h3>
  </header>

  <input type="text" id="session-id" placeholder="Leave empty to create" />
  <button id="create-session-btn">Create a session</button>
  <button id="connect-btn" disabled>Connect</button>
  <button id="disconnect-btn" disabled>Disconnect</button>

  <div class="hidden">
    <textarea
      name="markup"
      id="markup"
      cols="30"
      rows="10"
      placeholder="Markup code"
    ></textarea>
  </div>
</template>
