<script>
	import { onMount, afterUpdate } from 'svelte';
	import axios from 'axios';

	const apiURL = import.meta.env.VITE_API_ENDPOINT;
	let data = null;
	let loading = true;
	let audioElement = null;
	let showPlayer = false;
	let fetchingData = false;

	async function fetchData() {
		fetchingData = true;
		try {
			const response = await axios.get(apiURL);
			data = response.data;
		} catch (error) {
			console.log(error);
		} finally {	
			loading = false;
			fetchingData = false;
		}
	}

	onMount(fetchData);

	afterUpdate(() => {
		if (audioElement) {
			audioElement.load();
			audioElement.play();
		}
	});
	async function woohoosurprise() {
		showPlayer = true;
	}

	function stopSong() {
		if (audioElement.paused) {
			audioElement.play();
			document.getElementById('stop-button').style.backgroundColor = '#f44336';
			document.querySelector('#stop-button i').classList.remove('fa-play');
			document.querySelector('#stop-button i').classList.add('fa-stop');
		} else {
			audioElement.pause();
			document.getElementById('stop-button').style.backgroundColor = 'green';
			document.querySelector('#stop-button i').classList.remove('fa-stop');
			document.querySelector('#stop-button i').classList.add('fa-play');
		}
	}
</script>

<head>
	<title>Muffle</title>
	<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Lato&display=swap" />
	<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Lato&display=swap" />
	<link
		rel="stylesheet"
		href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/all.min.css"
	/>
</head>

<div class="container">
	{#if loading}
		<p>I love you ❤️, please wait</p>
	{:else if data == null}
		<p>Something went wrong, please try again :|</p>
	{:else}
		<div rel="preload" class="background" style="background-image: url({data.image})" />

		{#if showPlayer}
			<div class="player">
				<a href={data.spotifyurl} target="_blank">
					<img rel="preload" class="spotify-icon" src="/spotify.svg" alt="" />
				</a>
				<img rel="preload" class="cover" src={data.image} alt="cover" />

				<div class="track-info">
					<p class="track-name" title={data.track}>{data.track}</p>
					<p class="artist" title={data.artist}>{data.artist}</p>
					<p class="album" title={data.album}>{data.album}</p>
					<p class="year">{data.year}</p>
					<audio id="audioPlayer" loop bind:this={audioElement}>
						<source src={data.preview} />
					</audio>
				</div>
				{#if fetchingData}
					<div class="spinner">
						<i class="fas fa-spinner fa-spin" />
						<script>
							document.getElementById('audioPlayer').volume = 0;
						</script>
					</div>
				{:else}
					<script>
						document.getElementById('audioPlayer').volume = 0.4;
					</script>
					<button id="stop-button" on:click={stopSong}>
						<i class="fas fa-stop" />
					</button>
					<button class="roll-button" on:click={fetchData}>Roll the dice!🎲</button>
				{/if}
			</div>
		{:else}
			<button class="initial-roll-button" on:click={woohoosurprise}>Roll the dice!🎲</button>
		{/if}

		<footer class="footer">
			<div class="made-with-love">
				<p>Made with</p>
				<div class="icons">
					<img rel="preload" src="/heart.svg" alt="heart" />

					<a href="https://golang.org" target="_blank">
						<img rel="preload" src="/go.svg" class="go-icon" alt="golang" />
					</a>

					<a href="https://svelte.dev" target="_blank">
						<img rel="preload" src="/svelte.svg" alt="svelte" />
					</a>
				</div>
			</div>
			<a href="https://github.com/fvrrvg/" target="_blank">
				<img rel="preload" src="/github.svg" class="github-icon" alt="github" />
			</a>
		</footer>
	{/if}
</div>

<style>
	p {
		font-family: 'Lato';
	}

	.background {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 80%;
		z-index: -1;
		background-size: cover;
		background-position: center;
		filter: blur(70px) brightness(1);
	}

	.container {
		position: relative;
		overflow: hidden;
		margin-top: 30px;
		height: 90vh;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		text-align: center;
	}

	.player {
		position: fixed;
		width: 200px;
		height: 400px;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		background-color: #fff;
		padding: 20px;
		border-radius: 10px;
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
	}

	.spinner {
		margin-top: 10px;
		padding: 8px 16px;
		font-size: 20px;
		color: #000000;
	}

	.spinner i {
		animation: spin 1s linear infinite;
	}

	@keyframes spin {
		0% {
			transform: rotate(0deg);
		}
		100% {
			transform: rotate(360deg);
		}
	}

	.cover {
		width: 200px;
		height: 200px;
		object-fit: cover;
	}

	.track-name:hover {
		display: block;
	}

	.track-name {
		font-size: 20px;
		font-weight: bold;
		margin-bottom: 2px;
		margin-top: 10px;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	.artist,
	.album,
	.year {
		font-size: 15px;
		margin: 5px 0;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	button.roll-button {
		padding: 10px 20px;
		border-radius: 50px;
		width: 150px;
		height: 40px;
		background-color: rgb(0, 0, 0);
		color: #fff;
		border: none;
		cursor: pointer;
		font-size: 15px;
		font-family: 'Lato';
		position: absolute;
		bottom: 18px;
		left: 50%;
		transform: translateX(-50%);
	}

	#stop-button {
		width: 30px;
		height: 30px;
		border-radius: 50%;
		background-color: #f44336;
		color: #fff;
		border: none;
		cursor: pointer;
		font-size: 10px;
		margin-bottom: 10px;
		position: absolute;
		bottom: 60px;
		left: 50%;
		transform: translateX(-50%);
		display: flex;
		align-items: center;
		justify-content: center;
	}

	button.initial-roll-button {
		position: fixed;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		padding: 10px 20px;
		border-radius: 50px;
		background-color: rgb(0, 0, 0);
		color: #fff;
		border: none;
		cursor: pointer;
		font-size: 20px;
		font-family: 'Lato';
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.spotify-icon {
		width: 30px;
		height: 30px;
		background-size: cover;
		cursor: pointer;
		margin-bottom: 5px;
		position: absolute;
		top: -50px;
		left: 50%;
		transform: translateX(-50%);
	}

	.spotify-icon:hover {
		content: url('/spotifygreen.svg');
	}

	.footer {
		top: 600px;
		position: fixed;
		bottom: 20px;
		left: 20px;
		display: flex;
		align-items: flex-end;
		justify-content: flex-start;
		text-align: center;
		color: #000;
		font-size: 18px;
	}

	.made-with-love {
		display: flex;
		align-items: center;
	}

	.made-with-love p {
		margin-right: 5px;
	}

	.icons {
		display: flex;
		align-items: center;
	}

	.icons img {
		margin-right: 5px;
	}

	.github-icon {
		position: fixed;
		top: 20px;
		right: 20px;
		width: 20px;
		height: 20px;
		z-index: 1;
	}

	.go-icon {
		width: 60px;
		height: 55px;
	}

	img {
		width: 30px;
		height: 30px;
	}

	@media (max-width: 768px) {
		.player {
			padding: 10px;
			width: 200px;
			height: 390px;
		}
		.cover {
			width: 150px;
			height: 150px;
		}
		.track-name {
			font-size: 20px;
		}
		.artist,
		.album,
		.year {
			font-size: 15px;
		}
		button.roll-button {
			margin-bottom: 18px;
			padding: 8px 16px;
			font-size: 16px;
		}

		.github-icon {
			width: 30px;
			height: 30px;
			bottom: 700px;
		}

		#stop-button {
			width: 35px;
			height: 35px;
			font-size: 15px;
			margin-bottom: 25px;
		}

		.made-with-love {
			font-size: 15px;
		}
	}
</style>
