package fr.smartpark.navigator.adapters

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.databinding.DataBindingUtil
import androidx.navigation.findNavController
import androidx.recyclerview.widget.DiffUtil
import androidx.recyclerview.widget.ListAdapter
import androidx.recyclerview.widget.RecyclerView
import fr.smartpark.navigator.R
import fr.smartpark.navigator.TenantListFragmentDirections
import fr.smartpark.navigator.data.models.Tenant
import fr.smartpark.navigator.databinding.FragmentTenantItemBinding

class TenantAdapter :
    ListAdapter<Tenant, TenantAdapter.ViewHolder>(TenantDiffCallback()) {
    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ViewHolder {
        return ViewHolder(
            DataBindingUtil.inflate(
                LayoutInflater.from(parent.context),
                R.layout.fragment_tenant_item,
                parent,
                false
            )
        )
    }

    override fun onBindViewHolder(holder: ViewHolder, position: Int) {
        with(holder) {
            bind(getItem(position))
        }
    }

    class ViewHolder(
        private val binding: FragmentTenantItemBinding
    ) : RecyclerView.ViewHolder(binding.root) {
        private fun navigateToZoneList(tenant: Tenant, view: View) {
            val directions = TenantListFragmentDirections.actionTenantListToZones(tenant.tenantId)
            view.findNavController().navigate(directions)
        }

        fun bind(item: Tenant) {
            binding.apply {
                tenant = item

                setClickListener { view ->
                    navigateToZoneList(item, view)
                }

                executePendingBindings()
            }
        }
    }
}

private class TenantDiffCallback : DiffUtil.ItemCallback<Tenant>() {
    override fun areContentsTheSame(oldItem: Tenant, newItem: Tenant) =
        oldItem == newItem

    override fun areItemsTheSame(oldItem: Tenant, newItem: Tenant) =
        oldItem.tenantId == newItem.tenantId
}
