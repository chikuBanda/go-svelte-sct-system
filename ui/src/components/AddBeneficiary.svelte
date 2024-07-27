<script>
	import axios from 'axios';
	import { goto } from '$app/navigation';
	import { beneficiaries } from '../stores/beneficiaries'

	const backendUrl = import.meta.env.VITE_BACKEND_URL

	export let classes = ''
	let isModalOpen = false
	let loading = false

	let beneficiaryName = ''
	let beneficiaryAccountNumber = ''
	const addBeneficiary = () => {
		loading = true
		const payload = {
			name: beneficiaryName,
			account_number: beneficiaryAccountNumber
		}

		axios.post(`${backendUrl}/register_beneficiary`, payload)
			.then((response) => {
				loading = false
				isModalOpen = false
				beneficiaries.update((value) => {
					value.push(response.data)
					return value
				})
				goto("/")
			})
			.catch(() => {
				console.error("Error creating beneficiary")
				loading = false
			})
	}
</script>

<span class="text-slate-900">
	<button class={classes} on:click={() => isModalOpen = true}>
		Add a beneficiary
	</button>

	<dialog id="my_modal_2" class="modal" class:modal-open={isModalOpen}>
		<div class="modal-box">
			<h3 class="text-lg font-bold">Add beneficiary details</h3>
			<label class="form-control w-full py-4">
				<div class="label">
					<span class="label-text">Full name</span>
				</div>
				<input
					type="text"
					bind:value={beneficiaryName}
					placeholder="Enter name"
					class="input input-bordered w-full max-w-xs" />
			</label>

			<label class="form-control w-full py-4">
				<div class="label">
					<span class="label-text">Account number</span>
				</div>
				<input
					type="text"
					bind:value={beneficiaryAccountNumber}
					placeholder="Enter account number"
					class="input input-bordered w-full max-w-xs" />
			</label>

			<div class="modal-action">
				<form method="dialog">
					<button class="btn" on:click={() => isModalOpen = false}>Cancel</button>
					<button class="btn btn-accent" on:click={()=> addBeneficiary()}>
						{#if loading}
							<span class="loading loading-spinner loading-lg"></span>
						{:else}
							submit
						{/if}
					</button>
				</form>
			</div>
		</div>
	</dialog>
</span>