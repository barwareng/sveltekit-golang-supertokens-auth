<script lang="ts">
	import { VITE_SUPERTOKENS_APP_NAME } from '$lib/env';

	export let title = '';
	$: {
		if (title) {
			title = ' | ' + title;
		}
	}
</script>

<svelte:head><title>{VITE_SUPERTOKENS_APP_NAME}{title}</title></svelte:head>
