<script>
  import AnswerButton from "$lib/answerButton.svelte";

  let questions = [];
  async function fetchQuestions() {
    try {
      const response = await fetch("http://127.0.0.1:3000/questions");
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      const data = await response.json();
      questions = data;
    } catch (error) {
      console.error(error.message);
    }
  }

  function connectWebSocket() {
    const ws = new WebSocket("ws://127.0.0.1:3000/ws");
    ws.onopen = () => {
      console.log("Websocket connection opened");
      ws.send("Hello from the client!");
    };

    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      console.log(data);
    };
  }

  fetchQuestions();

  connectWebSocket();
</script>

<div>
  {#await questions}
    <p>Loading...</p>
  {:then questions}
    {#each questions as question}
      <h1 class="text-3xl font-bold mb-2">{question.title}</h1>
      <div class="grid grid-cols-2 grid-rows-2 gap-2">
        {#each question.answers as answer, i}
          <AnswerButton
            answer={answer.text}
            isCorrect={answer.isCorrect}
            index={i}
          />
        {/each}
      </div>
    {/each}
  {:catch error}
    <p>{error.message}</p>
  {/await}
</div>
