package fr.smartpark.navigator.adapters

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.databinding.DataBindingUtil
import androidx.navigation.findNavController
import androidx.navigation.fragment.FragmentNavigatorExtras
import androidx.recyclerview.widget.DiffUtil
import androidx.recyclerview.widget.ListAdapter
import androidx.recyclerview.widget.RecyclerView
import fr.smartpark.navigator.HomeFragmentDirections
import fr.smartpark.navigator.R
import fr.smartpark.navigator.data.ParkingZone
import fr.smartpark.navigator.databinding.FragmentParkingZoneItemBinding

class ParkingZoneAdapter :
    ListAdapter<ParkingZone, ParkingZoneAdapter.ViewHolder>(ParkingZoneDiffCallback()) {
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
        private fun navigateToZone(zone: ParkingZone, view: View) {
            val directions = HomeFragmentDirections.actionHomeToParkingZoneDetail(zone)

            val extras = FragmentNavigatorExtras(
                binding.zoneCard to binding.zoneCard.transitionName
            )

            view.findNavController().navigate(directions, extras)
        }

        fun bind(item: ParkingZone) {
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

private class ParkingZoneDiffCallback : DiffUtil.ItemCallback<ParkingZone>() {
    override fun areContentsTheSame(oldItem: ParkingZone, newItem: ParkingZone) =
        oldItem == newItem

    override fun areItemsTheSame(oldItem: ParkingZone, newItem: ParkingZone) =
        oldItem.zoneId == newItem.zoneId
}
