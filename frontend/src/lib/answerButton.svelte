<script>
  let { answer, isCorrect, index } = $props();

  function pickColor() {
    switch (index) {
      case 0:
        return "bg-green-500 hover:bg-green-600";
      case 1:
        return "bg-blue-500 hover:bg-blue-600";
      case 2:
        return "bg-yellow-400 hover:bg-yellow-500";
      case 3:
        return "bg-red-500 hover:bg-red-600";
      default:
        return "bg-gray-500 hover:bg-gray-600";
    }
  }

  // TODO: change this later
  function handleClick() {
    if (isCorrect) {
      alert("Correct!");
    } else {
      alert("Incorrect!");
    }
  }

  let color = pickColor();
</script>

<button
  type="button"
  class={`${color} rounded-xl flex justify-center items-center p-2 font-bold text-xl`}
  onclick={handleClick}
  aria-label={answer}
>
  {answer}
</button>
