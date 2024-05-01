<script lang="ts">
	import { tick } from 'svelte';
	import { cn } from '$lib/utils.js';
	import * as Avatar from '$lib/components/ui/avatar';
	import { Button } from '$lib/components/ui/button';
	import * as Command from '$lib/components/ui/command';
	import * as Popover from '$lib/components/ui/popover';
	import { CheckIcon, ChevronsUpDownIcon, PlusCircleIcon } from 'lucide-svelte';
	import { showCreateTeamDialog } from '$stores';
	import { page } from '$app/stores';
	import { getTeamCookie, setTeamCookie } from '$utils';
	import { invalidateAll } from '$app/navigation';
	import type { ITeam } from '$utils/interfaces/teams.interface';

	let className: string | undefined | null = undefined;
	export { className as class };

	let open = false;

	$: selectedTeam =
		$page.data.teams.find((team: Partial<ITeam>) => team.id == getTeamCookie()) ??
		$page.data.teams?.[0];

	const closeAndRefocusTrigger = (triggerId: string) => {
		open = false;

		tick().then(() => document.getElementById(triggerId)?.focus());
	};
	const changeTeam = (team: Partial<ITeam>) => {
		selectedTeam = team;
		setTeamCookie(team.id!);
		invalidateAll();
	};
</script>

{#if $page.data.teams && $page.data.teams.length > 0}
	<Popover.Root bind:open let:ids>
		<Popover.Trigger asChild let:builder>
			<Button
				builders={[builder]}
				variant="outline"
				role="combobox"
				aria-expanded={open}
				aria-label="Select a team"
				class={cn('mb-4 w-full justify-between rounded-full', className)}
			>
				<Avatar.Root class="mr-2 h-5 w-5">
					<Avatar.Image
						src="https://avatar.vercel.sh/${selectedTeam.id}.png"
						alt={selectedTeam.name}
						class="grayscale"
					/>
					<Avatar.Fallback>SC</Avatar.Fallback>
				</Avatar.Root>
				{selectedTeam.name}
				<ChevronsUpDownIcon class="ml-auto h-4 w-4 shrink-0 opacity-50" />
			</Button>
		</Popover.Trigger>
		<!-- TODO adjust with relative to trigger -->
		<Popover.Content class="p-0">
			<Command.Root>
				<Command.Input placeholder="Search team..." />
				<Command.List>
					<Command.Empty>No team found.</Command.Empty>
					<Command.Group heading="Teams">
						{#each $page.data.teams as team}
							<Command.Item
								onSelect={() => {
									changeTeam(team);
									closeAndRefocusTrigger(ids.trigger);
								}}
								value={team.name}
								class="text-sm"
							>
								<Avatar.Root class="mr-2 h-5 w-5">
									<Avatar.Image
										src="https://avatar.vercel.sh/${team.id}.png"
										alt={team.name}
										class="grayscale"
									/>
									<Avatar.Fallback></Avatar.Fallback>
								</Avatar.Root>
								{team.name}
								<CheckIcon
									class={cn('ml-auto h-4 w-4', selectedTeam.id !== team.id && 'text-transparent')}
								/>
							</Command.Item>
						{/each}
					</Command.Group>
				</Command.List>
				<Command.Separator />
				<Command.List>
					<Command.Group>
						<Command.Item
							onSelect={() => {
								open = false;
								$showCreateTeamDialog = true;
							}}
						>
							<PlusCircleIcon class="mr-2 h-5 w-5" />
							Create Team
						</Command.Item>
					</Command.Group>
				</Command.List>
			</Command.Root>
		</Popover.Content>
	</Popover.Root>
{:else}
	<Button class="mb-4 w-full" on:click={() => ($showCreateTeamDialog = true)}>
		<PlusCircleIcon class="mr-2 h-5 w-5" />
		Create Team</Button
	>
{/if}
