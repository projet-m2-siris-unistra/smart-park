package fr.smartpark.navigator.adapters

import android.util.Log
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.databinding.DataBindingUtil
import androidx.navigation.findNavController
import androidx.navigation.fragment.FragmentNavigatorExtras
import androidx.recyclerview.widget.DiffUtil
import androidx.recyclerview.widget.ListAdapter
import androidx.recyclerview.widget.RecyclerView
import fr.smartpark.navigator.ParkingZoneListFragmentDirections
import fr.smartpark.navigator.R
import fr.smartpark.navigator.data.models.Zone
import fr.smartpark.navigator.databinding.FragmentParkingZoneItemBinding

class ParkingZoneAdapter :
    ListAdapter<Zone, ParkingZoneAdapter.ViewHolder>(ParkingZoneDiffCallback()) {
    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ViewHolder {
        return ViewHolder(
            DataBindingUtil.inflate(
                LayoutInflater.from(parent.context),
                R.layout.fragment_parking_zone_item,
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
        private val binding: FragmentParkingZoneItemBinding
    ) : RecyclerView.ViewHolder(binding.root) {
        private fun navigateToZone(zone: Zone, view: View) {
            val directions = ParkingZoneListFragmentDirections.actionZoneListToDetail(zone)

            val extras = FragmentNavigatorExtras(
                binding.zoneCardContent.root to binding.zoneCardContent.root.transitionName
            )

            Log.d("GRAPH", view.findNavController().graph.toString())
            //view.findNavController().navigate(directions, extras)
        }

        fun bind(item: Zone) {
            binding.apply {
                zone = item

                setClickListener { view ->
                    navigateToZone(item, view)
                }

                executePendingBindings()
            }
        }
    }
}

private class ParkingZoneDiffCallback : DiffUtil.ItemCallback<Zone>() {
    override fun areContentsTheSame(oldItem: Zone, newItem: Zone) =
        oldItem == newItem

    override fun areItemsTheSame(oldItem: Zone, newItem: Zone) =
        oldItem.zoneId == newItem.zoneId
}
