<script>
    import { onMount } from "svelte";
    import Login from "./lib/login.svelte";
    import Content from "./lib/content.svelte";
    let todos;
    let showRegister = false;
    let showLogin = true;
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
    <link
        rel="stylesheet"
        href="https://cdnjs.cloudflare.com/ajax/libs/animate.css/4.1.1/animate.min.css"
    />
    <div class="container">
        <div class="title">
            <h1 class="animate__backInDown animate__animated">to do app</h1>
        </div>
        <div class="content">
            {#if showLogin || showRegister}
                <Login bind:showLogin bind:showRegister />
            {:else}
                <Content />
            {/if}
        </div>
    </div>
</main>

<style>
    *,
    *::before,
    *::after {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
    }
    .title {
        grid-area: 1/2/1/3;
        display: flex;
        justify-content: flex-start;
        align-items: flex-end;
    }
    .title > h1 {
        font-size: 50pt;
    }
    .container {
        display: grid;
        grid-template-rows: 1fr 4fr 1fr;
        grid-template-columns: 1fr 4fr 1fr;
        height: 100vh;
        width: 100vw;
        margin: 0px 0px 0px 0px;
    }
    .content {
        grid-area: 2/2/3/3;
    }
</style>
