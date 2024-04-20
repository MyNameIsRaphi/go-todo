<script>
    import App from "../App.svelte";

    export let showLogin = true;
    export let showRegister = false;
    let register_password;
    let firstName;
    let lastName;
    let register_password2;
    let registerEmail;
    let email;
    let password;
    let registerError = false;
    let loginError = false;
    async function register() {
        if (register_password !== register_password2) {
            console.log("invalid passwords");
            registerError = true;
            return;
        }

        const notificationsGranted = await Notification.requestPermission();
        const body = {
            Email: registerEmail,
            Password: register_password,
            ConfirmedPassword: register_password2,
            FirstName: firstName,
            LastName: lastName,
            NotificationsGranted: notificationsGranted !== "denied",
        };
        const url = "/register";
        try {
            const res = fetch(url, {
                method: "POST",
                body: JSON.stringify(body),
            });
            if (!res.ok) {
                registerError = true;
                return;
            }
            showRegister = false;
            showLogin = true;
        } catch (err) {
            console.log(err);
            registerError = true;
        }
    }
    async function login() {
        const url = "/login";
        const body = {
            Password: password,
            Email: email,
        };
        try {
            const res = await fetch(url, {
                method: "POST",
                body: JSON.stringify(body),
            });
            const resBody = await res.json();
            if (!res.ok || !resBody.Auth) {
                loginError = true;
            } else {
                loginError = false;
                showLogin = false;
                console.log(resBody);
            }
        } catch (err) {
            loginError = true;
        }
    }
</script>

<link rel="preconnect" href="https://fonts.googleapis.com" />
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
<link
    href="https://fonts.googleapis.com/css2?family=Roboto:ital,wght@0,100;0,300;0,400;0,500;0,700;0,900;1,100;1,300;1,400;1,500;1,700;1,900&display=swap"
    rel="stylesheet"
/>

<link
    href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css"
    rel="stylesheet"
    integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH"
    crossorigin="anonymous"
/>

{#if showLogin}
    <div class="login" style="visibility: {showLogin};">
        <form class="form">
            <div class="input-group input-group-sm mb-3 input">
                <span class="input-group-text" id="inputGroup-sizing-sm"
                    >email
                </span>
                <input
                    type="email"
                    name="email"
                    class="form-control"
                    aria-label="Sizing example input"
                    aria-describedby="inputGroup-sizing-sm"
                    bind:value={email}
                    required
                />
            </div>
            <div class="input-group input-group-sm mb-3 input input-password">
                <span class="input-group-text" id="inputGroup-sizing-sm"
                    >password</span
                >
                <input
                    class="form-control"
                    aria-label="Sizing example input"
                    aria-describedby="inputGroup-sizing-sm"
                    type="password"
                    bind:value={password}
                    required
                />
            </div>
            <button type="button" class="btn btn-primary" on:click={login}
                >login</button
            >
            {#if loginError}
                <p class="error">password and email don't match</p>
            {/if}
            <button
                type="button"
                class="btn btn-link"
                on:click={() => {
                    showLogin = false;
                    showRegister = true;
                }}>register</button
            >
        </form>
    </div>
{:else if showRegister}
    <div class="register">
        <form class="from">
            <button
                type="button"
                class="btn btn-link"
                on:click={() => {
                    showLogin = true;
                    showRegister = false;
                }}>login</button
            >
            <div class="input-group mb-3">
                <span class="input-group-text" id="basic-addon1">@</span>
                <input
                    type="email"
                    name="email"
                    class="form-control"
                    required
                    placeholder="email"
                    aria-label="Username"
                    aria-describedby="basic-addon1"
                    bind:value={registerEmail}
                />
            </div>
            <div class="input-group input-group-sm mb-3 input input-password">
                <span class="input-group-text" id="inputGroup-sizing-sm"
                    >password</span
                >
                <input
                    class="form-control"
                    aria-label="Sizing example input"
                    required
                    aria-describedby="inputGroup-sizing-sm"
                    type="password"
                    bind:value={register_password}
                />
                <div
                    class="input-group input-group-sm mb-3 input input-password"
                >
                    <span class="input-group-text" id="inputGroup-sizing-sm"
                        >confirm</span
                    >
                    <input
                        class="form-control"
                        aria-label="Sizing example input"
                        aria-describedby="inputGroup-sizing-sm"
                        type="password"
                        required
                        bind:value={register_password2}
                    />
                </div>
                <div
                    class="input-group input-group-sm mb-3 input input-password"
                >
                    <span class="input-group-text" id="inputGroup-sizing-sm"
                        >first name</span
                    >
                    <input
                        class="form-control"
                        aria-label="Sizing example input"
                        aria-describedby="inputGroup-sizing-sm"
                        type="password"
                        required
                        bind:value={firstName}
                    />
                </div>
                <div
                    class="input-group input-group-sm mb-3 input input-password"
                >
                    <span class="input-group-text" id="inputGroup-sizing-sm"
                        >last name</span
                    >
                    <input
                        class="form-control"
                        aria-label="Sizing example input"
                        aria-describedby="inputGroup-sizing-sm"
                        type="password"
                        required
                        bind:value={lastName}
                    />
                </div>
                {#if registerError}
                    <p class="error">
                        email is already used or passwords don't match
                    </p>
                {/if}
                <p>please confirm your email address after you register</p>
            </div>
            <button type="button" class="btn btn-primary" on:click={register}
                >register</button
            >
        </form>
    </div>
{/if}

<style>
    .form {
        width: 100%;
        height: 100%;
        display: felx;
        align-content: center;
        justify-content: flex-start;
    }
    .input-password {
        margin-top: 1em;
    }
    .login {
        height: 100%;
        width: 100%;
    }

    * {
        font-family: "Roboto", sans-serif;
        font-weight: 400;
        font-style: normal;
    }
    .error {
        color: red;
    }
</style>
