<script lang="ts">
	import * as Card from '$lib/components/ui/card/index.js';
	import * as Carousel from '$lib/components/ui/carousel/index.js';
	import * as Dialog from '$lib/components/ui/dialog/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { getRobot } from '$lib/api/api.js';
	import { onMount } from 'svelte';
	import type { Robot } from '$lib/types/robot.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Plus, Trash, Pencil } from 'lucide-svelte';

	let robots: Robot[] | null = null;

	onMount(async () => {
		try {
			robots = await getRobot();
			console.log(robots);
		} catch (error) {
			console.error('Failed to fetch robot data:', error);
		}
	});
</script>

<svelte:head>
	<title>Robot</title>
	<meta name="description" content="Robot configuration" />
</svelte:head>

<div class="relative flex flex-row items-center justify-center">
	<div class="flex h-full w-full flex-1 flex-col items-center justify-center">
		<div class="flex h-full w-[80%] flex-col justify-center">
			<Carousel.Root>
				{#if robots}
					<Carousel.Content>
						{#each robots as robot}
							<Carousel.Item>
								<Card.Root>
									<Card.Content class="flex aspect-square items-center justify-center p-6">
										<img
											src={'http://localhost:8080' + robot.filename}
											alt="Robot"
											class="h-fit w-fit"
										/>
									</Card.Content>
								</Card.Root>
							</Carousel.Item>
						{/each}
					</Carousel.Content>
				{:else}
					<p class="absolute left-1/2 top-1/2 origin-center">Loading...</p>
				{/if}
				<Carousel.Previous />
				<Carousel.Next />
			</Carousel.Root>
		</div>
	</div>

	<div class="absolute bottom-0 right-0">
		<div class="w-17 flex flex-col items-center justify-center gap-1 rounded-full border p-1">
			<Dialog.Root>
				<Dialog.Trigger>
					<Button
						variant="ghost"
						class="h-15 w-15 flex items-center justify-center rounded-full p-0"
					>
						<Plus class="!h-[25px] !w-[25px]" />
					</Button>
				</Dialog.Trigger>
				<Dialog.Content class="sm:max-w-[425px]">
					<Dialog.Header>
						<Dialog.Title>Create Robot</Dialog.Title>
						<Dialog.Description>Fill in informations of the new robot.</Dialog.Description>
					</Dialog.Header>
					<div class="grid grid-cols-4 items-center gap-4">
						<Label for="name" class="text-right">Name</Label>
					</div>
					<Dialog.Footer>
						<Button type="submit">Yes</Button>
						<Button variant="outline">Cancel</Button>
					</Dialog.Footer>
				</Dialog.Content>
			</Dialog.Root>

			<Dialog.Root>
				<Dialog.Trigger>
					<Button
						variant="ghost"
						class="h-15 w-15 flex items-center justify-center rounded-full p-0"
					>
						<Pencil class="!h-[25px] !w-[25px]" />
					</Button>
				</Dialog.Trigger>
				<Dialog.Content class="sm:max-w-[425px]">
					<Dialog.Header>
						<Dialog.Title>Edit Robot</Dialog.Title>
						<Dialog.Description>Edit informations of the robot.</Dialog.Description>
					</Dialog.Header>
					<div class="grid grid-cols-4 items-center gap-4">
						<Label for="name" class="text-right">Name</Label>
					</div>
					<Dialog.Footer>
						<Button type="submit">Yes</Button>
						<Button variant="outline">Cancel</Button>
					</Dialog.Footer>
				</Dialog.Content>
			</Dialog.Root>

			<Dialog.Root>
				<Dialog.Trigger>
					<Button
						variant="ghost"
						class="h-15 w-15 flex items-center justify-center rounded-full p-0"
					>
						<Trash class="!h-[25px] !w-[25px]" />
					</Button>
				</Dialog.Trigger>
				<Dialog.Content class="sm:max-w-[425px]">
					<Dialog.Header>
						<Dialog.Title>Delete Robot</Dialog.Title>
						<Dialog.Description>Make sure you want to delete this robot.</Dialog.Description>
					</Dialog.Header>
					<div class="grid grid-cols-4 items-center gap-4">
						<Label for="name" class="text-right">Name</Label>
					</div>
					<Dialog.Footer>
						<Button type="submit">Yes</Button>
						<Button variant="outline">Cancel</Button>
					</Dialog.Footer>
				</Dialog.Content>
			</Dialog.Root>
		</div>
	</div>
</div>
