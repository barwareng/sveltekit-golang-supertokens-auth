<script>
	import PageHeader from '$lib/components/reusable/PageHeader.svelte';
	import CreateTeam from '$lib/components/reusable/dialogs/CreateTeam.svelte';
	import '../app.pcss';
	import { ModeWatcher } from 'mode-watcher';
</script>

<PageHeader />
<ModeWatcher />
<CreateTeam />
<slot />
