<script>
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import axios from 'axios';
	import moment from 'moment'

	const backendUrl = import.meta.env.VITE_BACKEND_URL

	let transferLoading = false
	let isModalOpen = false
	let transferAmount = 0
	let beneficiary_id = $page.params.beneficiary_id
	let transfers = []
	let loading = true

	onMount(() => {
		let payload = {
			beneficiary_id: parseInt(beneficiary_id)
		}
		axios.post(`${backendUrl}/get_transaction_history`, payload)
			.then((response) => {
				transfers = response.data
				transfers.reverse()
				transfers = transfers
				loading = false
			}).catch((err) => {
				console.error(err)
				loading = false
			})
	})

	const initiateTransfer = () => {
		transferLoading = true
		const payload = {
			beneficiary_id: parseInt(beneficiary_id),
			amount: transferAmount
		}

		console.log("payload", payload)

		axios.post(`${backendUrl}/initiate_transfer`, payload)
			.then((response) => {
				transferLoading = false
				isModalOpen = false
				transfers.unshift(response.data)
				transfers = transfers
			})
			.catch(() => {
				console.error("Error creating transfer")
				transferLoading = false
			})
	}

	let isTransferStatusModalOpen = false
	let transferStatus = ''
	const checkStatus = (transferId) => {
		axios.get(`${backendUrl}/check_transfer_status/` + transferId)
			.then((response) => {
				transferStatus = response.data.status
				isTransferStatusModalOpen = true
			})
			.catch((err) => {
				console.log("Error checking status")
			})
	}
</script>

<div>
	{#if !loading}
		{#if (transfers.length !== 0)}
			<div class="flex justify-center mt-20">
				<div class="overflow-x-auto w-2/3">
					<table class="table">
					<!-- head -->
					<thead>
					<tr>
						<th></th>
						<th>Time of transaction</th>
						<th>Amount</th>
						<th>Transaction status</th>
						<th class="text-center">Actions</th>
					</tr>
					</thead>
					<tbody>
					<!-- row 1 -->
					{#each transfers as transfer, i}
						<tr>
							<th>{i+1}</th>
							<td>{moment(transfer.transaction_history.timestamp).format('Do MMMM YYYY, h:mm:ss a')}</td>
							<th>{transfer.amount}</th>
							<td>{transfer.transaction_history.status}</td>
							<td class="text-center">
								<button class="btn btn-sm btn-accent" on:click={() => checkStatus(transfer.id)}>
									Check transfer status
								</button>
							</td>
						</tr>
					{/each}
					</tbody>
				</table>
				</div>
			</div>
			<div class="flex justify-center mt-4">
				<button class="btn btn-primary" on:click={() => isModalOpen = true}>Initiate a transfer</button>
			</div>
		{:else }
			<h3 class="text-center mt-20 mb-4">You have no transfers with this beneficiary</h3>
			<div class="flex justify-center">
				<button class="btn btn-primary" on:click={() => isModalOpen = true}>Initiate a transfer</button>
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
		<dialog id="my_modal_3" class="modal" class:modal-open={isTransferStatusModalOpen}>
			<div class="modal-box">
				{#if transferStatus === 'INITIATED'}
					<h3 class="text-lg font-bold">Transfer status is: <span class="badge">INITIATED</span></h3>
					<p class="py-4">The transfer request has been initiated and is awaiting processing.</p>
				{:else if transferStatus === 'PENDING'}
					<h3 class="text-lg font-bold">Transfer status is: <span class="badge badge-neutral">PENDING</span> </h3>
					<p class="py-4">The transfer is in progress and awaiting finalization.</p>
				{:else if transferStatus === 'COMPLETED'}
					<h3 class="text-lg font-bold">Transfer status is: <span class="badge badge-success">COMPLETED</span> </h3>
					<p class="py-4">The transfer has been successfully completed.</p>
				{:else if transferStatus === 'ABORTED'}
					<h3 class="text-lg font-bold">Transfer status is: <span class="badge badge-error">ABORTED</span> </h3>
					<p class="py-4">The transfer has been aborted and will not be processed.</p>
				{/if}

				<div class="modal-action">
					<form method="dialog">
						<button class="btn" on:click={() => isTransferStatusModalOpen = false}>Cancel</button>
					</form>
				</div>
			</div>
		</dialog>
	{:else}
		<div class="flex justify-center align-middle">
			<span class="loading loading-spinner loading-lg"></span>
		</div>
	{/if}
</div>
