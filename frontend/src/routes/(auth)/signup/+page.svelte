<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	let email: string;
	let password: string;
	$: console.log(email, password);
</script>

<div class="mx-auto flex h-screen max-w-sm items-center justify-center">
	<Card.Root>
		<Card.Header class="space-y-1">
			<Card.Title class="text-2xl">Create an account</Card.Title>
			<Card.Description>Enter your email below to create your account</Card.Description>
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
		</Card.Content>
		<Card.Footer class="flex flex-col">
			<Button
				disabled={!password || !email}
				class="w-full"
				on:click={() => {
					// signupWithEmailAndPassword(email, password);
				}}>Create account</Button
			>
			<Button variant="link" href="/signin">Or Sign In</Button>
		</Card.Footer>
	</Card.Root>
</div>
