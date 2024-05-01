<script lang="ts">
	import * as Alert from '$lib/components/ui/alert/index.js';
	import { TriangleAlertIcon } from 'lucide-svelte';
</script>

<div class="min-w-screen flex min-h-screen items-center justify-center">
	<Alert.Root variant="destructive" class="max-w-xs">
		<TriangleAlertIcon class="h-4 w-4" />
		<Alert.Title>Session</Alert.Title>
		<Alert.Description class="mt-2 text-sm"
			>You sessions is invalid or has expired. We're refreshing it.</Alert.Description
		>
	</Alert.Root>
</div>
