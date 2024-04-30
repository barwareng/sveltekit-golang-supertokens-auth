<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import { Input } from '$lib/components/ui/input';
	import PageHead from '$lib/components/reusable/PageHead.svelte';
	let email: string;
</script>

<PageHead title="Signup" />
<div class="mx-auto flex h-screen max-w-sm items-center justify-center">
	<Card.Root>
		<Card.Header class="space-y-1">
			<Card.Title class="text-2xl">Reset Password</Card.Title>
			<Card.Description
				>We'll send a password reset link to your email address. Please click on this link to reset
				your password</Card.Description
			>
		</Card.Header>
		<Card.Content class="grid gap-4">
			<div class="grid gap-2">
				<Input id="email" type="email" bind:value={email} placeholder="m@example.com" />
			</div>
		</Card.Content>
		<Card.Footer class="flex flex-col">
			<Button
				disabled={!email}
				class="w-full"
				on:click={() => {
					// signupWithEmailAndPassword(email, password);
				}}>Send link</Button
			>
		</Card.Footer>
	</Card.Root>
</div>
