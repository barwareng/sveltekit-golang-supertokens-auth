<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import PageHead from '$lib/components/reusable/PageHead.svelte';
	let email: string;
	let password: string;
</script>

<PageHead title="Signin" />
<div class="mx-auto flex h-screen items-center justify-center">
	<Card.Root class="min-w-80 ">
		<Card.Header class="space-y-1">
			<Card.Title class="text-2xl">Signin to your account</Card.Title>
		</Card.Header>
		<Card.Content class="grid gap-4">
			<div class="grid gap-2">
				<Label for="email">Email</Label>
				<Input id="email" type="email" bind:value={email} placeholder="m@example.com" />
			</div>
			<div class="grid gap-2">
				<Label for="password">Password</Label>
				<Input id="password" type="password" bind:value={password} />
			</div>
			<div class="-my-6 flex justify-end">
				<Button variant="link" href="/reset-password">Forgot password?</Button>
			</div>
		</Card.Content>
		<Card.Footer class="flex flex-col">
			<Button
				disabled={!password || !email}
				class="w-full"
				on:click={() => {
					// signupWithEmailAndPassword(email, password);
				}}>Signin</Button
			>
			<Button variant="link" href="/signup">Or Sign up</Button>
		</Card.Footer>
	</Card.Root>
</div>
