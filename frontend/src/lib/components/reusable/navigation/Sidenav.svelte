<script lang="ts">
	import { page } from '$app/stores';
	import { Button } from '$lib/components/ui/button';
	import {
		Building,
		Cog,
		ConstructionIcon,
		HomeIcon,
		LogOutIcon,
		PieChartIcon,
		User
	} from 'lucide-svelte';
	import { logout } from '$utils/supertokens';
	import TeamSwitcher from './TeamSwitcher.svelte';

	$: routes = [
		{
			icon: PieChartIcon,
			name: 'Dashboard',
			path: `/`,
			position: 'top'
		},
		{
			icon: Building,
			name: 'Properties',
			path: `/properties`,
			position: 'top'
		},
		{
			icon: HomeIcon,
			name: 'Units',
			path: `/units`,
			position: 'top'
		},
		{
			icon: User,
			name: 'Tenants',
			path: `/tenants`,
			position: 'top'
		},
		{
			icon: ConstructionIcon,
			name: 'Maintenance',
			path: `/maintenance`,
			position: 'top'
		},
		{
			icon: Cog,
			name: 'Settings',
			path: `/settings`,
			position: 'bottom'
		}
	];
</script>

<nav class="flex h-full flex-col justify-between px-3 py-6">
	<div class="space-y-0.5">
		<TeamSwitcher />
		{#each routes.filter((route) => route.position == 'top') as route}
			<Button
				variant={route.path == $page.url.pathname ? 'secondary' : 'ghost'}
				class="w-full items-center justify-start rounded-full text-xs"
				href={route.path}
			>
				<svelte:component this={route.icon} class="mr-1 h-4 w-4" />
				<span>
					{route.name}
				</span>
			</Button>
		{/each}
	</div>
	<div class="space-y-0.5">
		{#each routes.filter((route) => route.position == 'bottom') as route}
			<Button
				variant={route.path == $page.url.pathname ? 'secondary' : 'ghost'}
				class="w-full items-center justify-start rounded-full text-xs"
				href={route.path}
			>
				<svelte:component this={route.icon} class="mr-1 h-4 w-4" />
				<span>
					{route.name}
				</span>
			</Button>
		{/each}
		<Button
			variant="ghost"
			class="w-full items-center justify-start rounded-full text-xs"
			on:click={logout}
		>
			<LogOutIcon class="mr-1 h-4 w-4" />
			Logout
		</Button>
		<div class="pt-2">
			<Button
				variant="link"
				href="https://acme-corp.com"
				target="_blank"
				class="flex items-center justify-start gap-x-2"
			>
				<img src="/images/png/logo.png" alt="acme-corp-logo" class="aspect-square h-6" />

				<h1 class="text-base font-bold text-[#5179A7]">AcmeCorp</h1>
				<p class="text-[10px]">&copy; {new Date().getFullYear()} All rights reserved</p>
			</Button>
		</div>
	</div>
</nav>
