<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { Button } from '$lib/components/ui/button';
	import { Textarea } from '$lib/components/ui/textarea';
	import { client } from '$lib/api/Client';
	import { showCreateTeamDialog } from '$stores';
	import { goto } from '$app/navigation';

	let name: string;
	let description: string;
	const createTeam = async () => {
		try {
			await client.teams.create({ name, description });
			name = '';
			description = '';
			$showCreateTeamDialog = false;
			goto('/', { invalidateAll: true });
		} catch (error) {
			// TODO catch error
		}
	};
</script>

<Dialog.Root bind:open={$showCreateTeamDialog}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Create team</Dialog.Title>
			<Dialog.Description>Add a new team to manage products and customers.</Dialog.Description>
		</Dialog.Header>
		<div>
			<div class="space-y-4 py-2 pb-4">
				<div class="space-y-2">
					<Label for="name">Team name</Label>
					<Input id="name" placeholder="Acme Realty." bind:value={name} />
				</div>
				<div class="space-y-2">
					<Label for="description">Description</Label>
					<Textarea class="resize-none" bind:value={description} />
				</div>
			</div>
		</div>
		<Dialog.Footer>
			<Button variant="outline" on:click={() => ($showCreateTeamDialog = false)}>Cancel</Button>
			<Button type="button" on:click={createTeam}>Continue</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
