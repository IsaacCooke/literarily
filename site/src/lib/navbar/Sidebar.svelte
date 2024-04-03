<script>
	import '../css/navbar.scss';
	export let open = false;
</script>

<aside class="absolute w-full h-full navColor border-r-2 shadow-lg" class:open>
	<nav class="p-12 text-xl">
		<a href="/" class="block">Home</a>
		<a href="/about" class="block">About</a>
		<a href="/blog" class="block">Blog</a>
    <a href="/my-top-picks" class="block">My Top Picks</a>
		<a href="/contact" class="block">Contact</a>
	</nav>
</aside>

<style lang="scss">
	aside {
		left: -100%;
		transition: left 0.3s ease-in-out;
		z-index: 2;
	}

	.open {
		left: 0;
	}
</style>
