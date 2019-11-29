package fr.smartpark.navigator.adapters

import android.view.LayoutInflater
import android.view.ViewGroup
import androidx.databinding.DataBindingUtil
import androidx.recyclerview.widget.RecyclerView
import androidx.recyclerview.widget.ListAdapter
import androidx.recyclerview.widget.DiffUtil
import fr.smartpark.navigator.R
import fr.smartpark.navigator.data.ParkingZone
import fr.smartpark.navigator.databinding.FragmentParkingZoneItemBinding

class ParkingZoneAdapter : ListAdapter<ParkingZone, ParkingZoneAdapter.ViewHolder>(ParkingZoneDiffCallback()) {
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
        fun bind(item: ParkingZone) {
            binding.apply {
                zone = item
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
