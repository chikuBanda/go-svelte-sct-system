<script>
	import axios from 'axios'
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { beneficiaries } from '../stores/beneficiaries'

	let isModalOpen = false
	let loading = true
	let selectedBeneficicaryId = null
	let transferAmount = 0
	let beneficiaries_value = []

	beneficiaries.subscribe((value) => {
		beneficiaries_value = value;
	});

	onMount(() => {
		axios.get("http://localhost:8080/get_beneficiaries")
			.then((response) => {
				beneficiaries.set(response.data)
				console.log(beneficiaries_value.length)
				console.log("beneficiaries", beneficiaries_value)
				loading = false
			}).catch((err) => {
			console.error(err)
			loading = false
		})
	})

	const openTransferModal = (beneficiary_id) => {
		isModalOpen = true
		selectedBeneficicaryId = beneficiary_id
	}

	let transferLoading = false
	const initiateTransfer = () => {
		transferLoading = true
		const payload = {
			beneficiary_id: selectedBeneficicaryId,
			amount: transferAmount
		}

		axios.post("http://localhost:8080/initiate_transfer", payload)
			.then((response) => {
					transferLoading = false
					isModalOpen = false
					goto("/transfers/" + selectedBeneficicaryId)
			})
			.catch(() => {
					console.error("Error creating transfer")
					transferLoading = false
			})
	}
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	{#if !loading}
		{#if (beneficiaries_value.length !== 0)}
			<div class="flex justify-center mt-20">
				<div class="overflow-x-auto w-2/3">
					<table class="table">
						<!-- head -->
						<thead>
						<tr>
							<th></th>
							<th>Name</th>
							<th>Account number</th>
							<th class="text-center">Action</th>
						</tr>
						</thead>
						<tbody>
						<!-- row 1 -->
						{#each beneficiaries_value as beneficiary, i}
							<tr>
								<th>{i+1}</th>
								<td>{beneficiary.name}</td>
								<td>{beneficiary.account_number}</td>
								<td class="text-center">
									<a href={'/transfers/' + beneficiary.id} class="btn btn-sm">View transactions</a>
									<button class="btn btn-sm btn-accent" on:click={() => openTransferModal(beneficiary.id)}>Initiate transaction</button>
								</td>
							</tr>
						{/each}
						</tbody>
					</table>
				</div>
			</div>
		{:else }
			<h3 class="text-center">You have no beneficiaries </h3>
			<div class="flex justify-center align-middle">
				<button class="btn btn-primary">Add a beneficiary</button>
			</div>
		{/if}

		<dialog id="my_modal_1" class="modal" class:modal-open={isModalOpen}>
			<div class="modal-box">
				<h3 class="text-lg font-bold">Initiate your transfer!</h3>
				<label class="form-control w-full py-4">
					<div class="label">
						<span class="label-text">Amount of transfer</span>
					</div>
					<input
						type="number"
						bind:value={transferAmount}
						placeholder="Enter amount"
						class="input input-bordered w-full max-w-xs" />
				</label>

				<div class="modal-action">
					<form method="dialog">
						<button class="btn" on:click={() => isModalOpen = false}>Cancel</button>
						<button class="btn btn-accent" on:click={()=> initiateTransfer()}>
							{#if transferLoading}
								<span class="loading loading-spinner loading-lg"></span>
							{:else}
								submit
							{/if}
						</button>
					</form>
				</div>
			</div>
		</dialog>
	{:else}
		<div class="flex justify-center align-middle">
			<span class="loading loading-spinner loading-lg"></span>
		</div>
	{/if}
</section>

<style>
</style>
