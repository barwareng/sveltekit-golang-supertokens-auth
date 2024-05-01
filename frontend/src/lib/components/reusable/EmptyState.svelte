<script lang="ts">
	export let title = '';
	export let subtitle = '';
</script>

<div class="flex h-96 w-full flex-col items-center justify-center text-center {$$props.class}">
	<slot name="leading-icon" />
	<h3 class="mt-2 text-sm font-semibold">{title}</h3>
	<p class="text-muted-foreground mt-1 max-w-md text-sm">{subtitle}</p>
	<div class="mt-6">
		<!-- {#if buttonName}
			<Button type="button" on:click>
				<Plus class="w-6 h-6" />
				{buttonName}
			</Button>
		{/if} -->
		<slot />
	</div>
</div>
