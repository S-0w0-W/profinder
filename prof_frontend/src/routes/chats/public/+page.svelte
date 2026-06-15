<script lang="ts">
  import { onMount } from "svelte";

  let currentPath = "";
  let socket: WebSocket | null = null;

  onMount(() => {
    // Browser APIs are completely safe to use here
    currentPath = window.location.pathname;
    console.log(currentPath);
    console.log(window.location);
    socket = new WebSocket(`ws://${location.host}/chat/public`);
    socket.onmessage = (e) => {
      // const msg = JSON.parse(e.data)
      console.log("received:", e);
    };
  });
  // const socket = new WebSocket(`ws://${location.hostname}/chat/public`);
  // const socket = new WebSocket(`ws://localhost:8080/chat/public`);

  function handleSubmit(event: Event) {
    // 1. Stop the browser from reloading the page
    event.preventDefault();

    // 2. Access form data easily
    const formData = new FormData(event.currentTarget as HTMLFormElement);
    var message = formData.get("comment") as string;
    // console.log(message);

    if (socket) {
      console.log("sending:", socket);
      socket.send(message);
    }
  }
</script>

<form id="myForm" onsubmit={handleSubmit} action="/submit-path">
  <label for="user-comment">Leave a Comment:</label><br />
  <!-- <textarea
    id="user-comment"
    name="comment"
    rows="4"
    cols="50"
    placeholder="Type your comment here..."
  ></textarea><br /> -->
  <input type="text" id="search" name="q" placeholder="Type here..." /><br />
  <button type="submit">Submit</button>
</form>
