<script>
    import { onMount } from "svelte";
    let todos;
    $: jwt = getCookie("JWT");
    const serverIP = location.host;

    function getTodos() {
        todos = fetch(serverIP, {
            method: "GET",
            headers: {
                authentication: jwt,
            },
        });
    }

    function getCookie(name) {
        const cookies = document.cookie.split(";");
        for (let cookie of cookies) {
            const [cookieName, cookieValue] = cookie
                .split("=")
                .map((c) => c.trim());
            if (cookieName === name) {
                return decodeURIComponent(cookieValue);
            }
        }
        return null;
    }
    onMount(getTodos);
</script>

<main>
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link
        rel="preconnect"
        href="https://fonts.gstatic.com"
        crossorigin="true"
    />
    <link
        href="https://fonts.googleapis.com/css2?family=Roboto&display=swap"
        rel="stylesheet"
    />
    <h1>To Do</h1>
    {#await todos}
        Loading todos...
    {:then resp}
        <ul class="todo-list">
            <h2 class="todo-titles">title &emsp date &esmp status</h2>
            {#each resp as todo}
                <li class="todo-element">
                    {todo.title} &esmp {todo.date} &esmp {todo.status}
                </li>
            {/each}
        </ul>
    {/await}
</main>

<style>
    main {
        background-color: #eadbc8;
    }
    .todo-element {
        font-family: "Roboto", sans-serif;
    }
</style>
