<script lang="ts">
	// import { Icons } from '$lib/components/docs/icons/index.js';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import { GithubIcon } from 'lucide-svelte';
	import { signinWithEmailAndPassword, signupWithEmailAndPassword } from '$utils/supertokens';
	let email: string;
	let password: string;
	$: console.log(email, password);
</script>

<div class="mx-auto flex h-screen max-w-md items-center justify-center">
	<Card.Root>
		<Card.Header class="space-y-1">
			<Card.Title class="text-2xl">Signin to your account</Card.Title>
			<!-- <Card.Description>Enter your email below to create your account</Card.Description> -->
		</Card.Header>
		<Card.Content class="grid gap-4">
			<div class="grid grid-cols-2 gap-6">
				<Button variant="outline">
					<GithubIcon class="mr-2 h-4 w-4" />
					GitHub
				</Button>
				<Button variant="outline">
					<!-- <Icons.google class="mr-2 h-4 w-4" /> -->
					Google
				</Button>
			</div>
			<div class="relative">
				<div class="absolute inset-0 flex items-center">
					<span class="w-full border-t" />
				</div>
				<div class="relative flex justify-center text-xs uppercase">
					<span class="bg-card text-muted-foreground px-2"> Or continue with </span>
				</div>
			</div>
			<div class="grid gap-2">
				<Label for="email">Email</Label>
				<Input id="email" type="email" bind:value={email} placeholder="m@example.com" />
			</div>
			<div class="grid gap-2">
				<Label for="password">Password</Label>
				<Input id="password" type="password" bind:value={password} />
			</div>
		</Card.Content>
		<Card.Footer>
			<Button
				disabled={!password || !email}
				class="w-full"
				on:click={() => {
					signinWithEmailAndPassword(email, password);
				}}>Signin</Button
			>
		</Card.Footer>
	</Card.Root>
</div>
