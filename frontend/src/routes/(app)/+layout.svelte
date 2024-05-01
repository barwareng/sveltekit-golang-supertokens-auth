<script lang="ts">
	import Sidenav from '$lib/components/reusable/navigation/Sidenav.svelte';
	import { Separator } from '$lib/components/ui/separator';
</script>

<main class="relative mx-auto flex h-screen w-screen max-w-screen-2xl">
	<aside class="sticky top-0 hidden max-h-screen min-w-[300px] shrink-0 basis-1/5 lg:block">
		<Sidenav />
	</aside>
	<Separator orientation="vertical" />
	<section class="min-h-screen flex-1 overflow-y-scroll p-10">
		<slot />
	</section>
</main>
