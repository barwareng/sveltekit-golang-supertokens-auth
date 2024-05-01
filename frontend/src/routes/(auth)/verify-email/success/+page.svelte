<script lang="ts">
	import * as Alert from '$lib/components/ui/alert';
	import { Button } from '$lib/components/ui/button';
	import { VITE_SUPERTOKENS_APP_NAME } from '$lib/env';
	import { showCreateTeamDialog } from '$stores';
</script>

<div class="flex h-screen w-full items-center justify-center">
	<Alert.Root class="max-w-sm sm:max-w-md">
		<Alert.Title class="text-base">Email Verified</Alert.Title>
		<Alert.Description
			>Thank you for verifying your email address. To use {VITE_SUPERTOKENS_APP_NAME} services, please
			create a team to help you manage your resources.</Alert.Description
		>
		<div class="mt-4 flex w-full justify-end">
			<Button on:click={() => ($showCreateTeamDialog = true)}>Create team</Button>
		</div>
	</Alert.Root>
</div>
